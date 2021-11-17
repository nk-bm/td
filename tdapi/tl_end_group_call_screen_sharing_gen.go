// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
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
)

// EndGroupCallScreenSharingRequest represents TL type `endGroupCallScreenSharing#85f41c4c`.
type EndGroupCallScreenSharingRequest struct {
	// Group call identifier
	GroupCallID int32
}

// EndGroupCallScreenSharingRequestTypeID is TL type id of EndGroupCallScreenSharingRequest.
const EndGroupCallScreenSharingRequestTypeID = 0x85f41c4c

// Ensuring interfaces in compile-time for EndGroupCallScreenSharingRequest.
var (
	_ bin.Encoder     = &EndGroupCallScreenSharingRequest{}
	_ bin.Decoder     = &EndGroupCallScreenSharingRequest{}
	_ bin.BareEncoder = &EndGroupCallScreenSharingRequest{}
	_ bin.BareDecoder = &EndGroupCallScreenSharingRequest{}
)

func (e *EndGroupCallScreenSharingRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.GroupCallID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EndGroupCallScreenSharingRequest) String() string {
	if e == nil {
		return "EndGroupCallScreenSharingRequest(nil)"
	}
	type Alias EndGroupCallScreenSharingRequest
	return fmt.Sprintf("EndGroupCallScreenSharingRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EndGroupCallScreenSharingRequest) TypeID() uint32 {
	return EndGroupCallScreenSharingRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EndGroupCallScreenSharingRequest) TypeName() string {
	return "endGroupCallScreenSharing"
}

// TypeInfo returns info about TL type.
func (e *EndGroupCallScreenSharingRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "endGroupCallScreenSharing",
		ID:   EndGroupCallScreenSharingRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EndGroupCallScreenSharingRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode endGroupCallScreenSharing#85f41c4c as nil")
	}
	b.PutID(EndGroupCallScreenSharingRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EndGroupCallScreenSharingRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode endGroupCallScreenSharing#85f41c4c as nil")
	}
	b.PutInt32(e.GroupCallID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EndGroupCallScreenSharingRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode endGroupCallScreenSharing#85f41c4c to nil")
	}
	if err := b.ConsumeID(EndGroupCallScreenSharingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode endGroupCallScreenSharing#85f41c4c: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EndGroupCallScreenSharingRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode endGroupCallScreenSharing#85f41c4c to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode endGroupCallScreenSharing#85f41c4c: field group_call_id: %w", err)
		}
		e.GroupCallID = value
	}
	return nil
}

// GetGroupCallID returns value of GroupCallID field.
func (e *EndGroupCallScreenSharingRequest) GetGroupCallID() (value int32) {
	return e.GroupCallID
}

// EndGroupCallScreenSharing invokes method endGroupCallScreenSharing#85f41c4c returning error if any.
func (c *Client) EndGroupCallScreenSharing(ctx context.Context, groupcallid int32) error {
	var ok Ok

	request := &EndGroupCallScreenSharingRequest{
		GroupCallID: groupcallid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
