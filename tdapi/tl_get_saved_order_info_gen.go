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

// GetSavedOrderInfoRequest represents TL type `getSavedOrderInfo#bb559edd`.
type GetSavedOrderInfoRequest struct {
}

// GetSavedOrderInfoRequestTypeID is TL type id of GetSavedOrderInfoRequest.
const GetSavedOrderInfoRequestTypeID = 0xbb559edd

// Ensuring interfaces in compile-time for GetSavedOrderInfoRequest.
var (
	_ bin.Encoder     = &GetSavedOrderInfoRequest{}
	_ bin.Decoder     = &GetSavedOrderInfoRequest{}
	_ bin.BareEncoder = &GetSavedOrderInfoRequest{}
	_ bin.BareDecoder = &GetSavedOrderInfoRequest{}
)

func (g *GetSavedOrderInfoRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSavedOrderInfoRequest) String() string {
	if g == nil {
		return "GetSavedOrderInfoRequest(nil)"
	}
	type Alias GetSavedOrderInfoRequest
	return fmt.Sprintf("GetSavedOrderInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSavedOrderInfoRequest) TypeID() uint32 {
	return GetSavedOrderInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSavedOrderInfoRequest) TypeName() string {
	return "getSavedOrderInfo"
}

// TypeInfo returns info about TL type.
func (g *GetSavedOrderInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSavedOrderInfo",
		ID:   GetSavedOrderInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSavedOrderInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSavedOrderInfo#bb559edd as nil")
	}
	b.PutID(GetSavedOrderInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSavedOrderInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSavedOrderInfo#bb559edd as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSavedOrderInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSavedOrderInfo#bb559edd to nil")
	}
	if err := b.ConsumeID(GetSavedOrderInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSavedOrderInfo#bb559edd: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSavedOrderInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSavedOrderInfo#bb559edd to nil")
	}
	return nil
}

// GetSavedOrderInfo invokes method getSavedOrderInfo#bb559edd returning error if any.
func (c *Client) GetSavedOrderInfo(ctx context.Context) (*OrderInfo, error) {
	var result OrderInfo

	request := &GetSavedOrderInfoRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
