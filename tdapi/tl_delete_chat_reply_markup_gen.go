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

// DeleteChatReplyMarkupRequest represents TL type `deleteChatReplyMarkup#5ff9b5b`.
type DeleteChatReplyMarkupRequest struct {
	// Chat identifier
	ChatID int64
	// The message identifier of the used keyboard
	MessageID int64
}

// DeleteChatReplyMarkupRequestTypeID is TL type id of DeleteChatReplyMarkupRequest.
const DeleteChatReplyMarkupRequestTypeID = 0x5ff9b5b

// Ensuring interfaces in compile-time for DeleteChatReplyMarkupRequest.
var (
	_ bin.Encoder     = &DeleteChatReplyMarkupRequest{}
	_ bin.Decoder     = &DeleteChatReplyMarkupRequest{}
	_ bin.BareEncoder = &DeleteChatReplyMarkupRequest{}
	_ bin.BareDecoder = &DeleteChatReplyMarkupRequest{}
)

func (d *DeleteChatReplyMarkupRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ChatID == 0) {
		return false
	}
	if !(d.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteChatReplyMarkupRequest) String() string {
	if d == nil {
		return "DeleteChatReplyMarkupRequest(nil)"
	}
	type Alias DeleteChatReplyMarkupRequest
	return fmt.Sprintf("DeleteChatReplyMarkupRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteChatReplyMarkupRequest) TypeID() uint32 {
	return DeleteChatReplyMarkupRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteChatReplyMarkupRequest) TypeName() string {
	return "deleteChatReplyMarkup"
}

// TypeInfo returns info about TL type.
func (d *DeleteChatReplyMarkupRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteChatReplyMarkup",
		ID:   DeleteChatReplyMarkupRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteChatReplyMarkupRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChatReplyMarkup#5ff9b5b as nil")
	}
	b.PutID(DeleteChatReplyMarkupRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteChatReplyMarkupRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChatReplyMarkup#5ff9b5b as nil")
	}
	b.PutLong(d.ChatID)
	b.PutLong(d.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteChatReplyMarkupRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChatReplyMarkup#5ff9b5b to nil")
	}
	if err := b.ConsumeID(DeleteChatReplyMarkupRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteChatReplyMarkup#5ff9b5b: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteChatReplyMarkupRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChatReplyMarkup#5ff9b5b to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode deleteChatReplyMarkup#5ff9b5b: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode deleteChatReplyMarkup#5ff9b5b: field message_id: %w", err)
		}
		d.MessageID = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (d *DeleteChatReplyMarkupRequest) GetChatID() (value int64) {
	return d.ChatID
}

// GetMessageID returns value of MessageID field.
func (d *DeleteChatReplyMarkupRequest) GetMessageID() (value int64) {
	return d.MessageID
}

// DeleteChatReplyMarkup invokes method deleteChatReplyMarkup#5ff9b5b returning error if any.
func (c *Client) DeleteChatReplyMarkup(ctx context.Context, request *DeleteChatReplyMarkupRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
