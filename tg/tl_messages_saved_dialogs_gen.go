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

// MessagesSavedDialogs represents TL type `messages.savedDialogs#f83ae221`.
type MessagesSavedDialogs struct {
	// Dialogs field of MessagesSavedDialogs.
	Dialogs []SavedDialog
	// Messages field of MessagesSavedDialogs.
	Messages []MessageClass
	// Chats field of MessagesSavedDialogs.
	Chats []ChatClass
	// Users field of MessagesSavedDialogs.
	Users []UserClass
}

// MessagesSavedDialogsTypeID is TL type id of MessagesSavedDialogs.
const MessagesSavedDialogsTypeID = 0xf83ae221

// construct implements constructor of MessagesSavedDialogsClass.
func (s MessagesSavedDialogs) construct() MessagesSavedDialogsClass { return &s }

// Ensuring interfaces in compile-time for MessagesSavedDialogs.
var (
	_ bin.Encoder     = &MessagesSavedDialogs{}
	_ bin.Decoder     = &MessagesSavedDialogs{}
	_ bin.BareEncoder = &MessagesSavedDialogs{}
	_ bin.BareDecoder = &MessagesSavedDialogs{}

	_ MessagesSavedDialogsClass = &MessagesSavedDialogs{}
)

func (s *MessagesSavedDialogs) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Dialogs == nil) {
		return false
	}
	if !(s.Messages == nil) {
		return false
	}
	if !(s.Chats == nil) {
		return false
	}
	if !(s.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSavedDialogs) String() string {
	if s == nil {
		return "MessagesSavedDialogs(nil)"
	}
	type Alias MessagesSavedDialogs
	return fmt.Sprintf("MessagesSavedDialogs%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSavedDialogs) TypeID() uint32 {
	return MessagesSavedDialogsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSavedDialogs) TypeName() string {
	return "messages.savedDialogs"
}

// TypeInfo returns info about TL type.
func (s *MessagesSavedDialogs) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.savedDialogs",
		ID:   MessagesSavedDialogsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Dialogs",
			SchemaName: "dialogs",
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
func (s *MessagesSavedDialogs) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.savedDialogs#f83ae221 as nil")
	}
	b.PutID(MessagesSavedDialogsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSavedDialogs) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.savedDialogs#f83ae221 as nil")
	}
	b.PutVectorHeader(len(s.Dialogs))
	for idx, v := range s.Dialogs {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.savedDialogs#f83ae221: field dialogs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Messages))
	for idx, v := range s.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.savedDialogs#f83ae221: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.savedDialogs#f83ae221: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Chats))
	for idx, v := range s.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.savedDialogs#f83ae221: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.savedDialogs#f83ae221: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Users))
	for idx, v := range s.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.savedDialogs#f83ae221: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.savedDialogs#f83ae221: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSavedDialogs) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.savedDialogs#f83ae221 to nil")
	}
	if err := b.ConsumeID(MessagesSavedDialogsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.savedDialogs#f83ae221: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSavedDialogs) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.savedDialogs#f83ae221 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedDialogs#f83ae221: field dialogs: %w", err)
		}

		if headerLen > 0 {
			s.Dialogs = make([]SavedDialog, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SavedDialog
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.savedDialogs#f83ae221: field dialogs: %w", err)
			}
			s.Dialogs = append(s.Dialogs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedDialogs#f83ae221: field messages: %w", err)
		}

		if headerLen > 0 {
			s.Messages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.savedDialogs#f83ae221: field messages: %w", err)
			}
			s.Messages = append(s.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedDialogs#f83ae221: field chats: %w", err)
		}

		if headerLen > 0 {
			s.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.savedDialogs#f83ae221: field chats: %w", err)
			}
			s.Chats = append(s.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedDialogs#f83ae221: field users: %w", err)
		}

		if headerLen > 0 {
			s.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.savedDialogs#f83ae221: field users: %w", err)
			}
			s.Users = append(s.Users, value)
		}
	}
	return nil
}

// GetDialogs returns value of Dialogs field.
func (s *MessagesSavedDialogs) GetDialogs() (value []SavedDialog) {
	if s == nil {
		return
	}
	return s.Dialogs
}

// GetMessages returns value of Messages field.
func (s *MessagesSavedDialogs) GetMessages() (value []MessageClass) {
	if s == nil {
		return
	}
	return s.Messages
}

// GetChats returns value of Chats field.
func (s *MessagesSavedDialogs) GetChats() (value []ChatClass) {
	if s == nil {
		return
	}
	return s.Chats
}

// GetUsers returns value of Users field.
func (s *MessagesSavedDialogs) GetUsers() (value []UserClass) {
	if s == nil {
		return
	}
	return s.Users
}

// MessagesSavedDialogsSlice represents TL type `messages.savedDialogsSlice#44ba9dd9`.
type MessagesSavedDialogsSlice struct {
	// Count field of MessagesSavedDialogsSlice.
	Count int
	// Dialogs field of MessagesSavedDialogsSlice.
	Dialogs []SavedDialog
	// Messages field of MessagesSavedDialogsSlice.
	Messages []MessageClass
	// Chats field of MessagesSavedDialogsSlice.
	Chats []ChatClass
	// Users field of MessagesSavedDialogsSlice.
	Users []UserClass
}

// MessagesSavedDialogsSliceTypeID is TL type id of MessagesSavedDialogsSlice.
const MessagesSavedDialogsSliceTypeID = 0x44ba9dd9

// construct implements constructor of MessagesSavedDialogsClass.
func (s MessagesSavedDialogsSlice) construct() MessagesSavedDialogsClass { return &s }

// Ensuring interfaces in compile-time for MessagesSavedDialogsSlice.
var (
	_ bin.Encoder     = &MessagesSavedDialogsSlice{}
	_ bin.Decoder     = &MessagesSavedDialogsSlice{}
	_ bin.BareEncoder = &MessagesSavedDialogsSlice{}
	_ bin.BareDecoder = &MessagesSavedDialogsSlice{}

	_ MessagesSavedDialogsClass = &MessagesSavedDialogsSlice{}
)

func (s *MessagesSavedDialogsSlice) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Count == 0) {
		return false
	}
	if !(s.Dialogs == nil) {
		return false
	}
	if !(s.Messages == nil) {
		return false
	}
	if !(s.Chats == nil) {
		return false
	}
	if !(s.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSavedDialogsSlice) String() string {
	if s == nil {
		return "MessagesSavedDialogsSlice(nil)"
	}
	type Alias MessagesSavedDialogsSlice
	return fmt.Sprintf("MessagesSavedDialogsSlice%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSavedDialogsSlice) TypeID() uint32 {
	return MessagesSavedDialogsSliceTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSavedDialogsSlice) TypeName() string {
	return "messages.savedDialogsSlice"
}

// TypeInfo returns info about TL type.
func (s *MessagesSavedDialogsSlice) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.savedDialogsSlice",
		ID:   MessagesSavedDialogsSliceTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Dialogs",
			SchemaName: "dialogs",
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
func (s *MessagesSavedDialogsSlice) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.savedDialogsSlice#44ba9dd9 as nil")
	}
	b.PutID(MessagesSavedDialogsSliceTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSavedDialogsSlice) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.savedDialogsSlice#44ba9dd9 as nil")
	}
	b.PutInt(s.Count)
	b.PutVectorHeader(len(s.Dialogs))
	for idx, v := range s.Dialogs {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.savedDialogsSlice#44ba9dd9: field dialogs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Messages))
	for idx, v := range s.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.savedDialogsSlice#44ba9dd9: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.savedDialogsSlice#44ba9dd9: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Chats))
	for idx, v := range s.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.savedDialogsSlice#44ba9dd9: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.savedDialogsSlice#44ba9dd9: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Users))
	for idx, v := range s.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.savedDialogsSlice#44ba9dd9: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.savedDialogsSlice#44ba9dd9: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSavedDialogsSlice) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.savedDialogsSlice#44ba9dd9 to nil")
	}
	if err := b.ConsumeID(MessagesSavedDialogsSliceTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.savedDialogsSlice#44ba9dd9: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSavedDialogsSlice) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.savedDialogsSlice#44ba9dd9 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedDialogsSlice#44ba9dd9: field count: %w", err)
		}
		s.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedDialogsSlice#44ba9dd9: field dialogs: %w", err)
		}

		if headerLen > 0 {
			s.Dialogs = make([]SavedDialog, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SavedDialog
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.savedDialogsSlice#44ba9dd9: field dialogs: %w", err)
			}
			s.Dialogs = append(s.Dialogs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedDialogsSlice#44ba9dd9: field messages: %w", err)
		}

		if headerLen > 0 {
			s.Messages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.savedDialogsSlice#44ba9dd9: field messages: %w", err)
			}
			s.Messages = append(s.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedDialogsSlice#44ba9dd9: field chats: %w", err)
		}

		if headerLen > 0 {
			s.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.savedDialogsSlice#44ba9dd9: field chats: %w", err)
			}
			s.Chats = append(s.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedDialogsSlice#44ba9dd9: field users: %w", err)
		}

		if headerLen > 0 {
			s.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.savedDialogsSlice#44ba9dd9: field users: %w", err)
			}
			s.Users = append(s.Users, value)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (s *MessagesSavedDialogsSlice) GetCount() (value int) {
	if s == nil {
		return
	}
	return s.Count
}

// GetDialogs returns value of Dialogs field.
func (s *MessagesSavedDialogsSlice) GetDialogs() (value []SavedDialog) {
	if s == nil {
		return
	}
	return s.Dialogs
}

// GetMessages returns value of Messages field.
func (s *MessagesSavedDialogsSlice) GetMessages() (value []MessageClass) {
	if s == nil {
		return
	}
	return s.Messages
}

// GetChats returns value of Chats field.
func (s *MessagesSavedDialogsSlice) GetChats() (value []ChatClass) {
	if s == nil {
		return
	}
	return s.Chats
}

// GetUsers returns value of Users field.
func (s *MessagesSavedDialogsSlice) GetUsers() (value []UserClass) {
	if s == nil {
		return
	}
	return s.Users
}

// MessagesSavedDialogsNotModified represents TL type `messages.savedDialogsNotModified#c01f6fe8`.
type MessagesSavedDialogsNotModified struct {
	// Count field of MessagesSavedDialogsNotModified.
	Count int
}

// MessagesSavedDialogsNotModifiedTypeID is TL type id of MessagesSavedDialogsNotModified.
const MessagesSavedDialogsNotModifiedTypeID = 0xc01f6fe8

// construct implements constructor of MessagesSavedDialogsClass.
func (s MessagesSavedDialogsNotModified) construct() MessagesSavedDialogsClass { return &s }

// Ensuring interfaces in compile-time for MessagesSavedDialogsNotModified.
var (
	_ bin.Encoder     = &MessagesSavedDialogsNotModified{}
	_ bin.Decoder     = &MessagesSavedDialogsNotModified{}
	_ bin.BareEncoder = &MessagesSavedDialogsNotModified{}
	_ bin.BareDecoder = &MessagesSavedDialogsNotModified{}

	_ MessagesSavedDialogsClass = &MessagesSavedDialogsNotModified{}
)

func (s *MessagesSavedDialogsNotModified) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSavedDialogsNotModified) String() string {
	if s == nil {
		return "MessagesSavedDialogsNotModified(nil)"
	}
	type Alias MessagesSavedDialogsNotModified
	return fmt.Sprintf("MessagesSavedDialogsNotModified%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSavedDialogsNotModified) TypeID() uint32 {
	return MessagesSavedDialogsNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSavedDialogsNotModified) TypeName() string {
	return "messages.savedDialogsNotModified"
}

// TypeInfo returns info about TL type.
func (s *MessagesSavedDialogsNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.savedDialogsNotModified",
		ID:   MessagesSavedDialogsNotModifiedTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSavedDialogsNotModified) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.savedDialogsNotModified#c01f6fe8 as nil")
	}
	b.PutID(MessagesSavedDialogsNotModifiedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSavedDialogsNotModified) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.savedDialogsNotModified#c01f6fe8 as nil")
	}
	b.PutInt(s.Count)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSavedDialogsNotModified) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.savedDialogsNotModified#c01f6fe8 to nil")
	}
	if err := b.ConsumeID(MessagesSavedDialogsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.savedDialogsNotModified#c01f6fe8: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSavedDialogsNotModified) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.savedDialogsNotModified#c01f6fe8 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedDialogsNotModified#c01f6fe8: field count: %w", err)
		}
		s.Count = value
	}
	return nil
}

// GetCount returns value of Count field.
func (s *MessagesSavedDialogsNotModified) GetCount() (value int) {
	if s == nil {
		return
	}
	return s.Count
}

// MessagesSavedDialogsClassName is schema name of MessagesSavedDialogsClass.
const MessagesSavedDialogsClassName = "messages.SavedDialogs"

// MessagesSavedDialogsClass represents messages.SavedDialogs generic type.
//
// Constructors:
//   - [MessagesSavedDialogs]
//   - [MessagesSavedDialogsSlice]
//   - [MessagesSavedDialogsNotModified]
//
// Example:
//
//	g, err := tg.DecodeMessagesSavedDialogs(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessagesSavedDialogs: // messages.savedDialogs#f83ae221
//	case *tg.MessagesSavedDialogsSlice: // messages.savedDialogsSlice#44ba9dd9
//	case *tg.MessagesSavedDialogsNotModified: // messages.savedDialogsNotModified#c01f6fe8
//	default: panic(v)
//	}
type MessagesSavedDialogsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesSavedDialogsClass

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

// DecodeMessagesSavedDialogs implements binary de-serialization for MessagesSavedDialogsClass.
func DecodeMessagesSavedDialogs(buf *bin.Buffer) (MessagesSavedDialogsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesSavedDialogsTypeID:
		// Decoding messages.savedDialogs#f83ae221.
		v := MessagesSavedDialogs{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesSavedDialogsClass: %w", err)
		}
		return &v, nil
	case MessagesSavedDialogsSliceTypeID:
		// Decoding messages.savedDialogsSlice#44ba9dd9.
		v := MessagesSavedDialogsSlice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesSavedDialogsClass: %w", err)
		}
		return &v, nil
	case MessagesSavedDialogsNotModifiedTypeID:
		// Decoding messages.savedDialogsNotModified#c01f6fe8.
		v := MessagesSavedDialogsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesSavedDialogsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesSavedDialogsClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesSavedDialogs boxes the MessagesSavedDialogsClass providing a helper.
type MessagesSavedDialogsBox struct {
	SavedDialogs MessagesSavedDialogsClass
}

// Decode implements bin.Decoder for MessagesSavedDialogsBox.
func (b *MessagesSavedDialogsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesSavedDialogsBox to nil")
	}
	v, err := DecodeMessagesSavedDialogs(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SavedDialogs = v
	return nil
}

// Encode implements bin.Encode for MessagesSavedDialogsBox.
func (b *MessagesSavedDialogsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SavedDialogs == nil {
		return fmt.Errorf("unable to encode MessagesSavedDialogsClass as nil")
	}
	return b.SavedDialogs.Encode(buf)
}
