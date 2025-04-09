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

// OutboxReadDate represents TL type `outboxReadDate#3bb842ac`.
type OutboxReadDate struct {
	// Date field of OutboxReadDate.
	Date int
}

// OutboxReadDateTypeID is TL type id of OutboxReadDate.
const OutboxReadDateTypeID = 0x3bb842ac

// Ensuring interfaces in compile-time for OutboxReadDate.
var (
	_ bin.Encoder     = &OutboxReadDate{}
	_ bin.Decoder     = &OutboxReadDate{}
	_ bin.BareEncoder = &OutboxReadDate{}
	_ bin.BareDecoder = &OutboxReadDate{}
)

func (o *OutboxReadDate) Zero() bool {
	if o == nil {
		return true
	}
	if !(o.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (o *OutboxReadDate) String() string {
	if o == nil {
		return "OutboxReadDate(nil)"
	}
	type Alias OutboxReadDate
	return fmt.Sprintf("OutboxReadDate%+v", Alias(*o))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*OutboxReadDate) TypeID() uint32 {
	return OutboxReadDateTypeID
}

// TypeName returns name of type in TL schema.
func (*OutboxReadDate) TypeName() string {
	return "outboxReadDate"
}

// TypeInfo returns info about TL type.
func (o *OutboxReadDate) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "outboxReadDate",
		ID:   OutboxReadDateTypeID,
	}
	if o == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (o *OutboxReadDate) Encode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode outboxReadDate#3bb842ac as nil")
	}
	b.PutID(OutboxReadDateTypeID)
	return o.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (o *OutboxReadDate) EncodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode outboxReadDate#3bb842ac as nil")
	}
	b.PutInt(o.Date)
	return nil
}

// Decode implements bin.Decoder.
func (o *OutboxReadDate) Decode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode outboxReadDate#3bb842ac to nil")
	}
	if err := b.ConsumeID(OutboxReadDateTypeID); err != nil {
		return fmt.Errorf("unable to decode outboxReadDate#3bb842ac: %w", err)
	}
	return o.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (o *OutboxReadDate) DecodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode outboxReadDate#3bb842ac to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode outboxReadDate#3bb842ac: field date: %w", err)
		}
		o.Date = value
	}
	return nil
}

// GetDate returns value of Date field.
func (o *OutboxReadDate) GetDate() (value int) {
	if o == nil {
		return
	}
	return o.Date
}
