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

// GetPassportElementRequest represents TL type `getPassportElement#8fcce17a`.
type GetPassportElementRequest struct {
	// Telegram Passport element type
	Type PassportElementTypeClass
	// Password of the current user
	Password string
}

// GetPassportElementRequestTypeID is TL type id of GetPassportElementRequest.
const GetPassportElementRequestTypeID = 0x8fcce17a

// Ensuring interfaces in compile-time for GetPassportElementRequest.
var (
	_ bin.Encoder     = &GetPassportElementRequest{}
	_ bin.Decoder     = &GetPassportElementRequest{}
	_ bin.BareEncoder = &GetPassportElementRequest{}
	_ bin.BareDecoder = &GetPassportElementRequest{}
)

func (g *GetPassportElementRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Type == nil) {
		return false
	}
	if !(g.Password == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetPassportElementRequest) String() string {
	if g == nil {
		return "GetPassportElementRequest(nil)"
	}
	type Alias GetPassportElementRequest
	return fmt.Sprintf("GetPassportElementRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetPassportElementRequest) TypeID() uint32 {
	return GetPassportElementRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetPassportElementRequest) TypeName() string {
	return "getPassportElement"
}

// TypeInfo returns info about TL type.
func (g *GetPassportElementRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getPassportElement",
		ID:   GetPassportElementRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetPassportElementRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPassportElement#8fcce17a as nil")
	}
	b.PutID(GetPassportElementRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetPassportElementRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPassportElement#8fcce17a as nil")
	}
	if g.Type == nil {
		return fmt.Errorf("unable to encode getPassportElement#8fcce17a: field type is nil")
	}
	if err := g.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getPassportElement#8fcce17a: field type: %w", err)
	}
	b.PutString(g.Password)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetPassportElementRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPassportElement#8fcce17a to nil")
	}
	if err := b.ConsumeID(GetPassportElementRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getPassportElement#8fcce17a: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetPassportElementRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPassportElement#8fcce17a to nil")
	}
	{
		value, err := DecodePassportElementType(b)
		if err != nil {
			return fmt.Errorf("unable to decode getPassportElement#8fcce17a: field type: %w", err)
		}
		g.Type = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getPassportElement#8fcce17a: field password: %w", err)
		}
		g.Password = value
	}
	return nil
}

// GetType returns value of Type field.
func (g *GetPassportElementRequest) GetType() (value PassportElementTypeClass) {
	return g.Type
}

// GetPassword returns value of Password field.
func (g *GetPassportElementRequest) GetPassword() (value string) {
	return g.Password
}

// GetPassportElement invokes method getPassportElement#8fcce17a returning error if any.
func (c *Client) GetPassportElement(ctx context.Context, request *GetPassportElementRequest) (PassportElementClass, error) {
	var result PassportElementBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.PassportElement, nil
}
