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

// NotificationGroupTypeMessages represents TL type `notificationGroupTypeMessages#9a86331d`.
type NotificationGroupTypeMessages struct {
}

// NotificationGroupTypeMessagesTypeID is TL type id of NotificationGroupTypeMessages.
const NotificationGroupTypeMessagesTypeID = 0x9a86331d

// construct implements constructor of NotificationGroupTypeClass.
func (n NotificationGroupTypeMessages) construct() NotificationGroupTypeClass { return &n }

// Ensuring interfaces in compile-time for NotificationGroupTypeMessages.
var (
	_ bin.Encoder     = &NotificationGroupTypeMessages{}
	_ bin.Decoder     = &NotificationGroupTypeMessages{}
	_ bin.BareEncoder = &NotificationGroupTypeMessages{}
	_ bin.BareDecoder = &NotificationGroupTypeMessages{}

	_ NotificationGroupTypeClass = &NotificationGroupTypeMessages{}
)

func (n *NotificationGroupTypeMessages) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotificationGroupTypeMessages) String() string {
	if n == nil {
		return "NotificationGroupTypeMessages(nil)"
	}
	type Alias NotificationGroupTypeMessages
	return fmt.Sprintf("NotificationGroupTypeMessages%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NotificationGroupTypeMessages) TypeID() uint32 {
	return NotificationGroupTypeMessagesTypeID
}

// TypeName returns name of type in TL schema.
func (*NotificationGroupTypeMessages) TypeName() string {
	return "notificationGroupTypeMessages"
}

// TypeInfo returns info about TL type.
func (n *NotificationGroupTypeMessages) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "notificationGroupTypeMessages",
		ID:   NotificationGroupTypeMessagesTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *NotificationGroupTypeMessages) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationGroupTypeMessages#9a86331d as nil")
	}
	b.PutID(NotificationGroupTypeMessagesTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NotificationGroupTypeMessages) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationGroupTypeMessages#9a86331d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NotificationGroupTypeMessages) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationGroupTypeMessages#9a86331d to nil")
	}
	if err := b.ConsumeID(NotificationGroupTypeMessagesTypeID); err != nil {
		return fmt.Errorf("unable to decode notificationGroupTypeMessages#9a86331d: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NotificationGroupTypeMessages) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationGroupTypeMessages#9a86331d to nil")
	}
	return nil
}

// NotificationGroupTypeMentions represents TL type `notificationGroupTypeMentions#85ca89ad`.
type NotificationGroupTypeMentions struct {
}

// NotificationGroupTypeMentionsTypeID is TL type id of NotificationGroupTypeMentions.
const NotificationGroupTypeMentionsTypeID = 0x85ca89ad

// construct implements constructor of NotificationGroupTypeClass.
func (n NotificationGroupTypeMentions) construct() NotificationGroupTypeClass { return &n }

// Ensuring interfaces in compile-time for NotificationGroupTypeMentions.
var (
	_ bin.Encoder     = &NotificationGroupTypeMentions{}
	_ bin.Decoder     = &NotificationGroupTypeMentions{}
	_ bin.BareEncoder = &NotificationGroupTypeMentions{}
	_ bin.BareDecoder = &NotificationGroupTypeMentions{}

	_ NotificationGroupTypeClass = &NotificationGroupTypeMentions{}
)

func (n *NotificationGroupTypeMentions) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotificationGroupTypeMentions) String() string {
	if n == nil {
		return "NotificationGroupTypeMentions(nil)"
	}
	type Alias NotificationGroupTypeMentions
	return fmt.Sprintf("NotificationGroupTypeMentions%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NotificationGroupTypeMentions) TypeID() uint32 {
	return NotificationGroupTypeMentionsTypeID
}

// TypeName returns name of type in TL schema.
func (*NotificationGroupTypeMentions) TypeName() string {
	return "notificationGroupTypeMentions"
}

// TypeInfo returns info about TL type.
func (n *NotificationGroupTypeMentions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "notificationGroupTypeMentions",
		ID:   NotificationGroupTypeMentionsTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *NotificationGroupTypeMentions) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationGroupTypeMentions#85ca89ad as nil")
	}
	b.PutID(NotificationGroupTypeMentionsTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NotificationGroupTypeMentions) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationGroupTypeMentions#85ca89ad as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NotificationGroupTypeMentions) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationGroupTypeMentions#85ca89ad to nil")
	}
	if err := b.ConsumeID(NotificationGroupTypeMentionsTypeID); err != nil {
		return fmt.Errorf("unable to decode notificationGroupTypeMentions#85ca89ad: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NotificationGroupTypeMentions) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationGroupTypeMentions#85ca89ad to nil")
	}
	return nil
}

// NotificationGroupTypeSecretChat represents TL type `notificationGroupTypeSecretChat#52e54e34`.
type NotificationGroupTypeSecretChat struct {
}

// NotificationGroupTypeSecretChatTypeID is TL type id of NotificationGroupTypeSecretChat.
const NotificationGroupTypeSecretChatTypeID = 0x52e54e34

// construct implements constructor of NotificationGroupTypeClass.
func (n NotificationGroupTypeSecretChat) construct() NotificationGroupTypeClass { return &n }

// Ensuring interfaces in compile-time for NotificationGroupTypeSecretChat.
var (
	_ bin.Encoder     = &NotificationGroupTypeSecretChat{}
	_ bin.Decoder     = &NotificationGroupTypeSecretChat{}
	_ bin.BareEncoder = &NotificationGroupTypeSecretChat{}
	_ bin.BareDecoder = &NotificationGroupTypeSecretChat{}

	_ NotificationGroupTypeClass = &NotificationGroupTypeSecretChat{}
)

func (n *NotificationGroupTypeSecretChat) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotificationGroupTypeSecretChat) String() string {
	if n == nil {
		return "NotificationGroupTypeSecretChat(nil)"
	}
	type Alias NotificationGroupTypeSecretChat
	return fmt.Sprintf("NotificationGroupTypeSecretChat%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NotificationGroupTypeSecretChat) TypeID() uint32 {
	return NotificationGroupTypeSecretChatTypeID
}

// TypeName returns name of type in TL schema.
func (*NotificationGroupTypeSecretChat) TypeName() string {
	return "notificationGroupTypeSecretChat"
}

// TypeInfo returns info about TL type.
func (n *NotificationGroupTypeSecretChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "notificationGroupTypeSecretChat",
		ID:   NotificationGroupTypeSecretChatTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *NotificationGroupTypeSecretChat) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationGroupTypeSecretChat#52e54e34 as nil")
	}
	b.PutID(NotificationGroupTypeSecretChatTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NotificationGroupTypeSecretChat) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationGroupTypeSecretChat#52e54e34 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NotificationGroupTypeSecretChat) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationGroupTypeSecretChat#52e54e34 to nil")
	}
	if err := b.ConsumeID(NotificationGroupTypeSecretChatTypeID); err != nil {
		return fmt.Errorf("unable to decode notificationGroupTypeSecretChat#52e54e34: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NotificationGroupTypeSecretChat) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationGroupTypeSecretChat#52e54e34 to nil")
	}
	return nil
}

// NotificationGroupTypeCalls represents TL type `notificationGroupTypeCalls#5233c152`.
type NotificationGroupTypeCalls struct {
}

// NotificationGroupTypeCallsTypeID is TL type id of NotificationGroupTypeCalls.
const NotificationGroupTypeCallsTypeID = 0x5233c152

// construct implements constructor of NotificationGroupTypeClass.
func (n NotificationGroupTypeCalls) construct() NotificationGroupTypeClass { return &n }

// Ensuring interfaces in compile-time for NotificationGroupTypeCalls.
var (
	_ bin.Encoder     = &NotificationGroupTypeCalls{}
	_ bin.Decoder     = &NotificationGroupTypeCalls{}
	_ bin.BareEncoder = &NotificationGroupTypeCalls{}
	_ bin.BareDecoder = &NotificationGroupTypeCalls{}

	_ NotificationGroupTypeClass = &NotificationGroupTypeCalls{}
)

func (n *NotificationGroupTypeCalls) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotificationGroupTypeCalls) String() string {
	if n == nil {
		return "NotificationGroupTypeCalls(nil)"
	}
	type Alias NotificationGroupTypeCalls
	return fmt.Sprintf("NotificationGroupTypeCalls%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NotificationGroupTypeCalls) TypeID() uint32 {
	return NotificationGroupTypeCallsTypeID
}

// TypeName returns name of type in TL schema.
func (*NotificationGroupTypeCalls) TypeName() string {
	return "notificationGroupTypeCalls"
}

// TypeInfo returns info about TL type.
func (n *NotificationGroupTypeCalls) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "notificationGroupTypeCalls",
		ID:   NotificationGroupTypeCallsTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *NotificationGroupTypeCalls) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationGroupTypeCalls#5233c152 as nil")
	}
	b.PutID(NotificationGroupTypeCallsTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NotificationGroupTypeCalls) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationGroupTypeCalls#5233c152 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NotificationGroupTypeCalls) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationGroupTypeCalls#5233c152 to nil")
	}
	if err := b.ConsumeID(NotificationGroupTypeCallsTypeID); err != nil {
		return fmt.Errorf("unable to decode notificationGroupTypeCalls#5233c152: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NotificationGroupTypeCalls) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationGroupTypeCalls#5233c152 to nil")
	}
	return nil
}

// NotificationGroupTypeClass represents NotificationGroupType generic type.
//
// Example:
//  g, err := tdapi.DecodeNotificationGroupType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.NotificationGroupTypeMessages: // notificationGroupTypeMessages#9a86331d
//  case *tdapi.NotificationGroupTypeMentions: // notificationGroupTypeMentions#85ca89ad
//  case *tdapi.NotificationGroupTypeSecretChat: // notificationGroupTypeSecretChat#52e54e34
//  case *tdapi.NotificationGroupTypeCalls: // notificationGroupTypeCalls#5233c152
//  default: panic(v)
//  }
type NotificationGroupTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() NotificationGroupTypeClass

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

// DecodeNotificationGroupType implements binary de-serialization for NotificationGroupTypeClass.
func DecodeNotificationGroupType(buf *bin.Buffer) (NotificationGroupTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case NotificationGroupTypeMessagesTypeID:
		// Decoding notificationGroupTypeMessages#9a86331d.
		v := NotificationGroupTypeMessages{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotificationGroupTypeClass: %w", err)
		}
		return &v, nil
	case NotificationGroupTypeMentionsTypeID:
		// Decoding notificationGroupTypeMentions#85ca89ad.
		v := NotificationGroupTypeMentions{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotificationGroupTypeClass: %w", err)
		}
		return &v, nil
	case NotificationGroupTypeSecretChatTypeID:
		// Decoding notificationGroupTypeSecretChat#52e54e34.
		v := NotificationGroupTypeSecretChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotificationGroupTypeClass: %w", err)
		}
		return &v, nil
	case NotificationGroupTypeCallsTypeID:
		// Decoding notificationGroupTypeCalls#5233c152.
		v := NotificationGroupTypeCalls{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotificationGroupTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode NotificationGroupTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// NotificationGroupType boxes the NotificationGroupTypeClass providing a helper.
type NotificationGroupTypeBox struct {
	NotificationGroupType NotificationGroupTypeClass
}

// Decode implements bin.Decoder for NotificationGroupTypeBox.
func (b *NotificationGroupTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode NotificationGroupTypeBox to nil")
	}
	v, err := DecodeNotificationGroupType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.NotificationGroupType = v
	return nil
}

// Encode implements bin.Encode for NotificationGroupTypeBox.
func (b *NotificationGroupTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.NotificationGroupType == nil {
		return fmt.Errorf("unable to encode NotificationGroupTypeClass as nil")
	}
	return b.NotificationGroupType.Encode(buf)
}
