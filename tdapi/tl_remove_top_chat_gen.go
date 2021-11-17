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

// RemoveTopChatRequest represents TL type `removeTopChat#8e481e55`.
type RemoveTopChatRequest struct {
	// Category of frequently used chats
	Category TopChatCategoryClass
	// Chat identifier
	ChatID int64
}

// RemoveTopChatRequestTypeID is TL type id of RemoveTopChatRequest.
const RemoveTopChatRequestTypeID = 0x8e481e55

// Ensuring interfaces in compile-time for RemoveTopChatRequest.
var (
	_ bin.Encoder     = &RemoveTopChatRequest{}
	_ bin.Decoder     = &RemoveTopChatRequest{}
	_ bin.BareEncoder = &RemoveTopChatRequest{}
	_ bin.BareDecoder = &RemoveTopChatRequest{}
)

func (r *RemoveTopChatRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Category == nil) {
		return false
	}
	if !(r.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RemoveTopChatRequest) String() string {
	if r == nil {
		return "RemoveTopChatRequest(nil)"
	}
	type Alias RemoveTopChatRequest
	return fmt.Sprintf("RemoveTopChatRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RemoveTopChatRequest) TypeID() uint32 {
	return RemoveTopChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RemoveTopChatRequest) TypeName() string {
	return "removeTopChat"
}

// TypeInfo returns info about TL type.
func (r *RemoveTopChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "removeTopChat",
		ID:   RemoveTopChatRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Category",
			SchemaName: "category",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RemoveTopChatRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeTopChat#8e481e55 as nil")
	}
	b.PutID(RemoveTopChatRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RemoveTopChatRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeTopChat#8e481e55 as nil")
	}
	if r.Category == nil {
		return fmt.Errorf("unable to encode removeTopChat#8e481e55: field category is nil")
	}
	if err := r.Category.Encode(b); err != nil {
		return fmt.Errorf("unable to encode removeTopChat#8e481e55: field category: %w", err)
	}
	b.PutLong(r.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (r *RemoveTopChatRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeTopChat#8e481e55 to nil")
	}
	if err := b.ConsumeID(RemoveTopChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode removeTopChat#8e481e55: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RemoveTopChatRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeTopChat#8e481e55 to nil")
	}
	{
		value, err := DecodeTopChatCategory(b)
		if err != nil {
			return fmt.Errorf("unable to decode removeTopChat#8e481e55: field category: %w", err)
		}
		r.Category = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode removeTopChat#8e481e55: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	return nil
}

// GetCategory returns value of Category field.
func (r *RemoveTopChatRequest) GetCategory() (value TopChatCategoryClass) {
	return r.Category
}

// GetChatID returns value of ChatID field.
func (r *RemoveTopChatRequest) GetChatID() (value int64) {
	return r.ChatID
}

// RemoveTopChat invokes method removeTopChat#8e481e55 returning error if any.
func (c *Client) RemoveTopChat(ctx context.Context, request *RemoveTopChatRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
