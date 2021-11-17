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

// PingProxyRequest represents TL type `pingProxy#c59b40b1`.
type PingProxyRequest struct {
	// Proxy identifier. Use 0 to ping a Telegram server without a proxy
	ProxyID int32
}

// PingProxyRequestTypeID is TL type id of PingProxyRequest.
const PingProxyRequestTypeID = 0xc59b40b1

// Ensuring interfaces in compile-time for PingProxyRequest.
var (
	_ bin.Encoder     = &PingProxyRequest{}
	_ bin.Decoder     = &PingProxyRequest{}
	_ bin.BareEncoder = &PingProxyRequest{}
	_ bin.BareDecoder = &PingProxyRequest{}
)

func (p *PingProxyRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ProxyID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PingProxyRequest) String() string {
	if p == nil {
		return "PingProxyRequest(nil)"
	}
	type Alias PingProxyRequest
	return fmt.Sprintf("PingProxyRequest%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PingProxyRequest) TypeID() uint32 {
	return PingProxyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PingProxyRequest) TypeName() string {
	return "pingProxy"
}

// TypeInfo returns info about TL type.
func (p *PingProxyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pingProxy",
		ID:   PingProxyRequestTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ProxyID",
			SchemaName: "proxy_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PingProxyRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pingProxy#c59b40b1 as nil")
	}
	b.PutID(PingProxyRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PingProxyRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pingProxy#c59b40b1 as nil")
	}
	b.PutInt32(p.ProxyID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PingProxyRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pingProxy#c59b40b1 to nil")
	}
	if err := b.ConsumeID(PingProxyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode pingProxy#c59b40b1: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PingProxyRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pingProxy#c59b40b1 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode pingProxy#c59b40b1: field proxy_id: %w", err)
		}
		p.ProxyID = value
	}
	return nil
}

// GetProxyID returns value of ProxyID field.
func (p *PingProxyRequest) GetProxyID() (value int32) {
	return p.ProxyID
}

// PingProxy invokes method pingProxy#c59b40b1 returning error if any.
func (c *Client) PingProxy(ctx context.Context, proxyid int32) (*Seconds, error) {
	var result Seconds

	request := &PingProxyRequest{
		ProxyID: proxyid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
