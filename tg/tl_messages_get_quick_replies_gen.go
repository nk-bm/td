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

// MessagesGetQuickRepliesRequest represents TL type `messages.getQuickReplies#d483f2a8`.
// Fetch basic info about all existing quick reply shortcuts¹.
//
// Links:
//  1. https://core.telegram.org/api/business#quick-reply-shortcuts
//
// See https://core.telegram.org/method/messages.getQuickReplies for reference.
type MessagesGetQuickRepliesRequest struct {
	// Hash for pagination, generated as specified here »¹ (not the usual algorithm used
	// for hash generation.)
	//
	// Links:
	//  1) https://core.telegram.org/api/business#quick-reply-shortcuts
	Hash int64
}

// MessagesGetQuickRepliesRequestTypeID is TL type id of MessagesGetQuickRepliesRequest.
const MessagesGetQuickRepliesRequestTypeID = 0xd483f2a8

// Ensuring interfaces in compile-time for MessagesGetQuickRepliesRequest.
var (
	_ bin.Encoder     = &MessagesGetQuickRepliesRequest{}
	_ bin.Decoder     = &MessagesGetQuickRepliesRequest{}
	_ bin.BareEncoder = &MessagesGetQuickRepliesRequest{}
	_ bin.BareDecoder = &MessagesGetQuickRepliesRequest{}
)

func (g *MessagesGetQuickRepliesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetQuickRepliesRequest) String() string {
	if g == nil {
		return "MessagesGetQuickRepliesRequest(nil)"
	}
	type Alias MessagesGetQuickRepliesRequest
	return fmt.Sprintf("MessagesGetQuickRepliesRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetQuickRepliesRequest from given interface.
func (g *MessagesGetQuickRepliesRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetQuickRepliesRequest) TypeID() uint32 {
	return MessagesGetQuickRepliesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetQuickRepliesRequest) TypeName() string {
	return "messages.getQuickReplies"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetQuickRepliesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getQuickReplies",
		ID:   MessagesGetQuickRepliesRequestTypeID,
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
func (g *MessagesGetQuickRepliesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getQuickReplies#d483f2a8 as nil")
	}
	b.PutID(MessagesGetQuickRepliesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetQuickRepliesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getQuickReplies#d483f2a8 as nil")
	}
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetQuickRepliesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getQuickReplies#d483f2a8 to nil")
	}
	if err := b.ConsumeID(MessagesGetQuickRepliesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getQuickReplies#d483f2a8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetQuickRepliesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getQuickReplies#d483f2a8 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getQuickReplies#d483f2a8: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *MessagesGetQuickRepliesRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetQuickReplies invokes method messages.getQuickReplies#d483f2a8 returning error if any.
// Fetch basic info about all existing quick reply shortcuts¹.
//
// Links:
//  1. https://core.telegram.org/api/business#quick-reply-shortcuts
//
// See https://core.telegram.org/method/messages.getQuickReplies for reference.
func (c *Client) MessagesGetQuickReplies(ctx context.Context, hash int64) (MessagesQuickRepliesClass, error) {
	var result MessagesQuickRepliesBox

	request := &MessagesGetQuickRepliesRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.QuickReplies, nil
}
