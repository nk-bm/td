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

// MessagesCreateChatRequest represents TL type `messages.createChat#92ceddd4`.
type MessagesCreateChatRequest struct {
	// Flags field of MessagesCreateChatRequest.
	Flags bin.Fields
	// Users field of MessagesCreateChatRequest.
	Users []InputUserClass
	// Title field of MessagesCreateChatRequest.
	Title string
	// TTLPeriod field of MessagesCreateChatRequest.
	//
	// Use SetTTLPeriod and GetTTLPeriod helpers.
	TTLPeriod int
}

// MessagesCreateChatRequestTypeID is TL type id of MessagesCreateChatRequest.
const MessagesCreateChatRequestTypeID = 0x92ceddd4

// Ensuring interfaces in compile-time for MessagesCreateChatRequest.
var (
	_ bin.Encoder     = &MessagesCreateChatRequest{}
	_ bin.Decoder     = &MessagesCreateChatRequest{}
	_ bin.BareEncoder = &MessagesCreateChatRequest{}
	_ bin.BareDecoder = &MessagesCreateChatRequest{}
)

func (c *MessagesCreateChatRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.TTLPeriod == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesCreateChatRequest) String() string {
	if c == nil {
		return "MessagesCreateChatRequest(nil)"
	}
	type Alias MessagesCreateChatRequest
	return fmt.Sprintf("MessagesCreateChatRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesCreateChatRequest) TypeID() uint32 {
	return MessagesCreateChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesCreateChatRequest) TypeName() string {
	return "messages.createChat"
}

// TypeInfo returns info about TL type.
func (c *MessagesCreateChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.createChat",
		ID:   MessagesCreateChatRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Users",
			SchemaName: "users",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "TTLPeriod",
			SchemaName: "ttl_period",
			Null:       !c.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *MessagesCreateChatRequest) SetFlags() {
	if !(c.TTLPeriod == 0) {
		c.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (c *MessagesCreateChatRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.createChat#92ceddd4 as nil")
	}
	b.PutID(MessagesCreateChatRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *MessagesCreateChatRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.createChat#92ceddd4 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.createChat#92ceddd4: field flags: %w", err)
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.createChat#92ceddd4: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.createChat#92ceddd4: field users element with index %d: %w", idx, err)
		}
	}
	b.PutString(c.Title)
	if c.Flags.Has(0) {
		b.PutInt(c.TTLPeriod)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesCreateChatRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.createChat#92ceddd4 to nil")
	}
	if err := b.ConsumeID(MessagesCreateChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.createChat#92ceddd4: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *MessagesCreateChatRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.createChat#92ceddd4 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.createChat#92ceddd4: field flags: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.createChat#92ceddd4: field users: %w", err)
		}

		if headerLen > 0 {
			c.Users = make([]InputUserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.createChat#92ceddd4: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.createChat#92ceddd4: field title: %w", err)
		}
		c.Title = value
	}
	if c.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.createChat#92ceddd4: field ttl_period: %w", err)
		}
		c.TTLPeriod = value
	}
	return nil
}

// GetUsers returns value of Users field.
func (c *MessagesCreateChatRequest) GetUsers() (value []InputUserClass) {
	if c == nil {
		return
	}
	return c.Users
}

// GetTitle returns value of Title field.
func (c *MessagesCreateChatRequest) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

// SetTTLPeriod sets value of TTLPeriod conditional field.
func (c *MessagesCreateChatRequest) SetTTLPeriod(value int) {
	c.Flags.Set(0)
	c.TTLPeriod = value
}

// GetTTLPeriod returns value of TTLPeriod conditional field and
// boolean which is true if field was set.
func (c *MessagesCreateChatRequest) GetTTLPeriod() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(0) {
		return value, false
	}
	return c.TTLPeriod, true
}

// MessagesCreateChat invokes method messages.createChat#92ceddd4 returning error if any.
func (c *Client) MessagesCreateChat(ctx context.Context, request *MessagesCreateChatRequest) (*MessagesInvitedUsers, error) {
	var result MessagesInvitedUsers

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
