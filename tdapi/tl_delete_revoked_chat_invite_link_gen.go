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

// DeleteRevokedChatInviteLinkRequest represents TL type `deleteRevokedChatInviteLink#91270c7f`.
type DeleteRevokedChatInviteLinkRequest struct {
	// Chat identifier
	ChatID int64
	// Invite link to revoke
	InviteLink string
}

// DeleteRevokedChatInviteLinkRequestTypeID is TL type id of DeleteRevokedChatInviteLinkRequest.
const DeleteRevokedChatInviteLinkRequestTypeID = 0x91270c7f

// Ensuring interfaces in compile-time for DeleteRevokedChatInviteLinkRequest.
var (
	_ bin.Encoder     = &DeleteRevokedChatInviteLinkRequest{}
	_ bin.Decoder     = &DeleteRevokedChatInviteLinkRequest{}
	_ bin.BareEncoder = &DeleteRevokedChatInviteLinkRequest{}
	_ bin.BareDecoder = &DeleteRevokedChatInviteLinkRequest{}
)

func (d *DeleteRevokedChatInviteLinkRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ChatID == 0) {
		return false
	}
	if !(d.InviteLink == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteRevokedChatInviteLinkRequest) String() string {
	if d == nil {
		return "DeleteRevokedChatInviteLinkRequest(nil)"
	}
	type Alias DeleteRevokedChatInviteLinkRequest
	return fmt.Sprintf("DeleteRevokedChatInviteLinkRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteRevokedChatInviteLinkRequest) TypeID() uint32 {
	return DeleteRevokedChatInviteLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteRevokedChatInviteLinkRequest) TypeName() string {
	return "deleteRevokedChatInviteLink"
}

// TypeInfo returns info about TL type.
func (d *DeleteRevokedChatInviteLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteRevokedChatInviteLink",
		ID:   DeleteRevokedChatInviteLinkRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteRevokedChatInviteLinkRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteRevokedChatInviteLink#91270c7f as nil")
	}
	b.PutID(DeleteRevokedChatInviteLinkRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteRevokedChatInviteLinkRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteRevokedChatInviteLink#91270c7f as nil")
	}
	b.PutLong(d.ChatID)
	b.PutString(d.InviteLink)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteRevokedChatInviteLinkRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteRevokedChatInviteLink#91270c7f to nil")
	}
	if err := b.ConsumeID(DeleteRevokedChatInviteLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteRevokedChatInviteLink#91270c7f: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteRevokedChatInviteLinkRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteRevokedChatInviteLink#91270c7f to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode deleteRevokedChatInviteLink#91270c7f: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode deleteRevokedChatInviteLink#91270c7f: field invite_link: %w", err)
		}
		d.InviteLink = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (d *DeleteRevokedChatInviteLinkRequest) GetChatID() (value int64) {
	return d.ChatID
}

// GetInviteLink returns value of InviteLink field.
func (d *DeleteRevokedChatInviteLinkRequest) GetInviteLink() (value string) {
	return d.InviteLink
}

// DeleteRevokedChatInviteLink invokes method deleteRevokedChatInviteLink#91270c7f returning error if any.
func (c *Client) DeleteRevokedChatInviteLink(ctx context.Context, request *DeleteRevokedChatInviteLinkRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
