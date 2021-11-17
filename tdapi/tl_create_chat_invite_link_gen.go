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

// CreateChatInviteLinkRequest represents TL type `createChatInviteLink#2f832dfe`.
type CreateChatInviteLinkRequest struct {
	// Chat identifier
	ChatID int64
	// Point in time (Unix timestamp) when the link will expire; pass 0 if never
	ExpireDate int32
	// The maximum number of chat members that can join the chat by the link simultaneously;
	// 0-99999; pass 0 if not limited
	MemberLimit int32
}

// CreateChatInviteLinkRequestTypeID is TL type id of CreateChatInviteLinkRequest.
const CreateChatInviteLinkRequestTypeID = 0x2f832dfe

// Ensuring interfaces in compile-time for CreateChatInviteLinkRequest.
var (
	_ bin.Encoder     = &CreateChatInviteLinkRequest{}
	_ bin.Decoder     = &CreateChatInviteLinkRequest{}
	_ bin.BareEncoder = &CreateChatInviteLinkRequest{}
	_ bin.BareDecoder = &CreateChatInviteLinkRequest{}
)

func (c *CreateChatInviteLinkRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ChatID == 0) {
		return false
	}
	if !(c.ExpireDate == 0) {
		return false
	}
	if !(c.MemberLimit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CreateChatInviteLinkRequest) String() string {
	if c == nil {
		return "CreateChatInviteLinkRequest(nil)"
	}
	type Alias CreateChatInviteLinkRequest
	return fmt.Sprintf("CreateChatInviteLinkRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CreateChatInviteLinkRequest) TypeID() uint32 {
	return CreateChatInviteLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CreateChatInviteLinkRequest) TypeName() string {
	return "createChatInviteLink"
}

// TypeInfo returns info about TL type.
func (c *CreateChatInviteLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "createChatInviteLink",
		ID:   CreateChatInviteLinkRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "ExpireDate",
			SchemaName: "expire_date",
		},
		{
			Name:       "MemberLimit",
			SchemaName: "member_limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CreateChatInviteLinkRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createChatInviteLink#2f832dfe as nil")
	}
	b.PutID(CreateChatInviteLinkRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CreateChatInviteLinkRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createChatInviteLink#2f832dfe as nil")
	}
	b.PutLong(c.ChatID)
	b.PutInt32(c.ExpireDate)
	b.PutInt32(c.MemberLimit)
	return nil
}

// Decode implements bin.Decoder.
func (c *CreateChatInviteLinkRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createChatInviteLink#2f832dfe to nil")
	}
	if err := b.ConsumeID(CreateChatInviteLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode createChatInviteLink#2f832dfe: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CreateChatInviteLinkRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createChatInviteLink#2f832dfe to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode createChatInviteLink#2f832dfe: field chat_id: %w", err)
		}
		c.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode createChatInviteLink#2f832dfe: field expire_date: %w", err)
		}
		c.ExpireDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode createChatInviteLink#2f832dfe: field member_limit: %w", err)
		}
		c.MemberLimit = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (c *CreateChatInviteLinkRequest) GetChatID() (value int64) {
	return c.ChatID
}

// GetExpireDate returns value of ExpireDate field.
func (c *CreateChatInviteLinkRequest) GetExpireDate() (value int32) {
	return c.ExpireDate
}

// GetMemberLimit returns value of MemberLimit field.
func (c *CreateChatInviteLinkRequest) GetMemberLimit() (value int32) {
	return c.MemberLimit
}

// CreateChatInviteLink invokes method createChatInviteLink#2f832dfe returning error if any.
func (c *Client) CreateChatInviteLink(ctx context.Context, request *CreateChatInviteLinkRequest) (*ChatInviteLink, error) {
	var result ChatInviteLink

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
