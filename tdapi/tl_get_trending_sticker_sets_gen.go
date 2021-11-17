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

// GetTrendingStickerSetsRequest represents TL type `getTrendingStickerSets#a6ea7d44`.
type GetTrendingStickerSetsRequest struct {
	// The offset from which to return the sticker sets; must be non-negative
	Offset int32
	// The maximum number of sticker sets to be returned; must be non-negative. For optimal
	// performance, the number of returned sticker sets is chosen by TDLib and can be smaller
	// than the specified limit, even if the end of the list has not been reached
	Limit int32
}

// GetTrendingStickerSetsRequestTypeID is TL type id of GetTrendingStickerSetsRequest.
const GetTrendingStickerSetsRequestTypeID = 0xa6ea7d44

// Ensuring interfaces in compile-time for GetTrendingStickerSetsRequest.
var (
	_ bin.Encoder     = &GetTrendingStickerSetsRequest{}
	_ bin.Decoder     = &GetTrendingStickerSetsRequest{}
	_ bin.BareEncoder = &GetTrendingStickerSetsRequest{}
	_ bin.BareDecoder = &GetTrendingStickerSetsRequest{}
)

func (g *GetTrendingStickerSetsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetTrendingStickerSetsRequest) String() string {
	if g == nil {
		return "GetTrendingStickerSetsRequest(nil)"
	}
	type Alias GetTrendingStickerSetsRequest
	return fmt.Sprintf("GetTrendingStickerSetsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetTrendingStickerSetsRequest) TypeID() uint32 {
	return GetTrendingStickerSetsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetTrendingStickerSetsRequest) TypeName() string {
	return "getTrendingStickerSets"
}

// TypeInfo returns info about TL type.
func (g *GetTrendingStickerSetsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getTrendingStickerSets",
		ID:   GetTrendingStickerSetsRequestTypeID,
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
func (g *GetTrendingStickerSetsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getTrendingStickerSets#a6ea7d44 as nil")
	}
	b.PutID(GetTrendingStickerSetsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetTrendingStickerSetsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getTrendingStickerSets#a6ea7d44 as nil")
	}
	b.PutInt32(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetTrendingStickerSetsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getTrendingStickerSets#a6ea7d44 to nil")
	}
	if err := b.ConsumeID(GetTrendingStickerSetsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getTrendingStickerSets#a6ea7d44: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetTrendingStickerSetsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getTrendingStickerSets#a6ea7d44 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getTrendingStickerSets#a6ea7d44: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getTrendingStickerSets#a6ea7d44: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetOffset returns value of Offset field.
func (g *GetTrendingStickerSetsRequest) GetOffset() (value int32) {
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetTrendingStickerSetsRequest) GetLimit() (value int32) {
	return g.Limit
}

// GetTrendingStickerSets invokes method getTrendingStickerSets#a6ea7d44 returning error if any.
func (c *Client) GetTrendingStickerSets(ctx context.Context, request *GetTrendingStickerSetsRequest) (*StickerSets, error) {
	var result StickerSets

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
