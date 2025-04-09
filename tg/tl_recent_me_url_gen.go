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

// RecentMeURLUnknown represents TL type `recentMeUrlUnknown#46e1d13d`.
type RecentMeURLUnknown struct {
	// URL field of RecentMeURLUnknown.
	URL string
}

// RecentMeURLUnknownTypeID is TL type id of RecentMeURLUnknown.
const RecentMeURLUnknownTypeID = 0x46e1d13d

// construct implements constructor of RecentMeURLClass.
func (r RecentMeURLUnknown) construct() RecentMeURLClass { return &r }

// Ensuring interfaces in compile-time for RecentMeURLUnknown.
var (
	_ bin.Encoder     = &RecentMeURLUnknown{}
	_ bin.Decoder     = &RecentMeURLUnknown{}
	_ bin.BareEncoder = &RecentMeURLUnknown{}
	_ bin.BareDecoder = &RecentMeURLUnknown{}

	_ RecentMeURLClass = &RecentMeURLUnknown{}
)

func (r *RecentMeURLUnknown) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeURLUnknown) String() string {
	if r == nil {
		return "RecentMeURLUnknown(nil)"
	}
	type Alias RecentMeURLUnknown
	return fmt.Sprintf("RecentMeURLUnknown%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeURLUnknown) TypeID() uint32 {
	return RecentMeURLUnknownTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeURLUnknown) TypeName() string {
	return "recentMeUrlUnknown"
}

// TypeInfo returns info about TL type.
func (r *RecentMeURLUnknown) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlUnknown",
		ID:   RecentMeURLUnknownTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeURLUnknown) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlUnknown#46e1d13d as nil")
	}
	b.PutID(RecentMeURLUnknownTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecentMeURLUnknown) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlUnknown#46e1d13d as nil")
	}
	b.PutString(r.URL)
	return nil
}

// Decode implements bin.Decoder.
func (r *RecentMeURLUnknown) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlUnknown#46e1d13d to nil")
	}
	if err := b.ConsumeID(RecentMeURLUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlUnknown#46e1d13d: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecentMeURLUnknown) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlUnknown#46e1d13d to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlUnknown#46e1d13d: field url: %w", err)
		}
		r.URL = value
	}
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeURLUnknown) GetURL() (value string) {
	if r == nil {
		return
	}
	return r.URL
}

// RecentMeURLUser represents TL type `recentMeUrlUser#b92c09e2`.
type RecentMeURLUser struct {
	// URL field of RecentMeURLUser.
	URL string
	// UserID field of RecentMeURLUser.
	UserID int64
}

// RecentMeURLUserTypeID is TL type id of RecentMeURLUser.
const RecentMeURLUserTypeID = 0xb92c09e2

// construct implements constructor of RecentMeURLClass.
func (r RecentMeURLUser) construct() RecentMeURLClass { return &r }

// Ensuring interfaces in compile-time for RecentMeURLUser.
var (
	_ bin.Encoder     = &RecentMeURLUser{}
	_ bin.Decoder     = &RecentMeURLUser{}
	_ bin.BareEncoder = &RecentMeURLUser{}
	_ bin.BareDecoder = &RecentMeURLUser{}

	_ RecentMeURLClass = &RecentMeURLUser{}
)

func (r *RecentMeURLUser) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeURLUser) String() string {
	if r == nil {
		return "RecentMeURLUser(nil)"
	}
	type Alias RecentMeURLUser
	return fmt.Sprintf("RecentMeURLUser%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeURLUser) TypeID() uint32 {
	return RecentMeURLUserTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeURLUser) TypeName() string {
	return "recentMeUrlUser"
}

// TypeInfo returns info about TL type.
func (r *RecentMeURLUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlUser",
		ID:   RecentMeURLUserTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeURLUser) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlUser#b92c09e2 as nil")
	}
	b.PutID(RecentMeURLUserTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecentMeURLUser) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlUser#b92c09e2 as nil")
	}
	b.PutString(r.URL)
	b.PutLong(r.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (r *RecentMeURLUser) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlUser#b92c09e2 to nil")
	}
	if err := b.ConsumeID(RecentMeURLUserTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlUser#b92c09e2: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecentMeURLUser) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlUser#b92c09e2 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlUser#b92c09e2: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlUser#b92c09e2: field user_id: %w", err)
		}
		r.UserID = value
	}
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeURLUser) GetURL() (value string) {
	if r == nil {
		return
	}
	return r.URL
}

// GetUserID returns value of UserID field.
func (r *RecentMeURLUser) GetUserID() (value int64) {
	if r == nil {
		return
	}
	return r.UserID
}

// RecentMeURLChat represents TL type `recentMeUrlChat#b2da71d2`.
type RecentMeURLChat struct {
	// URL field of RecentMeURLChat.
	URL string
	// ChatID field of RecentMeURLChat.
	ChatID int64
}

// RecentMeURLChatTypeID is TL type id of RecentMeURLChat.
const RecentMeURLChatTypeID = 0xb2da71d2

// construct implements constructor of RecentMeURLClass.
func (r RecentMeURLChat) construct() RecentMeURLClass { return &r }

// Ensuring interfaces in compile-time for RecentMeURLChat.
var (
	_ bin.Encoder     = &RecentMeURLChat{}
	_ bin.Decoder     = &RecentMeURLChat{}
	_ bin.BareEncoder = &RecentMeURLChat{}
	_ bin.BareDecoder = &RecentMeURLChat{}

	_ RecentMeURLClass = &RecentMeURLChat{}
)

func (r *RecentMeURLChat) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeURLChat) String() string {
	if r == nil {
		return "RecentMeURLChat(nil)"
	}
	type Alias RecentMeURLChat
	return fmt.Sprintf("RecentMeURLChat%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeURLChat) TypeID() uint32 {
	return RecentMeURLChatTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeURLChat) TypeName() string {
	return "recentMeUrlChat"
}

// TypeInfo returns info about TL type.
func (r *RecentMeURLChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlChat",
		ID:   RecentMeURLChatTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeURLChat) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlChat#b2da71d2 as nil")
	}
	b.PutID(RecentMeURLChatTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecentMeURLChat) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlChat#b2da71d2 as nil")
	}
	b.PutString(r.URL)
	b.PutLong(r.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (r *RecentMeURLChat) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlChat#b2da71d2 to nil")
	}
	if err := b.ConsumeID(RecentMeURLChatTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlChat#b2da71d2: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecentMeURLChat) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlChat#b2da71d2 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChat#b2da71d2: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChat#b2da71d2: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeURLChat) GetURL() (value string) {
	if r == nil {
		return
	}
	return r.URL
}

// GetChatID returns value of ChatID field.
func (r *RecentMeURLChat) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// RecentMeURLChatInvite represents TL type `recentMeUrlChatInvite#eb49081d`.
type RecentMeURLChatInvite struct {
	// URL field of RecentMeURLChatInvite.
	URL string
	// ChatInvite field of RecentMeURLChatInvite.
	ChatInvite ChatInviteClass
}

// RecentMeURLChatInviteTypeID is TL type id of RecentMeURLChatInvite.
const RecentMeURLChatInviteTypeID = 0xeb49081d

// construct implements constructor of RecentMeURLClass.
func (r RecentMeURLChatInvite) construct() RecentMeURLClass { return &r }

// Ensuring interfaces in compile-time for RecentMeURLChatInvite.
var (
	_ bin.Encoder     = &RecentMeURLChatInvite{}
	_ bin.Decoder     = &RecentMeURLChatInvite{}
	_ bin.BareEncoder = &RecentMeURLChatInvite{}
	_ bin.BareDecoder = &RecentMeURLChatInvite{}

	_ RecentMeURLClass = &RecentMeURLChatInvite{}
)

func (r *RecentMeURLChatInvite) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.ChatInvite == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeURLChatInvite) String() string {
	if r == nil {
		return "RecentMeURLChatInvite(nil)"
	}
	type Alias RecentMeURLChatInvite
	return fmt.Sprintf("RecentMeURLChatInvite%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeURLChatInvite) TypeID() uint32 {
	return RecentMeURLChatInviteTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeURLChatInvite) TypeName() string {
	return "recentMeUrlChatInvite"
}

// TypeInfo returns info about TL type.
func (r *RecentMeURLChatInvite) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlChatInvite",
		ID:   RecentMeURLChatInviteTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "ChatInvite",
			SchemaName: "chat_invite",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeURLChatInvite) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlChatInvite#eb49081d as nil")
	}
	b.PutID(RecentMeURLChatInviteTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecentMeURLChatInvite) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlChatInvite#eb49081d as nil")
	}
	b.PutString(r.URL)
	if r.ChatInvite == nil {
		return fmt.Errorf("unable to encode recentMeUrlChatInvite#eb49081d: field chat_invite is nil")
	}
	if err := r.ChatInvite.Encode(b); err != nil {
		return fmt.Errorf("unable to encode recentMeUrlChatInvite#eb49081d: field chat_invite: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RecentMeURLChatInvite) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlChatInvite#eb49081d to nil")
	}
	if err := b.ConsumeID(RecentMeURLChatInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlChatInvite#eb49081d: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecentMeURLChatInvite) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlChatInvite#eb49081d to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChatInvite#eb49081d: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := DecodeChatInvite(b)
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChatInvite#eb49081d: field chat_invite: %w", err)
		}
		r.ChatInvite = value
	}
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeURLChatInvite) GetURL() (value string) {
	if r == nil {
		return
	}
	return r.URL
}

// GetChatInvite returns value of ChatInvite field.
func (r *RecentMeURLChatInvite) GetChatInvite() (value ChatInviteClass) {
	if r == nil {
		return
	}
	return r.ChatInvite
}

// RecentMeURLStickerSet represents TL type `recentMeUrlStickerSet#bc0a57dc`.
type RecentMeURLStickerSet struct {
	// URL field of RecentMeURLStickerSet.
	URL string
	// Set field of RecentMeURLStickerSet.
	Set StickerSetCoveredClass
}

// RecentMeURLStickerSetTypeID is TL type id of RecentMeURLStickerSet.
const RecentMeURLStickerSetTypeID = 0xbc0a57dc

// construct implements constructor of RecentMeURLClass.
func (r RecentMeURLStickerSet) construct() RecentMeURLClass { return &r }

// Ensuring interfaces in compile-time for RecentMeURLStickerSet.
var (
	_ bin.Encoder     = &RecentMeURLStickerSet{}
	_ bin.Decoder     = &RecentMeURLStickerSet{}
	_ bin.BareEncoder = &RecentMeURLStickerSet{}
	_ bin.BareDecoder = &RecentMeURLStickerSet{}

	_ RecentMeURLClass = &RecentMeURLStickerSet{}
)

func (r *RecentMeURLStickerSet) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.Set == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeURLStickerSet) String() string {
	if r == nil {
		return "RecentMeURLStickerSet(nil)"
	}
	type Alias RecentMeURLStickerSet
	return fmt.Sprintf("RecentMeURLStickerSet%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeURLStickerSet) TypeID() uint32 {
	return RecentMeURLStickerSetTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeURLStickerSet) TypeName() string {
	return "recentMeUrlStickerSet"
}

// TypeInfo returns info about TL type.
func (r *RecentMeURLStickerSet) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlStickerSet",
		ID:   RecentMeURLStickerSetTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Set",
			SchemaName: "set",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeURLStickerSet) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlStickerSet#bc0a57dc as nil")
	}
	b.PutID(RecentMeURLStickerSetTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecentMeURLStickerSet) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlStickerSet#bc0a57dc as nil")
	}
	b.PutString(r.URL)
	if r.Set == nil {
		return fmt.Errorf("unable to encode recentMeUrlStickerSet#bc0a57dc: field set is nil")
	}
	if err := r.Set.Encode(b); err != nil {
		return fmt.Errorf("unable to encode recentMeUrlStickerSet#bc0a57dc: field set: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RecentMeURLStickerSet) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlStickerSet#bc0a57dc to nil")
	}
	if err := b.ConsumeID(RecentMeURLStickerSetTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlStickerSet#bc0a57dc: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecentMeURLStickerSet) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlStickerSet#bc0a57dc to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlStickerSet#bc0a57dc: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := DecodeStickerSetCovered(b)
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlStickerSet#bc0a57dc: field set: %w", err)
		}
		r.Set = value
	}
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeURLStickerSet) GetURL() (value string) {
	if r == nil {
		return
	}
	return r.URL
}

// GetSet returns value of Set field.
func (r *RecentMeURLStickerSet) GetSet() (value StickerSetCoveredClass) {
	if r == nil {
		return
	}
	return r.Set
}

// RecentMeURLClassName is schema name of RecentMeURLClass.
const RecentMeURLClassName = "RecentMeUrl"

// RecentMeURLClass represents RecentMeUrl generic type.
//
// Constructors:
//   - [RecentMeURLUnknown]
//   - [RecentMeURLUser]
//   - [RecentMeURLChat]
//   - [RecentMeURLChatInvite]
//   - [RecentMeURLStickerSet]
//
// Example:
//
//	g, err := tg.DecodeRecentMeURL(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.RecentMeURLUnknown: // recentMeUrlUnknown#46e1d13d
//	case *tg.RecentMeURLUser: // recentMeUrlUser#b92c09e2
//	case *tg.RecentMeURLChat: // recentMeUrlChat#b2da71d2
//	case *tg.RecentMeURLChatInvite: // recentMeUrlChatInvite#eb49081d
//	case *tg.RecentMeURLStickerSet: // recentMeUrlStickerSet#bc0a57dc
//	default: panic(v)
//	}
type RecentMeURLClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() RecentMeURLClass

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

	// URL field of RecentMeURLUnknown.
	GetURL() (value string)
}

// DecodeRecentMeURL implements binary de-serialization for RecentMeURLClass.
func DecodeRecentMeURL(buf *bin.Buffer) (RecentMeURLClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case RecentMeURLUnknownTypeID:
		// Decoding recentMeUrlUnknown#46e1d13d.
		v := RecentMeURLUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeURLClass: %w", err)
		}
		return &v, nil
	case RecentMeURLUserTypeID:
		// Decoding recentMeUrlUser#b92c09e2.
		v := RecentMeURLUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeURLClass: %w", err)
		}
		return &v, nil
	case RecentMeURLChatTypeID:
		// Decoding recentMeUrlChat#b2da71d2.
		v := RecentMeURLChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeURLClass: %w", err)
		}
		return &v, nil
	case RecentMeURLChatInviteTypeID:
		// Decoding recentMeUrlChatInvite#eb49081d.
		v := RecentMeURLChatInvite{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeURLClass: %w", err)
		}
		return &v, nil
	case RecentMeURLStickerSetTypeID:
		// Decoding recentMeUrlStickerSet#bc0a57dc.
		v := RecentMeURLStickerSet{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeURLClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode RecentMeURLClass: %w", bin.NewUnexpectedID(id))
	}
}

// RecentMeURL boxes the RecentMeURLClass providing a helper.
type RecentMeURLBox struct {
	RecentMeUrl RecentMeURLClass
}

// Decode implements bin.Decoder for RecentMeURLBox.
func (b *RecentMeURLBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode RecentMeURLBox to nil")
	}
	v, err := DecodeRecentMeURL(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.RecentMeUrl = v
	return nil
}

// Encode implements bin.Encode for RecentMeURLBox.
func (b *RecentMeURLBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.RecentMeUrl == nil {
		return fmt.Errorf("unable to encode RecentMeURLClass as nil")
	}
	return b.RecentMeUrl.Encode(buf)
}
