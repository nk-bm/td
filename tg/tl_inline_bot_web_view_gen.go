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

// InlineBotWebView represents TL type `inlineBotWebView#b57295d5`.
type InlineBotWebView struct {
	// Text field of InlineBotWebView.
	Text string
	// URL field of InlineBotWebView.
	URL string
}

// InlineBotWebViewTypeID is TL type id of InlineBotWebView.
const InlineBotWebViewTypeID = 0xb57295d5

// Ensuring interfaces in compile-time for InlineBotWebView.
var (
	_ bin.Encoder     = &InlineBotWebView{}
	_ bin.Decoder     = &InlineBotWebView{}
	_ bin.BareEncoder = &InlineBotWebView{}
	_ bin.BareDecoder = &InlineBotWebView{}
)

func (i *InlineBotWebView) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Text == "") {
		return false
	}
	if !(i.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineBotWebView) String() string {
	if i == nil {
		return "InlineBotWebView(nil)"
	}
	type Alias InlineBotWebView
	return fmt.Sprintf("InlineBotWebView%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineBotWebView) TypeID() uint32 {
	return InlineBotWebViewTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineBotWebView) TypeName() string {
	return "inlineBotWebView"
}

// TypeInfo returns info about TL type.
func (i *InlineBotWebView) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineBotWebView",
		ID:   InlineBotWebViewTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineBotWebView) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineBotWebView#b57295d5 as nil")
	}
	b.PutID(InlineBotWebViewTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineBotWebView) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineBotWebView#b57295d5 as nil")
	}
	b.PutString(i.Text)
	b.PutString(i.URL)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineBotWebView) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineBotWebView#b57295d5 to nil")
	}
	if err := b.ConsumeID(InlineBotWebViewTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineBotWebView#b57295d5: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineBotWebView) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineBotWebView#b57295d5 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineBotWebView#b57295d5: field text: %w", err)
		}
		i.Text = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineBotWebView#b57295d5: field url: %w", err)
		}
		i.URL = value
	}
	return nil
}

// GetText returns value of Text field.
func (i *InlineBotWebView) GetText() (value string) {
	if i == nil {
		return
	}
	return i.Text
}

// GetURL returns value of URL field.
func (i *InlineBotWebView) GetURL() (value string) {
	if i == nil {
		return
	}
	return i.URL
}
