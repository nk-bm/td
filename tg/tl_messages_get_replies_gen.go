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

// MessagesGetRepliesRequest represents TL type `messages.getReplies#22ddd30c`.
type MessagesGetRepliesRequest struct {
	// Peer field of MessagesGetRepliesRequest.
	Peer InputPeerClass
	// MsgID field of MessagesGetRepliesRequest.
	MsgID int
	// OffsetID field of MessagesGetRepliesRequest.
	OffsetID int
	// OffsetDate field of MessagesGetRepliesRequest.
	OffsetDate int
	// AddOffset field of MessagesGetRepliesRequest.
	AddOffset int
	// Limit field of MessagesGetRepliesRequest.
	Limit int
	// MaxID field of MessagesGetRepliesRequest.
	MaxID int
	// MinID field of MessagesGetRepliesRequest.
	MinID int
	// Hash field of MessagesGetRepliesRequest.
	Hash int64
}

// MessagesGetRepliesRequestTypeID is TL type id of MessagesGetRepliesRequest.
const MessagesGetRepliesRequestTypeID = 0x22ddd30c

// Ensuring interfaces in compile-time for MessagesGetRepliesRequest.
var (
	_ bin.Encoder     = &MessagesGetRepliesRequest{}
	_ bin.Decoder     = &MessagesGetRepliesRequest{}
	_ bin.BareEncoder = &MessagesGetRepliesRequest{}
	_ bin.BareDecoder = &MessagesGetRepliesRequest{}
)

func (g *MessagesGetRepliesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.OffsetDate == 0) {
		return false
	}
	if !(g.AddOffset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.MaxID == 0) {
		return false
	}
	if !(g.MinID == 0) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetRepliesRequest) String() string {
	if g == nil {
		return "MessagesGetRepliesRequest(nil)"
	}
	type Alias MessagesGetRepliesRequest
	return fmt.Sprintf("MessagesGetRepliesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetRepliesRequest) TypeID() uint32 {
	return MessagesGetRepliesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetRepliesRequest) TypeName() string {
	return "messages.getReplies"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetRepliesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getReplies",
		ID:   MessagesGetRepliesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "OffsetDate",
			SchemaName: "offset_date",
		},
		{
			Name:       "AddOffset",
			SchemaName: "add_offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
		{
			Name:       "MinID",
			SchemaName: "min_id",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetRepliesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getReplies#22ddd30c as nil")
	}
	b.PutID(MessagesGetRepliesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetRepliesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getReplies#22ddd30c as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getReplies#22ddd30c: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getReplies#22ddd30c: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	b.PutInt(g.OffsetID)
	b.PutInt(g.OffsetDate)
	b.PutInt(g.AddOffset)
	b.PutInt(g.Limit)
	b.PutInt(g.MaxID)
	b.PutInt(g.MinID)
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetRepliesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getReplies#22ddd30c to nil")
	}
	if err := b.ConsumeID(MessagesGetRepliesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getReplies#22ddd30c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetRepliesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getReplies#22ddd30c to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#22ddd30c: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#22ddd30c: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#22ddd30c: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#22ddd30c: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#22ddd30c: field add_offset: %w", err)
		}
		g.AddOffset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#22ddd30c: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#22ddd30c: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#22ddd30c: field min_id: %w", err)
		}
		g.MinID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#22ddd30c: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetRepliesRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetMsgID returns value of MsgID field.
func (g *MessagesGetRepliesRequest) GetMsgID() (value int) {
	if g == nil {
		return
	}
	return g.MsgID
}

// GetOffsetID returns value of OffsetID field.
func (g *MessagesGetRepliesRequest) GetOffsetID() (value int) {
	if g == nil {
		return
	}
	return g.OffsetID
}

// GetOffsetDate returns value of OffsetDate field.
func (g *MessagesGetRepliesRequest) GetOffsetDate() (value int) {
	if g == nil {
		return
	}
	return g.OffsetDate
}

// GetAddOffset returns value of AddOffset field.
func (g *MessagesGetRepliesRequest) GetAddOffset() (value int) {
	if g == nil {
		return
	}
	return g.AddOffset
}

// GetLimit returns value of Limit field.
func (g *MessagesGetRepliesRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetMaxID returns value of MaxID field.
func (g *MessagesGetRepliesRequest) GetMaxID() (value int) {
	if g == nil {
		return
	}
	return g.MaxID
}

// GetMinID returns value of MinID field.
func (g *MessagesGetRepliesRequest) GetMinID() (value int) {
	if g == nil {
		return
	}
	return g.MinID
}

// GetHash returns value of Hash field.
func (g *MessagesGetRepliesRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetReplies invokes method messages.getReplies#22ddd30c returning error if any.
func (c *Client) MessagesGetReplies(ctx context.Context, request *MessagesGetRepliesRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
