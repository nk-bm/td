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

// MessagesGetDiscussionMessageRequest represents TL type `messages.getDiscussionMessage#446972fd`.
type MessagesGetDiscussionMessageRequest struct {
	// Peer field of MessagesGetDiscussionMessageRequest.
	Peer InputPeerClass
	// MsgID field of MessagesGetDiscussionMessageRequest.
	MsgID int
}

// MessagesGetDiscussionMessageRequestTypeID is TL type id of MessagesGetDiscussionMessageRequest.
const MessagesGetDiscussionMessageRequestTypeID = 0x446972fd

// Ensuring interfaces in compile-time for MessagesGetDiscussionMessageRequest.
var (
	_ bin.Encoder     = &MessagesGetDiscussionMessageRequest{}
	_ bin.Decoder     = &MessagesGetDiscussionMessageRequest{}
	_ bin.BareEncoder = &MessagesGetDiscussionMessageRequest{}
	_ bin.BareDecoder = &MessagesGetDiscussionMessageRequest{}
)

func (g *MessagesGetDiscussionMessageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDiscussionMessageRequest) String() string {
	if g == nil {
		return "MessagesGetDiscussionMessageRequest(nil)"
	}
	type Alias MessagesGetDiscussionMessageRequest
	return fmt.Sprintf("MessagesGetDiscussionMessageRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetDiscussionMessageRequest) TypeID() uint32 {
	return MessagesGetDiscussionMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetDiscussionMessageRequest) TypeName() string {
	return "messages.getDiscussionMessage"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetDiscussionMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getDiscussionMessage",
		ID:   MessagesGetDiscussionMessageRequestTypeID,
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
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetDiscussionMessageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDiscussionMessage#446972fd as nil")
	}
	b.PutID(MessagesGetDiscussionMessageRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetDiscussionMessageRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDiscussionMessage#446972fd as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getDiscussionMessage#446972fd: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getDiscussionMessage#446972fd: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDiscussionMessageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDiscussionMessage#446972fd to nil")
	}
	if err := b.ConsumeID(MessagesGetDiscussionMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDiscussionMessage#446972fd: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetDiscussionMessageRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDiscussionMessage#446972fd to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDiscussionMessage#446972fd: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDiscussionMessage#446972fd: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetDiscussionMessageRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetMsgID returns value of MsgID field.
func (g *MessagesGetDiscussionMessageRequest) GetMsgID() (value int) {
	if g == nil {
		return
	}
	return g.MsgID
}

// MessagesGetDiscussionMessage invokes method messages.getDiscussionMessage#446972fd returning error if any.
func (c *Client) MessagesGetDiscussionMessage(ctx context.Context, request *MessagesGetDiscussionMessageRequest) (*MessagesDiscussionMessage, error) {
	var result MessagesDiscussionMessage

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
