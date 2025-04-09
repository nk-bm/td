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

// HelpRecentMeURLs represents TL type `help.recentMeUrls#e0310d7`.
type HelpRecentMeURLs struct {
	// URLs field of HelpRecentMeURLs.
	URLs []RecentMeURLClass
	// Chats field of HelpRecentMeURLs.
	Chats []ChatClass
	// Users field of HelpRecentMeURLs.
	Users []UserClass
}

// HelpRecentMeURLsTypeID is TL type id of HelpRecentMeURLs.
const HelpRecentMeURLsTypeID = 0xe0310d7

// Ensuring interfaces in compile-time for HelpRecentMeURLs.
var (
	_ bin.Encoder     = &HelpRecentMeURLs{}
	_ bin.Decoder     = &HelpRecentMeURLs{}
	_ bin.BareEncoder = &HelpRecentMeURLs{}
	_ bin.BareDecoder = &HelpRecentMeURLs{}
)

func (r *HelpRecentMeURLs) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URLs == nil) {
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
func (r *HelpRecentMeURLs) String() string {
	if r == nil {
		return "HelpRecentMeURLs(nil)"
	}
	type Alias HelpRecentMeURLs
	return fmt.Sprintf("HelpRecentMeURLs%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpRecentMeURLs) TypeID() uint32 {
	return HelpRecentMeURLsTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpRecentMeURLs) TypeName() string {
	return "help.recentMeUrls"
}

// TypeInfo returns info about TL type.
func (r *HelpRecentMeURLs) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.recentMeUrls",
		ID:   HelpRecentMeURLsTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URLs",
			SchemaName: "urls",
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
func (r *HelpRecentMeURLs) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode help.recentMeUrls#e0310d7 as nil")
	}
	b.PutID(HelpRecentMeURLsTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *HelpRecentMeURLs) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode help.recentMeUrls#e0310d7 as nil")
	}
	b.PutVectorHeader(len(r.URLs))
	for idx, v := range r.URLs {
		if v == nil {
			return fmt.Errorf("unable to encode help.recentMeUrls#e0310d7: field urls element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.recentMeUrls#e0310d7: field urls element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(r.Chats))
	for idx, v := range r.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode help.recentMeUrls#e0310d7: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.recentMeUrls#e0310d7: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(r.Users))
	for idx, v := range r.Users {
		if v == nil {
			return fmt.Errorf("unable to encode help.recentMeUrls#e0310d7: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.recentMeUrls#e0310d7: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *HelpRecentMeURLs) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode help.recentMeUrls#e0310d7 to nil")
	}
	if err := b.ConsumeID(HelpRecentMeURLsTypeID); err != nil {
		return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *HelpRecentMeURLs) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode help.recentMeUrls#e0310d7 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: field urls: %w", err)
		}

		if headerLen > 0 {
			r.URLs = make([]RecentMeURLClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeRecentMeURL(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: field urls: %w", err)
			}
			r.URLs = append(r.URLs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: field chats: %w", err)
		}

		if headerLen > 0 {
			r.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: field chats: %w", err)
			}
			r.Chats = append(r.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: field users: %w", err)
		}

		if headerLen > 0 {
			r.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: field users: %w", err)
			}
			r.Users = append(r.Users, value)
		}
	}
	return nil
}

// GetURLs returns value of URLs field.
func (r *HelpRecentMeURLs) GetURLs() (value []RecentMeURLClass) {
	if r == nil {
		return
	}
	return r.URLs
}

// GetChats returns value of Chats field.
func (r *HelpRecentMeURLs) GetChats() (value []ChatClass) {
	if r == nil {
		return
	}
	return r.Chats
}

// GetUsers returns value of Users field.
func (r *HelpRecentMeURLs) GetUsers() (value []UserClass) {
	if r == nil {
		return
	}
	return r.Users
}
