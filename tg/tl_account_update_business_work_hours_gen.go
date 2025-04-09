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

// AccountUpdateBusinessWorkHoursRequest represents TL type `account.updateBusinessWorkHours#4b00e066`.
type AccountUpdateBusinessWorkHoursRequest struct {
	// Flags field of AccountUpdateBusinessWorkHoursRequest.
	Flags bin.Fields
	// BusinessWorkHours field of AccountUpdateBusinessWorkHoursRequest.
	//
	// Use SetBusinessWorkHours and GetBusinessWorkHours helpers.
	BusinessWorkHours BusinessWorkHours
}

// AccountUpdateBusinessWorkHoursRequestTypeID is TL type id of AccountUpdateBusinessWorkHoursRequest.
const AccountUpdateBusinessWorkHoursRequestTypeID = 0x4b00e066

// Ensuring interfaces in compile-time for AccountUpdateBusinessWorkHoursRequest.
var (
	_ bin.Encoder     = &AccountUpdateBusinessWorkHoursRequest{}
	_ bin.Decoder     = &AccountUpdateBusinessWorkHoursRequest{}
	_ bin.BareEncoder = &AccountUpdateBusinessWorkHoursRequest{}
	_ bin.BareDecoder = &AccountUpdateBusinessWorkHoursRequest{}
)

func (u *AccountUpdateBusinessWorkHoursRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.BusinessWorkHours.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateBusinessWorkHoursRequest) String() string {
	if u == nil {
		return "AccountUpdateBusinessWorkHoursRequest(nil)"
	}
	type Alias AccountUpdateBusinessWorkHoursRequest
	return fmt.Sprintf("AccountUpdateBusinessWorkHoursRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateBusinessWorkHoursRequest) TypeID() uint32 {
	return AccountUpdateBusinessWorkHoursRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateBusinessWorkHoursRequest) TypeName() string {
	return "account.updateBusinessWorkHours"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateBusinessWorkHoursRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateBusinessWorkHours",
		ID:   AccountUpdateBusinessWorkHoursRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BusinessWorkHours",
			SchemaName: "business_work_hours",
			Null:       !u.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *AccountUpdateBusinessWorkHoursRequest) SetFlags() {
	if !(u.BusinessWorkHours.Zero()) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *AccountUpdateBusinessWorkHoursRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateBusinessWorkHours#4b00e066 as nil")
	}
	b.PutID(AccountUpdateBusinessWorkHoursRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdateBusinessWorkHoursRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateBusinessWorkHours#4b00e066 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateBusinessWorkHours#4b00e066: field flags: %w", err)
	}
	if u.Flags.Has(0) {
		if err := u.BusinessWorkHours.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.updateBusinessWorkHours#4b00e066: field business_work_hours: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateBusinessWorkHoursRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateBusinessWorkHours#4b00e066 to nil")
	}
	if err := b.ConsumeID(AccountUpdateBusinessWorkHoursRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateBusinessWorkHours#4b00e066: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdateBusinessWorkHoursRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateBusinessWorkHours#4b00e066 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateBusinessWorkHours#4b00e066: field flags: %w", err)
		}
	}
	if u.Flags.Has(0) {
		if err := u.BusinessWorkHours.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateBusinessWorkHours#4b00e066: field business_work_hours: %w", err)
		}
	}
	return nil
}

// SetBusinessWorkHours sets value of BusinessWorkHours conditional field.
func (u *AccountUpdateBusinessWorkHoursRequest) SetBusinessWorkHours(value BusinessWorkHours) {
	u.Flags.Set(0)
	u.BusinessWorkHours = value
}

// GetBusinessWorkHours returns value of BusinessWorkHours conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateBusinessWorkHoursRequest) GetBusinessWorkHours() (value BusinessWorkHours, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.BusinessWorkHours, true
}

// AccountUpdateBusinessWorkHours invokes method account.updateBusinessWorkHours#4b00e066 returning error if any.
func (c *Client) AccountUpdateBusinessWorkHours(ctx context.Context, request *AccountUpdateBusinessWorkHoursRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
