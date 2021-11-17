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

// SetRecoveryEmailAddressRequest represents TL type `setRecoveryEmailAddress#89df939f`.
type SetRecoveryEmailAddressRequest struct {
	// Password field of SetRecoveryEmailAddressRequest.
	Password string
	// NewRecoveryEmailAddress field of SetRecoveryEmailAddressRequest.
	NewRecoveryEmailAddress string
}

// SetRecoveryEmailAddressRequestTypeID is TL type id of SetRecoveryEmailAddressRequest.
const SetRecoveryEmailAddressRequestTypeID = 0x89df939f

// Ensuring interfaces in compile-time for SetRecoveryEmailAddressRequest.
var (
	_ bin.Encoder     = &SetRecoveryEmailAddressRequest{}
	_ bin.Decoder     = &SetRecoveryEmailAddressRequest{}
	_ bin.BareEncoder = &SetRecoveryEmailAddressRequest{}
	_ bin.BareDecoder = &SetRecoveryEmailAddressRequest{}
)

func (s *SetRecoveryEmailAddressRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Password == "") {
		return false
	}
	if !(s.NewRecoveryEmailAddress == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetRecoveryEmailAddressRequest) String() string {
	if s == nil {
		return "SetRecoveryEmailAddressRequest(nil)"
	}
	type Alias SetRecoveryEmailAddressRequest
	return fmt.Sprintf("SetRecoveryEmailAddressRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetRecoveryEmailAddressRequest) TypeID() uint32 {
	return SetRecoveryEmailAddressRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetRecoveryEmailAddressRequest) TypeName() string {
	return "setRecoveryEmailAddress"
}

// TypeInfo returns info about TL type.
func (s *SetRecoveryEmailAddressRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setRecoveryEmailAddress",
		ID:   SetRecoveryEmailAddressRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Password",
			SchemaName: "password",
		},
		{
			Name:       "NewRecoveryEmailAddress",
			SchemaName: "new_recovery_email_address",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetRecoveryEmailAddressRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setRecoveryEmailAddress#89df939f as nil")
	}
	b.PutID(SetRecoveryEmailAddressRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetRecoveryEmailAddressRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setRecoveryEmailAddress#89df939f as nil")
	}
	b.PutString(s.Password)
	b.PutString(s.NewRecoveryEmailAddress)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetRecoveryEmailAddressRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setRecoveryEmailAddress#89df939f to nil")
	}
	if err := b.ConsumeID(SetRecoveryEmailAddressRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setRecoveryEmailAddress#89df939f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetRecoveryEmailAddressRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setRecoveryEmailAddress#89df939f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setRecoveryEmailAddress#89df939f: field password: %w", err)
		}
		s.Password = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setRecoveryEmailAddress#89df939f: field new_recovery_email_address: %w", err)
		}
		s.NewRecoveryEmailAddress = value
	}
	return nil
}

// GetPassword returns value of Password field.
func (s *SetRecoveryEmailAddressRequest) GetPassword() (value string) {
	return s.Password
}

// GetNewRecoveryEmailAddress returns value of NewRecoveryEmailAddress field.
func (s *SetRecoveryEmailAddressRequest) GetNewRecoveryEmailAddress() (value string) {
	return s.NewRecoveryEmailAddress
}

// SetRecoveryEmailAddress invokes method setRecoveryEmailAddress#89df939f returning error if any.
func (c *Client) SetRecoveryEmailAddress(ctx context.Context, request *SetRecoveryEmailAddressRequest) (*PasswordState, error) {
	var result PasswordState

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
