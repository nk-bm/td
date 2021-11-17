// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
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
)

// LeaveChatRequest represents TL type `leaveChat#93377a61`.
type LeaveChatRequest struct {
	// Chat identifier
	ChatID int64
}

// LeaveChatRequestTypeID is TL type id of LeaveChatRequest.
const LeaveChatRequestTypeID = 0x93377a61

// Ensuring interfaces in compile-time for LeaveChatRequest.
var (
	_ bin.Encoder     = &LeaveChatRequest{}
	_ bin.Decoder     = &LeaveChatRequest{}
	_ bin.BareEncoder = &LeaveChatRequest{}
	_ bin.BareDecoder = &LeaveChatRequest{}
)

func (l *LeaveChatRequest) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LeaveChatRequest) String() string {
	if l == nil {
		return "LeaveChatRequest(nil)"
	}
	type Alias LeaveChatRequest
	return fmt.Sprintf("LeaveChatRequest%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LeaveChatRequest) TypeID() uint32 {
	return LeaveChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*LeaveChatRequest) TypeName() string {
	return "leaveChat"
}

// TypeInfo returns info about TL type.
func (l *LeaveChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "leaveChat",
		ID:   LeaveChatRequestTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LeaveChatRequest) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode leaveChat#93377a61 as nil")
	}
	b.PutID(LeaveChatRequestTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LeaveChatRequest) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode leaveChat#93377a61 as nil")
	}
	b.PutLong(l.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (l *LeaveChatRequest) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode leaveChat#93377a61 to nil")
	}
	if err := b.ConsumeID(LeaveChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode leaveChat#93377a61: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LeaveChatRequest) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode leaveChat#93377a61 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode leaveChat#93377a61: field chat_id: %w", err)
		}
		l.ChatID = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (l *LeaveChatRequest) GetChatID() (value int64) {
	return l.ChatID
}

// LeaveChat invokes method leaveChat#93377a61 returning error if any.
func (c *Client) LeaveChat(ctx context.Context, chatid int64) error {
	var ok Ok

	request := &LeaveChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
