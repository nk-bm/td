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

// AccountUpdateBusinessIntroRequest represents TL type `account.updateBusinessIntro#a614d034`.
// Set or remove the Telegram Business introduction »¹.
//
// Links:
//  1. https://core.telegram.org/api/business#business-introduction
//
// See https://core.telegram.org/method/account.updateBusinessIntro for reference.
type AccountUpdateBusinessIntroRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Telegram Business introduction, to remove it call the method without setting this flag.
	//
	// Use SetIntro and GetIntro helpers.
	Intro InputBusinessIntro
}

// AccountUpdateBusinessIntroRequestTypeID is TL type id of AccountUpdateBusinessIntroRequest.
const AccountUpdateBusinessIntroRequestTypeID = 0xa614d034

// Ensuring interfaces in compile-time for AccountUpdateBusinessIntroRequest.
var (
	_ bin.Encoder     = &AccountUpdateBusinessIntroRequest{}
	_ bin.Decoder     = &AccountUpdateBusinessIntroRequest{}
	_ bin.BareEncoder = &AccountUpdateBusinessIntroRequest{}
	_ bin.BareDecoder = &AccountUpdateBusinessIntroRequest{}
)

func (u *AccountUpdateBusinessIntroRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Intro.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateBusinessIntroRequest) String() string {
	if u == nil {
		return "AccountUpdateBusinessIntroRequest(nil)"
	}
	type Alias AccountUpdateBusinessIntroRequest
	return fmt.Sprintf("AccountUpdateBusinessIntroRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdateBusinessIntroRequest from given interface.
func (u *AccountUpdateBusinessIntroRequest) FillFrom(from interface {
	GetIntro() (value InputBusinessIntro, ok bool)
}) {
	if val, ok := from.GetIntro(); ok {
		u.Intro = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateBusinessIntroRequest) TypeID() uint32 {
	return AccountUpdateBusinessIntroRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateBusinessIntroRequest) TypeName() string {
	return "account.updateBusinessIntro"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateBusinessIntroRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateBusinessIntro",
		ID:   AccountUpdateBusinessIntroRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Intro",
			SchemaName: "intro",
			Null:       !u.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *AccountUpdateBusinessIntroRequest) SetFlags() {
	if !(u.Intro.Zero()) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *AccountUpdateBusinessIntroRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateBusinessIntro#a614d034 as nil")
	}
	b.PutID(AccountUpdateBusinessIntroRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdateBusinessIntroRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateBusinessIntro#a614d034 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateBusinessIntro#a614d034: field flags: %w", err)
	}
	if u.Flags.Has(0) {
		if err := u.Intro.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.updateBusinessIntro#a614d034: field intro: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateBusinessIntroRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateBusinessIntro#a614d034 to nil")
	}
	if err := b.ConsumeID(AccountUpdateBusinessIntroRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateBusinessIntro#a614d034: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdateBusinessIntroRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateBusinessIntro#a614d034 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateBusinessIntro#a614d034: field flags: %w", err)
		}
	}
	if u.Flags.Has(0) {
		if err := u.Intro.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateBusinessIntro#a614d034: field intro: %w", err)
		}
	}
	return nil
}

// SetIntro sets value of Intro conditional field.
func (u *AccountUpdateBusinessIntroRequest) SetIntro(value InputBusinessIntro) {
	u.Flags.Set(0)
	u.Intro = value
}

// GetIntro returns value of Intro conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateBusinessIntroRequest) GetIntro() (value InputBusinessIntro, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.Intro, true
}

// AccountUpdateBusinessIntro invokes method account.updateBusinessIntro#a614d034 returning error if any.
// Set or remove the Telegram Business introduction »¹.
//
// Links:
//  1. https://core.telegram.org/api/business#business-introduction
//
// See https://core.telegram.org/method/account.updateBusinessIntro for reference.
func (c *Client) AccountUpdateBusinessIntro(ctx context.Context, request *AccountUpdateBusinessIntroRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
