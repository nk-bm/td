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

// ChatlistsCheckChatlistInviteRequest represents TL type `chatlists.checkChatlistInvite#41c10fff`.
type ChatlistsCheckChatlistInviteRequest struct {
	// Slug field of ChatlistsCheckChatlistInviteRequest.
	Slug string
}

// ChatlistsCheckChatlistInviteRequestTypeID is TL type id of ChatlistsCheckChatlistInviteRequest.
const ChatlistsCheckChatlistInviteRequestTypeID = 0x41c10fff

// Ensuring interfaces in compile-time for ChatlistsCheckChatlistInviteRequest.
var (
	_ bin.Encoder     = &ChatlistsCheckChatlistInviteRequest{}
	_ bin.Decoder     = &ChatlistsCheckChatlistInviteRequest{}
	_ bin.BareEncoder = &ChatlistsCheckChatlistInviteRequest{}
	_ bin.BareDecoder = &ChatlistsCheckChatlistInviteRequest{}
)

func (c *ChatlistsCheckChatlistInviteRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Slug == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatlistsCheckChatlistInviteRequest) String() string {
	if c == nil {
		return "ChatlistsCheckChatlistInviteRequest(nil)"
	}
	type Alias ChatlistsCheckChatlistInviteRequest
	return fmt.Sprintf("ChatlistsCheckChatlistInviteRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatlistsCheckChatlistInviteRequest) TypeID() uint32 {
	return ChatlistsCheckChatlistInviteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatlistsCheckChatlistInviteRequest) TypeName() string {
	return "chatlists.checkChatlistInvite"
}

// TypeInfo returns info about TL type.
func (c *ChatlistsCheckChatlistInviteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatlists.checkChatlistInvite",
		ID:   ChatlistsCheckChatlistInviteRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Slug",
			SchemaName: "slug",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatlistsCheckChatlistInviteRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatlists.checkChatlistInvite#41c10fff as nil")
	}
	b.PutID(ChatlistsCheckChatlistInviteRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatlistsCheckChatlistInviteRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatlists.checkChatlistInvite#41c10fff as nil")
	}
	b.PutString(c.Slug)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatlistsCheckChatlistInviteRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatlists.checkChatlistInvite#41c10fff to nil")
	}
	if err := b.ConsumeID(ChatlistsCheckChatlistInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode chatlists.checkChatlistInvite#41c10fff: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatlistsCheckChatlistInviteRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatlists.checkChatlistInvite#41c10fff to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.checkChatlistInvite#41c10fff: field slug: %w", err)
		}
		c.Slug = value
	}
	return nil
}

// GetSlug returns value of Slug field.
func (c *ChatlistsCheckChatlistInviteRequest) GetSlug() (value string) {
	if c == nil {
		return
	}
	return c.Slug
}

// ChatlistsCheckChatlistInvite invokes method chatlists.checkChatlistInvite#41c10fff returning error if any.
func (c *Client) ChatlistsCheckChatlistInvite(ctx context.Context, slug string) (ChatlistsChatlistInviteClass, error) {
	var result ChatlistsChatlistInviteBox

	request := &ChatlistsCheckChatlistInviteRequest{
		Slug: slug,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ChatlistInvite, nil
}
