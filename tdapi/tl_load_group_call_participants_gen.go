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

// LoadGroupCallParticipantsRequest represents TL type `loadGroupCallParticipants#37f3bece`.
type LoadGroupCallParticipantsRequest struct {
	// Group call identifier. The group call must be previously received through getGroupCall
	// and must be joined or being joined
	GroupCallID int32
	// The maximum number of participants to load
	Limit int32
}

// LoadGroupCallParticipantsRequestTypeID is TL type id of LoadGroupCallParticipantsRequest.
const LoadGroupCallParticipantsRequestTypeID = 0x37f3bece

// Ensuring interfaces in compile-time for LoadGroupCallParticipantsRequest.
var (
	_ bin.Encoder     = &LoadGroupCallParticipantsRequest{}
	_ bin.Decoder     = &LoadGroupCallParticipantsRequest{}
	_ bin.BareEncoder = &LoadGroupCallParticipantsRequest{}
	_ bin.BareDecoder = &LoadGroupCallParticipantsRequest{}
)

func (l *LoadGroupCallParticipantsRequest) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.GroupCallID == 0) {
		return false
	}
	if !(l.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LoadGroupCallParticipantsRequest) String() string {
	if l == nil {
		return "LoadGroupCallParticipantsRequest(nil)"
	}
	type Alias LoadGroupCallParticipantsRequest
	return fmt.Sprintf("LoadGroupCallParticipantsRequest%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LoadGroupCallParticipantsRequest) TypeID() uint32 {
	return LoadGroupCallParticipantsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*LoadGroupCallParticipantsRequest) TypeName() string {
	return "loadGroupCallParticipants"
}

// TypeInfo returns info about TL type.
func (l *LoadGroupCallParticipantsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "loadGroupCallParticipants",
		ID:   LoadGroupCallParticipantsRequestTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LoadGroupCallParticipantsRequest) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode loadGroupCallParticipants#37f3bece as nil")
	}
	b.PutID(LoadGroupCallParticipantsRequestTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LoadGroupCallParticipantsRequest) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode loadGroupCallParticipants#37f3bece as nil")
	}
	b.PutInt32(l.GroupCallID)
	b.PutInt32(l.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (l *LoadGroupCallParticipantsRequest) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode loadGroupCallParticipants#37f3bece to nil")
	}
	if err := b.ConsumeID(LoadGroupCallParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode loadGroupCallParticipants#37f3bece: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LoadGroupCallParticipantsRequest) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode loadGroupCallParticipants#37f3bece to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode loadGroupCallParticipants#37f3bece: field group_call_id: %w", err)
		}
		l.GroupCallID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode loadGroupCallParticipants#37f3bece: field limit: %w", err)
		}
		l.Limit = value
	}
	return nil
}

// GetGroupCallID returns value of GroupCallID field.
func (l *LoadGroupCallParticipantsRequest) GetGroupCallID() (value int32) {
	return l.GroupCallID
}

// GetLimit returns value of Limit field.
func (l *LoadGroupCallParticipantsRequest) GetLimit() (value int32) {
	return l.Limit
}

// LoadGroupCallParticipants invokes method loadGroupCallParticipants#37f3bece returning error if any.
func (c *Client) LoadGroupCallParticipants(ctx context.Context, request *LoadGroupCallParticipantsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
