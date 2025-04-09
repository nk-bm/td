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

// InputStickerSetItem represents TL type `inputStickerSetItem#32da9e9c`.
// Sticker in a stickerset
//
// See https://core.telegram.org/constructor/inputStickerSetItem for reference.
type InputStickerSetItem struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// The sticker
	Document InputDocumentClass
	// Associated emoji
	Emoji string
	// Coordinates for mask sticker
	//
	// Use SetMaskCoords and GetMaskCoords helpers.
	MaskCoords MaskCoords
	// Set of keywords, separated by commas (can't be provided for mask stickers)
	//
	// Use SetKeywords and GetKeywords helpers.
	Keywords string
}

// InputStickerSetItemTypeID is TL type id of InputStickerSetItem.
const InputStickerSetItemTypeID = 0x32da9e9c

// Ensuring interfaces in compile-time for InputStickerSetItem.
var (
	_ bin.Encoder     = &InputStickerSetItem{}
	_ bin.Decoder     = &InputStickerSetItem{}
	_ bin.BareEncoder = &InputStickerSetItem{}
	_ bin.BareDecoder = &InputStickerSetItem{}
)

func (i *InputStickerSetItem) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Document == nil) {
		return false
	}
	if !(i.Emoji == "") {
		return false
	}
	if !(i.MaskCoords.Zero()) {
		return false
	}
	if !(i.Keywords == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStickerSetItem) String() string {
	if i == nil {
		return "InputStickerSetItem(nil)"
	}
	type Alias InputStickerSetItem
	return fmt.Sprintf("InputStickerSetItem%+v", Alias(*i))
}

// FillFrom fills InputStickerSetItem from given interface.
func (i *InputStickerSetItem) FillFrom(from interface {
	GetDocument() (value InputDocumentClass)
	GetEmoji() (value string)
	GetMaskCoords() (value MaskCoords, ok bool)
	GetKeywords() (value string, ok bool)
}) {
	i.Document = from.GetDocument()
	i.Emoji = from.GetEmoji()
	if val, ok := from.GetMaskCoords(); ok {
		i.MaskCoords = val
	}

	if val, ok := from.GetKeywords(); ok {
		i.Keywords = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStickerSetItem) TypeID() uint32 {
	return InputStickerSetItemTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStickerSetItem) TypeName() string {
	return "inputStickerSetItem"
}

// TypeInfo returns info about TL type.
func (i *InputStickerSetItem) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStickerSetItem",
		ID:   InputStickerSetItemTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Document",
			SchemaName: "document",
		},
		{
			Name:       "Emoji",
			SchemaName: "emoji",
		},
		{
			Name:       "MaskCoords",
			SchemaName: "mask_coords",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "Keywords",
			SchemaName: "keywords",
			Null:       !i.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (i *InputStickerSetItem) SetFlags() {
	if !(i.MaskCoords.Zero()) {
		i.Flags.Set(0)
	}
	if !(i.Keywords == "") {
		i.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (i *InputStickerSetItem) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetItem#32da9e9c as nil")
	}
	b.PutID(InputStickerSetItemTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStickerSetItem) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetItem#32da9e9c as nil")
	}
	i.SetFlags()
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerSetItem#32da9e9c: field flags: %w", err)
	}
	if i.Document == nil {
		return fmt.Errorf("unable to encode inputStickerSetItem#32da9e9c: field document is nil")
	}
	if err := i.Document.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerSetItem#32da9e9c: field document: %w", err)
	}
	b.PutString(i.Emoji)
	if i.Flags.Has(0) {
		if err := i.MaskCoords.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputStickerSetItem#32da9e9c: field mask_coords: %w", err)
		}
	}
	if i.Flags.Has(1) {
		b.PutString(i.Keywords)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickerSetItem) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetItem#32da9e9c to nil")
	}
	if err := b.ConsumeID(InputStickerSetItemTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetItem#32da9e9c: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStickerSetItem) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetItem#32da9e9c to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#32da9e9c: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#32da9e9c: field document: %w", err)
		}
		i.Document = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#32da9e9c: field emoji: %w", err)
		}
		i.Emoji = value
	}
	if i.Flags.Has(0) {
		if err := i.MaskCoords.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#32da9e9c: field mask_coords: %w", err)
		}
	}
	if i.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#32da9e9c: field keywords: %w", err)
		}
		i.Keywords = value
	}
	return nil
}

// GetDocument returns value of Document field.
func (i *InputStickerSetItem) GetDocument() (value InputDocumentClass) {
	if i == nil {
		return
	}
	return i.Document
}

// GetEmoji returns value of Emoji field.
func (i *InputStickerSetItem) GetEmoji() (value string) {
	if i == nil {
		return
	}
	return i.Emoji
}

// SetMaskCoords sets value of MaskCoords conditional field.
func (i *InputStickerSetItem) SetMaskCoords(value MaskCoords) {
	i.Flags.Set(0)
	i.MaskCoords = value
}

// GetMaskCoords returns value of MaskCoords conditional field and
// boolean which is true if field was set.
func (i *InputStickerSetItem) GetMaskCoords() (value MaskCoords, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.MaskCoords, true
}

// SetKeywords sets value of Keywords conditional field.
func (i *InputStickerSetItem) SetKeywords(value string) {
	i.Flags.Set(1)
	i.Keywords = value
}

// GetKeywords returns value of Keywords conditional field and
// boolean which is true if field was set.
func (i *InputStickerSetItem) GetKeywords() (value string, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Keywords, true
}

// GetDocumentAsNotEmpty returns mapped value of Document field.
func (i *InputStickerSetItem) GetDocumentAsNotEmpty() (*InputDocument, bool) {
	return i.Document.AsNotEmpty()
}
