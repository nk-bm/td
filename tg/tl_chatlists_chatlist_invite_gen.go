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

// ChatlistsChatlistInviteAlready represents TL type `chatlists.chatlistInviteAlready#fa87f659`.
// Updated info about a chat folder deep link »¹ we already imported.
//
// Links:
//  1. https://core.telegram.org/api/links#chat-folder-links
//
// See https://core.telegram.org/constructor/chatlists.chatlistInviteAlready for reference.
type ChatlistsChatlistInviteAlready struct {
	// ID of the imported folder
	FilterID int
	// New peers to be imported
	MissingPeers []PeerClass
	// Peers that were already imported
	AlreadyPeers []PeerClass
	// Related chat information
	Chats []ChatClass
	// Related user information
	Users []UserClass
}

// ChatlistsChatlistInviteAlreadyTypeID is TL type id of ChatlistsChatlistInviteAlready.
const ChatlistsChatlistInviteAlreadyTypeID = 0xfa87f659

// construct implements constructor of ChatlistsChatlistInviteClass.
func (c ChatlistsChatlistInviteAlready) construct() ChatlistsChatlistInviteClass { return &c }

// Ensuring interfaces in compile-time for ChatlistsChatlistInviteAlready.
var (
	_ bin.Encoder     = &ChatlistsChatlistInviteAlready{}
	_ bin.Decoder     = &ChatlistsChatlistInviteAlready{}
	_ bin.BareEncoder = &ChatlistsChatlistInviteAlready{}
	_ bin.BareDecoder = &ChatlistsChatlistInviteAlready{}

	_ ChatlistsChatlistInviteClass = &ChatlistsChatlistInviteAlready{}
)

func (c *ChatlistsChatlistInviteAlready) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.FilterID == 0) {
		return false
	}
	if !(c.MissingPeers == nil) {
		return false
	}
	if !(c.AlreadyPeers == nil) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatlistsChatlistInviteAlready) String() string {
	if c == nil {
		return "ChatlistsChatlistInviteAlready(nil)"
	}
	type Alias ChatlistsChatlistInviteAlready
	return fmt.Sprintf("ChatlistsChatlistInviteAlready%+v", Alias(*c))
}

// FillFrom fills ChatlistsChatlistInviteAlready from given interface.
func (c *ChatlistsChatlistInviteAlready) FillFrom(from interface {
	GetFilterID() (value int)
	GetMissingPeers() (value []PeerClass)
	GetAlreadyPeers() (value []PeerClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	c.FilterID = from.GetFilterID()
	c.MissingPeers = from.GetMissingPeers()
	c.AlreadyPeers = from.GetAlreadyPeers()
	c.Chats = from.GetChats()
	c.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatlistsChatlistInviteAlready) TypeID() uint32 {
	return ChatlistsChatlistInviteAlreadyTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatlistsChatlistInviteAlready) TypeName() string {
	return "chatlists.chatlistInviteAlready"
}

// TypeInfo returns info about TL type.
func (c *ChatlistsChatlistInviteAlready) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatlists.chatlistInviteAlready",
		ID:   ChatlistsChatlistInviteAlreadyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FilterID",
			SchemaName: "filter_id",
		},
		{
			Name:       "MissingPeers",
			SchemaName: "missing_peers",
		},
		{
			Name:       "AlreadyPeers",
			SchemaName: "already_peers",
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
func (c *ChatlistsChatlistInviteAlready) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatlists.chatlistInviteAlready#fa87f659 as nil")
	}
	b.PutID(ChatlistsChatlistInviteAlreadyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatlistsChatlistInviteAlready) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatlists.chatlistInviteAlready#fa87f659 as nil")
	}
	b.PutInt(c.FilterID)
	b.PutVectorHeader(len(c.MissingPeers))
	for idx, v := range c.MissingPeers {
		if v == nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInviteAlready#fa87f659: field missing_peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInviteAlready#fa87f659: field missing_peers element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.AlreadyPeers))
	for idx, v := range c.AlreadyPeers {
		if v == nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInviteAlready#fa87f659: field already_peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInviteAlready#fa87f659: field already_peers element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInviteAlready#fa87f659: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInviteAlready#fa87f659: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInviteAlready#fa87f659: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInviteAlready#fa87f659: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatlistsChatlistInviteAlready) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatlists.chatlistInviteAlready#fa87f659 to nil")
	}
	if err := b.ConsumeID(ChatlistsChatlistInviteAlreadyTypeID); err != nil {
		return fmt.Errorf("unable to decode chatlists.chatlistInviteAlready#fa87f659: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatlistsChatlistInviteAlready) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatlists.chatlistInviteAlready#fa87f659 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.chatlistInviteAlready#fa87f659: field filter_id: %w", err)
		}
		c.FilterID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.chatlistInviteAlready#fa87f659: field missing_peers: %w", err)
		}

		if headerLen > 0 {
			c.MissingPeers = make([]PeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatlists.chatlistInviteAlready#fa87f659: field missing_peers: %w", err)
			}
			c.MissingPeers = append(c.MissingPeers, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.chatlistInviteAlready#fa87f659: field already_peers: %w", err)
		}

		if headerLen > 0 {
			c.AlreadyPeers = make([]PeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatlists.chatlistInviteAlready#fa87f659: field already_peers: %w", err)
			}
			c.AlreadyPeers = append(c.AlreadyPeers, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.chatlistInviteAlready#fa87f659: field chats: %w", err)
		}

		if headerLen > 0 {
			c.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatlists.chatlistInviteAlready#fa87f659: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.chatlistInviteAlready#fa87f659: field users: %w", err)
		}

		if headerLen > 0 {
			c.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatlists.chatlistInviteAlready#fa87f659: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// GetFilterID returns value of FilterID field.
func (c *ChatlistsChatlistInviteAlready) GetFilterID() (value int) {
	if c == nil {
		return
	}
	return c.FilterID
}

// GetMissingPeers returns value of MissingPeers field.
func (c *ChatlistsChatlistInviteAlready) GetMissingPeers() (value []PeerClass) {
	if c == nil {
		return
	}
	return c.MissingPeers
}

// GetAlreadyPeers returns value of AlreadyPeers field.
func (c *ChatlistsChatlistInviteAlready) GetAlreadyPeers() (value []PeerClass) {
	if c == nil {
		return
	}
	return c.AlreadyPeers
}

// GetChats returns value of Chats field.
func (c *ChatlistsChatlistInviteAlready) GetChats() (value []ChatClass) {
	if c == nil {
		return
	}
	return c.Chats
}

// GetUsers returns value of Users field.
func (c *ChatlistsChatlistInviteAlready) GetUsers() (value []UserClass) {
	if c == nil {
		return
	}
	return c.Users
}

// MapMissingPeers returns field MissingPeers wrapped in PeerClassArray helper.
func (c *ChatlistsChatlistInviteAlready) MapMissingPeers() (value PeerClassArray) {
	return PeerClassArray(c.MissingPeers)
}

// MapAlreadyPeers returns field AlreadyPeers wrapped in PeerClassArray helper.
func (c *ChatlistsChatlistInviteAlready) MapAlreadyPeers() (value PeerClassArray) {
	return PeerClassArray(c.AlreadyPeers)
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (c *ChatlistsChatlistInviteAlready) MapChats() (value ChatClassArray) {
	return ChatClassArray(c.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (c *ChatlistsChatlistInviteAlready) MapUsers() (value UserClassArray) {
	return UserClassArray(c.Users)
}

// ChatlistsChatlistInvite represents TL type `chatlists.chatlistInvite#1dcd839d`.
// Info about a chat folder deep link »¹.
//
// Links:
//  1. https://core.telegram.org/api/links#chat-folder-links
//
// See https://core.telegram.org/constructor/chatlists.chatlistInvite for reference.
type ChatlistsChatlistInvite struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Name of the link
	Title string
	// Emoji to use as icon for the folder.
	//
	// Use SetEmoticon and GetEmoticon helpers.
	Emoticon string
	// Supergroups and channels to join
	Peers []PeerClass
	// Related chat information
	Chats []ChatClass
	// Related user information
	Users []UserClass
}

// ChatlistsChatlistInviteTypeID is TL type id of ChatlistsChatlistInvite.
const ChatlistsChatlistInviteTypeID = 0x1dcd839d

// construct implements constructor of ChatlistsChatlistInviteClass.
func (c ChatlistsChatlistInvite) construct() ChatlistsChatlistInviteClass { return &c }

// Ensuring interfaces in compile-time for ChatlistsChatlistInvite.
var (
	_ bin.Encoder     = &ChatlistsChatlistInvite{}
	_ bin.Decoder     = &ChatlistsChatlistInvite{}
	_ bin.BareEncoder = &ChatlistsChatlistInvite{}
	_ bin.BareDecoder = &ChatlistsChatlistInvite{}

	_ ChatlistsChatlistInviteClass = &ChatlistsChatlistInvite{}
)

func (c *ChatlistsChatlistInvite) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.Emoticon == "") {
		return false
	}
	if !(c.Peers == nil) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatlistsChatlistInvite) String() string {
	if c == nil {
		return "ChatlistsChatlistInvite(nil)"
	}
	type Alias ChatlistsChatlistInvite
	return fmt.Sprintf("ChatlistsChatlistInvite%+v", Alias(*c))
}

// FillFrom fills ChatlistsChatlistInvite from given interface.
func (c *ChatlistsChatlistInvite) FillFrom(from interface {
	GetTitle() (value string)
	GetEmoticon() (value string, ok bool)
	GetPeers() (value []PeerClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	c.Title = from.GetTitle()
	if val, ok := from.GetEmoticon(); ok {
		c.Emoticon = val
	}

	c.Peers = from.GetPeers()
	c.Chats = from.GetChats()
	c.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatlistsChatlistInvite) TypeID() uint32 {
	return ChatlistsChatlistInviteTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatlistsChatlistInvite) TypeName() string {
	return "chatlists.chatlistInvite"
}

// TypeInfo returns info about TL type.
func (c *ChatlistsChatlistInvite) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatlists.chatlistInvite",
		ID:   ChatlistsChatlistInviteTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Emoticon",
			SchemaName: "emoticon",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Peers",
			SchemaName: "peers",
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

// SetFlags sets flags for non-zero fields.
func (c *ChatlistsChatlistInvite) SetFlags() {
	if !(c.Emoticon == "") {
		c.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (c *ChatlistsChatlistInvite) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatlists.chatlistInvite#1dcd839d as nil")
	}
	b.PutID(ChatlistsChatlistInviteTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatlistsChatlistInvite) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatlists.chatlistInvite#1dcd839d as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatlists.chatlistInvite#1dcd839d: field flags: %w", err)
	}
	b.PutString(c.Title)
	if c.Flags.Has(0) {
		b.PutString(c.Emoticon)
	}
	b.PutVectorHeader(len(c.Peers))
	for idx, v := range c.Peers {
		if v == nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInvite#1dcd839d: field peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInvite#1dcd839d: field peers element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInvite#1dcd839d: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInvite#1dcd839d: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInvite#1dcd839d: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatlists.chatlistInvite#1dcd839d: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatlistsChatlistInvite) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatlists.chatlistInvite#1dcd839d to nil")
	}
	if err := b.ConsumeID(ChatlistsChatlistInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode chatlists.chatlistInvite#1dcd839d: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatlistsChatlistInvite) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatlists.chatlistInvite#1dcd839d to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatlists.chatlistInvite#1dcd839d: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.chatlistInvite#1dcd839d: field title: %w", err)
		}
		c.Title = value
	}
	if c.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.chatlistInvite#1dcd839d: field emoticon: %w", err)
		}
		c.Emoticon = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.chatlistInvite#1dcd839d: field peers: %w", err)
		}

		if headerLen > 0 {
			c.Peers = make([]PeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatlists.chatlistInvite#1dcd839d: field peers: %w", err)
			}
			c.Peers = append(c.Peers, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.chatlistInvite#1dcd839d: field chats: %w", err)
		}

		if headerLen > 0 {
			c.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatlists.chatlistInvite#1dcd839d: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.chatlistInvite#1dcd839d: field users: %w", err)
		}

		if headerLen > 0 {
			c.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatlists.chatlistInvite#1dcd839d: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// GetTitle returns value of Title field.
func (c *ChatlistsChatlistInvite) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

// SetEmoticon sets value of Emoticon conditional field.
func (c *ChatlistsChatlistInvite) SetEmoticon(value string) {
	c.Flags.Set(0)
	c.Emoticon = value
}

// GetEmoticon returns value of Emoticon conditional field and
// boolean which is true if field was set.
func (c *ChatlistsChatlistInvite) GetEmoticon() (value string, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(0) {
		return value, false
	}
	return c.Emoticon, true
}

// GetPeers returns value of Peers field.
func (c *ChatlistsChatlistInvite) GetPeers() (value []PeerClass) {
	if c == nil {
		return
	}
	return c.Peers
}

// GetChats returns value of Chats field.
func (c *ChatlistsChatlistInvite) GetChats() (value []ChatClass) {
	if c == nil {
		return
	}
	return c.Chats
}

// GetUsers returns value of Users field.
func (c *ChatlistsChatlistInvite) GetUsers() (value []UserClass) {
	if c == nil {
		return
	}
	return c.Users
}

// MapPeers returns field Peers wrapped in PeerClassArray helper.
func (c *ChatlistsChatlistInvite) MapPeers() (value PeerClassArray) {
	return PeerClassArray(c.Peers)
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (c *ChatlistsChatlistInvite) MapChats() (value ChatClassArray) {
	return ChatClassArray(c.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (c *ChatlistsChatlistInvite) MapUsers() (value UserClassArray) {
	return UserClassArray(c.Users)
}

// ChatlistsChatlistInviteClassName is schema name of ChatlistsChatlistInviteClass.
const ChatlistsChatlistInviteClassName = "chatlists.ChatlistInvite"

// ChatlistsChatlistInviteClass represents chatlists.ChatlistInvite generic type.
//
// See https://core.telegram.org/type/chatlists.ChatlistInvite for reference.
//
// Constructors:
//   - [ChatlistsChatlistInviteAlready]
//   - [ChatlistsChatlistInvite]
//
// Example:
//
//	g, err := tg.DecodeChatlistsChatlistInvite(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.ChatlistsChatlistInviteAlready: // chatlists.chatlistInviteAlready#fa87f659
//	case *tg.ChatlistsChatlistInvite: // chatlists.chatlistInvite#1dcd839d
//	default: panic(v)
//	}
type ChatlistsChatlistInviteClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatlistsChatlistInviteClass

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

	// Related chat information
	GetChats() (value []ChatClass)
	// Related chat information
	MapChats() (value ChatClassArray)
	// Related user information
	GetUsers() (value []UserClass)
	// Related user information
	MapUsers() (value UserClassArray)
}

// DecodeChatlistsChatlistInvite implements binary de-serialization for ChatlistsChatlistInviteClass.
func DecodeChatlistsChatlistInvite(buf *bin.Buffer) (ChatlistsChatlistInviteClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatlistsChatlistInviteAlreadyTypeID:
		// Decoding chatlists.chatlistInviteAlready#fa87f659.
		v := ChatlistsChatlistInviteAlready{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatlistsChatlistInviteClass: %w", err)
		}
		return &v, nil
	case ChatlistsChatlistInviteTypeID:
		// Decoding chatlists.chatlistInvite#1dcd839d.
		v := ChatlistsChatlistInvite{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatlistsChatlistInviteClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatlistsChatlistInviteClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatlistsChatlistInvite boxes the ChatlistsChatlistInviteClass providing a helper.
type ChatlistsChatlistInviteBox struct {
	ChatlistInvite ChatlistsChatlistInviteClass
}

// Decode implements bin.Decoder for ChatlistsChatlistInviteBox.
func (b *ChatlistsChatlistInviteBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatlistsChatlistInviteBox to nil")
	}
	v, err := DecodeChatlistsChatlistInvite(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatlistInvite = v
	return nil
}

// Encode implements bin.Encode for ChatlistsChatlistInviteBox.
func (b *ChatlistsChatlistInviteBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatlistInvite == nil {
		return fmt.Errorf("unable to encode ChatlistsChatlistInviteClass as nil")
	}
	return b.ChatlistInvite.Encode(buf)
}
