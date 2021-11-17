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

// Photo represents TL type `photo#105a0689`.
type Photo struct {
	// True, if stickers were added to the photo. The list of corresponding sticker sets can
	// be received using getAttachedStickerSets
	HasStickers bool
	// Photo minithumbnail; may be null
	Minithumbnail Minithumbnail
	// Available variants of the photo, in different sizes
	Sizes []PhotoSize
}

// PhotoTypeID is TL type id of Photo.
const PhotoTypeID = 0x105a0689

// Ensuring interfaces in compile-time for Photo.
var (
	_ bin.Encoder     = &Photo{}
	_ bin.Decoder     = &Photo{}
	_ bin.BareEncoder = &Photo{}
	_ bin.BareDecoder = &Photo{}
)

func (p *Photo) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.HasStickers == false) {
		return false
	}
	if !(p.Minithumbnail.Zero()) {
		return false
	}
	if !(p.Sizes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *Photo) String() string {
	if p == nil {
		return "Photo(nil)"
	}
	type Alias Photo
	return fmt.Sprintf("Photo%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Photo) TypeID() uint32 {
	return PhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*Photo) TypeName() string {
	return "photo"
}

// TypeInfo returns info about TL type.
func (p *Photo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photo",
		ID:   PhotoTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "HasStickers",
			SchemaName: "has_stickers",
		},
		{
			Name:       "Minithumbnail",
			SchemaName: "minithumbnail",
		},
		{
			Name:       "Sizes",
			SchemaName: "sizes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *Photo) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photo#105a0689 as nil")
	}
	b.PutID(PhotoTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *Photo) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photo#105a0689 as nil")
	}
	b.PutBool(p.HasStickers)
	if err := p.Minithumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photo#105a0689: field minithumbnail: %w", err)
	}
	b.PutInt(len(p.Sizes))
	for idx, v := range p.Sizes {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare photo#105a0689: field sizes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *Photo) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photo#105a0689 to nil")
	}
	if err := b.ConsumeID(PhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode photo#105a0689: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *Photo) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photo#105a0689 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode photo#105a0689: field has_stickers: %w", err)
		}
		p.HasStickers = value
	}
	{
		if err := p.Minithumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode photo#105a0689: field minithumbnail: %w", err)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photo#105a0689: field sizes: %w", err)
		}

		if headerLen > 0 {
			p.Sizes = make([]PhotoSize, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PhotoSize
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare photo#105a0689: field sizes: %w", err)
			}
			p.Sizes = append(p.Sizes, value)
		}
	}
	return nil
}

// GetHasStickers returns value of HasStickers field.
func (p *Photo) GetHasStickers() (value bool) {
	return p.HasStickers
}

// GetMinithumbnail returns value of Minithumbnail field.
func (p *Photo) GetMinithumbnail() (value Minithumbnail) {
	return p.Minithumbnail
}

// GetSizes returns value of Sizes field.
func (p *Photo) GetSizes() (value []PhotoSize) {
	return p.Sizes
}
