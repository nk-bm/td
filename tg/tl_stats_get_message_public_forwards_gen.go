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

// StatsGetMessagePublicForwardsRequest represents TL type `stats.getMessagePublicForwards#5f150144`.
type StatsGetMessagePublicForwardsRequest struct {
	// Channel field of StatsGetMessagePublicForwardsRequest.
	Channel InputChannelClass
	// MsgID field of StatsGetMessagePublicForwardsRequest.
	MsgID int
	// Offset field of StatsGetMessagePublicForwardsRequest.
	Offset string
	// Limit field of StatsGetMessagePublicForwardsRequest.
	Limit int
}

// StatsGetMessagePublicForwardsRequestTypeID is TL type id of StatsGetMessagePublicForwardsRequest.
const StatsGetMessagePublicForwardsRequestTypeID = 0x5f150144

// Ensuring interfaces in compile-time for StatsGetMessagePublicForwardsRequest.
var (
	_ bin.Encoder     = &StatsGetMessagePublicForwardsRequest{}
	_ bin.Decoder     = &StatsGetMessagePublicForwardsRequest{}
	_ bin.BareEncoder = &StatsGetMessagePublicForwardsRequest{}
	_ bin.BareDecoder = &StatsGetMessagePublicForwardsRequest{}
)

func (g *StatsGetMessagePublicForwardsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Channel == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StatsGetMessagePublicForwardsRequest) String() string {
	if g == nil {
		return "StatsGetMessagePublicForwardsRequest(nil)"
	}
	type Alias StatsGetMessagePublicForwardsRequest
	return fmt.Sprintf("StatsGetMessagePublicForwardsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGetMessagePublicForwardsRequest) TypeID() uint32 {
	return StatsGetMessagePublicForwardsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGetMessagePublicForwardsRequest) TypeName() string {
	return "stats.getMessagePublicForwards"
}

// TypeInfo returns info about TL type.
func (g *StatsGetMessagePublicForwardsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.getMessagePublicForwards",
		ID:   StatsGetMessagePublicForwardsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *StatsGetMessagePublicForwardsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getMessagePublicForwards#5f150144 as nil")
	}
	b.PutID(StatsGetMessagePublicForwardsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StatsGetMessagePublicForwardsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getMessagePublicForwards#5f150144 as nil")
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode stats.getMessagePublicForwards#5f150144: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getMessagePublicForwards#5f150144: field channel: %w", err)
	}
	b.PutInt(g.MsgID)
	b.PutString(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *StatsGetMessagePublicForwardsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getMessagePublicForwards#5f150144 to nil")
	}
	if err := b.ConsumeID(StatsGetMessagePublicForwardsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5f150144: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StatsGetMessagePublicForwardsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getMessagePublicForwards#5f150144 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5f150144: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5f150144: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5f150144: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5f150144: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (g *StatsGetMessagePublicForwardsRequest) GetChannel() (value InputChannelClass) {
	if g == nil {
		return
	}
	return g.Channel
}

// GetMsgID returns value of MsgID field.
func (g *StatsGetMessagePublicForwardsRequest) GetMsgID() (value int) {
	if g == nil {
		return
	}
	return g.MsgID
}

// GetOffset returns value of Offset field.
func (g *StatsGetMessagePublicForwardsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *StatsGetMessagePublicForwardsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// StatsGetMessagePublicForwards invokes method stats.getMessagePublicForwards#5f150144 returning error if any.
func (c *Client) StatsGetMessagePublicForwards(ctx context.Context, request *StatsGetMessagePublicForwardsRequest) (*StatsPublicForwards, error) {
	var result StatsPublicForwards

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
