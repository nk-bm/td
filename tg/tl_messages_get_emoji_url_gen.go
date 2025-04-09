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

// MessagesGetEmojiURLRequest represents TL type `messages.getEmojiURL#d5b10c26`.
type MessagesGetEmojiURLRequest struct {
	// LangCode field of MessagesGetEmojiURLRequest.
	LangCode string
}

// MessagesGetEmojiURLRequestTypeID is TL type id of MessagesGetEmojiURLRequest.
const MessagesGetEmojiURLRequestTypeID = 0xd5b10c26

// Ensuring interfaces in compile-time for MessagesGetEmojiURLRequest.
var (
	_ bin.Encoder     = &MessagesGetEmojiURLRequest{}
	_ bin.Decoder     = &MessagesGetEmojiURLRequest{}
	_ bin.BareEncoder = &MessagesGetEmojiURLRequest{}
	_ bin.BareDecoder = &MessagesGetEmojiURLRequest{}
)

func (g *MessagesGetEmojiURLRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetEmojiURLRequest) String() string {
	if g == nil {
		return "MessagesGetEmojiURLRequest(nil)"
	}
	type Alias MessagesGetEmojiURLRequest
	return fmt.Sprintf("MessagesGetEmojiURLRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetEmojiURLRequest) TypeID() uint32 {
	return MessagesGetEmojiURLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetEmojiURLRequest) TypeName() string {
	return "messages.getEmojiURL"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetEmojiURLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getEmojiURL",
		ID:   MessagesGetEmojiURLRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetEmojiURLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiURL#d5b10c26 as nil")
	}
	b.PutID(MessagesGetEmojiURLRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetEmojiURLRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiURL#d5b10c26 as nil")
	}
	b.PutString(g.LangCode)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetEmojiURLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiURL#d5b10c26 to nil")
	}
	if err := b.ConsumeID(MessagesGetEmojiURLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getEmojiURL#d5b10c26: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetEmojiURLRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiURL#d5b10c26 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiURL#d5b10c26: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	return nil
}

// GetLangCode returns value of LangCode field.
func (g *MessagesGetEmojiURLRequest) GetLangCode() (value string) {
	if g == nil {
		return
	}
	return g.LangCode
}

// MessagesGetEmojiURL invokes method messages.getEmojiURL#d5b10c26 returning error if any.
func (c *Client) MessagesGetEmojiURL(ctx context.Context, langcode string) (*EmojiURL, error) {
	var result EmojiURL

	request := &MessagesGetEmojiURLRequest{
		LangCode: langcode,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
