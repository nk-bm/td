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

// AuthSendCodeRequest represents TL type `auth.sendCode#a677244f`.
// Send the verification code for login
//
// See https://core.telegram.org/method/auth.sendCode for reference.
type AuthSendCodeRequest struct {
	// Phone number in international format
	PhoneNumber string
	// Application identifier (see App configuration¹)
	//
	// Links:
	//  1) https://core.telegram.org/myapp
	APIID int
	// Application secret hash (see App configuration¹)
	//
	// Links:
	//  1) https://core.telegram.org/myapp
	APIHash string
	// Settings for the code type to send
	Settings CodeSettings
}

// AuthSendCodeRequestTypeID is TL type id of AuthSendCodeRequest.
const AuthSendCodeRequestTypeID = 0xa677244f

// Ensuring interfaces in compile-time for AuthSendCodeRequest.
var (
	_ bin.Encoder     = &AuthSendCodeRequest{}
	_ bin.Decoder     = &AuthSendCodeRequest{}
	_ bin.BareEncoder = &AuthSendCodeRequest{}
	_ bin.BareDecoder = &AuthSendCodeRequest{}
)

func (s *AuthSendCodeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PhoneNumber == "") {
		return false
	}
	if !(s.APIID == 0) {
		return false
	}
	if !(s.APIHash == "") {
		return false
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AuthSendCodeRequest) String() string {
	if s == nil {
		return "AuthSendCodeRequest(nil)"
	}
	type Alias AuthSendCodeRequest
	return fmt.Sprintf("AuthSendCodeRequest%+v", Alias(*s))
}

// FillFrom fills AuthSendCodeRequest from given interface.
func (s *AuthSendCodeRequest) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetAPIID() (value int)
	GetAPIHash() (value string)
	GetSettings() (value CodeSettings)
}) {
	s.PhoneNumber = from.GetPhoneNumber()
	s.APIID = from.GetAPIID()
	s.APIHash = from.GetAPIHash()
	s.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthSendCodeRequest) TypeID() uint32 {
	return AuthSendCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthSendCodeRequest) TypeName() string {
	return "auth.sendCode"
}

// TypeInfo returns info about TL type.
func (s *AuthSendCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.sendCode",
		ID:   AuthSendCodeRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "APIID",
			SchemaName: "api_id",
		},
		{
			Name:       "APIHash",
			SchemaName: "api_hash",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AuthSendCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sendCode#a677244f as nil")
	}
	b.PutID(AuthSendCodeRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AuthSendCodeRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sendCode#a677244f as nil")
	}
	b.PutString(s.PhoneNumber)
	b.PutInt(s.APIID)
	b.PutString(s.APIHash)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.sendCode#a677244f: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AuthSendCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sendCode#a677244f to nil")
	}
	if err := b.ConsumeID(AuthSendCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sendCode#a677244f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AuthSendCodeRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sendCode#a677244f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sendCode#a677244f: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sendCode#a677244f: field api_id: %w", err)
		}
		s.APIID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sendCode#a677244f: field api_hash: %w", err)
		}
		s.APIHash = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.sendCode#a677244f: field settings: %w", err)
		}
	}
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (s *AuthSendCodeRequest) GetPhoneNumber() (value string) {
	if s == nil {
		return
	}
	return s.PhoneNumber
}

// GetAPIID returns value of APIID field.
func (s *AuthSendCodeRequest) GetAPIID() (value int) {
	if s == nil {
		return
	}
	return s.APIID
}

// GetAPIHash returns value of APIHash field.
func (s *AuthSendCodeRequest) GetAPIHash() (value string) {
	if s == nil {
		return
	}
	return s.APIHash
}

// GetSettings returns value of Settings field.
func (s *AuthSendCodeRequest) GetSettings() (value CodeSettings) {
	if s == nil {
		return
	}
	return s.Settings
}

// AuthSendCode invokes method auth.sendCode#a677244f returning error if any.
// Send the verification code for login
//
// Possible errors:
//
//	400 API_ID_INVALID: API ID invalid.
//	400 API_ID_PUBLISHED_FLOOD: This API id was published somewhere, you can't use it now.
//	500 AUTH_RESTART: Restart the authorization process.
//	500 AUTH_RESTART_%d: Internal error (debug info %d), please repeat the method call.
//	400 PHONE_NUMBER_APP_SIGNUP_FORBIDDEN: You can't sign up using this app.
//	400 PHONE_NUMBER_BANNED: The provided phone number is banned from telegram.
//	400 PHONE_NUMBER_FLOOD: You asked for the code too many times.
//	406 PHONE_NUMBER_INVALID: The phone number is invalid.
//	406 PHONE_PASSWORD_FLOOD: You have tried logging in too many times.
//	400 PHONE_PASSWORD_PROTECTED: This phone is password protected.
//	400 SMS_CODE_CREATE_FAILED: An error occurred while creating the SMS code.
//	406 UPDATE_APP_TO_LOGIN: Please update your client to login.
//
// See https://core.telegram.org/method/auth.sendCode for reference.
func (c *Client) AuthSendCode(ctx context.Context, request *AuthSendCodeRequest) (AuthSentCodeClass, error) {
	var result AuthSentCodeBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.SentCode, nil
}
