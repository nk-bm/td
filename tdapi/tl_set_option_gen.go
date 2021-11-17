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

// SetOptionRequest represents TL type `setOption#7e0b4ef2`.
type SetOptionRequest struct {
	// The name of the option
	Name string
	// The new value of the option
	Value OptionValueClass
}

// SetOptionRequestTypeID is TL type id of SetOptionRequest.
const SetOptionRequestTypeID = 0x7e0b4ef2

// Ensuring interfaces in compile-time for SetOptionRequest.
var (
	_ bin.Encoder     = &SetOptionRequest{}
	_ bin.Decoder     = &SetOptionRequest{}
	_ bin.BareEncoder = &SetOptionRequest{}
	_ bin.BareDecoder = &SetOptionRequest{}
)

func (s *SetOptionRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Name == "") {
		return false
	}
	if !(s.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetOptionRequest) String() string {
	if s == nil {
		return "SetOptionRequest(nil)"
	}
	type Alias SetOptionRequest
	return fmt.Sprintf("SetOptionRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetOptionRequest) TypeID() uint32 {
	return SetOptionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetOptionRequest) TypeName() string {
	return "setOption"
}

// TypeInfo returns info about TL type.
func (s *SetOptionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setOption",
		ID:   SetOptionRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetOptionRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setOption#7e0b4ef2 as nil")
	}
	b.PutID(SetOptionRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetOptionRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setOption#7e0b4ef2 as nil")
	}
	b.PutString(s.Name)
	if s.Value == nil {
		return fmt.Errorf("unable to encode setOption#7e0b4ef2: field value is nil")
	}
	if err := s.Value.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setOption#7e0b4ef2: field value: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetOptionRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setOption#7e0b4ef2 to nil")
	}
	if err := b.ConsumeID(SetOptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setOption#7e0b4ef2: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetOptionRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setOption#7e0b4ef2 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setOption#7e0b4ef2: field name: %w", err)
		}
		s.Name = value
	}
	{
		value, err := DecodeOptionValue(b)
		if err != nil {
			return fmt.Errorf("unable to decode setOption#7e0b4ef2: field value: %w", err)
		}
		s.Value = value
	}
	return nil
}

// GetName returns value of Name field.
func (s *SetOptionRequest) GetName() (value string) {
	return s.Name
}

// GetValue returns value of Value field.
func (s *SetOptionRequest) GetValue() (value OptionValueClass) {
	return s.Value
}

// SetOption invokes method setOption#7e0b4ef2 returning error if any.
func (c *Client) SetOption(ctx context.Context, request *SetOptionRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
