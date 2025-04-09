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

// PaymentSavedCredentialsCard represents TL type `paymentSavedCredentialsCard#cdc27a1f`.
type PaymentSavedCredentialsCard struct {
	// ID field of PaymentSavedCredentialsCard.
	ID string
	// Title field of PaymentSavedCredentialsCard.
	Title string
}

// PaymentSavedCredentialsCardTypeID is TL type id of PaymentSavedCredentialsCard.
const PaymentSavedCredentialsCardTypeID = 0xcdc27a1f

// Ensuring interfaces in compile-time for PaymentSavedCredentialsCard.
var (
	_ bin.Encoder     = &PaymentSavedCredentialsCard{}
	_ bin.Decoder     = &PaymentSavedCredentialsCard{}
	_ bin.BareEncoder = &PaymentSavedCredentialsCard{}
	_ bin.BareDecoder = &PaymentSavedCredentialsCard{}
)

func (p *PaymentSavedCredentialsCard) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == "") {
		return false
	}
	if !(p.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentSavedCredentialsCard) String() string {
	if p == nil {
		return "PaymentSavedCredentialsCard(nil)"
	}
	type Alias PaymentSavedCredentialsCard
	return fmt.Sprintf("PaymentSavedCredentialsCard%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentSavedCredentialsCard) TypeID() uint32 {
	return PaymentSavedCredentialsCardTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentSavedCredentialsCard) TypeName() string {
	return "paymentSavedCredentialsCard"
}

// TypeInfo returns info about TL type.
func (p *PaymentSavedCredentialsCard) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "paymentSavedCredentialsCard",
		ID:   PaymentSavedCredentialsCardTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentSavedCredentialsCard) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentSavedCredentialsCard#cdc27a1f as nil")
	}
	b.PutID(PaymentSavedCredentialsCardTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentSavedCredentialsCard) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentSavedCredentialsCard#cdc27a1f as nil")
	}
	b.PutString(p.ID)
	b.PutString(p.Title)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentSavedCredentialsCard) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentSavedCredentialsCard#cdc27a1f to nil")
	}
	if err := b.ConsumeID(PaymentSavedCredentialsCardTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentSavedCredentialsCard#cdc27a1f: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentSavedCredentialsCard) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentSavedCredentialsCard#cdc27a1f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentSavedCredentialsCard#cdc27a1f: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentSavedCredentialsCard#cdc27a1f: field title: %w", err)
		}
		p.Title = value
	}
	return nil
}

// GetID returns value of ID field.
func (p *PaymentSavedCredentialsCard) GetID() (value string) {
	if p == nil {
		return
	}
	return p.ID
}

// GetTitle returns value of Title field.
func (p *PaymentSavedCredentialsCard) GetTitle() (value string) {
	if p == nil {
		return
	}
	return p.Title
}
