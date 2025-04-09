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

// StoriesGetStoryReactionsListRequest represents TL type `stories.getStoryReactionsList#b9b2881f`.
type StoriesGetStoryReactionsListRequest struct {
	// Flags field of StoriesGetStoryReactionsListRequest.
	Flags bin.Fields
	// ForwardsFirst field of StoriesGetStoryReactionsListRequest.
	ForwardsFirst bool
	// Peer field of StoriesGetStoryReactionsListRequest.
	Peer InputPeerClass
	// ID field of StoriesGetStoryReactionsListRequest.
	ID int
	// Reaction field of StoriesGetStoryReactionsListRequest.
	//
	// Use SetReaction and GetReaction helpers.
	Reaction ReactionClass
	// Offset field of StoriesGetStoryReactionsListRequest.
	//
	// Use SetOffset and GetOffset helpers.
	Offset string
	// Limit field of StoriesGetStoryReactionsListRequest.
	Limit int
}

// StoriesGetStoryReactionsListRequestTypeID is TL type id of StoriesGetStoryReactionsListRequest.
const StoriesGetStoryReactionsListRequestTypeID = 0xb9b2881f

// Ensuring interfaces in compile-time for StoriesGetStoryReactionsListRequest.
var (
	_ bin.Encoder     = &StoriesGetStoryReactionsListRequest{}
	_ bin.Decoder     = &StoriesGetStoryReactionsListRequest{}
	_ bin.BareEncoder = &StoriesGetStoryReactionsListRequest{}
	_ bin.BareDecoder = &StoriesGetStoryReactionsListRequest{}
)

func (g *StoriesGetStoryReactionsListRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.ForwardsFirst == false) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.ID == 0) {
		return false
	}
	if !(g.Reaction == nil) {
		return false
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
func (g *StoriesGetStoryReactionsListRequest) String() string {
	if g == nil {
		return "StoriesGetStoryReactionsListRequest(nil)"
	}
	type Alias StoriesGetStoryReactionsListRequest
	return fmt.Sprintf("StoriesGetStoryReactionsListRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesGetStoryReactionsListRequest) TypeID() uint32 {
	return StoriesGetStoryReactionsListRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesGetStoryReactionsListRequest) TypeName() string {
	return "stories.getStoryReactionsList"
}

// TypeInfo returns info about TL type.
func (g *StoriesGetStoryReactionsListRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.getStoryReactionsList",
		ID:   StoriesGetStoryReactionsListRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ForwardsFirst",
			SchemaName: "forwards_first",
			Null:       !g.Flags.Has(2),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Reaction",
			SchemaName: "reaction",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *StoriesGetStoryReactionsListRequest) SetFlags() {
	if !(g.ForwardsFirst == false) {
		g.Flags.Set(2)
	}
	if !(g.Reaction == nil) {
		g.Flags.Set(0)
	}
	if !(g.Offset == "") {
		g.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (g *StoriesGetStoryReactionsListRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getStoryReactionsList#b9b2881f as nil")
	}
	b.PutID(StoriesGetStoryReactionsListRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StoriesGetStoryReactionsListRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getStoryReactionsList#b9b2881f as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.getStoryReactionsList#b9b2881f: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode stories.getStoryReactionsList#b9b2881f: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.getStoryReactionsList#b9b2881f: field peer: %w", err)
	}
	b.PutInt(g.ID)
	if g.Flags.Has(0) {
		if g.Reaction == nil {
			return fmt.Errorf("unable to encode stories.getStoryReactionsList#b9b2881f: field reaction is nil")
		}
		if err := g.Reaction.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.getStoryReactionsList#b9b2881f: field reaction: %w", err)
		}
	}
	if g.Flags.Has(1) {
		b.PutString(g.Offset)
	}
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *StoriesGetStoryReactionsListRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getStoryReactionsList#b9b2881f to nil")
	}
	if err := b.ConsumeID(StoriesGetStoryReactionsListRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.getStoryReactionsList#b9b2881f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StoriesGetStoryReactionsListRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getStoryReactionsList#b9b2881f to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.getStoryReactionsList#b9b2881f: field flags: %w", err)
		}
	}
	g.ForwardsFirst = g.Flags.Has(2)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.getStoryReactionsList#b9b2881f: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.getStoryReactionsList#b9b2881f: field id: %w", err)
		}
		g.ID = value
	}
	if g.Flags.Has(0) {
		value, err := DecodeReaction(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.getStoryReactionsList#b9b2881f: field reaction: %w", err)
		}
		g.Reaction = value
	}
	if g.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stories.getStoryReactionsList#b9b2881f: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.getStoryReactionsList#b9b2881f: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// SetForwardsFirst sets value of ForwardsFirst conditional field.
func (g *StoriesGetStoryReactionsListRequest) SetForwardsFirst(value bool) {
	if value {
		g.Flags.Set(2)
		g.ForwardsFirst = true
	} else {
		g.Flags.Unset(2)
		g.ForwardsFirst = false
	}
}

// GetForwardsFirst returns value of ForwardsFirst conditional field.
func (g *StoriesGetStoryReactionsListRequest) GetForwardsFirst() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(2)
}

// GetPeer returns value of Peer field.
func (g *StoriesGetStoryReactionsListRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetID returns value of ID field.
func (g *StoriesGetStoryReactionsListRequest) GetID() (value int) {
	if g == nil {
		return
	}
	return g.ID
}

// SetReaction sets value of Reaction conditional field.
func (g *StoriesGetStoryReactionsListRequest) SetReaction(value ReactionClass) {
	g.Flags.Set(0)
	g.Reaction = value
}

// GetReaction returns value of Reaction conditional field and
// boolean which is true if field was set.
func (g *StoriesGetStoryReactionsListRequest) GetReaction() (value ReactionClass, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Reaction, true
}

// SetOffset sets value of Offset conditional field.
func (g *StoriesGetStoryReactionsListRequest) SetOffset(value string) {
	g.Flags.Set(1)
	g.Offset = value
}

// GetOffset returns value of Offset conditional field and
// boolean which is true if field was set.
func (g *StoriesGetStoryReactionsListRequest) GetOffset() (value string, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(1) {
		return value, false
	}
	return g.Offset, true
}

// GetLimit returns value of Limit field.
func (g *StoriesGetStoryReactionsListRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// StoriesGetStoryReactionsList invokes method stories.getStoryReactionsList#b9b2881f returning error if any.
func (c *Client) StoriesGetStoryReactionsList(ctx context.Context, request *StoriesGetStoryReactionsListRequest) (*StoriesStoryReactionsList, error) {
	var result StoriesStoryReactionsList

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
