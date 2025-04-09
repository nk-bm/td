// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// MessagesGetScheduledHistoryRequest represents TL type `messages.getScheduledHistory#f516760b`.
type MessagesGetScheduledHistoryRequest struct {
	// Peer field of MessagesGetScheduledHistoryRequest.
	Peer InputPeerClass
	// Hash field of MessagesGetScheduledHistoryRequest.
	Hash int64
}

// MessagesGetScheduledHistoryRequestTypeID is TL type id of MessagesGetScheduledHistoryRequest.
const MessagesGetScheduledHistoryRequestTypeID = 0xf516760b

// Ensuring interfaces in compile-time for MessagesGetScheduledHistoryRequest.
var (
	_ bin.Encoder     = &MessagesGetScheduledHistoryRequest{}
	_ bin.Decoder     = &MessagesGetScheduledHistoryRequest{}
	_ bin.BareEncoder = &MessagesGetScheduledHistoryRequest{}
	_ bin.BareDecoder = &MessagesGetScheduledHistoryRequest{}
)

func (g *MessagesGetScheduledHistoryRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetScheduledHistoryRequest) String() string {
	if g == nil {
		return "MessagesGetScheduledHistoryRequest(nil)"
	}
	type Alias MessagesGetScheduledHistoryRequest
	return fmt.Sprintf("MessagesGetScheduledHistoryRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetScheduledHistoryRequest) TypeID() uint32 {
	return MessagesGetScheduledHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetScheduledHistoryRequest) TypeName() string {
	return "messages.getScheduledHistory"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetScheduledHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getScheduledHistory",
		ID:   MessagesGetScheduledHistoryRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetScheduledHistoryRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getScheduledHistory#f516760b as nil")
	}
	b.PutID(MessagesGetScheduledHistoryRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetScheduledHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getScheduledHistory#f516760b as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getScheduledHistory#f516760b: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getScheduledHistory#f516760b: field peer: %w", err)
	}
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetScheduledHistoryRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getScheduledHistory#f516760b to nil")
	}
	if err := b.ConsumeID(MessagesGetScheduledHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getScheduledHistory#f516760b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetScheduledHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getScheduledHistory#f516760b to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getScheduledHistory#f516760b: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getScheduledHistory#f516760b: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetScheduledHistoryRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetHash returns value of Hash field.
func (g *MessagesGetScheduledHistoryRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetScheduledHistory invokes method messages.getScheduledHistory#f516760b returning error if any.
func (c *Client) MessagesGetScheduledHistory(ctx context.Context, request *MessagesGetScheduledHistoryRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
