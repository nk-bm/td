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

// PaymentsRefundStarsChargeRequest represents TL type `payments.refundStarsCharge#25ae8f4a`.
type PaymentsRefundStarsChargeRequest struct {
	// UserID field of PaymentsRefundStarsChargeRequest.
	UserID InputUserClass
	// ChargeID field of PaymentsRefundStarsChargeRequest.
	ChargeID string
}

// PaymentsRefundStarsChargeRequestTypeID is TL type id of PaymentsRefundStarsChargeRequest.
const PaymentsRefundStarsChargeRequestTypeID = 0x25ae8f4a

// Ensuring interfaces in compile-time for PaymentsRefundStarsChargeRequest.
var (
	_ bin.Encoder     = &PaymentsRefundStarsChargeRequest{}
	_ bin.Decoder     = &PaymentsRefundStarsChargeRequest{}
	_ bin.BareEncoder = &PaymentsRefundStarsChargeRequest{}
	_ bin.BareDecoder = &PaymentsRefundStarsChargeRequest{}
)

func (r *PaymentsRefundStarsChargeRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.UserID == nil) {
		return false
	}
	if !(r.ChargeID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *PaymentsRefundStarsChargeRequest) String() string {
	if r == nil {
		return "PaymentsRefundStarsChargeRequest(nil)"
	}
	type Alias PaymentsRefundStarsChargeRequest
	return fmt.Sprintf("PaymentsRefundStarsChargeRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsRefundStarsChargeRequest) TypeID() uint32 {
	return PaymentsRefundStarsChargeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsRefundStarsChargeRequest) TypeName() string {
	return "payments.refundStarsCharge"
}

// TypeInfo returns info about TL type.
func (r *PaymentsRefundStarsChargeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.refundStarsCharge",
		ID:   PaymentsRefundStarsChargeRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "ChargeID",
			SchemaName: "charge_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *PaymentsRefundStarsChargeRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode payments.refundStarsCharge#25ae8f4a as nil")
	}
	b.PutID(PaymentsRefundStarsChargeRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *PaymentsRefundStarsChargeRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode payments.refundStarsCharge#25ae8f4a as nil")
	}
	if r.UserID == nil {
		return fmt.Errorf("unable to encode payments.refundStarsCharge#25ae8f4a: field user_id is nil")
	}
	if err := r.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.refundStarsCharge#25ae8f4a: field user_id: %w", err)
	}
	b.PutString(r.ChargeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *PaymentsRefundStarsChargeRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode payments.refundStarsCharge#25ae8f4a to nil")
	}
	if err := b.ConsumeID(PaymentsRefundStarsChargeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.refundStarsCharge#25ae8f4a: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *PaymentsRefundStarsChargeRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode payments.refundStarsCharge#25ae8f4a to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.refundStarsCharge#25ae8f4a: field user_id: %w", err)
		}
		r.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.refundStarsCharge#25ae8f4a: field charge_id: %w", err)
		}
		r.ChargeID = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (r *PaymentsRefundStarsChargeRequest) GetUserID() (value InputUserClass) {
	if r == nil {
		return
	}
	return r.UserID
}

// GetChargeID returns value of ChargeID field.
func (r *PaymentsRefundStarsChargeRequest) GetChargeID() (value string) {
	if r == nil {
		return
	}
	return r.ChargeID
}

// PaymentsRefundStarsCharge invokes method payments.refundStarsCharge#25ae8f4a returning error if any.
func (c *Client) PaymentsRefundStarsCharge(ctx context.Context, request *PaymentsRefundStarsChargeRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
