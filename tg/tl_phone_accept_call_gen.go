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

// PhoneAcceptCallRequest represents TL type `phone.acceptCall#3bd2b4a0`.
type PhoneAcceptCallRequest struct {
	// Peer field of PhoneAcceptCallRequest.
	Peer InputPhoneCall
	// GB field of PhoneAcceptCallRequest.
	GB []byte
	// Protocol field of PhoneAcceptCallRequest.
	Protocol PhoneCallProtocol
}

// PhoneAcceptCallRequestTypeID is TL type id of PhoneAcceptCallRequest.
const PhoneAcceptCallRequestTypeID = 0x3bd2b4a0

// Ensuring interfaces in compile-time for PhoneAcceptCallRequest.
var (
	_ bin.Encoder     = &PhoneAcceptCallRequest{}
	_ bin.Decoder     = &PhoneAcceptCallRequest{}
	_ bin.BareEncoder = &PhoneAcceptCallRequest{}
	_ bin.BareDecoder = &PhoneAcceptCallRequest{}
)

func (a *PhoneAcceptCallRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Peer.Zero()) {
		return false
	}
	if !(a.GB == nil) {
		return false
	}
	if !(a.Protocol.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *PhoneAcceptCallRequest) String() string {
	if a == nil {
		return "PhoneAcceptCallRequest(nil)"
	}
	type Alias PhoneAcceptCallRequest
	return fmt.Sprintf("PhoneAcceptCallRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneAcceptCallRequest) TypeID() uint32 {
	return PhoneAcceptCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneAcceptCallRequest) TypeName() string {
	return "phone.acceptCall"
}

// TypeInfo returns info about TL type.
func (a *PhoneAcceptCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.acceptCall",
		ID:   PhoneAcceptCallRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "GB",
			SchemaName: "g_b",
		},
		{
			Name:       "Protocol",
			SchemaName: "protocol",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *PhoneAcceptCallRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode phone.acceptCall#3bd2b4a0 as nil")
	}
	b.PutID(PhoneAcceptCallRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *PhoneAcceptCallRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode phone.acceptCall#3bd2b4a0 as nil")
	}
	if err := a.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.acceptCall#3bd2b4a0: field peer: %w", err)
	}
	b.PutBytes(a.GB)
	if err := a.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.acceptCall#3bd2b4a0: field protocol: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *PhoneAcceptCallRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode phone.acceptCall#3bd2b4a0 to nil")
	}
	if err := b.ConsumeID(PhoneAcceptCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *PhoneAcceptCallRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode phone.acceptCall#3bd2b4a0 to nil")
	}
	{
		if err := a.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: field peer: %w", err)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: field g_b: %w", err)
		}
		a.GB = value
	}
	{
		if err := a.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: field protocol: %w", err)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (a *PhoneAcceptCallRequest) GetPeer() (value InputPhoneCall) {
	if a == nil {
		return
	}
	return a.Peer
}

// GetGB returns value of GB field.
func (a *PhoneAcceptCallRequest) GetGB() (value []byte) {
	if a == nil {
		return
	}
	return a.GB
}

// GetProtocol returns value of Protocol field.
func (a *PhoneAcceptCallRequest) GetProtocol() (value PhoneCallProtocol) {
	if a == nil {
		return
	}
	return a.Protocol
}

// PhoneAcceptCall invokes method phone.acceptCall#3bd2b4a0 returning error if any.
func (c *Client) PhoneAcceptCall(ctx context.Context, request *PhoneAcceptCallRequest) (*PhonePhoneCall, error) {
	var result PhonePhoneCall

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
