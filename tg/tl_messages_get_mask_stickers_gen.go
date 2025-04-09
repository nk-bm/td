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

// MessagesGetMaskStickersRequest represents TL type `messages.getMaskStickers#640f82b8`.
type MessagesGetMaskStickersRequest struct {
	// Hash field of MessagesGetMaskStickersRequest.
	Hash int64
}

// MessagesGetMaskStickersRequestTypeID is TL type id of MessagesGetMaskStickersRequest.
const MessagesGetMaskStickersRequestTypeID = 0x640f82b8

// Ensuring interfaces in compile-time for MessagesGetMaskStickersRequest.
var (
	_ bin.Encoder     = &MessagesGetMaskStickersRequest{}
	_ bin.Decoder     = &MessagesGetMaskStickersRequest{}
	_ bin.BareEncoder = &MessagesGetMaskStickersRequest{}
	_ bin.BareDecoder = &MessagesGetMaskStickersRequest{}
)

func (g *MessagesGetMaskStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetMaskStickersRequest) String() string {
	if g == nil {
		return "MessagesGetMaskStickersRequest(nil)"
	}
	type Alias MessagesGetMaskStickersRequest
	return fmt.Sprintf("MessagesGetMaskStickersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetMaskStickersRequest) TypeID() uint32 {
	return MessagesGetMaskStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetMaskStickersRequest) TypeName() string {
	return "messages.getMaskStickers"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetMaskStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getMaskStickers",
		ID:   MessagesGetMaskStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetMaskStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMaskStickers#640f82b8 as nil")
	}
	b.PutID(MessagesGetMaskStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetMaskStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMaskStickers#640f82b8 as nil")
	}
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetMaskStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMaskStickers#640f82b8 to nil")
	}
	if err := b.ConsumeID(MessagesGetMaskStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getMaskStickers#640f82b8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetMaskStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMaskStickers#640f82b8 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMaskStickers#640f82b8: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *MessagesGetMaskStickersRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetMaskStickers invokes method messages.getMaskStickers#640f82b8 returning error if any.
func (c *Client) MessagesGetMaskStickers(ctx context.Context, hash int64) (MessagesAllStickersClass, error) {
	var result MessagesAllStickersBox

	request := &MessagesGetMaskStickersRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.AllStickers, nil
}
