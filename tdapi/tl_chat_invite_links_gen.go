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

// ChatInviteLinks represents TL type `chatInviteLinks#9b1eddfa`.
type ChatInviteLinks struct {
	// Approximate total count of chat invite links found
	TotalCount int32
	// List of invite links
	InviteLinks []ChatInviteLink
}

// ChatInviteLinksTypeID is TL type id of ChatInviteLinks.
const ChatInviteLinksTypeID = 0x9b1eddfa

// Ensuring interfaces in compile-time for ChatInviteLinks.
var (
	_ bin.Encoder     = &ChatInviteLinks{}
	_ bin.Decoder     = &ChatInviteLinks{}
	_ bin.BareEncoder = &ChatInviteLinks{}
	_ bin.BareDecoder = &ChatInviteLinks{}
)

func (c *ChatInviteLinks) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.TotalCount == 0) {
		return false
	}
	if !(c.InviteLinks == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteLinks) String() string {
	if c == nil {
		return "ChatInviteLinks(nil)"
	}
	type Alias ChatInviteLinks
	return fmt.Sprintf("ChatInviteLinks%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInviteLinks) TypeID() uint32 {
	return ChatInviteLinksTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInviteLinks) TypeName() string {
	return "chatInviteLinks"
}

// TypeInfo returns info about TL type.
func (c *ChatInviteLinks) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInviteLinks",
		ID:   ChatInviteLinksTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "InviteLinks",
			SchemaName: "invite_links",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatInviteLinks) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinks#9b1eddfa as nil")
	}
	b.PutID(ChatInviteLinksTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInviteLinks) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinks#9b1eddfa as nil")
	}
	b.PutInt32(c.TotalCount)
	b.PutInt(len(c.InviteLinks))
	for idx, v := range c.InviteLinks {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatInviteLinks#9b1eddfa: field invite_links element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInviteLinks) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinks#9b1eddfa to nil")
	}
	if err := b.ConsumeID(ChatInviteLinksTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteLinks#9b1eddfa: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInviteLinks) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinks#9b1eddfa to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinks#9b1eddfa: field total_count: %w", err)
		}
		c.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinks#9b1eddfa: field invite_links: %w", err)
		}

		if headerLen > 0 {
			c.InviteLinks = make([]ChatInviteLink, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatInviteLink
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare chatInviteLinks#9b1eddfa: field invite_links: %w", err)
			}
			c.InviteLinks = append(c.InviteLinks, value)
		}
	}
	return nil
}

// GetTotalCount returns value of TotalCount field.
func (c *ChatInviteLinks) GetTotalCount() (value int32) {
	return c.TotalCount
}

// GetInviteLinks returns value of InviteLinks field.
func (c *ChatInviteLinks) GetInviteLinks() (value []ChatInviteLink) {
	return c.InviteLinks
}
