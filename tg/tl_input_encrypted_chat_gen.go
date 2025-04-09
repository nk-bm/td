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

// InputEncryptedChat represents TL type `inputEncryptedChat#f141b5e1`.
type InputEncryptedChat struct {
	// ChatID field of InputEncryptedChat.
	ChatID int
	// AccessHash field of InputEncryptedChat.
	AccessHash int64
}

// InputEncryptedChatTypeID is TL type id of InputEncryptedChat.
const InputEncryptedChatTypeID = 0xf141b5e1

// Ensuring interfaces in compile-time for InputEncryptedChat.
var (
	_ bin.Encoder     = &InputEncryptedChat{}
	_ bin.Decoder     = &InputEncryptedChat{}
	_ bin.BareEncoder = &InputEncryptedChat{}
	_ bin.BareDecoder = &InputEncryptedChat{}
)

func (i *InputEncryptedChat) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChatID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputEncryptedChat) String() string {
	if i == nil {
		return "InputEncryptedChat(nil)"
	}
	type Alias InputEncryptedChat
	return fmt.Sprintf("InputEncryptedChat%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputEncryptedChat) TypeID() uint32 {
	return InputEncryptedChatTypeID
}

// TypeName returns name of type in TL schema.
func (*InputEncryptedChat) TypeName() string {
	return "inputEncryptedChat"
}

// TypeInfo returns info about TL type.
func (i *InputEncryptedChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputEncryptedChat",
		ID:   InputEncryptedChatTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputEncryptedChat) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedChat#f141b5e1 as nil")
	}
	b.PutID(InputEncryptedChatTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputEncryptedChat) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedChat#f141b5e1 as nil")
	}
	b.PutInt(i.ChatID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedChat) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedChat#f141b5e1 to nil")
	}
	if err := b.ConsumeID(InputEncryptedChatTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedChat#f141b5e1: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputEncryptedChat) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedChat#f141b5e1 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedChat#f141b5e1: field chat_id: %w", err)
		}
		i.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedChat#f141b5e1: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (i *InputEncryptedChat) GetChatID() (value int) {
	if i == nil {
		return
	}
	return i.ChatID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputEncryptedChat) GetAccessHash() (value int64) {
	if i == nil {
		return
	}
	return i.AccessHash
}
