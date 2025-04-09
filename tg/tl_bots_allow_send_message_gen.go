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

// BotsAllowSendMessageRequest represents TL type `bots.allowSendMessage#f132e3ef`.
type BotsAllowSendMessageRequest struct {
	// Bot field of BotsAllowSendMessageRequest.
	Bot InputUserClass
}

// BotsAllowSendMessageRequestTypeID is TL type id of BotsAllowSendMessageRequest.
const BotsAllowSendMessageRequestTypeID = 0xf132e3ef

// Ensuring interfaces in compile-time for BotsAllowSendMessageRequest.
var (
	_ bin.Encoder     = &BotsAllowSendMessageRequest{}
	_ bin.Decoder     = &BotsAllowSendMessageRequest{}
	_ bin.BareEncoder = &BotsAllowSendMessageRequest{}
	_ bin.BareDecoder = &BotsAllowSendMessageRequest{}
)

func (a *BotsAllowSendMessageRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Bot == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *BotsAllowSendMessageRequest) String() string {
	if a == nil {
		return "BotsAllowSendMessageRequest(nil)"
	}
	type Alias BotsAllowSendMessageRequest
	return fmt.Sprintf("BotsAllowSendMessageRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsAllowSendMessageRequest) TypeID() uint32 {
	return BotsAllowSendMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsAllowSendMessageRequest) TypeName() string {
	return "bots.allowSendMessage"
}

// TypeInfo returns info about TL type.
func (a *BotsAllowSendMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.allowSendMessage",
		ID:   BotsAllowSendMessageRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *BotsAllowSendMessageRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode bots.allowSendMessage#f132e3ef as nil")
	}
	b.PutID(BotsAllowSendMessageRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *BotsAllowSendMessageRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode bots.allowSendMessage#f132e3ef as nil")
	}
	if a.Bot == nil {
		return fmt.Errorf("unable to encode bots.allowSendMessage#f132e3ef: field bot is nil")
	}
	if err := a.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.allowSendMessage#f132e3ef: field bot: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *BotsAllowSendMessageRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode bots.allowSendMessage#f132e3ef to nil")
	}
	if err := b.ConsumeID(BotsAllowSendMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.allowSendMessage#f132e3ef: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *BotsAllowSendMessageRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode bots.allowSendMessage#f132e3ef to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.allowSendMessage#f132e3ef: field bot: %w", err)
		}
		a.Bot = value
	}
	return nil
}

// GetBot returns value of Bot field.
func (a *BotsAllowSendMessageRequest) GetBot() (value InputUserClass) {
	if a == nil {
		return
	}
	return a.Bot
}

// BotsAllowSendMessage invokes method bots.allowSendMessage#f132e3ef returning error if any.
func (c *Client) BotsAllowSendMessage(ctx context.Context, bot InputUserClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &BotsAllowSendMessageRequest{
		Bot: bot,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
