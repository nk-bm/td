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

// PhoneJoinGroupCallPresentationRequest represents TL type `phone.joinGroupCallPresentation#cbea6bc4`.
type PhoneJoinGroupCallPresentationRequest struct {
	// Call field of PhoneJoinGroupCallPresentationRequest.
	Call InputGroupCall
	// Params field of PhoneJoinGroupCallPresentationRequest.
	Params DataJSON
}

// PhoneJoinGroupCallPresentationRequestTypeID is TL type id of PhoneJoinGroupCallPresentationRequest.
const PhoneJoinGroupCallPresentationRequestTypeID = 0xcbea6bc4

// Ensuring interfaces in compile-time for PhoneJoinGroupCallPresentationRequest.
var (
	_ bin.Encoder     = &PhoneJoinGroupCallPresentationRequest{}
	_ bin.Decoder     = &PhoneJoinGroupCallPresentationRequest{}
	_ bin.BareEncoder = &PhoneJoinGroupCallPresentationRequest{}
	_ bin.BareDecoder = &PhoneJoinGroupCallPresentationRequest{}
)

func (j *PhoneJoinGroupCallPresentationRequest) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.Call.Zero()) {
		return false
	}
	if !(j.Params.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *PhoneJoinGroupCallPresentationRequest) String() string {
	if j == nil {
		return "PhoneJoinGroupCallPresentationRequest(nil)"
	}
	type Alias PhoneJoinGroupCallPresentationRequest
	return fmt.Sprintf("PhoneJoinGroupCallPresentationRequest%+v", Alias(*j))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneJoinGroupCallPresentationRequest) TypeID() uint32 {
	return PhoneJoinGroupCallPresentationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneJoinGroupCallPresentationRequest) TypeName() string {
	return "phone.joinGroupCallPresentation"
}

// TypeInfo returns info about TL type.
func (j *PhoneJoinGroupCallPresentationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.joinGroupCallPresentation",
		ID:   PhoneJoinGroupCallPresentationRequestTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "Params",
			SchemaName: "params",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *PhoneJoinGroupCallPresentationRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode phone.joinGroupCallPresentation#cbea6bc4 as nil")
	}
	b.PutID(PhoneJoinGroupCallPresentationRequestTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *PhoneJoinGroupCallPresentationRequest) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode phone.joinGroupCallPresentation#cbea6bc4 as nil")
	}
	if err := j.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.joinGroupCallPresentation#cbea6bc4: field call: %w", err)
	}
	if err := j.Params.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.joinGroupCallPresentation#cbea6bc4: field params: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *PhoneJoinGroupCallPresentationRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode phone.joinGroupCallPresentation#cbea6bc4 to nil")
	}
	if err := b.ConsumeID(PhoneJoinGroupCallPresentationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.joinGroupCallPresentation#cbea6bc4: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *PhoneJoinGroupCallPresentationRequest) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode phone.joinGroupCallPresentation#cbea6bc4 to nil")
	}
	{
		if err := j.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.joinGroupCallPresentation#cbea6bc4: field call: %w", err)
		}
	}
	{
		if err := j.Params.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.joinGroupCallPresentation#cbea6bc4: field params: %w", err)
		}
	}
	return nil
}

// GetCall returns value of Call field.
func (j *PhoneJoinGroupCallPresentationRequest) GetCall() (value InputGroupCall) {
	if j == nil {
		return
	}
	return j.Call
}

// GetParams returns value of Params field.
func (j *PhoneJoinGroupCallPresentationRequest) GetParams() (value DataJSON) {
	if j == nil {
		return
	}
	return j.Params
}

// PhoneJoinGroupCallPresentation invokes method phone.joinGroupCallPresentation#cbea6bc4 returning error if any.
func (c *Client) PhoneJoinGroupCallPresentation(ctx context.Context, request *PhoneJoinGroupCallPresentationRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
