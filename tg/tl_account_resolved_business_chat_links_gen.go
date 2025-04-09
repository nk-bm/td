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

// AccountResolvedBusinessChatLinks represents TL type `account.resolvedBusinessChatLinks#9a23af21`.
type AccountResolvedBusinessChatLinks struct {
	// Flags field of AccountResolvedBusinessChatLinks.
	Flags bin.Fields
	// Peer field of AccountResolvedBusinessChatLinks.
	Peer PeerClass
	// Message field of AccountResolvedBusinessChatLinks.
	Message string
	// Entities field of AccountResolvedBusinessChatLinks.
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// Chats field of AccountResolvedBusinessChatLinks.
	Chats []ChatClass
	// Users field of AccountResolvedBusinessChatLinks.
	Users []UserClass
}

// AccountResolvedBusinessChatLinksTypeID is TL type id of AccountResolvedBusinessChatLinks.
const AccountResolvedBusinessChatLinksTypeID = 0x9a23af21

// Ensuring interfaces in compile-time for AccountResolvedBusinessChatLinks.
var (
	_ bin.Encoder     = &AccountResolvedBusinessChatLinks{}
	_ bin.Decoder     = &AccountResolvedBusinessChatLinks{}
	_ bin.BareEncoder = &AccountResolvedBusinessChatLinks{}
	_ bin.BareDecoder = &AccountResolvedBusinessChatLinks{}
)

func (r *AccountResolvedBusinessChatLinks) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.Message == "") {
		return false
	}
	if !(r.Entities == nil) {
		return false
	}
	if !(r.Chats == nil) {
		return false
	}
	if !(r.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResolvedBusinessChatLinks) String() string {
	if r == nil {
		return "AccountResolvedBusinessChatLinks(nil)"
	}
	type Alias AccountResolvedBusinessChatLinks
	return fmt.Sprintf("AccountResolvedBusinessChatLinks%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountResolvedBusinessChatLinks) TypeID() uint32 {
	return AccountResolvedBusinessChatLinksTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountResolvedBusinessChatLinks) TypeName() string {
	return "account.resolvedBusinessChatLinks"
}

// TypeInfo returns info about TL type.
func (r *AccountResolvedBusinessChatLinks) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.resolvedBusinessChatLinks",
		ID:   AccountResolvedBusinessChatLinksTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !r.Flags.Has(0),
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
func (r *AccountResolvedBusinessChatLinks) SetFlags() {
	if !(r.Entities == nil) {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *AccountResolvedBusinessChatLinks) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resolvedBusinessChatLinks#9a23af21 as nil")
	}
	b.PutID(AccountResolvedBusinessChatLinksTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountResolvedBusinessChatLinks) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resolvedBusinessChatLinks#9a23af21 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.resolvedBusinessChatLinks#9a23af21: field flags: %w", err)
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode account.resolvedBusinessChatLinks#9a23af21: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.resolvedBusinessChatLinks#9a23af21: field peer: %w", err)
	}
	b.PutString(r.Message)
	if r.Flags.Has(0) {
		b.PutVectorHeader(len(r.Entities))
		for idx, v := range r.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode account.resolvedBusinessChatLinks#9a23af21: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode account.resolvedBusinessChatLinks#9a23af21: field entities element with index %d: %w", idx, err)
			}
		}
	}
	b.PutVectorHeader(len(r.Chats))
	for idx, v := range r.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode account.resolvedBusinessChatLinks#9a23af21: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.resolvedBusinessChatLinks#9a23af21: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(r.Users))
	for idx, v := range r.Users {
		if v == nil {
			return fmt.Errorf("unable to encode account.resolvedBusinessChatLinks#9a23af21: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.resolvedBusinessChatLinks#9a23af21: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResolvedBusinessChatLinks) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resolvedBusinessChatLinks#9a23af21 to nil")
	}
	if err := b.ConsumeID(AccountResolvedBusinessChatLinksTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resolvedBusinessChatLinks#9a23af21: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountResolvedBusinessChatLinks) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resolvedBusinessChatLinks#9a23af21 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.resolvedBusinessChatLinks#9a23af21: field flags: %w", err)
		}
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.resolvedBusinessChatLinks#9a23af21: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.resolvedBusinessChatLinks#9a23af21: field message: %w", err)
		}
		r.Message = value
	}
	if r.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.resolvedBusinessChatLinks#9a23af21: field entities: %w", err)
		}

		if headerLen > 0 {
			r.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.resolvedBusinessChatLinks#9a23af21: field entities: %w", err)
			}
			r.Entities = append(r.Entities, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.resolvedBusinessChatLinks#9a23af21: field chats: %w", err)
		}

		if headerLen > 0 {
			r.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.resolvedBusinessChatLinks#9a23af21: field chats: %w", err)
			}
			r.Chats = append(r.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.resolvedBusinessChatLinks#9a23af21: field users: %w", err)
		}

		if headerLen > 0 {
			r.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.resolvedBusinessChatLinks#9a23af21: field users: %w", err)
			}
			r.Users = append(r.Users, value)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *AccountResolvedBusinessChatLinks) GetPeer() (value PeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// GetMessage returns value of Message field.
func (r *AccountResolvedBusinessChatLinks) GetMessage() (value string) {
	if r == nil {
		return
	}
	return r.Message
}

// SetEntities sets value of Entities conditional field.
func (r *AccountResolvedBusinessChatLinks) SetEntities(value []MessageEntityClass) {
	r.Flags.Set(0)
	r.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (r *AccountResolvedBusinessChatLinks) GetEntities() (value []MessageEntityClass, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.Entities, true
}

// GetChats returns value of Chats field.
func (r *AccountResolvedBusinessChatLinks) GetChats() (value []ChatClass) {
	if r == nil {
		return
	}
	return r.Chats
}

// GetUsers returns value of Users field.
func (r *AccountResolvedBusinessChatLinks) GetUsers() (value []UserClass) {
	if r == nil {
		return
	}
	return r.Users
}
