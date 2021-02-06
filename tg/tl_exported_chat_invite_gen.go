// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// ChatInviteEmpty represents TL type `chatInviteEmpty#69df3769`.
// No info is associated to the chat invite
//
// See https://core.telegram.org/constructor/chatInviteEmpty for reference.
type ChatInviteEmpty struct {
}

// ChatInviteEmptyTypeID is TL type id of ChatInviteEmpty.
const ChatInviteEmptyTypeID = 0x69df3769

func (c *ChatInviteEmpty) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteEmpty) String() string {
	if c == nil {
		return "ChatInviteEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChatInviteEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ChatInviteEmpty) TypeID() uint32 {
	return ChatInviteEmptyTypeID
}

// Encode implements bin.Encoder.
func (c *ChatInviteEmpty) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteEmpty#69df3769 as nil")
	}
	b.PutID(ChatInviteEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInviteEmpty) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteEmpty#69df3769 to nil")
	}
	if err := b.ConsumeID(ChatInviteEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteEmpty#69df3769: %w", err)
	}
	return nil
}

// construct implements constructor of ExportedChatInviteClass.
func (c ChatInviteEmpty) construct() ExportedChatInviteClass { return &c }

// Ensuring interfaces in compile-time for ChatInviteEmpty.
var (
	_ bin.Encoder = &ChatInviteEmpty{}
	_ bin.Decoder = &ChatInviteEmpty{}

	_ ExportedChatInviteClass = &ChatInviteEmpty{}
)

// ChatInviteExported represents TL type `chatInviteExported#fc2e05bc`.
// Exported chat invite
//
// See https://core.telegram.org/constructor/chatInviteExported for reference.
type ChatInviteExported struct {
	// Chat invitation link
	Link string
}

// ChatInviteExportedTypeID is TL type id of ChatInviteExported.
const ChatInviteExportedTypeID = 0xfc2e05bc

func (c *ChatInviteExported) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Link == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteExported) String() string {
	if c == nil {
		return "ChatInviteExported(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChatInviteExported")
	sb.WriteString("{\n")
	sb.WriteString("\tLink: ")
	sb.WriteString(fmt.Sprint(c.Link))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ChatInviteExported) TypeID() uint32 {
	return ChatInviteExportedTypeID
}

// Encode implements bin.Encoder.
func (c *ChatInviteExported) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteExported#fc2e05bc as nil")
	}
	b.PutID(ChatInviteExportedTypeID)
	b.PutString(c.Link)
	return nil
}

// GetLink returns value of Link field.
func (c *ChatInviteExported) GetLink() (value string) {
	return c.Link
}

// Decode implements bin.Decoder.
func (c *ChatInviteExported) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteExported#fc2e05bc to nil")
	}
	if err := b.ConsumeID(ChatInviteExportedTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteExported#fc2e05bc: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteExported#fc2e05bc: field link: %w", err)
		}
		c.Link = value
	}
	return nil
}

// construct implements constructor of ExportedChatInviteClass.
func (c ChatInviteExported) construct() ExportedChatInviteClass { return &c }

// Ensuring interfaces in compile-time for ChatInviteExported.
var (
	_ bin.Encoder = &ChatInviteExported{}
	_ bin.Decoder = &ChatInviteExported{}

	_ ExportedChatInviteClass = &ChatInviteExported{}
)

// ExportedChatInviteClass represents ExportedChatInvite generic type.
//
// See https://core.telegram.org/type/ExportedChatInvite for reference.
//
// Example:
//  g, err := DecodeExportedChatInvite(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *ChatInviteEmpty: // chatInviteEmpty#69df3769
//  case *ChatInviteExported: // chatInviteExported#fc2e05bc
//  default: panic(v)
//  }
type ExportedChatInviteClass interface {
	bin.Encoder
	bin.Decoder
	construct() ExportedChatInviteClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeExportedChatInvite implements binary de-serialization for ExportedChatInviteClass.
func DecodeExportedChatInvite(buf *bin.Buffer) (ExportedChatInviteClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatInviteEmptyTypeID:
		// Decoding chatInviteEmpty#69df3769.
		v := ChatInviteEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ExportedChatInviteClass: %w", err)
		}
		return &v, nil
	case ChatInviteExportedTypeID:
		// Decoding chatInviteExported#fc2e05bc.
		v := ChatInviteExported{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ExportedChatInviteClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ExportedChatInviteClass: %w", bin.NewUnexpectedID(id))
	}
}

// ExportedChatInvite boxes the ExportedChatInviteClass providing a helper.
type ExportedChatInviteBox struct {
	ExportedChatInvite ExportedChatInviteClass
}

// Decode implements bin.Decoder for ExportedChatInviteBox.
func (b *ExportedChatInviteBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ExportedChatInviteBox to nil")
	}
	v, err := DecodeExportedChatInvite(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ExportedChatInvite = v
	return nil
}

// Encode implements bin.Encode for ExportedChatInviteBox.
func (b *ExportedChatInviteBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ExportedChatInvite == nil {
		return fmt.Errorf("unable to encode ExportedChatInviteClass as nil")
	}
	return b.ExportedChatInvite.Encode(buf)
}
