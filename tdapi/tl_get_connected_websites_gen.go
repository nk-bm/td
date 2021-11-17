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

// GetConnectedWebsitesRequest represents TL type `getConnectedWebsites#f5d5d352`.
type GetConnectedWebsitesRequest struct {
}

// GetConnectedWebsitesRequestTypeID is TL type id of GetConnectedWebsitesRequest.
const GetConnectedWebsitesRequestTypeID = 0xf5d5d352

// Ensuring interfaces in compile-time for GetConnectedWebsitesRequest.
var (
	_ bin.Encoder     = &GetConnectedWebsitesRequest{}
	_ bin.Decoder     = &GetConnectedWebsitesRequest{}
	_ bin.BareEncoder = &GetConnectedWebsitesRequest{}
	_ bin.BareDecoder = &GetConnectedWebsitesRequest{}
)

func (g *GetConnectedWebsitesRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetConnectedWebsitesRequest) String() string {
	if g == nil {
		return "GetConnectedWebsitesRequest(nil)"
	}
	type Alias GetConnectedWebsitesRequest
	return fmt.Sprintf("GetConnectedWebsitesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetConnectedWebsitesRequest) TypeID() uint32 {
	return GetConnectedWebsitesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetConnectedWebsitesRequest) TypeName() string {
	return "getConnectedWebsites"
}

// TypeInfo returns info about TL type.
func (g *GetConnectedWebsitesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getConnectedWebsites",
		ID:   GetConnectedWebsitesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetConnectedWebsitesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getConnectedWebsites#f5d5d352 as nil")
	}
	b.PutID(GetConnectedWebsitesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetConnectedWebsitesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getConnectedWebsites#f5d5d352 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetConnectedWebsitesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getConnectedWebsites#f5d5d352 to nil")
	}
	if err := b.ConsumeID(GetConnectedWebsitesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getConnectedWebsites#f5d5d352: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetConnectedWebsitesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getConnectedWebsites#f5d5d352 to nil")
	}
	return nil
}

// GetConnectedWebsites invokes method getConnectedWebsites#f5d5d352 returning error if any.
func (c *Client) GetConnectedWebsites(ctx context.Context) (*ConnectedWebsites, error) {
	var result ConnectedWebsites

	request := &GetConnectedWebsitesRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
