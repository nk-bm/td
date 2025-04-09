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

// AuthCheckPasswordRequest represents TL type `auth.checkPassword#d18b4d16`.
type AuthCheckPasswordRequest struct {
	// Password field of AuthCheckPasswordRequest.
	Password InputCheckPasswordSRPClass
}

// AuthCheckPasswordRequestTypeID is TL type id of AuthCheckPasswordRequest.
const AuthCheckPasswordRequestTypeID = 0xd18b4d16

// Ensuring interfaces in compile-time for AuthCheckPasswordRequest.
var (
	_ bin.Encoder     = &AuthCheckPasswordRequest{}
	_ bin.Decoder     = &AuthCheckPasswordRequest{}
	_ bin.BareEncoder = &AuthCheckPasswordRequest{}
	_ bin.BareDecoder = &AuthCheckPasswordRequest{}
)

func (c *AuthCheckPasswordRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Password == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AuthCheckPasswordRequest) String() string {
	if c == nil {
		return "AuthCheckPasswordRequest(nil)"
	}
	type Alias AuthCheckPasswordRequest
	return fmt.Sprintf("AuthCheckPasswordRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthCheckPasswordRequest) TypeID() uint32 {
	return AuthCheckPasswordRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthCheckPasswordRequest) TypeName() string {
	return "auth.checkPassword"
}

// TypeInfo returns info about TL type.
func (c *AuthCheckPasswordRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.checkPassword",
		ID:   AuthCheckPasswordRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *AuthCheckPasswordRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.checkPassword#d18b4d16 as nil")
	}
	b.PutID(AuthCheckPasswordRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *AuthCheckPasswordRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.checkPassword#d18b4d16 as nil")
	}
	if c.Password == nil {
		return fmt.Errorf("unable to encode auth.checkPassword#d18b4d16: field password is nil")
	}
	if err := c.Password.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.checkPassword#d18b4d16: field password: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *AuthCheckPasswordRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.checkPassword#d18b4d16 to nil")
	}
	if err := b.ConsumeID(AuthCheckPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.checkPassword#d18b4d16: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *AuthCheckPasswordRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.checkPassword#d18b4d16 to nil")
	}
	{
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode auth.checkPassword#d18b4d16: field password: %w", err)
		}
		c.Password = value
	}
	return nil
}

// GetPassword returns value of Password field.
func (c *AuthCheckPasswordRequest) GetPassword() (value InputCheckPasswordSRPClass) {
	if c == nil {
		return
	}
	return c.Password
}

// AuthCheckPassword invokes method auth.checkPassword#d18b4d16 returning error if any.
func (c *Client) AuthCheckPassword(ctx context.Context, password InputCheckPasswordSRPClass) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox

	request := &AuthCheckPasswordRequest{
		Password: password,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}
