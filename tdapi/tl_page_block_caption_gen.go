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

// PageBlockCaption represents TL type `pageBlockCaption#b9a9a476`.
type PageBlockCaption struct {
	// Content of the caption
	Text RichTextClass
	// Block credit (like HTML tag <cite>)
	Credit RichTextClass
}

// PageBlockCaptionTypeID is TL type id of PageBlockCaption.
const PageBlockCaptionTypeID = 0xb9a9a476

// Ensuring interfaces in compile-time for PageBlockCaption.
var (
	_ bin.Encoder     = &PageBlockCaption{}
	_ bin.Decoder     = &PageBlockCaption{}
	_ bin.BareEncoder = &PageBlockCaption{}
	_ bin.BareDecoder = &PageBlockCaption{}
)

func (p *PageBlockCaption) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Text == nil) {
		return false
	}
	if !(p.Credit == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageBlockCaption) String() string {
	if p == nil {
		return "PageBlockCaption(nil)"
	}
	type Alias PageBlockCaption
	return fmt.Sprintf("PageBlockCaption%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PageBlockCaption) TypeID() uint32 {
	return PageBlockCaptionTypeID
}

// TypeName returns name of type in TL schema.
func (*PageBlockCaption) TypeName() string {
	return "pageBlockCaption"
}

// TypeInfo returns info about TL type.
func (p *PageBlockCaption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pageBlockCaption",
		ID:   PageBlockCaptionTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "Credit",
			SchemaName: "credit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PageBlockCaption) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockCaption#b9a9a476 as nil")
	}
	b.PutID(PageBlockCaptionTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PageBlockCaption) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockCaption#b9a9a476 as nil")
	}
	if p.Text == nil {
		return fmt.Errorf("unable to encode pageBlockCaption#b9a9a476: field text is nil")
	}
	if err := p.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageBlockCaption#b9a9a476: field text: %w", err)
	}
	if p.Credit == nil {
		return fmt.Errorf("unable to encode pageBlockCaption#b9a9a476: field credit is nil")
	}
	if err := p.Credit.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageBlockCaption#b9a9a476: field credit: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageBlockCaption) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockCaption#b9a9a476 to nil")
	}
	if err := b.ConsumeID(PageBlockCaptionTypeID); err != nil {
		return fmt.Errorf("unable to decode pageBlockCaption#b9a9a476: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PageBlockCaption) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockCaption#b9a9a476 to nil")
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockCaption#b9a9a476: field text: %w", err)
		}
		p.Text = value
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockCaption#b9a9a476: field credit: %w", err)
		}
		p.Credit = value
	}
	return nil
}

// GetText returns value of Text field.
func (p *PageBlockCaption) GetText() (value RichTextClass) {
	return p.Text
}

// GetCredit returns value of Credit field.
func (p *PageBlockCaption) GetCredit() (value RichTextClass) {
	return p.Credit
}
