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

// MessagesGetPinnedSavedDialogsRequest represents TL type `messages.getPinnedSavedDialogs#d63d94e0`.
type MessagesGetPinnedSavedDialogsRequest struct {
}

// MessagesGetPinnedSavedDialogsRequestTypeID is TL type id of MessagesGetPinnedSavedDialogsRequest.
const MessagesGetPinnedSavedDialogsRequestTypeID = 0xd63d94e0

// Ensuring interfaces in compile-time for MessagesGetPinnedSavedDialogsRequest.
var (
	_ bin.Encoder     = &MessagesGetPinnedSavedDialogsRequest{}
	_ bin.Decoder     = &MessagesGetPinnedSavedDialogsRequest{}
	_ bin.BareEncoder = &MessagesGetPinnedSavedDialogsRequest{}
	_ bin.BareDecoder = &MessagesGetPinnedSavedDialogsRequest{}
)

func (g *MessagesGetPinnedSavedDialogsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetPinnedSavedDialogsRequest) String() string {
	if g == nil {
		return "MessagesGetPinnedSavedDialogsRequest(nil)"
	}
	type Alias MessagesGetPinnedSavedDialogsRequest
	return fmt.Sprintf("MessagesGetPinnedSavedDialogsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetPinnedSavedDialogsRequest) TypeID() uint32 {
	return MessagesGetPinnedSavedDialogsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetPinnedSavedDialogsRequest) TypeName() string {
	return "messages.getPinnedSavedDialogs"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetPinnedSavedDialogsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getPinnedSavedDialogs",
		ID:   MessagesGetPinnedSavedDialogsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetPinnedSavedDialogsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPinnedSavedDialogs#d63d94e0 as nil")
	}
	b.PutID(MessagesGetPinnedSavedDialogsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetPinnedSavedDialogsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPinnedSavedDialogs#d63d94e0 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetPinnedSavedDialogsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPinnedSavedDialogs#d63d94e0 to nil")
	}
	if err := b.ConsumeID(MessagesGetPinnedSavedDialogsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getPinnedSavedDialogs#d63d94e0: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetPinnedSavedDialogsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPinnedSavedDialogs#d63d94e0 to nil")
	}
	return nil
}

// MessagesGetPinnedSavedDialogs invokes method messages.getPinnedSavedDialogs#d63d94e0 returning error if any.
func (c *Client) MessagesGetPinnedSavedDialogs(ctx context.Context) (MessagesSavedDialogsClass, error) {
	var result MessagesSavedDialogsBox

	request := &MessagesGetPinnedSavedDialogsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.SavedDialogs, nil
}
