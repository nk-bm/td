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

// SetBotUpdatesStatusRequest represents TL type `setBotUpdatesStatus#bb293991`.
type SetBotUpdatesStatusRequest struct {
	// The number of pending updates
	PendingUpdateCount int32
	// The last error message
	ErrorMessage string
}

// SetBotUpdatesStatusRequestTypeID is TL type id of SetBotUpdatesStatusRequest.
const SetBotUpdatesStatusRequestTypeID = 0xbb293991

// Ensuring interfaces in compile-time for SetBotUpdatesStatusRequest.
var (
	_ bin.Encoder     = &SetBotUpdatesStatusRequest{}
	_ bin.Decoder     = &SetBotUpdatesStatusRequest{}
	_ bin.BareEncoder = &SetBotUpdatesStatusRequest{}
	_ bin.BareDecoder = &SetBotUpdatesStatusRequest{}
)

func (s *SetBotUpdatesStatusRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PendingUpdateCount == 0) {
		return false
	}
	if !(s.ErrorMessage == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetBotUpdatesStatusRequest) String() string {
	if s == nil {
		return "SetBotUpdatesStatusRequest(nil)"
	}
	type Alias SetBotUpdatesStatusRequest
	return fmt.Sprintf("SetBotUpdatesStatusRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetBotUpdatesStatusRequest) TypeID() uint32 {
	return SetBotUpdatesStatusRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetBotUpdatesStatusRequest) TypeName() string {
	return "setBotUpdatesStatus"
}

// TypeInfo returns info about TL type.
func (s *SetBotUpdatesStatusRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setBotUpdatesStatus",
		ID:   SetBotUpdatesStatusRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PendingUpdateCount",
			SchemaName: "pending_update_count",
		},
		{
			Name:       "ErrorMessage",
			SchemaName: "error_message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetBotUpdatesStatusRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotUpdatesStatus#bb293991 as nil")
	}
	b.PutID(SetBotUpdatesStatusRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetBotUpdatesStatusRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotUpdatesStatus#bb293991 as nil")
	}
	b.PutInt32(s.PendingUpdateCount)
	b.PutString(s.ErrorMessage)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetBotUpdatesStatusRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotUpdatesStatus#bb293991 to nil")
	}
	if err := b.ConsumeID(SetBotUpdatesStatusRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setBotUpdatesStatus#bb293991: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetBotUpdatesStatusRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotUpdatesStatus#bb293991 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setBotUpdatesStatus#bb293991: field pending_update_count: %w", err)
		}
		s.PendingUpdateCount = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setBotUpdatesStatus#bb293991: field error_message: %w", err)
		}
		s.ErrorMessage = value
	}
	return nil
}

// GetPendingUpdateCount returns value of PendingUpdateCount field.
func (s *SetBotUpdatesStatusRequest) GetPendingUpdateCount() (value int32) {
	return s.PendingUpdateCount
}

// GetErrorMessage returns value of ErrorMessage field.
func (s *SetBotUpdatesStatusRequest) GetErrorMessage() (value string) {
	return s.ErrorMessage
}

// SetBotUpdatesStatus invokes method setBotUpdatesStatus#bb293991 returning error if any.
func (c *Client) SetBotUpdatesStatus(ctx context.Context, request *SetBotUpdatesStatusRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
