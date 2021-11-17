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

// Int53 represents TL type `int53#6781c7ee`.
type Int53 struct {
}

// Int53TypeID is TL type id of Int53.
const Int53TypeID = 0x6781c7ee

// Ensuring interfaces in compile-time for Int53.
var (
	_ bin.Encoder     = &Int53{}
	_ bin.Decoder     = &Int53{}
	_ bin.BareEncoder = &Int53{}
	_ bin.BareDecoder = &Int53{}
)

func (i *Int53) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *Int53) String() string {
	if i == nil {
		return "Int53(nil)"
	}
	type Alias Int53
	return fmt.Sprintf("Int53%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Int53) TypeID() uint32 {
	return Int53TypeID
}

// TypeName returns name of type in TL schema.
func (*Int53) TypeName() string {
	return "int53"
}

// TypeInfo returns info about TL type.
func (i *Int53) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "int53",
		ID:   Int53TypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *Int53) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode int53#6781c7ee as nil")
	}
	b.PutID(Int53TypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *Int53) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode int53#6781c7ee as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *Int53) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode int53#6781c7ee to nil")
	}
	if err := b.ConsumeID(Int53TypeID); err != nil {
		return fmt.Errorf("unable to decode int53#6781c7ee: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *Int53) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode int53#6781c7ee to nil")
	}
	return nil
}
