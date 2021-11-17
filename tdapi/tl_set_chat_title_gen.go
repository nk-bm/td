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

// SetChatTitleRequest represents TL type `setChatTitle#9cabebf`.
type SetChatTitleRequest struct {
	// Chat identifier
	ChatID int64
	// New title of the chat; 1-128 characters
	Title string
}

// SetChatTitleRequestTypeID is TL type id of SetChatTitleRequest.
const SetChatTitleRequestTypeID = 0x9cabebf

// Ensuring interfaces in compile-time for SetChatTitleRequest.
var (
	_ bin.Encoder     = &SetChatTitleRequest{}
	_ bin.Decoder     = &SetChatTitleRequest{}
	_ bin.BareEncoder = &SetChatTitleRequest{}
	_ bin.BareDecoder = &SetChatTitleRequest{}
)

func (s *SetChatTitleRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatTitleRequest) String() string {
	if s == nil {
		return "SetChatTitleRequest(nil)"
	}
	type Alias SetChatTitleRequest
	return fmt.Sprintf("SetChatTitleRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatTitleRequest) TypeID() uint32 {
	return SetChatTitleRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatTitleRequest) TypeName() string {
	return "setChatTitle"
}

// TypeInfo returns info about TL type.
func (s *SetChatTitleRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatTitle",
		ID:   SetChatTitleRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatTitleRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatTitle#9cabebf as nil")
	}
	b.PutID(SetChatTitleRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatTitleRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatTitle#9cabebf as nil")
	}
	b.PutLong(s.ChatID)
	b.PutString(s.Title)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatTitleRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatTitle#9cabebf to nil")
	}
	if err := b.ConsumeID(SetChatTitleRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatTitle#9cabebf: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatTitleRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatTitle#9cabebf to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode setChatTitle#9cabebf: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setChatTitle#9cabebf: field title: %w", err)
		}
		s.Title = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (s *SetChatTitleRequest) GetChatID() (value int64) {
	return s.ChatID
}

// GetTitle returns value of Title field.
func (s *SetChatTitleRequest) GetTitle() (value string) {
	return s.Title
}

// SetChatTitle invokes method setChatTitle#9cabebf returning error if any.
func (c *Client) SetChatTitle(ctx context.Context, request *SetChatTitleRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
