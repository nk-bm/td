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

// InputBotAppID represents TL type `inputBotAppID#a920bd7a`.
type InputBotAppID struct {
	// ID field of InputBotAppID.
	ID int64
	// AccessHash field of InputBotAppID.
	AccessHash int64
}

// InputBotAppIDTypeID is TL type id of InputBotAppID.
const InputBotAppIDTypeID = 0xa920bd7a

// construct implements constructor of InputBotAppClass.
func (i InputBotAppID) construct() InputBotAppClass { return &i }

// Ensuring interfaces in compile-time for InputBotAppID.
var (
	_ bin.Encoder     = &InputBotAppID{}
	_ bin.Decoder     = &InputBotAppID{}
	_ bin.BareEncoder = &InputBotAppID{}
	_ bin.BareDecoder = &InputBotAppID{}

	_ InputBotAppClass = &InputBotAppID{}
)

func (i *InputBotAppID) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputBotAppID) String() string {
	if i == nil {
		return "InputBotAppID(nil)"
	}
	type Alias InputBotAppID
	return fmt.Sprintf("InputBotAppID%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputBotAppID) TypeID() uint32 {
	return InputBotAppIDTypeID
}

// TypeName returns name of type in TL schema.
func (*InputBotAppID) TypeName() string {
	return "inputBotAppID"
}

// TypeInfo returns info about TL type.
func (i *InputBotAppID) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputBotAppID",
		ID:   InputBotAppIDTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputBotAppID) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotAppID#a920bd7a as nil")
	}
	b.PutID(InputBotAppIDTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputBotAppID) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotAppID#a920bd7a as nil")
	}
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputBotAppID) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotAppID#a920bd7a to nil")
	}
	if err := b.ConsumeID(InputBotAppIDTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotAppID#a920bd7a: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputBotAppID) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotAppID#a920bd7a to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotAppID#a920bd7a: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotAppID#a920bd7a: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputBotAppID) GetID() (value int64) {
	if i == nil {
		return
	}
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputBotAppID) GetAccessHash() (value int64) {
	if i == nil {
		return
	}
	return i.AccessHash
}

// InputBotAppShortName represents TL type `inputBotAppShortName#908c0407`.
type InputBotAppShortName struct {
	// BotID field of InputBotAppShortName.
	BotID InputUserClass
	// ShortName field of InputBotAppShortName.
	ShortName string
}

// InputBotAppShortNameTypeID is TL type id of InputBotAppShortName.
const InputBotAppShortNameTypeID = 0x908c0407

// construct implements constructor of InputBotAppClass.
func (i InputBotAppShortName) construct() InputBotAppClass { return &i }

// Ensuring interfaces in compile-time for InputBotAppShortName.
var (
	_ bin.Encoder     = &InputBotAppShortName{}
	_ bin.Decoder     = &InputBotAppShortName{}
	_ bin.BareEncoder = &InputBotAppShortName{}
	_ bin.BareDecoder = &InputBotAppShortName{}

	_ InputBotAppClass = &InputBotAppShortName{}
)

func (i *InputBotAppShortName) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.BotID == nil) {
		return false
	}
	if !(i.ShortName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputBotAppShortName) String() string {
	if i == nil {
		return "InputBotAppShortName(nil)"
	}
	type Alias InputBotAppShortName
	return fmt.Sprintf("InputBotAppShortName%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputBotAppShortName) TypeID() uint32 {
	return InputBotAppShortNameTypeID
}

// TypeName returns name of type in TL schema.
func (*InputBotAppShortName) TypeName() string {
	return "inputBotAppShortName"
}

// TypeInfo returns info about TL type.
func (i *InputBotAppShortName) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputBotAppShortName",
		ID:   InputBotAppShortNameTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotID",
			SchemaName: "bot_id",
		},
		{
			Name:       "ShortName",
			SchemaName: "short_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputBotAppShortName) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotAppShortName#908c0407 as nil")
	}
	b.PutID(InputBotAppShortNameTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputBotAppShortName) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBotAppShortName#908c0407 as nil")
	}
	if i.BotID == nil {
		return fmt.Errorf("unable to encode inputBotAppShortName#908c0407: field bot_id is nil")
	}
	if err := i.BotID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBotAppShortName#908c0407: field bot_id: %w", err)
	}
	b.PutString(i.ShortName)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputBotAppShortName) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotAppShortName#908c0407 to nil")
	}
	if err := b.ConsumeID(InputBotAppShortNameTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBotAppShortName#908c0407: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputBotAppShortName) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBotAppShortName#908c0407 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputBotAppShortName#908c0407: field bot_id: %w", err)
		}
		i.BotID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputBotAppShortName#908c0407: field short_name: %w", err)
		}
		i.ShortName = value
	}
	return nil
}

// GetBotID returns value of BotID field.
func (i *InputBotAppShortName) GetBotID() (value InputUserClass) {
	if i == nil {
		return
	}
	return i.BotID
}

// GetShortName returns value of ShortName field.
func (i *InputBotAppShortName) GetShortName() (value string) {
	if i == nil {
		return
	}
	return i.ShortName
}

// InputBotAppClassName is schema name of InputBotAppClass.
const InputBotAppClassName = "InputBotApp"

// InputBotAppClass represents InputBotApp generic type.
//
// Constructors:
//   - [InputBotAppID]
//   - [InputBotAppShortName]
//
// Example:
//
//	g, err := tg.DecodeInputBotApp(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputBotAppID: // inputBotAppID#a920bd7a
//	case *tg.InputBotAppShortName: // inputBotAppShortName#908c0407
//	default: panic(v)
//	}
type InputBotAppClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputBotAppClass

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

// DecodeInputBotApp implements binary de-serialization for InputBotAppClass.
func DecodeInputBotApp(buf *bin.Buffer) (InputBotAppClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputBotAppIDTypeID:
		// Decoding inputBotAppID#a920bd7a.
		v := InputBotAppID{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputBotAppClass: %w", err)
		}
		return &v, nil
	case InputBotAppShortNameTypeID:
		// Decoding inputBotAppShortName#908c0407.
		v := InputBotAppShortName{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputBotAppClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputBotAppClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputBotApp boxes the InputBotAppClass providing a helper.
type InputBotAppBox struct {
	InputBotApp InputBotAppClass
}

// Decode implements bin.Decoder for InputBotAppBox.
func (b *InputBotAppBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputBotAppBox to nil")
	}
	v, err := DecodeInputBotApp(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputBotApp = v
	return nil
}

// Encode implements bin.Encode for InputBotAppBox.
func (b *InputBotAppBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputBotApp == nil {
		return fmt.Errorf("unable to encode InputBotAppClass as nil")
	}
	return b.InputBotApp.Encode(buf)
}
