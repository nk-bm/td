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

// SendCustomRequestRequest represents TL type `sendCustomRequest#10fd71a1`.
type SendCustomRequestRequest struct {
	// The method name
	Method string
	// JSON-serialized method parameters
	Parameters string
}

// SendCustomRequestRequestTypeID is TL type id of SendCustomRequestRequest.
const SendCustomRequestRequestTypeID = 0x10fd71a1

// Ensuring interfaces in compile-time for SendCustomRequestRequest.
var (
	_ bin.Encoder     = &SendCustomRequestRequest{}
	_ bin.Decoder     = &SendCustomRequestRequest{}
	_ bin.BareEncoder = &SendCustomRequestRequest{}
	_ bin.BareDecoder = &SendCustomRequestRequest{}
)

func (s *SendCustomRequestRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Method == "") {
		return false
	}
	if !(s.Parameters == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendCustomRequestRequest) String() string {
	if s == nil {
		return "SendCustomRequestRequest(nil)"
	}
	type Alias SendCustomRequestRequest
	return fmt.Sprintf("SendCustomRequestRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendCustomRequestRequest) TypeID() uint32 {
	return SendCustomRequestRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendCustomRequestRequest) TypeName() string {
	return "sendCustomRequest"
}

// TypeInfo returns info about TL type.
func (s *SendCustomRequestRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendCustomRequest",
		ID:   SendCustomRequestRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Method",
			SchemaName: "method",
		},
		{
			Name:       "Parameters",
			SchemaName: "parameters",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendCustomRequestRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendCustomRequest#10fd71a1 as nil")
	}
	b.PutID(SendCustomRequestRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendCustomRequestRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendCustomRequest#10fd71a1 as nil")
	}
	b.PutString(s.Method)
	b.PutString(s.Parameters)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendCustomRequestRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendCustomRequest#10fd71a1 to nil")
	}
	if err := b.ConsumeID(SendCustomRequestRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendCustomRequest#10fd71a1: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendCustomRequestRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendCustomRequest#10fd71a1 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sendCustomRequest#10fd71a1: field method: %w", err)
		}
		s.Method = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sendCustomRequest#10fd71a1: field parameters: %w", err)
		}
		s.Parameters = value
	}
	return nil
}

// GetMethod returns value of Method field.
func (s *SendCustomRequestRequest) GetMethod() (value string) {
	return s.Method
}

// GetParameters returns value of Parameters field.
func (s *SendCustomRequestRequest) GetParameters() (value string) {
	return s.Parameters
}

// SendCustomRequest invokes method sendCustomRequest#10fd71a1 returning error if any.
func (c *Client) SendCustomRequest(ctx context.Context, request *SendCustomRequestRequest) (*CustomRequestResult, error) {
	var result CustomRequestResult

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
