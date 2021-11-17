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

// MessageFileTypePrivate represents TL type `messageFileTypePrivate#e0e44ed4`.
type MessageFileTypePrivate struct {
	// Name of the other party; may be empty if unrecognized
	Name string
}

// MessageFileTypePrivateTypeID is TL type id of MessageFileTypePrivate.
const MessageFileTypePrivateTypeID = 0xe0e44ed4

// construct implements constructor of MessageFileTypeClass.
func (m MessageFileTypePrivate) construct() MessageFileTypeClass { return &m }

// Ensuring interfaces in compile-time for MessageFileTypePrivate.
var (
	_ bin.Encoder     = &MessageFileTypePrivate{}
	_ bin.Decoder     = &MessageFileTypePrivate{}
	_ bin.BareEncoder = &MessageFileTypePrivate{}
	_ bin.BareDecoder = &MessageFileTypePrivate{}

	_ MessageFileTypeClass = &MessageFileTypePrivate{}
)

func (m *MessageFileTypePrivate) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageFileTypePrivate) String() string {
	if m == nil {
		return "MessageFileTypePrivate(nil)"
	}
	type Alias MessageFileTypePrivate
	return fmt.Sprintf("MessageFileTypePrivate%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageFileTypePrivate) TypeID() uint32 {
	return MessageFileTypePrivateTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageFileTypePrivate) TypeName() string {
	return "messageFileTypePrivate"
}

// TypeInfo returns info about TL type.
func (m *MessageFileTypePrivate) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageFileTypePrivate",
		ID:   MessageFileTypePrivateTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageFileTypePrivate) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageFileTypePrivate#e0e44ed4 as nil")
	}
	b.PutID(MessageFileTypePrivateTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageFileTypePrivate) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageFileTypePrivate#e0e44ed4 as nil")
	}
	b.PutString(m.Name)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageFileTypePrivate) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageFileTypePrivate#e0e44ed4 to nil")
	}
	if err := b.ConsumeID(MessageFileTypePrivateTypeID); err != nil {
		return fmt.Errorf("unable to decode messageFileTypePrivate#e0e44ed4: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageFileTypePrivate) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageFileTypePrivate#e0e44ed4 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageFileTypePrivate#e0e44ed4: field name: %w", err)
		}
		m.Name = value
	}
	return nil
}

// GetName returns value of Name field.
func (m *MessageFileTypePrivate) GetName() (value string) {
	return m.Name
}

// MessageFileTypeGroup represents TL type `messageFileTypeGroup#f2e58f68`.
type MessageFileTypeGroup struct {
	// Title of the group chat; may be empty if unrecognized
	Title string
}

// MessageFileTypeGroupTypeID is TL type id of MessageFileTypeGroup.
const MessageFileTypeGroupTypeID = 0xf2e58f68

// construct implements constructor of MessageFileTypeClass.
func (m MessageFileTypeGroup) construct() MessageFileTypeClass { return &m }

// Ensuring interfaces in compile-time for MessageFileTypeGroup.
var (
	_ bin.Encoder     = &MessageFileTypeGroup{}
	_ bin.Decoder     = &MessageFileTypeGroup{}
	_ bin.BareEncoder = &MessageFileTypeGroup{}
	_ bin.BareDecoder = &MessageFileTypeGroup{}

	_ MessageFileTypeClass = &MessageFileTypeGroup{}
)

func (m *MessageFileTypeGroup) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageFileTypeGroup) String() string {
	if m == nil {
		return "MessageFileTypeGroup(nil)"
	}
	type Alias MessageFileTypeGroup
	return fmt.Sprintf("MessageFileTypeGroup%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageFileTypeGroup) TypeID() uint32 {
	return MessageFileTypeGroupTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageFileTypeGroup) TypeName() string {
	return "messageFileTypeGroup"
}

// TypeInfo returns info about TL type.
func (m *MessageFileTypeGroup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageFileTypeGroup",
		ID:   MessageFileTypeGroupTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageFileTypeGroup) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageFileTypeGroup#f2e58f68 as nil")
	}
	b.PutID(MessageFileTypeGroupTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageFileTypeGroup) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageFileTypeGroup#f2e58f68 as nil")
	}
	b.PutString(m.Title)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageFileTypeGroup) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageFileTypeGroup#f2e58f68 to nil")
	}
	if err := b.ConsumeID(MessageFileTypeGroupTypeID); err != nil {
		return fmt.Errorf("unable to decode messageFileTypeGroup#f2e58f68: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageFileTypeGroup) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageFileTypeGroup#f2e58f68 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageFileTypeGroup#f2e58f68: field title: %w", err)
		}
		m.Title = value
	}
	return nil
}

// GetTitle returns value of Title field.
func (m *MessageFileTypeGroup) GetTitle() (value string) {
	return m.Title
}

// MessageFileTypeUnknown represents TL type `messageFileTypeUnknown#461dbab2`.
type MessageFileTypeUnknown struct {
}

// MessageFileTypeUnknownTypeID is TL type id of MessageFileTypeUnknown.
const MessageFileTypeUnknownTypeID = 0x461dbab2

// construct implements constructor of MessageFileTypeClass.
func (m MessageFileTypeUnknown) construct() MessageFileTypeClass { return &m }

// Ensuring interfaces in compile-time for MessageFileTypeUnknown.
var (
	_ bin.Encoder     = &MessageFileTypeUnknown{}
	_ bin.Decoder     = &MessageFileTypeUnknown{}
	_ bin.BareEncoder = &MessageFileTypeUnknown{}
	_ bin.BareDecoder = &MessageFileTypeUnknown{}

	_ MessageFileTypeClass = &MessageFileTypeUnknown{}
)

func (m *MessageFileTypeUnknown) Zero() bool {
	if m == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageFileTypeUnknown) String() string {
	if m == nil {
		return "MessageFileTypeUnknown(nil)"
	}
	type Alias MessageFileTypeUnknown
	return fmt.Sprintf("MessageFileTypeUnknown%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageFileTypeUnknown) TypeID() uint32 {
	return MessageFileTypeUnknownTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageFileTypeUnknown) TypeName() string {
	return "messageFileTypeUnknown"
}

// TypeInfo returns info about TL type.
func (m *MessageFileTypeUnknown) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageFileTypeUnknown",
		ID:   MessageFileTypeUnknownTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageFileTypeUnknown) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageFileTypeUnknown#461dbab2 as nil")
	}
	b.PutID(MessageFileTypeUnknownTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageFileTypeUnknown) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageFileTypeUnknown#461dbab2 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageFileTypeUnknown) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageFileTypeUnknown#461dbab2 to nil")
	}
	if err := b.ConsumeID(MessageFileTypeUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode messageFileTypeUnknown#461dbab2: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageFileTypeUnknown) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageFileTypeUnknown#461dbab2 to nil")
	}
	return nil
}

// MessageFileTypeClass represents MessageFileType generic type.
//
// Example:
//  g, err := tdapi.DecodeMessageFileType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.MessageFileTypePrivate: // messageFileTypePrivate#e0e44ed4
//  case *tdapi.MessageFileTypeGroup: // messageFileTypeGroup#f2e58f68
//  case *tdapi.MessageFileTypeUnknown: // messageFileTypeUnknown#461dbab2
//  default: panic(v)
//  }
type MessageFileTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessageFileTypeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeMessageFileType implements binary de-serialization for MessageFileTypeClass.
func DecodeMessageFileType(buf *bin.Buffer) (MessageFileTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageFileTypePrivateTypeID:
		// Decoding messageFileTypePrivate#e0e44ed4.
		v := MessageFileTypePrivate{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageFileTypeClass: %w", err)
		}
		return &v, nil
	case MessageFileTypeGroupTypeID:
		// Decoding messageFileTypeGroup#f2e58f68.
		v := MessageFileTypeGroup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageFileTypeClass: %w", err)
		}
		return &v, nil
	case MessageFileTypeUnknownTypeID:
		// Decoding messageFileTypeUnknown#461dbab2.
		v := MessageFileTypeUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageFileTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageFileTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessageFileType boxes the MessageFileTypeClass providing a helper.
type MessageFileTypeBox struct {
	MessageFileType MessageFileTypeClass
}

// Decode implements bin.Decoder for MessageFileTypeBox.
func (b *MessageFileTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageFileTypeBox to nil")
	}
	v, err := DecodeMessageFileType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageFileType = v
	return nil
}

// Encode implements bin.Encode for MessageFileTypeBox.
func (b *MessageFileTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessageFileType == nil {
		return fmt.Errorf("unable to encode MessageFileTypeClass as nil")
	}
	return b.MessageFileType.Encode(buf)
}
