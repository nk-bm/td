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

// AccountSentEmailCode represents TL type `account.sentEmailCode#811f854f`.
type AccountSentEmailCode struct {
	// EmailPattern field of AccountSentEmailCode.
	EmailPattern string
	// Length field of AccountSentEmailCode.
	Length int
}

// AccountSentEmailCodeTypeID is TL type id of AccountSentEmailCode.
const AccountSentEmailCodeTypeID = 0x811f854f

// Ensuring interfaces in compile-time for AccountSentEmailCode.
var (
	_ bin.Encoder     = &AccountSentEmailCode{}
	_ bin.Decoder     = &AccountSentEmailCode{}
	_ bin.BareEncoder = &AccountSentEmailCode{}
	_ bin.BareDecoder = &AccountSentEmailCode{}
)

func (s *AccountSentEmailCode) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.EmailPattern == "") {
		return false
	}
	if !(s.Length == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSentEmailCode) String() string {
	if s == nil {
		return "AccountSentEmailCode(nil)"
	}
	type Alias AccountSentEmailCode
	return fmt.Sprintf("AccountSentEmailCode%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSentEmailCode) TypeID() uint32 {
	return AccountSentEmailCodeTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSentEmailCode) TypeName() string {
	return "account.sentEmailCode"
}

// TypeInfo returns info about TL type.
func (s *AccountSentEmailCode) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.sentEmailCode",
		ID:   AccountSentEmailCodeTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EmailPattern",
			SchemaName: "email_pattern",
		},
		{
			Name:       "Length",
			SchemaName: "length",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSentEmailCode) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sentEmailCode#811f854f as nil")
	}
	b.PutID(AccountSentEmailCodeTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSentEmailCode) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sentEmailCode#811f854f as nil")
	}
	b.PutString(s.EmailPattern)
	b.PutInt(s.Length)
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSentEmailCode) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sentEmailCode#811f854f to nil")
	}
	if err := b.ConsumeID(AccountSentEmailCodeTypeID); err != nil {
		return fmt.Errorf("unable to decode account.sentEmailCode#811f854f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSentEmailCode) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sentEmailCode#811f854f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.sentEmailCode#811f854f: field email_pattern: %w", err)
		}
		s.EmailPattern = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.sentEmailCode#811f854f: field length: %w", err)
		}
		s.Length = value
	}
	return nil
}

// GetEmailPattern returns value of EmailPattern field.
func (s *AccountSentEmailCode) GetEmailPattern() (value string) {
	if s == nil {
		return
	}
	return s.EmailPattern
}

// GetLength returns value of Length field.
func (s *AccountSentEmailCode) GetLength() (value int) {
	if s == nil {
		return
	}
	return s.Length
}
