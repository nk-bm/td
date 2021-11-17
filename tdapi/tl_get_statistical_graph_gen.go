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

// GetStatisticalGraphRequest represents TL type `getStatisticalGraph#419f8d9b`.
type GetStatisticalGraphRequest struct {
	// Chat identifier
	ChatID int64
	// The token for graph loading
	Token string
	// X-value for zoomed in graph or 0 otherwise
	X int64
}

// GetStatisticalGraphRequestTypeID is TL type id of GetStatisticalGraphRequest.
const GetStatisticalGraphRequestTypeID = 0x419f8d9b

// Ensuring interfaces in compile-time for GetStatisticalGraphRequest.
var (
	_ bin.Encoder     = &GetStatisticalGraphRequest{}
	_ bin.Decoder     = &GetStatisticalGraphRequest{}
	_ bin.BareEncoder = &GetStatisticalGraphRequest{}
	_ bin.BareDecoder = &GetStatisticalGraphRequest{}
)

func (g *GetStatisticalGraphRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.Token == "") {
		return false
	}
	if !(g.X == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetStatisticalGraphRequest) String() string {
	if g == nil {
		return "GetStatisticalGraphRequest(nil)"
	}
	type Alias GetStatisticalGraphRequest
	return fmt.Sprintf("GetStatisticalGraphRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetStatisticalGraphRequest) TypeID() uint32 {
	return GetStatisticalGraphRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetStatisticalGraphRequest) TypeName() string {
	return "getStatisticalGraph"
}

// TypeInfo returns info about TL type.
func (g *GetStatisticalGraphRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getStatisticalGraph",
		ID:   GetStatisticalGraphRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Token",
			SchemaName: "token",
		},
		{
			Name:       "X",
			SchemaName: "x",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetStatisticalGraphRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStatisticalGraph#419f8d9b as nil")
	}
	b.PutID(GetStatisticalGraphRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetStatisticalGraphRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStatisticalGraph#419f8d9b as nil")
	}
	b.PutLong(g.ChatID)
	b.PutString(g.Token)
	b.PutLong(g.X)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetStatisticalGraphRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStatisticalGraph#419f8d9b to nil")
	}
	if err := b.ConsumeID(GetStatisticalGraphRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getStatisticalGraph#419f8d9b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetStatisticalGraphRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStatisticalGraph#419f8d9b to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getStatisticalGraph#419f8d9b: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getStatisticalGraph#419f8d9b: field token: %w", err)
		}
		g.Token = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getStatisticalGraph#419f8d9b: field x: %w", err)
		}
		g.X = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (g *GetStatisticalGraphRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetToken returns value of Token field.
func (g *GetStatisticalGraphRequest) GetToken() (value string) {
	return g.Token
}

// GetX returns value of X field.
func (g *GetStatisticalGraphRequest) GetX() (value int64) {
	return g.X
}

// GetStatisticalGraph invokes method getStatisticalGraph#419f8d9b returning error if any.
func (c *Client) GetStatisticalGraph(ctx context.Context, request *GetStatisticalGraphRequest) (StatisticalGraphClass, error) {
	var result StatisticalGraphBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.StatisticalGraph, nil
}
