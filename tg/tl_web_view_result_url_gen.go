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

// WebViewResultURL represents TL type `webViewResultUrl#4d22ff98`.
type WebViewResultURL struct {
	// Flags field of WebViewResultURL.
	Flags bin.Fields
	// Fullsize field of WebViewResultURL.
	Fullsize bool
	// Fullscreen field of WebViewResultURL.
	Fullscreen bool
	// QueryID field of WebViewResultURL.
	//
	// Use SetQueryID and GetQueryID helpers.
	QueryID int64
	// URL field of WebViewResultURL.
	URL string
}

// WebViewResultURLTypeID is TL type id of WebViewResultURL.
const WebViewResultURLTypeID = 0x4d22ff98

// Ensuring interfaces in compile-time for WebViewResultURL.
var (
	_ bin.Encoder     = &WebViewResultURL{}
	_ bin.Decoder     = &WebViewResultURL{}
	_ bin.BareEncoder = &WebViewResultURL{}
	_ bin.BareDecoder = &WebViewResultURL{}
)

func (w *WebViewResultURL) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.Fullsize == false) {
		return false
	}
	if !(w.Fullscreen == false) {
		return false
	}
	if !(w.QueryID == 0) {
		return false
	}
	if !(w.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebViewResultURL) String() string {
	if w == nil {
		return "WebViewResultURL(nil)"
	}
	type Alias WebViewResultURL
	return fmt.Sprintf("WebViewResultURL%+v", Alias(*w))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebViewResultURL) TypeID() uint32 {
	return WebViewResultURLTypeID
}

// TypeName returns name of type in TL schema.
func (*WebViewResultURL) TypeName() string {
	return "webViewResultUrl"
}

// TypeInfo returns info about TL type.
func (w *WebViewResultURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webViewResultUrl",
		ID:   WebViewResultURLTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Fullsize",
			SchemaName: "fullsize",
			Null:       !w.Flags.Has(1),
		},
		{
			Name:       "Fullscreen",
			SchemaName: "fullscreen",
			Null:       !w.Flags.Has(2),
		},
		{
			Name:       "QueryID",
			SchemaName: "query_id",
			Null:       !w.Flags.Has(0),
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (w *WebViewResultURL) SetFlags() {
	if !(w.Fullsize == false) {
		w.Flags.Set(1)
	}
	if !(w.Fullscreen == false) {
		w.Flags.Set(2)
	}
	if !(w.QueryID == 0) {
		w.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (w *WebViewResultURL) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webViewResultUrl#4d22ff98 as nil")
	}
	b.PutID(WebViewResultURLTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebViewResultURL) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webViewResultUrl#4d22ff98 as nil")
	}
	w.SetFlags()
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webViewResultUrl#4d22ff98: field flags: %w", err)
	}
	if w.Flags.Has(0) {
		b.PutLong(w.QueryID)
	}
	b.PutString(w.URL)
	return nil
}

// Decode implements bin.Decoder.
func (w *WebViewResultURL) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webViewResultUrl#4d22ff98 to nil")
	}
	if err := b.ConsumeID(WebViewResultURLTypeID); err != nil {
		return fmt.Errorf("unable to decode webViewResultUrl#4d22ff98: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebViewResultURL) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webViewResultUrl#4d22ff98 to nil")
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webViewResultUrl#4d22ff98: field flags: %w", err)
		}
	}
	w.Fullsize = w.Flags.Has(1)
	w.Fullscreen = w.Flags.Has(2)
	if w.Flags.Has(0) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode webViewResultUrl#4d22ff98: field query_id: %w", err)
		}
		w.QueryID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webViewResultUrl#4d22ff98: field url: %w", err)
		}
		w.URL = value
	}
	return nil
}

// SetFullsize sets value of Fullsize conditional field.
func (w *WebViewResultURL) SetFullsize(value bool) {
	if value {
		w.Flags.Set(1)
		w.Fullsize = true
	} else {
		w.Flags.Unset(1)
		w.Fullsize = false
	}
}

// GetFullsize returns value of Fullsize conditional field.
func (w *WebViewResultURL) GetFullsize() (value bool) {
	if w == nil {
		return
	}
	return w.Flags.Has(1)
}

// SetFullscreen sets value of Fullscreen conditional field.
func (w *WebViewResultURL) SetFullscreen(value bool) {
	if value {
		w.Flags.Set(2)
		w.Fullscreen = true
	} else {
		w.Flags.Unset(2)
		w.Fullscreen = false
	}
}

// GetFullscreen returns value of Fullscreen conditional field.
func (w *WebViewResultURL) GetFullscreen() (value bool) {
	if w == nil {
		return
	}
	return w.Flags.Has(2)
}

// SetQueryID sets value of QueryID conditional field.
func (w *WebViewResultURL) SetQueryID(value int64) {
	w.Flags.Set(0)
	w.QueryID = value
}

// GetQueryID returns value of QueryID conditional field and
// boolean which is true if field was set.
func (w *WebViewResultURL) GetQueryID() (value int64, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(0) {
		return value, false
	}
	return w.QueryID, true
}

// GetURL returns value of URL field.
func (w *WebViewResultURL) GetURL() (value string) {
	if w == nil {
		return
	}
	return w.URL
}
