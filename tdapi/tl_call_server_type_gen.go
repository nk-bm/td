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

// CallServerTypeTelegramReflector represents TL type `callServerTypeTelegramReflector#a6200634`.
type CallServerTypeTelegramReflector struct {
	// A peer tag to be used with the reflector
	PeerTag []byte
}

// CallServerTypeTelegramReflectorTypeID is TL type id of CallServerTypeTelegramReflector.
const CallServerTypeTelegramReflectorTypeID = 0xa6200634

// construct implements constructor of CallServerTypeClass.
func (c CallServerTypeTelegramReflector) construct() CallServerTypeClass { return &c }

// Ensuring interfaces in compile-time for CallServerTypeTelegramReflector.
var (
	_ bin.Encoder     = &CallServerTypeTelegramReflector{}
	_ bin.Decoder     = &CallServerTypeTelegramReflector{}
	_ bin.BareEncoder = &CallServerTypeTelegramReflector{}
	_ bin.BareDecoder = &CallServerTypeTelegramReflector{}

	_ CallServerTypeClass = &CallServerTypeTelegramReflector{}
)

func (c *CallServerTypeTelegramReflector) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.PeerTag == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CallServerTypeTelegramReflector) String() string {
	if c == nil {
		return "CallServerTypeTelegramReflector(nil)"
	}
	type Alias CallServerTypeTelegramReflector
	return fmt.Sprintf("CallServerTypeTelegramReflector%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CallServerTypeTelegramReflector) TypeID() uint32 {
	return CallServerTypeTelegramReflectorTypeID
}

// TypeName returns name of type in TL schema.
func (*CallServerTypeTelegramReflector) TypeName() string {
	return "callServerTypeTelegramReflector"
}

// TypeInfo returns info about TL type.
func (c *CallServerTypeTelegramReflector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "callServerTypeTelegramReflector",
		ID:   CallServerTypeTelegramReflectorTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PeerTag",
			SchemaName: "peer_tag",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CallServerTypeTelegramReflector) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callServerTypeTelegramReflector#a6200634 as nil")
	}
	b.PutID(CallServerTypeTelegramReflectorTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CallServerTypeTelegramReflector) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callServerTypeTelegramReflector#a6200634 as nil")
	}
	b.PutBytes(c.PeerTag)
	return nil
}

// Decode implements bin.Decoder.
func (c *CallServerTypeTelegramReflector) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callServerTypeTelegramReflector#a6200634 to nil")
	}
	if err := b.ConsumeID(CallServerTypeTelegramReflectorTypeID); err != nil {
		return fmt.Errorf("unable to decode callServerTypeTelegramReflector#a6200634: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CallServerTypeTelegramReflector) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callServerTypeTelegramReflector#a6200634 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode callServerTypeTelegramReflector#a6200634: field peer_tag: %w", err)
		}
		c.PeerTag = value
	}
	return nil
}

// GetPeerTag returns value of PeerTag field.
func (c *CallServerTypeTelegramReflector) GetPeerTag() (value []byte) {
	return c.PeerTag
}

// CallServerTypeWebrtc represents TL type `callServerTypeWebrtc#4a8afd65`.
type CallServerTypeWebrtc struct {
	// Username to be used for authentication
	Username string
	// Authentication password
	Password string
	// True, if the server supports TURN
	SupportsTurn bool
	// True, if the server supports STUN
	SupportsStun bool
}

// CallServerTypeWebrtcTypeID is TL type id of CallServerTypeWebrtc.
const CallServerTypeWebrtcTypeID = 0x4a8afd65

// construct implements constructor of CallServerTypeClass.
func (c CallServerTypeWebrtc) construct() CallServerTypeClass { return &c }

// Ensuring interfaces in compile-time for CallServerTypeWebrtc.
var (
	_ bin.Encoder     = &CallServerTypeWebrtc{}
	_ bin.Decoder     = &CallServerTypeWebrtc{}
	_ bin.BareEncoder = &CallServerTypeWebrtc{}
	_ bin.BareDecoder = &CallServerTypeWebrtc{}

	_ CallServerTypeClass = &CallServerTypeWebrtc{}
)

func (c *CallServerTypeWebrtc) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Username == "") {
		return false
	}
	if !(c.Password == "") {
		return false
	}
	if !(c.SupportsTurn == false) {
		return false
	}
	if !(c.SupportsStun == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CallServerTypeWebrtc) String() string {
	if c == nil {
		return "CallServerTypeWebrtc(nil)"
	}
	type Alias CallServerTypeWebrtc
	return fmt.Sprintf("CallServerTypeWebrtc%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CallServerTypeWebrtc) TypeID() uint32 {
	return CallServerTypeWebrtcTypeID
}

// TypeName returns name of type in TL schema.
func (*CallServerTypeWebrtc) TypeName() string {
	return "callServerTypeWebrtc"
}

// TypeInfo returns info about TL type.
func (c *CallServerTypeWebrtc) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "callServerTypeWebrtc",
		ID:   CallServerTypeWebrtcTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Username",
			SchemaName: "username",
		},
		{
			Name:       "Password",
			SchemaName: "password",
		},
		{
			Name:       "SupportsTurn",
			SchemaName: "supports_turn",
		},
		{
			Name:       "SupportsStun",
			SchemaName: "supports_stun",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CallServerTypeWebrtc) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callServerTypeWebrtc#4a8afd65 as nil")
	}
	b.PutID(CallServerTypeWebrtcTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CallServerTypeWebrtc) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callServerTypeWebrtc#4a8afd65 as nil")
	}
	b.PutString(c.Username)
	b.PutString(c.Password)
	b.PutBool(c.SupportsTurn)
	b.PutBool(c.SupportsStun)
	return nil
}

// Decode implements bin.Decoder.
func (c *CallServerTypeWebrtc) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callServerTypeWebrtc#4a8afd65 to nil")
	}
	if err := b.ConsumeID(CallServerTypeWebrtcTypeID); err != nil {
		return fmt.Errorf("unable to decode callServerTypeWebrtc#4a8afd65: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CallServerTypeWebrtc) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callServerTypeWebrtc#4a8afd65 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode callServerTypeWebrtc#4a8afd65: field username: %w", err)
		}
		c.Username = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode callServerTypeWebrtc#4a8afd65: field password: %w", err)
		}
		c.Password = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode callServerTypeWebrtc#4a8afd65: field supports_turn: %w", err)
		}
		c.SupportsTurn = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode callServerTypeWebrtc#4a8afd65: field supports_stun: %w", err)
		}
		c.SupportsStun = value
	}
	return nil
}

// GetUsername returns value of Username field.
func (c *CallServerTypeWebrtc) GetUsername() (value string) {
	return c.Username
}

// GetPassword returns value of Password field.
func (c *CallServerTypeWebrtc) GetPassword() (value string) {
	return c.Password
}

// GetSupportsTurn returns value of SupportsTurn field.
func (c *CallServerTypeWebrtc) GetSupportsTurn() (value bool) {
	return c.SupportsTurn
}

// GetSupportsStun returns value of SupportsStun field.
func (c *CallServerTypeWebrtc) GetSupportsStun() (value bool) {
	return c.SupportsStun
}

// CallServerTypeClass represents CallServerType generic type.
//
// Example:
//  g, err := tdapi.DecodeCallServerType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.CallServerTypeTelegramReflector: // callServerTypeTelegramReflector#a6200634
//  case *tdapi.CallServerTypeWebrtc: // callServerTypeWebrtc#4a8afd65
//  default: panic(v)
//  }
type CallServerTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() CallServerTypeClass

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

// DecodeCallServerType implements binary de-serialization for CallServerTypeClass.
func DecodeCallServerType(buf *bin.Buffer) (CallServerTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case CallServerTypeTelegramReflectorTypeID:
		// Decoding callServerTypeTelegramReflector#a6200634.
		v := CallServerTypeTelegramReflector{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CallServerTypeClass: %w", err)
		}
		return &v, nil
	case CallServerTypeWebrtcTypeID:
		// Decoding callServerTypeWebrtc#4a8afd65.
		v := CallServerTypeWebrtc{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CallServerTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode CallServerTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// CallServerType boxes the CallServerTypeClass providing a helper.
type CallServerTypeBox struct {
	CallServerType CallServerTypeClass
}

// Decode implements bin.Decoder for CallServerTypeBox.
func (b *CallServerTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode CallServerTypeBox to nil")
	}
	v, err := DecodeCallServerType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CallServerType = v
	return nil
}

// Encode implements bin.Encode for CallServerTypeBox.
func (b *CallServerTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.CallServerType == nil {
		return fmt.Errorf("unable to encode CallServerTypeClass as nil")
	}
	return b.CallServerType.Encode(buf)
}
