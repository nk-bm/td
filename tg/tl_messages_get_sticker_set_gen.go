// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// MessagesGetStickerSetRequest represents TL type `messages.getStickerSet#2619a90e`.
// Get info about a stickerset
//
// See https://core.telegram.org/method/messages.getStickerSet for reference.
type MessagesGetStickerSetRequest struct {
	// Stickerset
	Stickerset InputStickerSetClass
}

// MessagesGetStickerSetRequestTypeID is TL type id of MessagesGetStickerSetRequest.
const MessagesGetStickerSetRequestTypeID = 0x2619a90e

func (g *MessagesGetStickerSetRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Stickerset == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetStickerSetRequest) String() string {
	if g == nil {
		return "MessagesGetStickerSetRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetStickerSetRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tStickerset: ")
	sb.WriteString(fmt.Sprint(g.Stickerset))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetStickerSetRequest) TypeID() uint32 {
	return MessagesGetStickerSetRequestTypeID
}

// Encode implements bin.Encoder.
func (g *MessagesGetStickerSetRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getStickerSet#2619a90e as nil")
	}
	b.PutID(MessagesGetStickerSetRequestTypeID)
	if g.Stickerset == nil {
		return fmt.Errorf("unable to encode messages.getStickerSet#2619a90e: field stickerset is nil")
	}
	if err := g.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getStickerSet#2619a90e: field stickerset: %w", err)
	}
	return nil
}

// GetStickerset returns value of Stickerset field.
func (g *MessagesGetStickerSetRequest) GetStickerset() (value InputStickerSetClass) {
	return g.Stickerset
}

// Decode implements bin.Decoder.
func (g *MessagesGetStickerSetRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getStickerSet#2619a90e to nil")
	}
	if err := b.ConsumeID(MessagesGetStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getStickerSet#2619a90e: %w", err)
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getStickerSet#2619a90e: field stickerset: %w", err)
		}
		g.Stickerset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetStickerSetRequest.
var (
	_ bin.Encoder = &MessagesGetStickerSetRequest{}
	_ bin.Decoder = &MessagesGetStickerSetRequest{}
)

// MessagesGetStickerSet invokes method messages.getStickerSet#2619a90e returning error if any.
// Get info about a stickerset
//
// Possible errors:
//  400 STICKERSET_INVALID: The provided sticker set is invalid
//
// See https://core.telegram.org/method/messages.getStickerSet for reference.
// Can be used by bots.
func (c *Client) MessagesGetStickerSet(ctx context.Context, stickerset InputStickerSetClass) (*MessagesStickerSet, error) {
	var result MessagesStickerSet

	request := &MessagesGetStickerSetRequest{
		Stickerset: stickerset,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
