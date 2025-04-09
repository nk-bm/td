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

// PaymentsGetStarsRevenueStatsRequest represents TL type `payments.getStarsRevenueStats#d91ffad6`.
type PaymentsGetStarsRevenueStatsRequest struct {
	// Flags field of PaymentsGetStarsRevenueStatsRequest.
	Flags bin.Fields
	// Dark field of PaymentsGetStarsRevenueStatsRequest.
	Dark bool
	// Peer field of PaymentsGetStarsRevenueStatsRequest.
	Peer InputPeerClass
}

// PaymentsGetStarsRevenueStatsRequestTypeID is TL type id of PaymentsGetStarsRevenueStatsRequest.
const PaymentsGetStarsRevenueStatsRequestTypeID = 0xd91ffad6

// Ensuring interfaces in compile-time for PaymentsGetStarsRevenueStatsRequest.
var (
	_ bin.Encoder     = &PaymentsGetStarsRevenueStatsRequest{}
	_ bin.Decoder     = &PaymentsGetStarsRevenueStatsRequest{}
	_ bin.BareEncoder = &PaymentsGetStarsRevenueStatsRequest{}
	_ bin.BareDecoder = &PaymentsGetStarsRevenueStatsRequest{}
)

func (g *PaymentsGetStarsRevenueStatsRequest) Zero() bool {
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
func (g *PaymentsGetStarsRevenueStatsRequest) String() string {
	if g == nil {
		return "PaymentsGetStarsRevenueStatsRequest(nil)"
	}
	type Alias PaymentsGetStarsRevenueStatsRequest
	return fmt.Sprintf("PaymentsGetStarsRevenueStatsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsGetStarsRevenueStatsRequest) TypeID() uint32 {
	return PaymentsGetStarsRevenueStatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsGetStarsRevenueStatsRequest) TypeName() string {
	return "payments.getStarsRevenueStats"
}

// TypeInfo returns info about TL type.
func (g *PaymentsGetStarsRevenueStatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.getStarsRevenueStats",
		ID:   PaymentsGetStarsRevenueStatsRequestTypeID,
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
func (g *PaymentsGetStarsRevenueStatsRequest) SetFlags() {
	if !(g.Dark == false) {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *PaymentsGetStarsRevenueStatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getStarsRevenueStats#d91ffad6 as nil")
	}
	b.PutID(PaymentsGetStarsRevenueStatsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PaymentsGetStarsRevenueStatsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getStarsRevenueStats#d91ffad6 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.getStarsRevenueStats#d91ffad6: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode payments.getStarsRevenueStats#d91ffad6: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.getStarsRevenueStats#d91ffad6: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetStarsRevenueStatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getStarsRevenueStats#d91ffad6 to nil")
	}
	if err := b.ConsumeID(PaymentsGetStarsRevenueStatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getStarsRevenueStats#d91ffad6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PaymentsGetStarsRevenueStatsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getStarsRevenueStats#d91ffad6 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.getStarsRevenueStats#d91ffad6: field flags: %w", err)
		}
	}
	g.Dark = g.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.getStarsRevenueStats#d91ffad6: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// SetDark sets value of Dark conditional field.
func (g *PaymentsGetStarsRevenueStatsRequest) SetDark(value bool) {
	if value {
		g.Flags.Set(0)
		g.Dark = true
	} else {
		g.Flags.Unset(0)
		g.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (g *PaymentsGetStarsRevenueStatsRequest) GetDark() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (g *PaymentsGetStarsRevenueStatsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// PaymentsGetStarsRevenueStats invokes method payments.getStarsRevenueStats#d91ffad6 returning error if any.
func (c *Client) PaymentsGetStarsRevenueStats(ctx context.Context, request *PaymentsGetStarsRevenueStatsRequest) (*PaymentsStarsRevenueStats, error) {
	var result PaymentsStarsRevenueStats

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
