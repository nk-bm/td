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

// ChatInviteLinkCount represents TL type `chatInviteLinkCount#c03da0b9`.
type ChatInviteLinkCount struct {
	// Administrator's user identifier
	UserID int32
	// Number of active invite links
	InviteLinkCount int32
	// Number of revoked invite links
	RevokedInviteLinkCount int32
}

// ChatInviteLinkCountTypeID is TL type id of ChatInviteLinkCount.
const ChatInviteLinkCountTypeID = 0xc03da0b9

// Ensuring interfaces in compile-time for ChatInviteLinkCount.
var (
	_ bin.Encoder     = &ChatInviteLinkCount{}
	_ bin.Decoder     = &ChatInviteLinkCount{}
	_ bin.BareEncoder = &ChatInviteLinkCount{}
	_ bin.BareDecoder = &ChatInviteLinkCount{}
)

func (c *ChatInviteLinkCount) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.InviteLinkCount == 0) {
		return false
	}
	if !(c.RevokedInviteLinkCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteLinkCount) String() string {
	if c == nil {
		return "ChatInviteLinkCount(nil)"
	}
	type Alias ChatInviteLinkCount
	return fmt.Sprintf("ChatInviteLinkCount%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInviteLinkCount) TypeID() uint32 {
	return ChatInviteLinkCountTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInviteLinkCount) TypeName() string {
	return "chatInviteLinkCount"
}

// TypeInfo returns info about TL type.
func (c *ChatInviteLinkCount) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInviteLinkCount",
		ID:   ChatInviteLinkCountTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "InviteLinkCount",
			SchemaName: "invite_link_count",
		},
		{
			Name:       "RevokedInviteLinkCount",
			SchemaName: "revoked_invite_link_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatInviteLinkCount) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkCount#c03da0b9 as nil")
	}
	b.PutID(ChatInviteLinkCountTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInviteLinkCount) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkCount#c03da0b9 as nil")
	}
	b.PutInt32(c.UserID)
	b.PutInt32(c.InviteLinkCount)
	b.PutInt32(c.RevokedInviteLinkCount)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInviteLinkCount) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkCount#c03da0b9 to nil")
	}
	if err := b.ConsumeID(ChatInviteLinkCountTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteLinkCount#c03da0b9: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInviteLinkCount) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkCount#c03da0b9 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkCount#c03da0b9: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkCount#c03da0b9: field invite_link_count: %w", err)
		}
		c.InviteLinkCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkCount#c03da0b9: field revoked_invite_link_count: %w", err)
		}
		c.RevokedInviteLinkCount = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (c *ChatInviteLinkCount) GetUserID() (value int32) {
	return c.UserID
}

// GetInviteLinkCount returns value of InviteLinkCount field.
func (c *ChatInviteLinkCount) GetInviteLinkCount() (value int32) {
	return c.InviteLinkCount
}

// GetRevokedInviteLinkCount returns value of RevokedInviteLinkCount field.
func (c *ChatInviteLinkCount) GetRevokedInviteLinkCount() (value int32) {
	return c.RevokedInviteLinkCount
}
