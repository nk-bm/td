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

// ChatlistsGetExportedInvitesRequest represents TL type `chatlists.getExportedInvites#ce03da83`.
// List all chat folder deep links »¹ associated to a folder
//
// Links:
//  1. https://core.telegram.org/api/links#chat-folder-links
//
// See https://core.telegram.org/method/chatlists.getExportedInvites for reference.
type ChatlistsGetExportedInvitesRequest struct {
	// The folder
	Chatlist InputChatlistDialogFilter
}

// ChatlistsGetExportedInvitesRequestTypeID is TL type id of ChatlistsGetExportedInvitesRequest.
const ChatlistsGetExportedInvitesRequestTypeID = 0xce03da83

// Ensuring interfaces in compile-time for ChatlistsGetExportedInvitesRequest.
var (
	_ bin.Encoder     = &ChatlistsGetExportedInvitesRequest{}
	_ bin.Decoder     = &ChatlistsGetExportedInvitesRequest{}
	_ bin.BareEncoder = &ChatlistsGetExportedInvitesRequest{}
	_ bin.BareDecoder = &ChatlistsGetExportedInvitesRequest{}
)

func (g *ChatlistsGetExportedInvitesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Chatlist.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChatlistsGetExportedInvitesRequest) String() string {
	if g == nil {
		return "ChatlistsGetExportedInvitesRequest(nil)"
	}
	type Alias ChatlistsGetExportedInvitesRequest
	return fmt.Sprintf("ChatlistsGetExportedInvitesRequest%+v", Alias(*g))
}

// FillFrom fills ChatlistsGetExportedInvitesRequest from given interface.
func (g *ChatlistsGetExportedInvitesRequest) FillFrom(from interface {
	GetChatlist() (value InputChatlistDialogFilter)
}) {
	g.Chatlist = from.GetChatlist()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatlistsGetExportedInvitesRequest) TypeID() uint32 {
	return ChatlistsGetExportedInvitesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatlistsGetExportedInvitesRequest) TypeName() string {
	return "chatlists.getExportedInvites"
}

// TypeInfo returns info about TL type.
func (g *ChatlistsGetExportedInvitesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatlists.getExportedInvites",
		ID:   ChatlistsGetExportedInvitesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Chatlist",
			SchemaName: "chatlist",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *ChatlistsGetExportedInvitesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode chatlists.getExportedInvites#ce03da83 as nil")
	}
	b.PutID(ChatlistsGetExportedInvitesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChatlistsGetExportedInvitesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode chatlists.getExportedInvites#ce03da83 as nil")
	}
	if err := g.Chatlist.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatlists.getExportedInvites#ce03da83: field chatlist: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChatlistsGetExportedInvitesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode chatlists.getExportedInvites#ce03da83 to nil")
	}
	if err := b.ConsumeID(ChatlistsGetExportedInvitesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode chatlists.getExportedInvites#ce03da83: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChatlistsGetExportedInvitesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode chatlists.getExportedInvites#ce03da83 to nil")
	}
	{
		if err := g.Chatlist.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatlists.getExportedInvites#ce03da83: field chatlist: %w", err)
		}
	}
	return nil
}

// GetChatlist returns value of Chatlist field.
func (g *ChatlistsGetExportedInvitesRequest) GetChatlist() (value InputChatlistDialogFilter) {
	if g == nil {
		return
	}
	return g.Chatlist
}

// ChatlistsGetExportedInvites invokes method chatlists.getExportedInvites#ce03da83 returning error if any.
// List all chat folder deep links »¹ associated to a folder
//
// Links:
//  1. https://core.telegram.org/api/links#chat-folder-links
//
// Possible errors:
//
//	400 FILTER_ID_INVALID: The specified filter ID is invalid.
//
// See https://core.telegram.org/method/chatlists.getExportedInvites for reference.
func (c *Client) ChatlistsGetExportedInvites(ctx context.Context, chatlist InputChatlistDialogFilter) (*ChatlistsExportedInvites, error) {
	var result ChatlistsExportedInvites

	request := &ChatlistsGetExportedInvitesRequest{
		Chatlist: chatlist,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
