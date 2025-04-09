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

// StatsGetBroadcastRevenueStatsRequest represents TL type `stats.getBroadcastRevenueStats#f788ee19`.
type StatsGetBroadcastRevenueStatsRequest struct {
	// Flags field of StatsGetBroadcastRevenueStatsRequest.
	Flags bin.Fields
	// Dark field of StatsGetBroadcastRevenueStatsRequest.
	Dark bool
	// Peer field of StatsGetBroadcastRevenueStatsRequest.
	Peer InputPeerClass
}

// StatsGetBroadcastRevenueStatsRequestTypeID is TL type id of StatsGetBroadcastRevenueStatsRequest.
const StatsGetBroadcastRevenueStatsRequestTypeID = 0xf788ee19

// Ensuring interfaces in compile-time for StatsGetBroadcastRevenueStatsRequest.
var (
	_ bin.Encoder     = &StatsGetBroadcastRevenueStatsRequest{}
	_ bin.Decoder     = &StatsGetBroadcastRevenueStatsRequest{}
	_ bin.BareEncoder = &StatsGetBroadcastRevenueStatsRequest{}
	_ bin.BareDecoder = &StatsGetBroadcastRevenueStatsRequest{}
)

func (g *StatsGetBroadcastRevenueStatsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Dark == false) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StatsGetBroadcastRevenueStatsRequest) String() string {
	if g == nil {
		return "StatsGetBroadcastRevenueStatsRequest(nil)"
	}
	type Alias StatsGetBroadcastRevenueStatsRequest
	return fmt.Sprintf("StatsGetBroadcastRevenueStatsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGetBroadcastRevenueStatsRequest) TypeID() uint32 {
	return StatsGetBroadcastRevenueStatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGetBroadcastRevenueStatsRequest) TypeName() string {
	return "stats.getBroadcastRevenueStats"
}

// TypeInfo returns info about TL type.
func (g *StatsGetBroadcastRevenueStatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.getBroadcastRevenueStats",
		ID:   StatsGetBroadcastRevenueStatsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Dark",
			SchemaName: "dark",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *StatsGetBroadcastRevenueStatsRequest) SetFlags() {
	if !(g.Dark == false) {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *StatsGetBroadcastRevenueStatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getBroadcastRevenueStats#f788ee19 as nil")
	}
	b.PutID(StatsGetBroadcastRevenueStatsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StatsGetBroadcastRevenueStatsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getBroadcastRevenueStats#f788ee19 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getBroadcastRevenueStats#f788ee19: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode stats.getBroadcastRevenueStats#f788ee19: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getBroadcastRevenueStats#f788ee19: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *StatsGetBroadcastRevenueStatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getBroadcastRevenueStats#f788ee19 to nil")
	}
	if err := b.ConsumeID(StatsGetBroadcastRevenueStatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.getBroadcastRevenueStats#f788ee19: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StatsGetBroadcastRevenueStatsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getBroadcastRevenueStats#f788ee19 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.getBroadcastRevenueStats#f788ee19: field flags: %w", err)
		}
	}
	g.Dark = g.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getBroadcastRevenueStats#f788ee19: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// SetDark sets value of Dark conditional field.
func (g *StatsGetBroadcastRevenueStatsRequest) SetDark(value bool) {
	if value {
		g.Flags.Set(0)
		g.Dark = true
	} else {
		g.Flags.Unset(0)
		g.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (g *StatsGetBroadcastRevenueStatsRequest) GetDark() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (g *StatsGetBroadcastRevenueStatsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// StatsGetBroadcastRevenueStats invokes method stats.getBroadcastRevenueStats#f788ee19 returning error if any.
func (c *Client) StatsGetBroadcastRevenueStats(ctx context.Context, request *StatsGetBroadcastRevenueStatsRequest) (*StatsBroadcastRevenueStats, error) {
	var result StatsBroadcastRevenueStats

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
