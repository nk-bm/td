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

// GetChatInviteLinksRequest represents TL type `getChatInviteLinks#94a0fd90`.
type GetChatInviteLinksRequest struct {
	// Chat identifier
	ChatID int64
	// User identifier of a chat administrator. Must be an identifier of the current user for
	// non-owner
	CreatorUserID int32
	// Pass true if revoked links needs to be returned instead of active or expired
	IsRevoked bool
	// Creation date of an invite link starting after which to return invite links; use 0 to
	// get results from the beginning
	OffsetDate int32
	// Invite link starting after which to return invite links; use empty string to get
	// results from the beginning
	OffsetInviteLink string
	// The maximum number of invite links to return
	Limit int32
}

// GetChatInviteLinksRequestTypeID is TL type id of GetChatInviteLinksRequest.
const GetChatInviteLinksRequestTypeID = 0x94a0fd90

// Ensuring interfaces in compile-time for GetChatInviteLinksRequest.
var (
	_ bin.Encoder     = &GetChatInviteLinksRequest{}
	_ bin.Decoder     = &GetChatInviteLinksRequest{}
	_ bin.BareEncoder = &GetChatInviteLinksRequest{}
	_ bin.BareDecoder = &GetChatInviteLinksRequest{}
)

func (g *GetChatInviteLinksRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.CreatorUserID == 0) {
		return false
	}
	if !(g.IsRevoked == false) {
		return false
	}
	if !(g.OffsetDate == 0) {
		return false
	}
	if !(g.OffsetInviteLink == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatInviteLinksRequest) String() string {
	if g == nil {
		return "GetChatInviteLinksRequest(nil)"
	}
	type Alias GetChatInviteLinksRequest
	return fmt.Sprintf("GetChatInviteLinksRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatInviteLinksRequest) TypeID() uint32 {
	return GetChatInviteLinksRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatInviteLinksRequest) TypeName() string {
	return "getChatInviteLinks"
}

// TypeInfo returns info about TL type.
func (g *GetChatInviteLinksRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatInviteLinks",
		ID:   GetChatInviteLinksRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "CreatorUserID",
			SchemaName: "creator_user_id",
		},
		{
			Name:       "IsRevoked",
			SchemaName: "is_revoked",
		},
		{
			Name:       "OffsetDate",
			SchemaName: "offset_date",
		},
		{
			Name:       "OffsetInviteLink",
			SchemaName: "offset_invite_link",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatInviteLinksRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLinks#94a0fd90 as nil")
	}
	b.PutID(GetChatInviteLinksRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatInviteLinksRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLinks#94a0fd90 as nil")
	}
	b.PutLong(g.ChatID)
	b.PutInt32(g.CreatorUserID)
	b.PutBool(g.IsRevoked)
	b.PutInt32(g.OffsetDate)
	b.PutString(g.OffsetInviteLink)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatInviteLinksRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLinks#94a0fd90 to nil")
	}
	if err := b.ConsumeID(GetChatInviteLinksRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatInviteLinks#94a0fd90: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatInviteLinksRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLinks#94a0fd90 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinks#94a0fd90: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinks#94a0fd90: field creator_user_id: %w", err)
		}
		g.CreatorUserID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinks#94a0fd90: field is_revoked: %w", err)
		}
		g.IsRevoked = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinks#94a0fd90: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinks#94a0fd90: field offset_invite_link: %w", err)
		}
		g.OffsetInviteLink = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinks#94a0fd90: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (g *GetChatInviteLinksRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetCreatorUserID returns value of CreatorUserID field.
func (g *GetChatInviteLinksRequest) GetCreatorUserID() (value int32) {
	return g.CreatorUserID
}

// GetIsRevoked returns value of IsRevoked field.
func (g *GetChatInviteLinksRequest) GetIsRevoked() (value bool) {
	return g.IsRevoked
}

// GetOffsetDate returns value of OffsetDate field.
func (g *GetChatInviteLinksRequest) GetOffsetDate() (value int32) {
	return g.OffsetDate
}

// GetOffsetInviteLink returns value of OffsetInviteLink field.
func (g *GetChatInviteLinksRequest) GetOffsetInviteLink() (value string) {
	return g.OffsetInviteLink
}

// GetLimit returns value of Limit field.
func (g *GetChatInviteLinksRequest) GetLimit() (value int32) {
	return g.Limit
}

// GetChatInviteLinks invokes method getChatInviteLinks#94a0fd90 returning error if any.
func (c *Client) GetChatInviteLinks(ctx context.Context, request *GetChatInviteLinksRequest) (*ChatInviteLinks, error) {
	var result ChatInviteLinks

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
