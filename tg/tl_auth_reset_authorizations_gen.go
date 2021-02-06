// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// AuthResetAuthorizationsRequest represents TL type `auth.resetAuthorizations#9fab0d1a`.
// Terminates all user's authorized sessions except for the current one.
// After calling this method it is necessary to reregister the current device using the method account.registerDevice¹
//
// Links:
//  1) https://core.telegram.org/method/account.registerDevice
//
// See https://core.telegram.org/method/auth.resetAuthorizations for reference.
type AuthResetAuthorizationsRequest struct {
}

// AuthResetAuthorizationsRequestTypeID is TL type id of AuthResetAuthorizationsRequest.
const AuthResetAuthorizationsRequestTypeID = 0x9fab0d1a

func (r *AuthResetAuthorizationsRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *AuthResetAuthorizationsRequest) String() string {
	if r == nil {
		return "AuthResetAuthorizationsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthResetAuthorizationsRequest")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *AuthResetAuthorizationsRequest) TypeID() uint32 {
	return AuthResetAuthorizationsRequestTypeID
}

// Encode implements bin.Encoder.
func (r *AuthResetAuthorizationsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.resetAuthorizations#9fab0d1a as nil")
	}
	b.PutID(AuthResetAuthorizationsRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *AuthResetAuthorizationsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.resetAuthorizations#9fab0d1a to nil")
	}
	if err := b.ConsumeID(AuthResetAuthorizationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.resetAuthorizations#9fab0d1a: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthResetAuthorizationsRequest.
var (
	_ bin.Encoder = &AuthResetAuthorizationsRequest{}
	_ bin.Decoder = &AuthResetAuthorizationsRequest{}
)

// AuthResetAuthorizations invokes method auth.resetAuthorizations#9fab0d1a returning error if any.
// Terminates all user's authorized sessions except for the current one.
// After calling this method it is necessary to reregister the current device using the method account.registerDevice¹
//
// Links:
//  1) https://core.telegram.org/method/account.registerDevice
//
// Possible errors:
//  406 FRESH_RESET_AUTHORISATION_FORBIDDEN: You can't logout other sessions if less than 24 hours have passed since you logged on the current session
//
// See https://core.telegram.org/method/auth.resetAuthorizations for reference.
func (c *Client) AuthResetAuthorizations(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &AuthResetAuthorizationsRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
