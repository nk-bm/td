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

// InviteGroupCallParticipantsRequest represents TL type `inviteGroupCallParticipants#36d83295`.
type InviteGroupCallParticipantsRequest struct {
	// Group call identifier
	GroupCallID int32
	// User identifiers. At most 10 users can be invited simultaneously
	UserIDs []int32
}

// InviteGroupCallParticipantsRequestTypeID is TL type id of InviteGroupCallParticipantsRequest.
const InviteGroupCallParticipantsRequestTypeID = 0x36d83295

// Ensuring interfaces in compile-time for InviteGroupCallParticipantsRequest.
var (
	_ bin.Encoder     = &InviteGroupCallParticipantsRequest{}
	_ bin.Decoder     = &InviteGroupCallParticipantsRequest{}
	_ bin.BareEncoder = &InviteGroupCallParticipantsRequest{}
	_ bin.BareDecoder = &InviteGroupCallParticipantsRequest{}
)

func (i *InviteGroupCallParticipantsRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.GroupCallID == 0) {
		return false
	}
	if !(i.UserIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InviteGroupCallParticipantsRequest) String() string {
	if i == nil {
		return "InviteGroupCallParticipantsRequest(nil)"
	}
	type Alias InviteGroupCallParticipantsRequest
	return fmt.Sprintf("InviteGroupCallParticipantsRequest%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InviteGroupCallParticipantsRequest) TypeID() uint32 {
	return InviteGroupCallParticipantsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*InviteGroupCallParticipantsRequest) TypeName() string {
	return "inviteGroupCallParticipants"
}

// TypeInfo returns info about TL type.
func (i *InviteGroupCallParticipantsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inviteGroupCallParticipants",
		ID:   InviteGroupCallParticipantsRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
		{
			Name:       "UserIDs",
			SchemaName: "user_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InviteGroupCallParticipantsRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inviteGroupCallParticipants#36d83295 as nil")
	}
	b.PutID(InviteGroupCallParticipantsRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InviteGroupCallParticipantsRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inviteGroupCallParticipants#36d83295 as nil")
	}
	b.PutInt32(i.GroupCallID)
	b.PutInt(len(i.UserIDs))
	for _, v := range i.UserIDs {
		b.PutInt32(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InviteGroupCallParticipantsRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inviteGroupCallParticipants#36d83295 to nil")
	}
	if err := b.ConsumeID(InviteGroupCallParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode inviteGroupCallParticipants#36d83295: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InviteGroupCallParticipantsRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inviteGroupCallParticipants#36d83295 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inviteGroupCallParticipants#36d83295: field group_call_id: %w", err)
		}
		i.GroupCallID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inviteGroupCallParticipants#36d83295: field user_ids: %w", err)
		}

		if headerLen > 0 {
			i.UserIDs = make([]int32, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode inviteGroupCallParticipants#36d83295: field user_ids: %w", err)
			}
			i.UserIDs = append(i.UserIDs, value)
		}
	}
	return nil
}

// GetGroupCallID returns value of GroupCallID field.
func (i *InviteGroupCallParticipantsRequest) GetGroupCallID() (value int32) {
	return i.GroupCallID
}

// GetUserIDs returns value of UserIDs field.
func (i *InviteGroupCallParticipantsRequest) GetUserIDs() (value []int32) {
	return i.UserIDs
}

// InviteGroupCallParticipants invokes method inviteGroupCallParticipants#36d83295 returning error if any.
func (c *Client) InviteGroupCallParticipants(ctx context.Context, request *InviteGroupCallParticipantsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
