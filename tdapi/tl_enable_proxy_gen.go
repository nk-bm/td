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

// EnableProxyRequest represents TL type `enableProxy#59138296`.
type EnableProxyRequest struct {
	// Proxy identifier
	ProxyID int32
}

// EnableProxyRequestTypeID is TL type id of EnableProxyRequest.
const EnableProxyRequestTypeID = 0x59138296

// Ensuring interfaces in compile-time for EnableProxyRequest.
var (
	_ bin.Encoder     = &EnableProxyRequest{}
	_ bin.Decoder     = &EnableProxyRequest{}
	_ bin.BareEncoder = &EnableProxyRequest{}
	_ bin.BareDecoder = &EnableProxyRequest{}
)

func (e *EnableProxyRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ProxyID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EnableProxyRequest) String() string {
	if e == nil {
		return "EnableProxyRequest(nil)"
	}
	type Alias EnableProxyRequest
	return fmt.Sprintf("EnableProxyRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EnableProxyRequest) TypeID() uint32 {
	return EnableProxyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EnableProxyRequest) TypeName() string {
	return "enableProxy"
}

// TypeInfo returns info about TL type.
func (e *EnableProxyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "enableProxy",
		ID:   EnableProxyRequestTypeID,
	}
	if e == nil {
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
func (e *EnableProxyRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode enableProxy#59138296 as nil")
	}
	b.PutID(EnableProxyRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EnableProxyRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode enableProxy#59138296 as nil")
	}
	b.PutInt32(e.ProxyID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EnableProxyRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode enableProxy#59138296 to nil")
	}
	if err := b.ConsumeID(EnableProxyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode enableProxy#59138296: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EnableProxyRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode enableProxy#59138296 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode enableProxy#59138296: field proxy_id: %w", err)
		}
		e.ProxyID = value
	}
	return nil
}

// GetProxyID returns value of ProxyID field.
func (e *EnableProxyRequest) GetProxyID() (value int32) {
	return e.ProxyID
}

// EnableProxy invokes method enableProxy#59138296 returning error if any.
func (c *Client) EnableProxy(ctx context.Context, proxyid int32) error {
	var ok Ok

	request := &EnableProxyRequest{
		ProxyID: proxyid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
