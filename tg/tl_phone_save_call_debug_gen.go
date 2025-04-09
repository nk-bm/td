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

// PhoneSaveCallDebugRequest represents TL type `phone.saveCallDebug#277add7e`.
type PhoneSaveCallDebugRequest struct {
	// Peer field of PhoneSaveCallDebugRequest.
	Peer InputPhoneCall
	// Debug field of PhoneSaveCallDebugRequest.
	Debug DataJSON
}

// PhoneSaveCallDebugRequestTypeID is TL type id of PhoneSaveCallDebugRequest.
const PhoneSaveCallDebugRequestTypeID = 0x277add7e

// Ensuring interfaces in compile-time for PhoneSaveCallDebugRequest.
var (
	_ bin.Encoder     = &PhoneSaveCallDebugRequest{}
	_ bin.Decoder     = &PhoneSaveCallDebugRequest{}
	_ bin.BareEncoder = &PhoneSaveCallDebugRequest{}
	_ bin.BareDecoder = &PhoneSaveCallDebugRequest{}
)

func (s *PhoneSaveCallDebugRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer.Zero()) {
		return false
	}
	if !(s.Debug.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PhoneSaveCallDebugRequest) String() string {
	if s == nil {
		return "PhoneSaveCallDebugRequest(nil)"
	}
	type Alias PhoneSaveCallDebugRequest
	return fmt.Sprintf("PhoneSaveCallDebugRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneSaveCallDebugRequest) TypeID() uint32 {
	return PhoneSaveCallDebugRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneSaveCallDebugRequest) TypeName() string {
	return "phone.saveCallDebug"
}

// TypeInfo returns info about TL type.
func (s *PhoneSaveCallDebugRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.saveCallDebug",
		ID:   PhoneSaveCallDebugRequestTypeID,
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
			Name:       "Debug",
			SchemaName: "debug",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *PhoneSaveCallDebugRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.saveCallDebug#277add7e as nil")
	}
	b.PutID(PhoneSaveCallDebugRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *PhoneSaveCallDebugRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.saveCallDebug#277add7e as nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.saveCallDebug#277add7e: field peer: %w", err)
	}
	if err := s.Debug.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.saveCallDebug#277add7e: field debug: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *PhoneSaveCallDebugRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.saveCallDebug#277add7e to nil")
	}
	if err := b.ConsumeID(PhoneSaveCallDebugRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.saveCallDebug#277add7e: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *PhoneSaveCallDebugRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.saveCallDebug#277add7e to nil")
	}
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.saveCallDebug#277add7e: field peer: %w", err)
		}
	}
	{
		if err := s.Debug.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.saveCallDebug#277add7e: field debug: %w", err)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *PhoneSaveCallDebugRequest) GetPeer() (value InputPhoneCall) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetDebug returns value of Debug field.
func (s *PhoneSaveCallDebugRequest) GetDebug() (value DataJSON) {
	if s == nil {
		return
	}
	return s.Debug
}

// PhoneSaveCallDebug invokes method phone.saveCallDebug#277add7e returning error if any.
func (c *Client) PhoneSaveCallDebug(ctx context.Context, request *PhoneSaveCallDebugRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
