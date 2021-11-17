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

// GetStickerSetRequest represents TL type `getStickerSet#3eb91bc3`.
type GetStickerSetRequest struct {
	// Identifier of the sticker set
	SetID Int64
}

// GetStickerSetRequestTypeID is TL type id of GetStickerSetRequest.
const GetStickerSetRequestTypeID = 0x3eb91bc3

// Ensuring interfaces in compile-time for GetStickerSetRequest.
var (
	_ bin.Encoder     = &GetStickerSetRequest{}
	_ bin.Decoder     = &GetStickerSetRequest{}
	_ bin.BareEncoder = &GetStickerSetRequest{}
	_ bin.BareDecoder = &GetStickerSetRequest{}
)

func (g *GetStickerSetRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.SetID.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetStickerSetRequest) String() string {
	if g == nil {
		return "GetStickerSetRequest(nil)"
	}
	type Alias GetStickerSetRequest
	return fmt.Sprintf("GetStickerSetRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetStickerSetRequest) TypeID() uint32 {
	return GetStickerSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetStickerSetRequest) TypeName() string {
	return "getStickerSet"
}

// TypeInfo returns info about TL type.
func (g *GetStickerSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getStickerSet",
		ID:   GetStickerSetRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SetID",
			SchemaName: "set_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetStickerSetRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStickerSet#3eb91bc3 as nil")
	}
	b.PutID(GetStickerSetRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetStickerSetRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStickerSet#3eb91bc3 as nil")
	}
	if err := g.SetID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getStickerSet#3eb91bc3: field set_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetStickerSetRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStickerSet#3eb91bc3 to nil")
	}
	if err := b.ConsumeID(GetStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getStickerSet#3eb91bc3: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetStickerSetRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStickerSet#3eb91bc3 to nil")
	}
	{
		if err := g.SetID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode getStickerSet#3eb91bc3: field set_id: %w", err)
		}
	}
	return nil
}

// GetSetID returns value of SetID field.
func (g *GetStickerSetRequest) GetSetID() (value Int64) {
	return g.SetID
}

// GetStickerSet invokes method getStickerSet#3eb91bc3 returning error if any.
func (c *Client) GetStickerSet(ctx context.Context, setid Int64) (*StickerSet, error) {
	var result StickerSet

	request := &GetStickerSetRequest{
		SetID: setid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
