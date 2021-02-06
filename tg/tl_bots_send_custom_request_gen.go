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

// BotsSendCustomRequestRequest represents TL type `bots.sendCustomRequest#aa2769ed`.
// Sends a custom request; for bots only
//
// See https://core.telegram.org/method/bots.sendCustomRequest for reference.
type BotsSendCustomRequestRequest struct {
	// The method name
	CustomMethod string
	// JSON-serialized method parameters
	Params DataJSON
}

// BotsSendCustomRequestRequestTypeID is TL type id of BotsSendCustomRequestRequest.
const BotsSendCustomRequestRequestTypeID = 0xaa2769ed

func (s *BotsSendCustomRequestRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.CustomMethod == "") {
		return false
	}
	if !(s.Params.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *BotsSendCustomRequestRequest) String() string {
	if s == nil {
		return "BotsSendCustomRequestRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("BotsSendCustomRequestRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tCustomMethod: ")
	sb.WriteString(fmt.Sprint(s.CustomMethod))
	sb.WriteString(",\n")
	sb.WriteString("\tParams: ")
	sb.WriteString(fmt.Sprint(s.Params))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *BotsSendCustomRequestRequest) TypeID() uint32 {
	return BotsSendCustomRequestRequestTypeID
}

// Encode implements bin.Encoder.
func (s *BotsSendCustomRequestRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode bots.sendCustomRequest#aa2769ed as nil")
	}
	b.PutID(BotsSendCustomRequestRequestTypeID)
	b.PutString(s.CustomMethod)
	if err := s.Params.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.sendCustomRequest#aa2769ed: field params: %w", err)
	}
	return nil
}

// GetCustomMethod returns value of CustomMethod field.
func (s *BotsSendCustomRequestRequest) GetCustomMethod() (value string) {
	return s.CustomMethod
}

// GetParams returns value of Params field.
func (s *BotsSendCustomRequestRequest) GetParams() (value DataJSON) {
	return s.Params
}

// Decode implements bin.Decoder.
func (s *BotsSendCustomRequestRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode bots.sendCustomRequest#aa2769ed to nil")
	}
	if err := b.ConsumeID(BotsSendCustomRequestRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.sendCustomRequest#aa2769ed: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.sendCustomRequest#aa2769ed: field custom_method: %w", err)
		}
		s.CustomMethod = value
	}
	{
		if err := s.Params.Decode(b); err != nil {
			return fmt.Errorf("unable to decode bots.sendCustomRequest#aa2769ed: field params: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for BotsSendCustomRequestRequest.
var (
	_ bin.Encoder = &BotsSendCustomRequestRequest{}
	_ bin.Decoder = &BotsSendCustomRequestRequest{}
)

// BotsSendCustomRequest invokes method bots.sendCustomRequest#aa2769ed returning error if any.
// Sends a custom request; for bots only
//
// Possible errors:
//  400 METHOD_INVALID: The specified method is invalid
//  400 USER_BOT_INVALID: This method can only be called by a bot
//
// See https://core.telegram.org/method/bots.sendCustomRequest for reference.
// Can be used by bots.
func (c *Client) BotsSendCustomRequest(ctx context.Context, request *BotsSendCustomRequestRequest) (*DataJSON, error) {
	var result DataJSON

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
