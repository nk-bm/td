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

// BotsGetPopularAppBotsRequest represents TL type `bots.getPopularAppBots#c2510192`.
type BotsGetPopularAppBotsRequest struct {
	// Offset field of BotsGetPopularAppBotsRequest.
	Offset string
	// Limit field of BotsGetPopularAppBotsRequest.
	Limit int
}

// BotsGetPopularAppBotsRequestTypeID is TL type id of BotsGetPopularAppBotsRequest.
const BotsGetPopularAppBotsRequestTypeID = 0xc2510192

// Ensuring interfaces in compile-time for BotsGetPopularAppBotsRequest.
var (
	_ bin.Encoder     = &BotsGetPopularAppBotsRequest{}
	_ bin.Decoder     = &BotsGetPopularAppBotsRequest{}
	_ bin.BareEncoder = &BotsGetPopularAppBotsRequest{}
	_ bin.BareDecoder = &BotsGetPopularAppBotsRequest{}
)

func (g *BotsGetPopularAppBotsRequest) Zero() bool {
	if g == nil {
		return true
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
func (g *BotsGetPopularAppBotsRequest) String() string {
	if g == nil {
		return "BotsGetPopularAppBotsRequest(nil)"
	}
	type Alias BotsGetPopularAppBotsRequest
	return fmt.Sprintf("BotsGetPopularAppBotsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsGetPopularAppBotsRequest) TypeID() uint32 {
	return BotsGetPopularAppBotsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsGetPopularAppBotsRequest) TypeName() string {
	return "bots.getPopularAppBots"
}

// TypeInfo returns info about TL type.
func (g *BotsGetPopularAppBotsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.getPopularAppBots",
		ID:   BotsGetPopularAppBotsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
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
func (g *BotsGetPopularAppBotsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode bots.getPopularAppBots#c2510192 as nil")
	}
	b.PutID(BotsGetPopularAppBotsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *BotsGetPopularAppBotsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode bots.getPopularAppBots#c2510192 as nil")
	}
	b.PutString(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *BotsGetPopularAppBotsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode bots.getPopularAppBots#c2510192 to nil")
	}
	if err := b.ConsumeID(BotsGetPopularAppBotsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.getPopularAppBots#c2510192: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *BotsGetPopularAppBotsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode bots.getPopularAppBots#c2510192 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.getPopularAppBots#c2510192: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode bots.getPopularAppBots#c2510192: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetOffset returns value of Offset field.
func (g *BotsGetPopularAppBotsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *BotsGetPopularAppBotsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// BotsGetPopularAppBots invokes method bots.getPopularAppBots#c2510192 returning error if any.
func (c *Client) BotsGetPopularAppBots(ctx context.Context, request *BotsGetPopularAppBotsRequest) (*BotsPopularAppBots, error) {
	var result BotsPopularAppBots

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
