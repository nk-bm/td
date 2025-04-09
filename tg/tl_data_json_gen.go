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

// DataJSON represents TL type `dataJSON#7d748d04`.
type DataJSON struct {
	// Data field of DataJSON.
	Data string
}

// DataJSONTypeID is TL type id of DataJSON.
const DataJSONTypeID = 0x7d748d04

// Ensuring interfaces in compile-time for DataJSON.
var (
	_ bin.Encoder     = &DataJSON{}
	_ bin.Decoder     = &DataJSON{}
	_ bin.BareEncoder = &DataJSON{}
	_ bin.BareDecoder = &DataJSON{}
)

func (d *DataJSON) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Data == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DataJSON) String() string {
	if d == nil {
		return "DataJSON(nil)"
	}
	type Alias DataJSON
	return fmt.Sprintf("DataJSON%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DataJSON) TypeID() uint32 {
	return DataJSONTypeID
}

// TypeName returns name of type in TL schema.
func (*DataJSON) TypeName() string {
	return "dataJSON"
}

// TypeInfo returns info about TL type.
func (d *DataJSON) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dataJSON",
		ID:   DataJSONTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DataJSON) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dataJSON#7d748d04 as nil")
	}
	b.PutID(DataJSONTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DataJSON) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dataJSON#7d748d04 as nil")
	}
	b.PutString(d.Data)
	return nil
}

// Decode implements bin.Decoder.
func (d *DataJSON) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dataJSON#7d748d04 to nil")
	}
	if err := b.ConsumeID(DataJSONTypeID); err != nil {
		return fmt.Errorf("unable to decode dataJSON#7d748d04: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DataJSON) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dataJSON#7d748d04 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dataJSON#7d748d04: field data: %w", err)
		}
		d.Data = value
	}
	return nil
}

// GetData returns value of Data field.
func (d *DataJSON) GetData() (value string) {
	if d == nil {
		return
	}
	return d.Data
}
