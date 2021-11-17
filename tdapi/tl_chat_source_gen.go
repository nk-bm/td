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

// ChatSourceMtprotoProxy represents TL type `chatSourceMtprotoProxy#177d1803`.
type ChatSourceMtprotoProxy struct {
}

// ChatSourceMtprotoProxyTypeID is TL type id of ChatSourceMtprotoProxy.
const ChatSourceMtprotoProxyTypeID = 0x177d1803

// construct implements constructor of ChatSourceClass.
func (c ChatSourceMtprotoProxy) construct() ChatSourceClass { return &c }

// Ensuring interfaces in compile-time for ChatSourceMtprotoProxy.
var (
	_ bin.Encoder     = &ChatSourceMtprotoProxy{}
	_ bin.Decoder     = &ChatSourceMtprotoProxy{}
	_ bin.BareEncoder = &ChatSourceMtprotoProxy{}
	_ bin.BareDecoder = &ChatSourceMtprotoProxy{}

	_ ChatSourceClass = &ChatSourceMtprotoProxy{}
)

func (c *ChatSourceMtprotoProxy) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatSourceMtprotoProxy) String() string {
	if c == nil {
		return "ChatSourceMtprotoProxy(nil)"
	}
	type Alias ChatSourceMtprotoProxy
	return fmt.Sprintf("ChatSourceMtprotoProxy%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatSourceMtprotoProxy) TypeID() uint32 {
	return ChatSourceMtprotoProxyTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatSourceMtprotoProxy) TypeName() string {
	return "chatSourceMtprotoProxy"
}

// TypeInfo returns info about TL type.
func (c *ChatSourceMtprotoProxy) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatSourceMtprotoProxy",
		ID:   ChatSourceMtprotoProxyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatSourceMtprotoProxy) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatSourceMtprotoProxy#177d1803 as nil")
	}
	b.PutID(ChatSourceMtprotoProxyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatSourceMtprotoProxy) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatSourceMtprotoProxy#177d1803 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatSourceMtprotoProxy) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatSourceMtprotoProxy#177d1803 to nil")
	}
	if err := b.ConsumeID(ChatSourceMtprotoProxyTypeID); err != nil {
		return fmt.Errorf("unable to decode chatSourceMtprotoProxy#177d1803: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatSourceMtprotoProxy) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatSourceMtprotoProxy#177d1803 to nil")
	}
	return nil
}

// ChatSourcePublicServiceAnnouncement represents TL type `chatSourcePublicServiceAnnouncement#ec6a6694`.
type ChatSourcePublicServiceAnnouncement struct {
	// The type of the announcement
	Type string
	// The text of the announcement
	Text string
}

// ChatSourcePublicServiceAnnouncementTypeID is TL type id of ChatSourcePublicServiceAnnouncement.
const ChatSourcePublicServiceAnnouncementTypeID = 0xec6a6694

// construct implements constructor of ChatSourceClass.
func (c ChatSourcePublicServiceAnnouncement) construct() ChatSourceClass { return &c }

// Ensuring interfaces in compile-time for ChatSourcePublicServiceAnnouncement.
var (
	_ bin.Encoder     = &ChatSourcePublicServiceAnnouncement{}
	_ bin.Decoder     = &ChatSourcePublicServiceAnnouncement{}
	_ bin.BareEncoder = &ChatSourcePublicServiceAnnouncement{}
	_ bin.BareDecoder = &ChatSourcePublicServiceAnnouncement{}

	_ ChatSourceClass = &ChatSourcePublicServiceAnnouncement{}
)

func (c *ChatSourcePublicServiceAnnouncement) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Type == "") {
		return false
	}
	if !(c.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatSourcePublicServiceAnnouncement) String() string {
	if c == nil {
		return "ChatSourcePublicServiceAnnouncement(nil)"
	}
	type Alias ChatSourcePublicServiceAnnouncement
	return fmt.Sprintf("ChatSourcePublicServiceAnnouncement%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatSourcePublicServiceAnnouncement) TypeID() uint32 {
	return ChatSourcePublicServiceAnnouncementTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatSourcePublicServiceAnnouncement) TypeName() string {
	return "chatSourcePublicServiceAnnouncement"
}

// TypeInfo returns info about TL type.
func (c *ChatSourcePublicServiceAnnouncement) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatSourcePublicServiceAnnouncement",
		ID:   ChatSourcePublicServiceAnnouncementTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatSourcePublicServiceAnnouncement) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatSourcePublicServiceAnnouncement#ec6a6694 as nil")
	}
	b.PutID(ChatSourcePublicServiceAnnouncementTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatSourcePublicServiceAnnouncement) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatSourcePublicServiceAnnouncement#ec6a6694 as nil")
	}
	b.PutString(c.Type)
	b.PutString(c.Text)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatSourcePublicServiceAnnouncement) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatSourcePublicServiceAnnouncement#ec6a6694 to nil")
	}
	if err := b.ConsumeID(ChatSourcePublicServiceAnnouncementTypeID); err != nil {
		return fmt.Errorf("unable to decode chatSourcePublicServiceAnnouncement#ec6a6694: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatSourcePublicServiceAnnouncement) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatSourcePublicServiceAnnouncement#ec6a6694 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatSourcePublicServiceAnnouncement#ec6a6694: field type: %w", err)
		}
		c.Type = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatSourcePublicServiceAnnouncement#ec6a6694: field text: %w", err)
		}
		c.Text = value
	}
	return nil
}

// GetType returns value of Type field.
func (c *ChatSourcePublicServiceAnnouncement) GetType() (value string) {
	return c.Type
}

// GetText returns value of Text field.
func (c *ChatSourcePublicServiceAnnouncement) GetText() (value string) {
	return c.Text
}

// ChatSourceClass represents ChatSource generic type.
//
// Example:
//  g, err := tdapi.DecodeChatSource(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.ChatSourceMtprotoProxy: // chatSourceMtprotoProxy#177d1803
//  case *tdapi.ChatSourcePublicServiceAnnouncement: // chatSourcePublicServiceAnnouncement#ec6a6694
//  default: panic(v)
//  }
type ChatSourceClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatSourceClass

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

// DecodeChatSource implements binary de-serialization for ChatSourceClass.
func DecodeChatSource(buf *bin.Buffer) (ChatSourceClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatSourceMtprotoProxyTypeID:
		// Decoding chatSourceMtprotoProxy#177d1803.
		v := ChatSourceMtprotoProxy{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatSourceClass: %w", err)
		}
		return &v, nil
	case ChatSourcePublicServiceAnnouncementTypeID:
		// Decoding chatSourcePublicServiceAnnouncement#ec6a6694.
		v := ChatSourcePublicServiceAnnouncement{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatSourceClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatSourceClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatSource boxes the ChatSourceClass providing a helper.
type ChatSourceBox struct {
	ChatSource ChatSourceClass
}

// Decode implements bin.Decoder for ChatSourceBox.
func (b *ChatSourceBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatSourceBox to nil")
	}
	v, err := DecodeChatSource(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatSource = v
	return nil
}

// Encode implements bin.Encode for ChatSourceBox.
func (b *ChatSourceBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatSource == nil {
		return fmt.Errorf("unable to encode ChatSourceClass as nil")
	}
	return b.ChatSource.Encode(buf)
}
