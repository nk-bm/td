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

// PhoneDiscardGroupCallRequest represents TL type `phone.discardGroupCall#7a777135`.
type PhoneDiscardGroupCallRequest struct {
	// Call field of PhoneDiscardGroupCallRequest.
	Call InputGroupCall
}

// PhoneDiscardGroupCallRequestTypeID is TL type id of PhoneDiscardGroupCallRequest.
const PhoneDiscardGroupCallRequestTypeID = 0x7a777135

// Ensuring interfaces in compile-time for PhoneDiscardGroupCallRequest.
var (
	_ bin.Encoder     = &PhoneDiscardGroupCallRequest{}
	_ bin.Decoder     = &PhoneDiscardGroupCallRequest{}
	_ bin.BareEncoder = &PhoneDiscardGroupCallRequest{}
	_ bin.BareDecoder = &PhoneDiscardGroupCallRequest{}
)

func (d *PhoneDiscardGroupCallRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Call.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *PhoneDiscardGroupCallRequest) String() string {
	if d == nil {
		return "PhoneDiscardGroupCallRequest(nil)"
	}
	type Alias PhoneDiscardGroupCallRequest
	return fmt.Sprintf("PhoneDiscardGroupCallRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneDiscardGroupCallRequest) TypeID() uint32 {
	return PhoneDiscardGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneDiscardGroupCallRequest) TypeName() string {
	return "phone.discardGroupCall"
}

// TypeInfo returns info about TL type.
func (d *PhoneDiscardGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.discardGroupCall",
		ID:   PhoneDiscardGroupCallRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *PhoneDiscardGroupCallRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode phone.discardGroupCall#7a777135 as nil")
	}
	b.PutID(PhoneDiscardGroupCallRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *PhoneDiscardGroupCallRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode phone.discardGroupCall#7a777135 as nil")
	}
	if err := d.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.discardGroupCall#7a777135: field call: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *PhoneDiscardGroupCallRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode phone.discardGroupCall#7a777135 to nil")
	}
	if err := b.ConsumeID(PhoneDiscardGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.discardGroupCall#7a777135: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *PhoneDiscardGroupCallRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode phone.discardGroupCall#7a777135 to nil")
	}
	{
		if err := d.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.discardGroupCall#7a777135: field call: %w", err)
		}
	}
	return nil
}

// GetCall returns value of Call field.
func (d *PhoneDiscardGroupCallRequest) GetCall() (value InputGroupCall) {
	if d == nil {
		return
	}
	return d.Call
}

// PhoneDiscardGroupCall invokes method phone.discardGroupCall#7a777135 returning error if any.
func (c *Client) PhoneDiscardGroupCall(ctx context.Context, call InputGroupCall) (UpdatesClass, error) {
	var result UpdatesBox

	request := &PhoneDiscardGroupCallRequest{
		Call: call,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
