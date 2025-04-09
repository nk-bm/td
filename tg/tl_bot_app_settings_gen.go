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

// BotAppSettings represents TL type `botAppSettings#c99b1950`.
type BotAppSettings struct {
	// Flags field of BotAppSettings.
	Flags bin.Fields
	// PlaceholderPath field of BotAppSettings.
	//
	// Use SetPlaceholderPath and GetPlaceholderPath helpers.
	PlaceholderPath []byte
	// BackgroundColor field of BotAppSettings.
	//
	// Use SetBackgroundColor and GetBackgroundColor helpers.
	BackgroundColor int
	// BackgroundDarkColor field of BotAppSettings.
	//
	// Use SetBackgroundDarkColor and GetBackgroundDarkColor helpers.
	BackgroundDarkColor int
	// HeaderColor field of BotAppSettings.
	//
	// Use SetHeaderColor and GetHeaderColor helpers.
	HeaderColor int
	// HeaderDarkColor field of BotAppSettings.
	//
	// Use SetHeaderDarkColor and GetHeaderDarkColor helpers.
	HeaderDarkColor int
}

// BotAppSettingsTypeID is TL type id of BotAppSettings.
const BotAppSettingsTypeID = 0xc99b1950

// Ensuring interfaces in compile-time for BotAppSettings.
var (
	_ bin.Encoder     = &BotAppSettings{}
	_ bin.Decoder     = &BotAppSettings{}
	_ bin.BareEncoder = &BotAppSettings{}
	_ bin.BareDecoder = &BotAppSettings{}
)

func (b *BotAppSettings) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Flags.Zero()) {
		return false
	}
	if !(b.PlaceholderPath == nil) {
		return false
	}
	if !(b.BackgroundColor == 0) {
		return false
	}
	if !(b.BackgroundDarkColor == 0) {
		return false
	}
	if !(b.HeaderColor == 0) {
		return false
	}
	if !(b.HeaderDarkColor == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotAppSettings) String() string {
	if b == nil {
		return "BotAppSettings(nil)"
	}
	type Alias BotAppSettings
	return fmt.Sprintf("BotAppSettings%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotAppSettings) TypeID() uint32 {
	return BotAppSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*BotAppSettings) TypeName() string {
	return "botAppSettings"
}

// TypeInfo returns info about TL type.
func (b *BotAppSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botAppSettings",
		ID:   BotAppSettingsTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PlaceholderPath",
			SchemaName: "placeholder_path",
			Null:       !b.Flags.Has(0),
		},
		{
			Name:       "BackgroundColor",
			SchemaName: "background_color",
			Null:       !b.Flags.Has(1),
		},
		{
			Name:       "BackgroundDarkColor",
			SchemaName: "background_dark_color",
			Null:       !b.Flags.Has(2),
		},
		{
			Name:       "HeaderColor",
			SchemaName: "header_color",
			Null:       !b.Flags.Has(3),
		},
		{
			Name:       "HeaderDarkColor",
			SchemaName: "header_dark_color",
			Null:       !b.Flags.Has(4),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (b *BotAppSettings) SetFlags() {
	if !(b.PlaceholderPath == nil) {
		b.Flags.Set(0)
	}
	if !(b.BackgroundColor == 0) {
		b.Flags.Set(1)
	}
	if !(b.BackgroundDarkColor == 0) {
		b.Flags.Set(2)
	}
	if !(b.HeaderColor == 0) {
		b.Flags.Set(3)
	}
	if !(b.HeaderDarkColor == 0) {
		b.Flags.Set(4)
	}
}

// Encode implements bin.Encoder.
func (b *BotAppSettings) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botAppSettings#c99b1950 as nil")
	}
	buf.PutID(BotAppSettingsTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotAppSettings) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botAppSettings#c99b1950 as nil")
	}
	b.SetFlags()
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botAppSettings#c99b1950: field flags: %w", err)
	}
	if b.Flags.Has(0) {
		buf.PutBytes(b.PlaceholderPath)
	}
	if b.Flags.Has(1) {
		buf.PutInt(b.BackgroundColor)
	}
	if b.Flags.Has(2) {
		buf.PutInt(b.BackgroundDarkColor)
	}
	if b.Flags.Has(3) {
		buf.PutInt(b.HeaderColor)
	}
	if b.Flags.Has(4) {
		buf.PutInt(b.HeaderDarkColor)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotAppSettings) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botAppSettings#c99b1950 to nil")
	}
	if err := buf.ConsumeID(BotAppSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode botAppSettings#c99b1950: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotAppSettings) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botAppSettings#c99b1950 to nil")
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botAppSettings#c99b1950: field flags: %w", err)
		}
	}
	if b.Flags.Has(0) {
		value, err := buf.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode botAppSettings#c99b1950: field placeholder_path: %w", err)
		}
		b.PlaceholderPath = value
	}
	if b.Flags.Has(1) {
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode botAppSettings#c99b1950: field background_color: %w", err)
		}
		b.BackgroundColor = value
	}
	if b.Flags.Has(2) {
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode botAppSettings#c99b1950: field background_dark_color: %w", err)
		}
		b.BackgroundDarkColor = value
	}
	if b.Flags.Has(3) {
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode botAppSettings#c99b1950: field header_color: %w", err)
		}
		b.HeaderColor = value
	}
	if b.Flags.Has(4) {
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode botAppSettings#c99b1950: field header_dark_color: %w", err)
		}
		b.HeaderDarkColor = value
	}
	return nil
}

// SetPlaceholderPath sets value of PlaceholderPath conditional field.
func (b *BotAppSettings) SetPlaceholderPath(value []byte) {
	b.Flags.Set(0)
	b.PlaceholderPath = value
}

// GetPlaceholderPath returns value of PlaceholderPath conditional field and
// boolean which is true if field was set.
func (b *BotAppSettings) GetPlaceholderPath() (value []byte, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(0) {
		return value, false
	}
	return b.PlaceholderPath, true
}

// SetBackgroundColor sets value of BackgroundColor conditional field.
func (b *BotAppSettings) SetBackgroundColor(value int) {
	b.Flags.Set(1)
	b.BackgroundColor = value
}

// GetBackgroundColor returns value of BackgroundColor conditional field and
// boolean which is true if field was set.
func (b *BotAppSettings) GetBackgroundColor() (value int, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(1) {
		return value, false
	}
	return b.BackgroundColor, true
}

// SetBackgroundDarkColor sets value of BackgroundDarkColor conditional field.
func (b *BotAppSettings) SetBackgroundDarkColor(value int) {
	b.Flags.Set(2)
	b.BackgroundDarkColor = value
}

// GetBackgroundDarkColor returns value of BackgroundDarkColor conditional field and
// boolean which is true if field was set.
func (b *BotAppSettings) GetBackgroundDarkColor() (value int, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.BackgroundDarkColor, true
}

// SetHeaderColor sets value of HeaderColor conditional field.
func (b *BotAppSettings) SetHeaderColor(value int) {
	b.Flags.Set(3)
	b.HeaderColor = value
}

// GetHeaderColor returns value of HeaderColor conditional field and
// boolean which is true if field was set.
func (b *BotAppSettings) GetHeaderColor() (value int, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(3) {
		return value, false
	}
	return b.HeaderColor, true
}

// SetHeaderDarkColor sets value of HeaderDarkColor conditional field.
func (b *BotAppSettings) SetHeaderDarkColor(value int) {
	b.Flags.Set(4)
	b.HeaderDarkColor = value
}

// GetHeaderDarkColor returns value of HeaderDarkColor conditional field and
// boolean which is true if field was set.
func (b *BotAppSettings) GetHeaderDarkColor() (value int, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(4) {
		return value, false
	}
	return b.HeaderDarkColor, true
}
