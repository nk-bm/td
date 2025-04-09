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

// InputPhoneCall represents TL type `inputPhoneCall#1e36fded`.
type InputPhoneCall struct {
	// ID field of InputPhoneCall.
	ID int64
	// AccessHash field of InputPhoneCall.
	AccessHash int64
}

// InputPhoneCallTypeID is TL type id of InputPhoneCall.
const InputPhoneCallTypeID = 0x1e36fded

// Ensuring interfaces in compile-time for InputPhoneCall.
var (
	_ bin.Encoder     = &InputPhoneCall{}
	_ bin.Decoder     = &InputPhoneCall{}
	_ bin.BareEncoder = &InputPhoneCall{}
	_ bin.BareDecoder = &InputPhoneCall{}
)

func (i *InputPhoneCall) Zero() bool {
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
func (i *InputPhoneCall) String() string {
	if i == nil {
		return "InputPhoneCall(nil)"
	}
	type Alias InputPhoneCall
	return fmt.Sprintf("InputPhoneCall%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPhoneCall) TypeID() uint32 {
	return InputPhoneCallTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPhoneCall) TypeName() string {
	return "inputPhoneCall"
}

// TypeInfo returns info about TL type.
func (i *InputPhoneCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPhoneCall",
		ID:   InputPhoneCallTypeID,
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
func (i *InputPhoneCall) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhoneCall#1e36fded as nil")
	}
	b.PutID(InputPhoneCallTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPhoneCall) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhoneCall#1e36fded as nil")
	}
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPhoneCall) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhoneCall#1e36fded to nil")
	}
	if err := b.ConsumeID(InputPhoneCallTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhoneCall#1e36fded: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPhoneCall) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhoneCall#1e36fded to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneCall#1e36fded: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneCall#1e36fded: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputPhoneCall) GetID() (value int64) {
	if i == nil {
		return
	}
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputPhoneCall) GetAccessHash() (value int64) {
	if i == nil {
		return
	}
	return i.AccessHash
}
