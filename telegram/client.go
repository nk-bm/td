package telegram

import (
	"context"
	"errors"
	"fmt"
	"io"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"go.uber.org/atomic"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/clock"
	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/internal/mt"
	"github.com/gotd/td/internal/mtproto"
	"github.com/gotd/td/internal/pool"
	"github.com/gotd/td/internal/proto"
	"github.com/gotd/td/internal/tdsync"
	"github.com/gotd/td/internal/tmap"
	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/telegram/internal/manager"
	"github.com/gotd/td/tg"
)

// UpdateHandler will be called on received updates from Telegram.
type UpdateHandler interface {
	Handle(ctx context.Context, u *tg.Updates) error
	HandleShort(ctx context.Context, u *tg.UpdateShort) error
}

type clientStorage interface {
	Load(ctx context.Context) (*session.Data, error)
	Save(ctx context.Context, data *session.Data) error
}

type clientConn interface {
	Run(ctx context.Context) error
	InvokeRaw(ctx context.Context, input bin.Encoder, output bin.Decoder) error
}

// Client represents a MTProto client to Telegram.
type Client struct {
	// tg provides RPC calls via Client.
	tg *tg.Client // immutable

	// Telegram device information.
	device DeviceConfig // immutable

	// MTProto options.
	opts mtproto.Options // immutable

	// Connection state. Guarded by connMux.
	session *pool.SyncSession
	cfg     *manager.AtomicConfig
	conn    clientConn
	connMux sync.Mutex
	// Connection factory fields.
	create       connConstructor        // immutable
	resolver     dcs.Resolver           // immutable
	connBackoff  func() backoff.BackOff // immutable
	defaultMode  manager.ConnMode       // immutable
	connsCounter atomic.Int64

	// Restart signal channel.
	restart chan struct{} // immutable
	// Migration state.
	exported         chan *tg.AuthExportedAuthorization // immutable
	migrationTimeout time.Duration                      // immutable
	migration        chan struct{}

	// Connections to non-primary DC.
	subConns    map[int]CloseInvoker
	subConnsMux sync.Mutex
	sessions    map[int]*pool.SyncSession
	sessionsMux sync.Mutex

	// Wrappers for external world, like logs or PRNG.
	rand  io.Reader   // immutable
	log   *zap.Logger // immutable
	clock clock.Clock // immutable

	// Client context. Will be canceled by Run on exit.
	ctx    context.Context    // immutable
	cancel context.CancelFunc // immutable

	// Client config.
	appID   int    // immutable
	appHash string // immutable
	// Session storage.
	storage clientStorage // immutable, nillable

	// Ready signal channel, sends signal when client connection is ready.
	// Resets on reconnect.
	ready *tdsync.ResetReady // immutable

	// Telegram updates handler.
	updateHandler UpdateHandler // immutable
}

func (c *Client) onMessage(b *bin.Buffer) error {
	return c.handleUpdates(b)
}

// getVersion optimistically gets current client version.
//
// Does not handle replace directives.
func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	// Hard-coded package name. Probably we can generate this via parsing
	// the go.mod file.
	const pkg = "github.com/gotd/td"
	for _, d := range info.Deps {
		if strings.HasPrefix(d.Path, pkg) {
			return d.Version
		}
	}
	return ""
}

// Port is default port used by telegram.
const Port = 443

// NewClient creates new unstarted client.
func NewClient(appID int, appHash string, opt Options) *Client {
	opt.setDefaults()

	mode := manager.ConnModeUpdates
	if opt.NoUpdates {
		mode = manager.ConnModeData
	}
	clientCtx, clientCancel := context.WithCancel(context.Background())
	client := &Client{
		rand:          opt.Random,
		log:           opt.Logger,
		ctx:           clientCtx,
		cancel:        clientCancel,
		appID:         appID,
		appHash:       appHash,
		updateHandler: opt.UpdateHandler,
		session: pool.NewSyncSession(pool.Session{
			DC: opt.DC,
		}),
		cfg: manager.NewAtomicConfig(tg.Config{
			DCOptions: opt.DCList,
		}),
		create:           defaultConstructor(),
		resolver:         opt.Resolver,
		defaultMode:      mode,
		connBackoff:      opt.ReconnectionBackoff,
		clock:            opt.Clock,
		device:           opt.Device,
		migrationTimeout: opt.MigrationTimeout,
	}
	client.init()

	// Including version into client logger to help with debugging.
	if v := getVersion(); v != "" {
		client.log = client.log.With(zap.String("v", v))
	}

	if opt.SessionStorage != nil {
		client.storage = &session.Loader{
			Storage: opt.SessionStorage,
		}
	}

	client.opts = mtproto.Options{
		PublicKeys:    opt.PublicKeys,
		Random:        opt.Random,
		Logger:        opt.Logger,
		AckBatchSize:  opt.AckBatchSize,
		AckInterval:   opt.AckInterval,
		RetryInterval: opt.RetryInterval,
		MaxRetries:    opt.MaxRetries,
		MessageID:     opt.MessageID,
		Clock:         opt.Clock,

		Types: tmap.New(
			tg.TypesMap(),
			mt.TypesMap(),
			proto.TypesMap(),
		),
	}
	client.conn = client.createPrimaryConn(nil)

	return client
}

// init sets fields which needs explicit initialization, like maps or channels.
func (c *Client) init() {
	if c.cfg == nil {
		c.cfg = manager.NewAtomicConfig(tg.Config{})
	}
	c.ready = tdsync.NewResetReady()
	c.restart = make(chan struct{})
	c.migration = make(chan struct{}, 1)
	c.exported = make(chan *tg.AuthExportedAuthorization, 1)
	c.sessions = map[int]*pool.SyncSession{}
	c.subConns = map[int]CloseInvoker{}
	// Initializing internal RPC caller.
	c.tg = tg.NewClient(c)
}

func (c *Client) restoreConnection(ctx context.Context) error {
	if c.storage == nil {
		return nil
	}

	data, err := c.storage.Load(ctx)
	if errors.Is(err, session.ErrNotFound) {
		return nil
	}
	if err != nil {
		return xerrors.Errorf("load: %w", err)
	}

	// If file does not contain DC ID, so we use DC from options.
	prev := c.session.Load()
	if data.DC == 0 {
		data.DC = prev.DC
	}

	// Restoring persisted auth key.
	var key crypto.AuthKey
	copy(key.Value[:], data.AuthKey)
	copy(key.ID[:], data.AuthKeyID)

	if key.Value.ID() != key.ID {
		return xerrors.New("corrupted key")
	}

	// Re-initializing connection from persisted state.
	c.log.Info("Connection restored from state",
		zap.String("addr", data.Addr),
		zap.String("key_id", fmt.Sprintf("%x", data.AuthKeyID)),
	)

	c.connMux.Lock()
	c.session.Store(pool.Session{
		DC:      data.DC,
		AuthKey: key,
		Salt:    data.Salt,
	})
	c.conn = c.createPrimaryConn(nil)
	c.connMux.Unlock()

	return nil
}

func (c *Client) runUntilRestart(ctx context.Context) error {
	g := tdsync.NewCancellableGroup(ctx)
	g.Go(c.conn.Run)
	g.Go(func(ctx context.Context) error {
		self, err := c.Self(ctx)
		if err != nil {
			// Ignore unauthorized errors.
			if !unauthorized(err) {
				c.log.Warn("Got error on self", zap.Error(err))
			}
			return nil
		}

		c.log.Info("Got self", zap.String("username", self.Username))
		return nil
	})
	g.Go(func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.restart:
			c.log.Debug("Restart triggered")
			// Should call cancel() to cancel group.
			g.Cancel()

			return nil
		}
	})

	return g.Wait()
}

func (c *Client) reconnectUntilClosed(ctx context.Context) error {
	// Note that we currently have no timeout on connection, so this is
	// potentially eternal.
	b := tdsync.SyncBackoff(backoff.WithContext(c.connBackoff(), ctx))

	return backoff.Retry(func() error {
		err := c.runUntilRestart(ctx)
		if err == nil {
			return nil
		}
		c.log.Info("Restarting connection", zap.Error(err))

		c.connMux.Lock()
		setup := func(ctx context.Context, invoker tg.Invoker) error {
			// Setup function call means successful connection
			// initialization, so we can reset backoff.
			b.Reset()

			raw := tg.NewClient(invoker)
			select {
			case export := <-c.exported:
				_, err := raw.AuthImportAuthorization(ctx, &tg.AuthImportAuthorizationRequest{
					ID:    export.ID,
					Bytes: export.Bytes,
				})
				return err
			default:
			}
			return nil
		}
		c.conn = c.createPrimaryConn(setup)
		c.connMux.Unlock()
		return err
	}, b)
}

func (c *Client) onReady() {
	c.log.Debug("Ready")
	c.ready.Signal()
}

func (c *Client) resetReady() {
	c.ready.Reset()
}

// Run starts client session and block until connection close.
// The f callback is called on successful session initialization and Run
// will return on f() result.
//
// Context of callback will be canceled if fatal error is detected.
func (c *Client) Run(ctx context.Context, f func(ctx context.Context) error) (err error) {
	select {
	case <-c.ctx.Done():
		return xerrors.Errorf("client already closed: %w", c.ctx.Err())
	default:
	}

	c.log.Info("Starting")
	defer c.log.Info("Closed")
	// Cancel client on exit.
	defer c.cancel()
	defer func() {
		c.subConnsMux.Lock()
		defer c.subConnsMux.Unlock()

		for _, conn := range c.subConns {
			if closeErr := conn.Close(); !xerrors.Is(closeErr, context.Canceled) {
				multierr.AppendInto(&err, closeErr)
			}
		}
	}()

	c.resetReady()
	if err := c.restoreConnection(ctx); err != nil {
		return err
	}

	g := tdsync.NewCancellableGroup(ctx)
	g.Go(c.reconnectUntilClosed)
	g.Go(func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			c.cancel()
			return ctx.Err()
		case <-c.ctx.Done():
			return c.ctx.Err()
		}
	})
	g.Go(func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.ready.Ready():
			if err := f(ctx); err != nil {
				return xerrors.Errorf("callback: %w", err)
			}
			// Should call cancel() to cancel ctx.
			// This will terminate c.conn.Run().
			c.log.Debug("Callback returned, stopping")
			g.Cancel()
			return nil
		}
	})
	if err := g.Wait(); !xerrors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (c *Client) saveSession(cfg tg.Config, s mtproto.Session) error {
	if c.storage == nil {
		return nil
	}

	data, err := c.storage.Load(c.ctx)
	if errors.Is(err, session.ErrNotFound) {
		// Initializing new state.
		err = nil
		data = &session.Data{}
	}
	if err != nil {
		return xerrors.Errorf("load: %w", err)
	}

	// Updating previous data.
	data.Config = cfg
	data.AuthKey = s.Key.Value[:]
	data.AuthKeyID = s.Key.ID[:]
	data.DC = cfg.ThisDC
	data.Salt = s.Salt

	if err := c.storage.Save(c.ctx, data); err != nil {
		return xerrors.Errorf("save: %w", err)
	}

	c.log.Debug("Data saved",
		zap.String("key_id", fmt.Sprintf("%x", data.AuthKeyID)),
	)
	return nil
}

func (c *Client) onSession(cfg tg.Config, s mtproto.Session) error {
	c.sessionsMux.Lock()
	c.sessions[cfg.ThisDC] = pool.NewSyncSession(pool.Session{
		DC:      cfg.ThisDC,
		Salt:    s.Salt,
		AuthKey: s.Key,
	})
	c.sessionsMux.Unlock()

	primaryDC := c.session.Load().DC
	// Do not save session for non-primary DC.
	if cfg.ThisDC != 0 && primaryDC != 0 && primaryDC != cfg.ThisDC {
		return nil
	}

	if err := c.saveSession(cfg, s); err != nil {
		return xerrors.Errorf("save: %w", err)
	}

	c.connMux.Lock()
	c.session.Store(pool.Session{
		DC:      cfg.ThisDC,
		Salt:    s.Salt,
		AuthKey: s.Key,
	})
	c.cfg.Store(cfg)
	c.onReady()
	c.connMux.Unlock()

	return nil
}
