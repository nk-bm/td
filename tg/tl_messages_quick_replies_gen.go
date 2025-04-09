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

// MessagesQuickReplies represents TL type `messages.quickReplies#c68d6695`.
type MessagesQuickReplies struct {
	// QuickReplies field of MessagesQuickReplies.
	QuickReplies []QuickReply
	// Messages field of MessagesQuickReplies.
	Messages []MessageClass
	// Chats field of MessagesQuickReplies.
	Chats []ChatClass
	// Users field of MessagesQuickReplies.
	Users []UserClass
}

// MessagesQuickRepliesTypeID is TL type id of MessagesQuickReplies.
const MessagesQuickRepliesTypeID = 0xc68d6695

// construct implements constructor of MessagesQuickRepliesClass.
func (q MessagesQuickReplies) construct() MessagesQuickRepliesClass { return &q }

// Ensuring interfaces in compile-time for MessagesQuickReplies.
var (
	_ bin.Encoder     = &MessagesQuickReplies{}
	_ bin.Decoder     = &MessagesQuickReplies{}
	_ bin.BareEncoder = &MessagesQuickReplies{}
	_ bin.BareDecoder = &MessagesQuickReplies{}

	_ MessagesQuickRepliesClass = &MessagesQuickReplies{}
)

func (q *MessagesQuickReplies) Zero() bool {
	if q == nil {
		return true
	}
	if !(q.QuickReplies == nil) {
		return false
	}
	if !(q.Messages == nil) {
		return false
	}
	if !(q.Chats == nil) {
		return false
	}
	if !(q.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (q *MessagesQuickReplies) String() string {
	if q == nil {
		return "MessagesQuickReplies(nil)"
	}
	type Alias MessagesQuickReplies
	return fmt.Sprintf("MessagesQuickReplies%+v", Alias(*q))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesQuickReplies) TypeID() uint32 {
	return MessagesQuickRepliesTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesQuickReplies) TypeName() string {
	return "messages.quickReplies"
}

// TypeInfo returns info about TL type.
func (q *MessagesQuickReplies) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.quickReplies",
		ID:   MessagesQuickRepliesTypeID,
	}
	if q == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "QuickReplies",
			SchemaName: "quick_replies",
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (q *MessagesQuickReplies) Encode(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't encode messages.quickReplies#c68d6695 as nil")
	}
	b.PutID(MessagesQuickRepliesTypeID)
	return q.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (q *MessagesQuickReplies) EncodeBare(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't encode messages.quickReplies#c68d6695 as nil")
	}
	b.PutVectorHeader(len(q.QuickReplies))
	for idx, v := range q.QuickReplies {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.quickReplies#c68d6695: field quick_replies element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(q.Messages))
	for idx, v := range q.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.quickReplies#c68d6695: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.quickReplies#c68d6695: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(q.Chats))
	for idx, v := range q.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.quickReplies#c68d6695: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.quickReplies#c68d6695: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(q.Users))
	for idx, v := range q.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.quickReplies#c68d6695: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.quickReplies#c68d6695: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (q *MessagesQuickReplies) Decode(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't decode messages.quickReplies#c68d6695 to nil")
	}
	if err := b.ConsumeID(MessagesQuickRepliesTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.quickReplies#c68d6695: %w", err)
	}
	return q.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (q *MessagesQuickReplies) DecodeBare(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't decode messages.quickReplies#c68d6695 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.quickReplies#c68d6695: field quick_replies: %w", err)
		}

		if headerLen > 0 {
			q.QuickReplies = make([]QuickReply, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value QuickReply
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.quickReplies#c68d6695: field quick_replies: %w", err)
			}
			q.QuickReplies = append(q.QuickReplies, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.quickReplies#c68d6695: field messages: %w", err)
		}

		if headerLen > 0 {
			q.Messages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.quickReplies#c68d6695: field messages: %w", err)
			}
			q.Messages = append(q.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.quickReplies#c68d6695: field chats: %w", err)
		}

		if headerLen > 0 {
			q.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.quickReplies#c68d6695: field chats: %w", err)
			}
			q.Chats = append(q.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.quickReplies#c68d6695: field users: %w", err)
		}

		if headerLen > 0 {
			q.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.quickReplies#c68d6695: field users: %w", err)
			}
			q.Users = append(q.Users, value)
		}
	}
	return nil
}

// GetQuickReplies returns value of QuickReplies field.
func (q *MessagesQuickReplies) GetQuickReplies() (value []QuickReply) {
	if q == nil {
		return
	}
	return q.QuickReplies
}

// GetMessages returns value of Messages field.
func (q *MessagesQuickReplies) GetMessages() (value []MessageClass) {
	if q == nil {
		return
	}
	return q.Messages
}

// GetChats returns value of Chats field.
func (q *MessagesQuickReplies) GetChats() (value []ChatClass) {
	if q == nil {
		return
	}
	return q.Chats
}

// GetUsers returns value of Users field.
func (q *MessagesQuickReplies) GetUsers() (value []UserClass) {
	if q == nil {
		return
	}
	return q.Users
}

// MessagesQuickRepliesNotModified represents TL type `messages.quickRepliesNotModified#5f91eb5b`.
type MessagesQuickRepliesNotModified struct {
}

// MessagesQuickRepliesNotModifiedTypeID is TL type id of MessagesQuickRepliesNotModified.
const MessagesQuickRepliesNotModifiedTypeID = 0x5f91eb5b

// construct implements constructor of MessagesQuickRepliesClass.
func (q MessagesQuickRepliesNotModified) construct() MessagesQuickRepliesClass { return &q }

// Ensuring interfaces in compile-time for MessagesQuickRepliesNotModified.
var (
	_ bin.Encoder     = &MessagesQuickRepliesNotModified{}
	_ bin.Decoder     = &MessagesQuickRepliesNotModified{}
	_ bin.BareEncoder = &MessagesQuickRepliesNotModified{}
	_ bin.BareDecoder = &MessagesQuickRepliesNotModified{}

	_ MessagesQuickRepliesClass = &MessagesQuickRepliesNotModified{}
)

func (q *MessagesQuickRepliesNotModified) Zero() bool {
	if q == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (q *MessagesQuickRepliesNotModified) String() string {
	if q == nil {
		return "MessagesQuickRepliesNotModified(nil)"
	}
	type Alias MessagesQuickRepliesNotModified
	return fmt.Sprintf("MessagesQuickRepliesNotModified%+v", Alias(*q))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesQuickRepliesNotModified) TypeID() uint32 {
	return MessagesQuickRepliesNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesQuickRepliesNotModified) TypeName() string {
	return "messages.quickRepliesNotModified"
}

// TypeInfo returns info about TL type.
func (q *MessagesQuickRepliesNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.quickRepliesNotModified",
		ID:   MessagesQuickRepliesNotModifiedTypeID,
	}
	if q == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (q *MessagesQuickRepliesNotModified) Encode(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't encode messages.quickRepliesNotModified#5f91eb5b as nil")
	}
	b.PutID(MessagesQuickRepliesNotModifiedTypeID)
	return q.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (q *MessagesQuickRepliesNotModified) EncodeBare(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't encode messages.quickRepliesNotModified#5f91eb5b as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (q *MessagesQuickRepliesNotModified) Decode(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't decode messages.quickRepliesNotModified#5f91eb5b to nil")
	}
	if err := b.ConsumeID(MessagesQuickRepliesNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.quickRepliesNotModified#5f91eb5b: %w", err)
	}
	return q.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (q *MessagesQuickRepliesNotModified) DecodeBare(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't decode messages.quickRepliesNotModified#5f91eb5b to nil")
	}
	return nil
}

// MessagesQuickRepliesClassName is schema name of MessagesQuickRepliesClass.
const MessagesQuickRepliesClassName = "messages.QuickReplies"

// MessagesQuickRepliesClass represents messages.QuickReplies generic type.
//
// Constructors:
//   - [MessagesQuickReplies]
//   - [MessagesQuickRepliesNotModified]
//
// Example:
//
//	g, err := tg.DecodeMessagesQuickReplies(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessagesQuickReplies: // messages.quickReplies#c68d6695
//	case *tg.MessagesQuickRepliesNotModified: // messages.quickRepliesNotModified#5f91eb5b
//	default: panic(v)
//	}
type MessagesQuickRepliesClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesQuickRepliesClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeMessagesQuickReplies implements binary de-serialization for MessagesQuickRepliesClass.
func DecodeMessagesQuickReplies(buf *bin.Buffer) (MessagesQuickRepliesClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesQuickRepliesTypeID:
		// Decoding messages.quickReplies#c68d6695.
		v := MessagesQuickReplies{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesQuickRepliesClass: %w", err)
		}
		return &v, nil
	case MessagesQuickRepliesNotModifiedTypeID:
		// Decoding messages.quickRepliesNotModified#5f91eb5b.
		v := MessagesQuickRepliesNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesQuickRepliesClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesQuickRepliesClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesQuickReplies boxes the MessagesQuickRepliesClass providing a helper.
type MessagesQuickRepliesBox struct {
	QuickReplies MessagesQuickRepliesClass
}

// Decode implements bin.Decoder for MessagesQuickRepliesBox.
func (b *MessagesQuickRepliesBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesQuickRepliesBox to nil")
	}
	v, err := DecodeMessagesQuickReplies(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.QuickReplies = v
	return nil
}

// Encode implements bin.Encode for MessagesQuickRepliesBox.
func (b *MessagesQuickRepliesBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.QuickReplies == nil {
		return fmt.Errorf("unable to encode MessagesQuickRepliesClass as nil")
	}
	return b.QuickReplies.Encode(buf)
}
