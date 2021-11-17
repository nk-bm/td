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

// GetGroupCallStreamSegmentRequest represents TL type `getGroupCallStreamSegment#6bbcb0a7`.
type GetGroupCallStreamSegmentRequest struct {
	// Group call identifier
	GroupCallID int32
	// Point in time when the stream segment begins; Unix timestamp in milliseconds
	TimeOffset int64
	// Segment duration scale; 0-1. Segment's duration is 1000/(2**scale) milliseconds
	Scale int32
}

// GetGroupCallStreamSegmentRequestTypeID is TL type id of GetGroupCallStreamSegmentRequest.
const GetGroupCallStreamSegmentRequestTypeID = 0x6bbcb0a7

// Ensuring interfaces in compile-time for GetGroupCallStreamSegmentRequest.
var (
	_ bin.Encoder     = &GetGroupCallStreamSegmentRequest{}
	_ bin.Decoder     = &GetGroupCallStreamSegmentRequest{}
	_ bin.BareEncoder = &GetGroupCallStreamSegmentRequest{}
	_ bin.BareDecoder = &GetGroupCallStreamSegmentRequest{}
)

func (g *GetGroupCallStreamSegmentRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.GroupCallID == 0) {
		return false
	}
	if !(g.TimeOffset == 0) {
		return false
	}
	if !(g.Scale == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetGroupCallStreamSegmentRequest) String() string {
	if g == nil {
		return "GetGroupCallStreamSegmentRequest(nil)"
	}
	type Alias GetGroupCallStreamSegmentRequest
	return fmt.Sprintf("GetGroupCallStreamSegmentRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetGroupCallStreamSegmentRequest) TypeID() uint32 {
	return GetGroupCallStreamSegmentRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetGroupCallStreamSegmentRequest) TypeName() string {
	return "getGroupCallStreamSegment"
}

// TypeInfo returns info about TL type.
func (g *GetGroupCallStreamSegmentRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getGroupCallStreamSegment",
		ID:   GetGroupCallStreamSegmentRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
		{
			Name:       "TimeOffset",
			SchemaName: "time_offset",
		},
		{
			Name:       "Scale",
			SchemaName: "scale",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetGroupCallStreamSegmentRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCallStreamSegment#6bbcb0a7 as nil")
	}
	b.PutID(GetGroupCallStreamSegmentRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetGroupCallStreamSegmentRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCallStreamSegment#6bbcb0a7 as nil")
	}
	b.PutInt32(g.GroupCallID)
	b.PutLong(g.TimeOffset)
	b.PutInt32(g.Scale)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetGroupCallStreamSegmentRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCallStreamSegment#6bbcb0a7 to nil")
	}
	if err := b.ConsumeID(GetGroupCallStreamSegmentRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getGroupCallStreamSegment#6bbcb0a7: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetGroupCallStreamSegmentRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCallStreamSegment#6bbcb0a7 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getGroupCallStreamSegment#6bbcb0a7: field group_call_id: %w", err)
		}
		g.GroupCallID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getGroupCallStreamSegment#6bbcb0a7: field time_offset: %w", err)
		}
		g.TimeOffset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getGroupCallStreamSegment#6bbcb0a7: field scale: %w", err)
		}
		g.Scale = value
	}
	return nil
}

// GetGroupCallID returns value of GroupCallID field.
func (g *GetGroupCallStreamSegmentRequest) GetGroupCallID() (value int32) {
	return g.GroupCallID
}

// GetTimeOffset returns value of TimeOffset field.
func (g *GetGroupCallStreamSegmentRequest) GetTimeOffset() (value int64) {
	return g.TimeOffset
}

// GetScale returns value of Scale field.
func (g *GetGroupCallStreamSegmentRequest) GetScale() (value int32) {
	return g.Scale
}

// GetGroupCallStreamSegment invokes method getGroupCallStreamSegment#6bbcb0a7 returning error if any.
func (c *Client) GetGroupCallStreamSegment(ctx context.Context, request *GetGroupCallStreamSegmentRequest) (*FilePart, error) {
	var result FilePart

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
