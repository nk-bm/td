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

// BankCardOpenURL represents TL type `bankCardOpenUrl#f568028a`.
type BankCardOpenURL struct {
	// URL field of BankCardOpenURL.
	URL string
	// Name field of BankCardOpenURL.
	Name string
}

// BankCardOpenURLTypeID is TL type id of BankCardOpenURL.
const BankCardOpenURLTypeID = 0xf568028a

// Ensuring interfaces in compile-time for BankCardOpenURL.
var (
	_ bin.Encoder     = &BankCardOpenURL{}
	_ bin.Decoder     = &BankCardOpenURL{}
	_ bin.BareEncoder = &BankCardOpenURL{}
	_ bin.BareDecoder = &BankCardOpenURL{}
)

func (b *BankCardOpenURL) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.URL == "") {
		return false
	}
	if !(b.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BankCardOpenURL) String() string {
	if b == nil {
		return "BankCardOpenURL(nil)"
	}
	type Alias BankCardOpenURL
	return fmt.Sprintf("BankCardOpenURL%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BankCardOpenURL) TypeID() uint32 {
	return BankCardOpenURLTypeID
}

// TypeName returns name of type in TL schema.
func (*BankCardOpenURL) TypeName() string {
	return "bankCardOpenUrl"
}

// TypeInfo returns info about TL type.
func (b *BankCardOpenURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bankCardOpenUrl",
		ID:   BankCardOpenURLTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BankCardOpenURL) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bankCardOpenUrl#f568028a as nil")
	}
	buf.PutID(BankCardOpenURLTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BankCardOpenURL) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bankCardOpenUrl#f568028a as nil")
	}
	buf.PutString(b.URL)
	buf.PutString(b.Name)
	return nil
}

// Decode implements bin.Decoder.
func (b *BankCardOpenURL) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bankCardOpenUrl#f568028a to nil")
	}
	if err := buf.ConsumeID(BankCardOpenURLTypeID); err != nil {
		return fmt.Errorf("unable to decode bankCardOpenUrl#f568028a: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BankCardOpenURL) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bankCardOpenUrl#f568028a to nil")
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode bankCardOpenUrl#f568028a: field url: %w", err)
		}
		b.URL = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode bankCardOpenUrl#f568028a: field name: %w", err)
		}
		b.Name = value
	}
	return nil
}

// GetURL returns value of URL field.
func (b *BankCardOpenURL) GetURL() (value string) {
	if b == nil {
		return
	}
	return b.URL
}

// GetName returns value of Name field.
func (b *BankCardOpenURL) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}
