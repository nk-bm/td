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

// StickersSuggestedShortName represents TL type `stickers.suggestedShortName#85fea03f`.
type StickersSuggestedShortName struct {
	// ShortName field of StickersSuggestedShortName.
	ShortName string
}

// StickersSuggestedShortNameTypeID is TL type id of StickersSuggestedShortName.
const StickersSuggestedShortNameTypeID = 0x85fea03f

// Ensuring interfaces in compile-time for StickersSuggestedShortName.
var (
	_ bin.Encoder     = &StickersSuggestedShortName{}
	_ bin.Decoder     = &StickersSuggestedShortName{}
	_ bin.BareEncoder = &StickersSuggestedShortName{}
	_ bin.BareDecoder = &StickersSuggestedShortName{}
)

func (s *StickersSuggestedShortName) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ShortName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickersSuggestedShortName) String() string {
	if s == nil {
		return "StickersSuggestedShortName(nil)"
	}
	type Alias StickersSuggestedShortName
	return fmt.Sprintf("StickersSuggestedShortName%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickersSuggestedShortName) TypeID() uint32 {
	return StickersSuggestedShortNameTypeID
}

// TypeName returns name of type in TL schema.
func (*StickersSuggestedShortName) TypeName() string {
	return "stickers.suggestedShortName"
}

// TypeInfo returns info about TL type.
func (s *StickersSuggestedShortName) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickers.suggestedShortName",
		ID:   StickersSuggestedShortNameTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShortName",
			SchemaName: "short_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StickersSuggestedShortName) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickers.suggestedShortName#85fea03f as nil")
	}
	b.PutID(StickersSuggestedShortNameTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StickersSuggestedShortName) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickers.suggestedShortName#85fea03f as nil")
	}
	b.PutString(s.ShortName)
	return nil
}

// Decode implements bin.Decoder.
func (s *StickersSuggestedShortName) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickers.suggestedShortName#85fea03f to nil")
	}
	if err := b.ConsumeID(StickersSuggestedShortNameTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.suggestedShortName#85fea03f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StickersSuggestedShortName) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickers.suggestedShortName#85fea03f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickers.suggestedShortName#85fea03f: field short_name: %w", err)
		}
		s.ShortName = value
	}
	return nil
}

// GetShortName returns value of ShortName field.
func (s *StickersSuggestedShortName) GetShortName() (value string) {
	if s == nil {
		return
	}
	return s.ShortName
}
