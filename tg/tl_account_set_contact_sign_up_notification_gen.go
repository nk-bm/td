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

// AccountSetContactSignUpNotificationRequest represents TL type `account.setContactSignUpNotification#cff43f61`.
// Toggle contact sign up notifications
//
// See https://core.telegram.org/method/account.setContactSignUpNotification for reference.
type AccountSetContactSignUpNotificationRequest struct {
	// Whether to disable contact sign up notifications
	Silent bool
}

// AccountSetContactSignUpNotificationRequestTypeID is TL type id of AccountSetContactSignUpNotificationRequest.
const AccountSetContactSignUpNotificationRequestTypeID = 0xcff43f61

func (s *AccountSetContactSignUpNotificationRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Silent == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSetContactSignUpNotificationRequest) String() string {
	if s == nil {
		return "AccountSetContactSignUpNotificationRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountSetContactSignUpNotificationRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tSilent: ")
	sb.WriteString(fmt.Sprint(s.Silent))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AccountSetContactSignUpNotificationRequest) TypeID() uint32 {
	return AccountSetContactSignUpNotificationRequestTypeID
}

// Encode implements bin.Encoder.
func (s *AccountSetContactSignUpNotificationRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setContactSignUpNotification#cff43f61 as nil")
	}
	b.PutID(AccountSetContactSignUpNotificationRequestTypeID)
	b.PutBool(s.Silent)
	return nil
}

// GetSilent returns value of Silent field.
func (s *AccountSetContactSignUpNotificationRequest) GetSilent() (value bool) {
	return s.Silent
}

// Decode implements bin.Decoder.
func (s *AccountSetContactSignUpNotificationRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setContactSignUpNotification#cff43f61 to nil")
	}
	if err := b.ConsumeID(AccountSetContactSignUpNotificationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.setContactSignUpNotification#cff43f61: %w", err)
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.setContactSignUpNotification#cff43f61: field silent: %w", err)
		}
		s.Silent = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSetContactSignUpNotificationRequest.
var (
	_ bin.Encoder = &AccountSetContactSignUpNotificationRequest{}
	_ bin.Decoder = &AccountSetContactSignUpNotificationRequest{}
)

// AccountSetContactSignUpNotification invokes method account.setContactSignUpNotification#cff43f61 returning error if any.
// Toggle contact sign up notifications
//
// See https://core.telegram.org/method/account.setContactSignUpNotification for reference.
func (c *Client) AccountSetContactSignUpNotification(ctx context.Context, silent bool) (bool, error) {
	var result BoolBox

	request := &AccountSetContactSignUpNotificationRequest{
		Silent: silent,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
