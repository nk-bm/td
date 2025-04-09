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

// PhoneSendSignalingDataRequest represents TL type `phone.sendSignalingData#ff7a9383`.
type PhoneSendSignalingDataRequest struct {
	// Peer field of PhoneSendSignalingDataRequest.
	Peer InputPhoneCall
	// Data field of PhoneSendSignalingDataRequest.
	Data []byte
}

// PhoneSendSignalingDataRequestTypeID is TL type id of PhoneSendSignalingDataRequest.
const PhoneSendSignalingDataRequestTypeID = 0xff7a9383

// Ensuring interfaces in compile-time for PhoneSendSignalingDataRequest.
var (
	_ bin.Encoder     = &PhoneSendSignalingDataRequest{}
	_ bin.Decoder     = &PhoneSendSignalingDataRequest{}
	_ bin.BareEncoder = &PhoneSendSignalingDataRequest{}
	_ bin.BareDecoder = &PhoneSendSignalingDataRequest{}
)

func (s *PhoneSendSignalingDataRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer.Zero()) {
		return false
	}
	if !(s.Data == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PhoneSendSignalingDataRequest) String() string {
	if s == nil {
		return "PhoneSendSignalingDataRequest(nil)"
	}
	type Alias PhoneSendSignalingDataRequest
	return fmt.Sprintf("PhoneSendSignalingDataRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneSendSignalingDataRequest) TypeID() uint32 {
	return PhoneSendSignalingDataRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneSendSignalingDataRequest) TypeName() string {
	return "phone.sendSignalingData"
}

// TypeInfo returns info about TL type.
func (s *PhoneSendSignalingDataRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.sendSignalingData",
		ID:   PhoneSendSignalingDataRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *PhoneSendSignalingDataRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.sendSignalingData#ff7a9383 as nil")
	}
	b.PutID(PhoneSendSignalingDataRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *PhoneSendSignalingDataRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.sendSignalingData#ff7a9383 as nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.sendSignalingData#ff7a9383: field peer: %w", err)
	}
	b.PutBytes(s.Data)
	return nil
}

// Decode implements bin.Decoder.
func (s *PhoneSendSignalingDataRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.sendSignalingData#ff7a9383 to nil")
	}
	if err := b.ConsumeID(PhoneSendSignalingDataRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.sendSignalingData#ff7a9383: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *PhoneSendSignalingDataRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.sendSignalingData#ff7a9383 to nil")
	}
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.sendSignalingData#ff7a9383: field peer: %w", err)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phone.sendSignalingData#ff7a9383: field data: %w", err)
		}
		s.Data = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *PhoneSendSignalingDataRequest) GetPeer() (value InputPhoneCall) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetData returns value of Data field.
func (s *PhoneSendSignalingDataRequest) GetData() (value []byte) {
	if s == nil {
		return
	}
	return s.Data
}

// PhoneSendSignalingData invokes method phone.sendSignalingData#ff7a9383 returning error if any.
func (c *Client) PhoneSendSignalingData(ctx context.Context, request *PhoneSendSignalingDataRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
