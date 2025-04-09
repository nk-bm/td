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

// StoriesGetAllStoriesRequest represents TL type `stories.getAllStories#eeb0d625`.
type StoriesGetAllStoriesRequest struct {
	// Flags field of StoriesGetAllStoriesRequest.
	Flags bin.Fields
	// Next field of StoriesGetAllStoriesRequest.
	Next bool
	// Hidden field of StoriesGetAllStoriesRequest.
	Hidden bool
	// State field of StoriesGetAllStoriesRequest.
	//
	// Use SetState and GetState helpers.
	State string
}

// StoriesGetAllStoriesRequestTypeID is TL type id of StoriesGetAllStoriesRequest.
const StoriesGetAllStoriesRequestTypeID = 0xeeb0d625

// Ensuring interfaces in compile-time for StoriesGetAllStoriesRequest.
var (
	_ bin.Encoder     = &StoriesGetAllStoriesRequest{}
	_ bin.Decoder     = &StoriesGetAllStoriesRequest{}
	_ bin.BareEncoder = &StoriesGetAllStoriesRequest{}
	_ bin.BareDecoder = &StoriesGetAllStoriesRequest{}
)

func (g *StoriesGetAllStoriesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Next == false) {
		return false
	}
	if !(g.Hidden == false) {
		return false
	}
	if !(g.State == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StoriesGetAllStoriesRequest) String() string {
	if g == nil {
		return "StoriesGetAllStoriesRequest(nil)"
	}
	type Alias StoriesGetAllStoriesRequest
	return fmt.Sprintf("StoriesGetAllStoriesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesGetAllStoriesRequest) TypeID() uint32 {
	return StoriesGetAllStoriesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesGetAllStoriesRequest) TypeName() string {
	return "stories.getAllStories"
}

// TypeInfo returns info about TL type.
func (g *StoriesGetAllStoriesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.getAllStories",
		ID:   StoriesGetAllStoriesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Next",
			SchemaName: "next",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "Hidden",
			SchemaName: "hidden",
			Null:       !g.Flags.Has(2),
		},
		{
			Name:       "State",
			SchemaName: "state",
			Null:       !g.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *StoriesGetAllStoriesRequest) SetFlags() {
	if !(g.Next == false) {
		g.Flags.Set(1)
	}
	if !(g.Hidden == false) {
		g.Flags.Set(2)
	}
	if !(g.State == "") {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *StoriesGetAllStoriesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getAllStories#eeb0d625 as nil")
	}
	b.PutID(StoriesGetAllStoriesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StoriesGetAllStoriesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getAllStories#eeb0d625 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.getAllStories#eeb0d625: field flags: %w", err)
	}
	if g.Flags.Has(0) {
		b.PutString(g.State)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *StoriesGetAllStoriesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getAllStories#eeb0d625 to nil")
	}
	if err := b.ConsumeID(StoriesGetAllStoriesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.getAllStories#eeb0d625: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StoriesGetAllStoriesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getAllStories#eeb0d625 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.getAllStories#eeb0d625: field flags: %w", err)
		}
	}
	g.Next = g.Flags.Has(1)
	g.Hidden = g.Flags.Has(2)
	if g.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stories.getAllStories#eeb0d625: field state: %w", err)
		}
		g.State = value
	}
	return nil
}

// SetNext sets value of Next conditional field.
func (g *StoriesGetAllStoriesRequest) SetNext(value bool) {
	if value {
		g.Flags.Set(1)
		g.Next = true
	} else {
		g.Flags.Unset(1)
		g.Next = false
	}
}

// GetNext returns value of Next conditional field.
func (g *StoriesGetAllStoriesRequest) GetNext() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(1)
}

// SetHidden sets value of Hidden conditional field.
func (g *StoriesGetAllStoriesRequest) SetHidden(value bool) {
	if value {
		g.Flags.Set(2)
		g.Hidden = true
	} else {
		g.Flags.Unset(2)
		g.Hidden = false
	}
}

// GetHidden returns value of Hidden conditional field.
func (g *StoriesGetAllStoriesRequest) GetHidden() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(2)
}

// SetState sets value of State conditional field.
func (g *StoriesGetAllStoriesRequest) SetState(value string) {
	g.Flags.Set(0)
	g.State = value
}

// GetState returns value of State conditional field and
// boolean which is true if field was set.
func (g *StoriesGetAllStoriesRequest) GetState() (value string, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.State, true
}

// StoriesGetAllStories invokes method stories.getAllStories#eeb0d625 returning error if any.
func (c *Client) StoriesGetAllStories(ctx context.Context, request *StoriesGetAllStoriesRequest) (StoriesAllStoriesClass, error) {
	var result StoriesAllStoriesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.AllStories, nil
}
