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

// StoriesGetPinnedStoriesRequest represents TL type `stories.getPinnedStories#5821a5dc`.
// Fetch the stories¹ pinned on a peer's profile.
//
// Links:
//  1. https://core.telegram.org/api/stories#pinned-or-archived-stories
//
// See https://core.telegram.org/method/stories.getPinnedStories for reference.
type StoriesGetPinnedStoriesRequest struct {
	// Peer whose pinned stories should be fetched
	Peer InputPeerClass
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// StoriesGetPinnedStoriesRequestTypeID is TL type id of StoriesGetPinnedStoriesRequest.
const StoriesGetPinnedStoriesRequestTypeID = 0x5821a5dc

// Ensuring interfaces in compile-time for StoriesGetPinnedStoriesRequest.
var (
	_ bin.Encoder     = &StoriesGetPinnedStoriesRequest{}
	_ bin.Decoder     = &StoriesGetPinnedStoriesRequest{}
	_ bin.BareEncoder = &StoriesGetPinnedStoriesRequest{}
	_ bin.BareDecoder = &StoriesGetPinnedStoriesRequest{}
)

func (g *StoriesGetPinnedStoriesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StoriesGetPinnedStoriesRequest) String() string {
	if g == nil {
		return "StoriesGetPinnedStoriesRequest(nil)"
	}
	type Alias StoriesGetPinnedStoriesRequest
	return fmt.Sprintf("StoriesGetPinnedStoriesRequest%+v", Alias(*g))
}

// FillFrom fills StoriesGetPinnedStoriesRequest from given interface.
func (g *StoriesGetPinnedStoriesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetOffsetID() (value int)
	GetLimit() (value int)
}) {
	g.Peer = from.GetPeer()
	g.OffsetID = from.GetOffsetID()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesGetPinnedStoriesRequest) TypeID() uint32 {
	return StoriesGetPinnedStoriesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesGetPinnedStoriesRequest) TypeName() string {
	return "stories.getPinnedStories"
}

// TypeInfo returns info about TL type.
func (g *StoriesGetPinnedStoriesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.getPinnedStories",
		ID:   StoriesGetPinnedStoriesRequestTypeID,
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
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *StoriesGetPinnedStoriesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getPinnedStories#5821a5dc as nil")
	}
	b.PutID(StoriesGetPinnedStoriesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StoriesGetPinnedStoriesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getPinnedStories#5821a5dc as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode stories.getPinnedStories#5821a5dc: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.getPinnedStories#5821a5dc: field peer: %w", err)
	}
	b.PutInt(g.OffsetID)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *StoriesGetPinnedStoriesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getPinnedStories#5821a5dc to nil")
	}
	if err := b.ConsumeID(StoriesGetPinnedStoriesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.getPinnedStories#5821a5dc: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StoriesGetPinnedStoriesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getPinnedStories#5821a5dc to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.getPinnedStories#5821a5dc: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.getPinnedStories#5821a5dc: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.getPinnedStories#5821a5dc: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *StoriesGetPinnedStoriesRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetOffsetID returns value of OffsetID field.
func (g *StoriesGetPinnedStoriesRequest) GetOffsetID() (value int) {
	if g == nil {
		return
	}
	return g.OffsetID
}

// GetLimit returns value of Limit field.
func (g *StoriesGetPinnedStoriesRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// StoriesGetPinnedStories invokes method stories.getPinnedStories#5821a5dc returning error if any.
// Fetch the stories¹ pinned on a peer's profile.
//
// Links:
//  1. https://core.telegram.org/api/stories#pinned-or-archived-stories
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 USER_ID_INVALID: The provided user ID is invalid.
//
// See https://core.telegram.org/method/stories.getPinnedStories for reference.
func (c *Client) StoriesGetPinnedStories(ctx context.Context, request *StoriesGetPinnedStoriesRequest) (*StoriesStories, error) {
	var result StoriesStories

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
