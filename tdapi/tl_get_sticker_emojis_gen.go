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

// GetStickerEmojisRequest represents TL type `getStickerEmojis#8f04d547`.
type GetStickerEmojisRequest struct {
	// Sticker file identifier
	Sticker InputFileClass
}

// GetStickerEmojisRequestTypeID is TL type id of GetStickerEmojisRequest.
const GetStickerEmojisRequestTypeID = 0x8f04d547

// Ensuring interfaces in compile-time for GetStickerEmojisRequest.
var (
	_ bin.Encoder     = &GetStickerEmojisRequest{}
	_ bin.Decoder     = &GetStickerEmojisRequest{}
	_ bin.BareEncoder = &GetStickerEmojisRequest{}
	_ bin.BareDecoder = &GetStickerEmojisRequest{}
)

func (g *GetStickerEmojisRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Sticker == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetStickerEmojisRequest) String() string {
	if g == nil {
		return "GetStickerEmojisRequest(nil)"
	}
	type Alias GetStickerEmojisRequest
	return fmt.Sprintf("GetStickerEmojisRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetStickerEmojisRequest) TypeID() uint32 {
	return GetStickerEmojisRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetStickerEmojisRequest) TypeName() string {
	return "getStickerEmojis"
}

// TypeInfo returns info about TL type.
func (g *GetStickerEmojisRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getStickerEmojis",
		ID:   GetStickerEmojisRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetStickerEmojisRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStickerEmojis#8f04d547 as nil")
	}
	b.PutID(GetStickerEmojisRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetStickerEmojisRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStickerEmojis#8f04d547 as nil")
	}
	if g.Sticker == nil {
		return fmt.Errorf("unable to encode getStickerEmojis#8f04d547: field sticker is nil")
	}
	if err := g.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getStickerEmojis#8f04d547: field sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetStickerEmojisRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStickerEmojis#8f04d547 to nil")
	}
	if err := b.ConsumeID(GetStickerEmojisRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getStickerEmojis#8f04d547: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetStickerEmojisRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStickerEmojis#8f04d547 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode getStickerEmojis#8f04d547: field sticker: %w", err)
		}
		g.Sticker = value
	}
	return nil
}

// GetSticker returns value of Sticker field.
func (g *GetStickerEmojisRequest) GetSticker() (value InputFileClass) {
	return g.Sticker
}

// GetStickerEmojis invokes method getStickerEmojis#8f04d547 returning error if any.
func (c *Client) GetStickerEmojis(ctx context.Context, sticker InputFileClass) (*Emojis, error) {
	var result Emojis

	request := &GetStickerEmojisRequest{
		Sticker: sticker,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
