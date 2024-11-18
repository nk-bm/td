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

// MessagesGetPreparedInlineMessageRequest represents TL type `messages.getPreparedInlineMessage#857ebdb8`.
//
// See https://core.telegram.org/method/messages.getPreparedInlineMessage for reference.
type MessagesGetPreparedInlineMessageRequest struct {
	// Bot field of MessagesGetPreparedInlineMessageRequest.
	Bot InputUserClass
	// ID field of MessagesGetPreparedInlineMessageRequest.
	ID string
}

// MessagesGetPreparedInlineMessageRequestTypeID is TL type id of MessagesGetPreparedInlineMessageRequest.
const MessagesGetPreparedInlineMessageRequestTypeID = 0x857ebdb8

// Ensuring interfaces in compile-time for MessagesGetPreparedInlineMessageRequest.
var (
	_ bin.Encoder     = &MessagesGetPreparedInlineMessageRequest{}
	_ bin.Decoder     = &MessagesGetPreparedInlineMessageRequest{}
	_ bin.BareEncoder = &MessagesGetPreparedInlineMessageRequest{}
	_ bin.BareDecoder = &MessagesGetPreparedInlineMessageRequest{}
)

func (g *MessagesGetPreparedInlineMessageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Bot == nil) {
		return false
	}
	if !(g.ID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetPreparedInlineMessageRequest) String() string {
	if g == nil {
		return "MessagesGetPreparedInlineMessageRequest(nil)"
	}
	type Alias MessagesGetPreparedInlineMessageRequest
	return fmt.Sprintf("MessagesGetPreparedInlineMessageRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetPreparedInlineMessageRequest from given interface.
func (g *MessagesGetPreparedInlineMessageRequest) FillFrom(from interface {
	GetBot() (value InputUserClass)
	GetID() (value string)
}) {
	g.Bot = from.GetBot()
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetPreparedInlineMessageRequest) TypeID() uint32 {
	return MessagesGetPreparedInlineMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetPreparedInlineMessageRequest) TypeName() string {
	return "messages.getPreparedInlineMessage"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetPreparedInlineMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getPreparedInlineMessage",
		ID:   MessagesGetPreparedInlineMessageRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetPreparedInlineMessageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPreparedInlineMessage#857ebdb8 as nil")
	}
	b.PutID(MessagesGetPreparedInlineMessageRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetPreparedInlineMessageRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPreparedInlineMessage#857ebdb8 as nil")
	}
	if g.Bot == nil {
		return fmt.Errorf("unable to encode messages.getPreparedInlineMessage#857ebdb8: field bot is nil")
	}
	if err := g.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getPreparedInlineMessage#857ebdb8: field bot: %w", err)
	}
	b.PutString(g.ID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetPreparedInlineMessageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPreparedInlineMessage#857ebdb8 to nil")
	}
	if err := b.ConsumeID(MessagesGetPreparedInlineMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getPreparedInlineMessage#857ebdb8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetPreparedInlineMessageRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPreparedInlineMessage#857ebdb8 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPreparedInlineMessage#857ebdb8: field bot: %w", err)
		}
		g.Bot = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPreparedInlineMessage#857ebdb8: field id: %w", err)
		}
		g.ID = value
	}
	return nil
}

// GetBot returns value of Bot field.
func (g *MessagesGetPreparedInlineMessageRequest) GetBot() (value InputUserClass) {
	if g == nil {
		return
	}
	return g.Bot
}

// GetID returns value of ID field.
func (g *MessagesGetPreparedInlineMessageRequest) GetID() (value string) {
	if g == nil {
		return
	}
	return g.ID
}

// MessagesGetPreparedInlineMessage invokes method messages.getPreparedInlineMessage#857ebdb8 returning error if any.
//
// See https://core.telegram.org/method/messages.getPreparedInlineMessage for reference.
func (c *Client) MessagesGetPreparedInlineMessage(ctx context.Context, request *MessagesGetPreparedInlineMessageRequest) (*MessagesPreparedInlineMessage, error) {
	var result MessagesPreparedInlineMessage

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
