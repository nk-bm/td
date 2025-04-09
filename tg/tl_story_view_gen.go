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

// StoryView represents TL type `storyView#b0bdeac5`.
type StoryView struct {
	// Flags field of StoryView.
	Flags bin.Fields
	// Blocked field of StoryView.
	Blocked bool
	// BlockedMyStoriesFrom field of StoryView.
	BlockedMyStoriesFrom bool
	// UserID field of StoryView.
	UserID int64
	// Date field of StoryView.
	Date int
	// Reaction field of StoryView.
	//
	// Use SetReaction and GetReaction helpers.
	Reaction ReactionClass
}

// StoryViewTypeID is TL type id of StoryView.
const StoryViewTypeID = 0xb0bdeac5

// construct implements constructor of StoryViewClass.
func (s StoryView) construct() StoryViewClass { return &s }

// Ensuring interfaces in compile-time for StoryView.
var (
	_ bin.Encoder     = &StoryView{}
	_ bin.Decoder     = &StoryView{}
	_ bin.BareEncoder = &StoryView{}
	_ bin.BareDecoder = &StoryView{}

	_ StoryViewClass = &StoryView{}
)

func (s *StoryView) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Blocked == false) {
		return false
	}
	if !(s.BlockedMyStoriesFrom == false) {
		return false
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.Reaction == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryView) String() string {
	if s == nil {
		return "StoryView(nil)"
	}
	type Alias StoryView
	return fmt.Sprintf("StoryView%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryView) TypeID() uint32 {
	return StoryViewTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryView) TypeName() string {
	return "storyView"
}

// TypeInfo returns info about TL type.
func (s *StoryView) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyView",
		ID:   StoryViewTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Blocked",
			SchemaName: "blocked",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "BlockedMyStoriesFrom",
			SchemaName: "blocked_my_stories_from",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Reaction",
			SchemaName: "reaction",
			Null:       !s.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StoryView) SetFlags() {
	if !(s.Blocked == false) {
		s.Flags.Set(0)
	}
	if !(s.BlockedMyStoriesFrom == false) {
		s.Flags.Set(1)
	}
	if !(s.Reaction == nil) {
		s.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (s *StoryView) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyView#b0bdeac5 as nil")
	}
	b.PutID(StoryViewTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryView) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyView#b0bdeac5 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyView#b0bdeac5: field flags: %w", err)
	}
	b.PutLong(s.UserID)
	b.PutInt(s.Date)
	if s.Flags.Has(2) {
		if s.Reaction == nil {
			return fmt.Errorf("unable to encode storyView#b0bdeac5: field reaction is nil")
		}
		if err := s.Reaction.Encode(b); err != nil {
			return fmt.Errorf("unable to encode storyView#b0bdeac5: field reaction: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryView) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyView#b0bdeac5 to nil")
	}
	if err := b.ConsumeID(StoryViewTypeID); err != nil {
		return fmt.Errorf("unable to decode storyView#b0bdeac5: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryView) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyView#b0bdeac5 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyView#b0bdeac5: field flags: %w", err)
		}
	}
	s.Blocked = s.Flags.Has(0)
	s.BlockedMyStoriesFrom = s.Flags.Has(1)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode storyView#b0bdeac5: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyView#b0bdeac5: field date: %w", err)
		}
		s.Date = value
	}
	if s.Flags.Has(2) {
		value, err := DecodeReaction(b)
		if err != nil {
			return fmt.Errorf("unable to decode storyView#b0bdeac5: field reaction: %w", err)
		}
		s.Reaction = value
	}
	return nil
}

// SetBlocked sets value of Blocked conditional field.
func (s *StoryView) SetBlocked(value bool) {
	if value {
		s.Flags.Set(0)
		s.Blocked = true
	} else {
		s.Flags.Unset(0)
		s.Blocked = false
	}
}

// GetBlocked returns value of Blocked conditional field.
func (s *StoryView) GetBlocked() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// SetBlockedMyStoriesFrom sets value of BlockedMyStoriesFrom conditional field.
func (s *StoryView) SetBlockedMyStoriesFrom(value bool) {
	if value {
		s.Flags.Set(1)
		s.BlockedMyStoriesFrom = true
	} else {
		s.Flags.Unset(1)
		s.BlockedMyStoriesFrom = false
	}
}

// GetBlockedMyStoriesFrom returns value of BlockedMyStoriesFrom conditional field.
func (s *StoryView) GetBlockedMyStoriesFrom() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// GetUserID returns value of UserID field.
func (s *StoryView) GetUserID() (value int64) {
	if s == nil {
		return
	}
	return s.UserID
}

// GetDate returns value of Date field.
func (s *StoryView) GetDate() (value int) {
	if s == nil {
		return
	}
	return s.Date
}

// SetReaction sets value of Reaction conditional field.
func (s *StoryView) SetReaction(value ReactionClass) {
	s.Flags.Set(2)
	s.Reaction = value
}

// GetReaction returns value of Reaction conditional field and
// boolean which is true if field was set.
func (s *StoryView) GetReaction() (value ReactionClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.Reaction, true
}

// StoryViewPublicForward represents TL type `storyViewPublicForward#9083670b`.
type StoryViewPublicForward struct {
	// Flags field of StoryViewPublicForward.
	Flags bin.Fields
	// Blocked field of StoryViewPublicForward.
	Blocked bool
	// BlockedMyStoriesFrom field of StoryViewPublicForward.
	BlockedMyStoriesFrom bool
	// Message field of StoryViewPublicForward.
	Message MessageClass
}

// StoryViewPublicForwardTypeID is TL type id of StoryViewPublicForward.
const StoryViewPublicForwardTypeID = 0x9083670b

// construct implements constructor of StoryViewClass.
func (s StoryViewPublicForward) construct() StoryViewClass { return &s }

// Ensuring interfaces in compile-time for StoryViewPublicForward.
var (
	_ bin.Encoder     = &StoryViewPublicForward{}
	_ bin.Decoder     = &StoryViewPublicForward{}
	_ bin.BareEncoder = &StoryViewPublicForward{}
	_ bin.BareDecoder = &StoryViewPublicForward{}

	_ StoryViewClass = &StoryViewPublicForward{}
)

func (s *StoryViewPublicForward) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Blocked == false) {
		return false
	}
	if !(s.BlockedMyStoriesFrom == false) {
		return false
	}
	if !(s.Message == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryViewPublicForward) String() string {
	if s == nil {
		return "StoryViewPublicForward(nil)"
	}
	type Alias StoryViewPublicForward
	return fmt.Sprintf("StoryViewPublicForward%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryViewPublicForward) TypeID() uint32 {
	return StoryViewPublicForwardTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryViewPublicForward) TypeName() string {
	return "storyViewPublicForward"
}

// TypeInfo returns info about TL type.
func (s *StoryViewPublicForward) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyViewPublicForward",
		ID:   StoryViewPublicForwardTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Blocked",
			SchemaName: "blocked",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "BlockedMyStoriesFrom",
			SchemaName: "blocked_my_stories_from",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StoryViewPublicForward) SetFlags() {
	if !(s.Blocked == false) {
		s.Flags.Set(0)
	}
	if !(s.BlockedMyStoriesFrom == false) {
		s.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (s *StoryViewPublicForward) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyViewPublicForward#9083670b as nil")
	}
	b.PutID(StoryViewPublicForwardTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryViewPublicForward) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyViewPublicForward#9083670b as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyViewPublicForward#9083670b: field flags: %w", err)
	}
	if s.Message == nil {
		return fmt.Errorf("unable to encode storyViewPublicForward#9083670b: field message is nil")
	}
	if err := s.Message.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyViewPublicForward#9083670b: field message: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryViewPublicForward) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyViewPublicForward#9083670b to nil")
	}
	if err := b.ConsumeID(StoryViewPublicForwardTypeID); err != nil {
		return fmt.Errorf("unable to decode storyViewPublicForward#9083670b: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryViewPublicForward) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyViewPublicForward#9083670b to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyViewPublicForward#9083670b: field flags: %w", err)
		}
	}
	s.Blocked = s.Flags.Has(0)
	s.BlockedMyStoriesFrom = s.Flags.Has(1)
	{
		value, err := DecodeMessage(b)
		if err != nil {
			return fmt.Errorf("unable to decode storyViewPublicForward#9083670b: field message: %w", err)
		}
		s.Message = value
	}
	return nil
}

// SetBlocked sets value of Blocked conditional field.
func (s *StoryViewPublicForward) SetBlocked(value bool) {
	if value {
		s.Flags.Set(0)
		s.Blocked = true
	} else {
		s.Flags.Unset(0)
		s.Blocked = false
	}
}

// GetBlocked returns value of Blocked conditional field.
func (s *StoryViewPublicForward) GetBlocked() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// SetBlockedMyStoriesFrom sets value of BlockedMyStoriesFrom conditional field.
func (s *StoryViewPublicForward) SetBlockedMyStoriesFrom(value bool) {
	if value {
		s.Flags.Set(1)
		s.BlockedMyStoriesFrom = true
	} else {
		s.Flags.Unset(1)
		s.BlockedMyStoriesFrom = false
	}
}

// GetBlockedMyStoriesFrom returns value of BlockedMyStoriesFrom conditional field.
func (s *StoryViewPublicForward) GetBlockedMyStoriesFrom() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// GetMessage returns value of Message field.
func (s *StoryViewPublicForward) GetMessage() (value MessageClass) {
	if s == nil {
		return
	}
	return s.Message
}

// StoryViewPublicRepost represents TL type `storyViewPublicRepost#bd74cf49`.
type StoryViewPublicRepost struct {
	// Flags field of StoryViewPublicRepost.
	Flags bin.Fields
	// Blocked field of StoryViewPublicRepost.
	Blocked bool
	// BlockedMyStoriesFrom field of StoryViewPublicRepost.
	BlockedMyStoriesFrom bool
	// PeerID field of StoryViewPublicRepost.
	PeerID PeerClass
	// Story field of StoryViewPublicRepost.
	Story StoryItemClass
}

// StoryViewPublicRepostTypeID is TL type id of StoryViewPublicRepost.
const StoryViewPublicRepostTypeID = 0xbd74cf49

// construct implements constructor of StoryViewClass.
func (s StoryViewPublicRepost) construct() StoryViewClass { return &s }

// Ensuring interfaces in compile-time for StoryViewPublicRepost.
var (
	_ bin.Encoder     = &StoryViewPublicRepost{}
	_ bin.Decoder     = &StoryViewPublicRepost{}
	_ bin.BareEncoder = &StoryViewPublicRepost{}
	_ bin.BareDecoder = &StoryViewPublicRepost{}

	_ StoryViewClass = &StoryViewPublicRepost{}
)

func (s *StoryViewPublicRepost) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Blocked == false) {
		return false
	}
	if !(s.BlockedMyStoriesFrom == false) {
		return false
	}
	if !(s.PeerID == nil) {
		return false
	}
	if !(s.Story == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryViewPublicRepost) String() string {
	if s == nil {
		return "StoryViewPublicRepost(nil)"
	}
	type Alias StoryViewPublicRepost
	return fmt.Sprintf("StoryViewPublicRepost%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryViewPublicRepost) TypeID() uint32 {
	return StoryViewPublicRepostTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryViewPublicRepost) TypeName() string {
	return "storyViewPublicRepost"
}

// TypeInfo returns info about TL type.
func (s *StoryViewPublicRepost) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyViewPublicRepost",
		ID:   StoryViewPublicRepostTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Blocked",
			SchemaName: "blocked",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "BlockedMyStoriesFrom",
			SchemaName: "blocked_my_stories_from",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "PeerID",
			SchemaName: "peer_id",
		},
		{
			Name:       "Story",
			SchemaName: "story",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StoryViewPublicRepost) SetFlags() {
	if !(s.Blocked == false) {
		s.Flags.Set(0)
	}
	if !(s.BlockedMyStoriesFrom == false) {
		s.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (s *StoryViewPublicRepost) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyViewPublicRepost#bd74cf49 as nil")
	}
	b.PutID(StoryViewPublicRepostTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryViewPublicRepost) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyViewPublicRepost#bd74cf49 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyViewPublicRepost#bd74cf49: field flags: %w", err)
	}
	if s.PeerID == nil {
		return fmt.Errorf("unable to encode storyViewPublicRepost#bd74cf49: field peer_id is nil")
	}
	if err := s.PeerID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyViewPublicRepost#bd74cf49: field peer_id: %w", err)
	}
	if s.Story == nil {
		return fmt.Errorf("unable to encode storyViewPublicRepost#bd74cf49: field story is nil")
	}
	if err := s.Story.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyViewPublicRepost#bd74cf49: field story: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryViewPublicRepost) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyViewPublicRepost#bd74cf49 to nil")
	}
	if err := b.ConsumeID(StoryViewPublicRepostTypeID); err != nil {
		return fmt.Errorf("unable to decode storyViewPublicRepost#bd74cf49: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryViewPublicRepost) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyViewPublicRepost#bd74cf49 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyViewPublicRepost#bd74cf49: field flags: %w", err)
		}
	}
	s.Blocked = s.Flags.Has(0)
	s.BlockedMyStoriesFrom = s.Flags.Has(1)
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode storyViewPublicRepost#bd74cf49: field peer_id: %w", err)
		}
		s.PeerID = value
	}
	{
		value, err := DecodeStoryItem(b)
		if err != nil {
			return fmt.Errorf("unable to decode storyViewPublicRepost#bd74cf49: field story: %w", err)
		}
		s.Story = value
	}
	return nil
}

// SetBlocked sets value of Blocked conditional field.
func (s *StoryViewPublicRepost) SetBlocked(value bool) {
	if value {
		s.Flags.Set(0)
		s.Blocked = true
	} else {
		s.Flags.Unset(0)
		s.Blocked = false
	}
}

// GetBlocked returns value of Blocked conditional field.
func (s *StoryViewPublicRepost) GetBlocked() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// SetBlockedMyStoriesFrom sets value of BlockedMyStoriesFrom conditional field.
func (s *StoryViewPublicRepost) SetBlockedMyStoriesFrom(value bool) {
	if value {
		s.Flags.Set(1)
		s.BlockedMyStoriesFrom = true
	} else {
		s.Flags.Unset(1)
		s.BlockedMyStoriesFrom = false
	}
}

// GetBlockedMyStoriesFrom returns value of BlockedMyStoriesFrom conditional field.
func (s *StoryViewPublicRepost) GetBlockedMyStoriesFrom() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// GetPeerID returns value of PeerID field.
func (s *StoryViewPublicRepost) GetPeerID() (value PeerClass) {
	if s == nil {
		return
	}
	return s.PeerID
}

// GetStory returns value of Story field.
func (s *StoryViewPublicRepost) GetStory() (value StoryItemClass) {
	if s == nil {
		return
	}
	return s.Story
}

// StoryViewClassName is schema name of StoryViewClass.
const StoryViewClassName = "StoryView"

// StoryViewClass represents StoryView generic type.
//
// Constructors:
//   - [StoryView]
//   - [StoryViewPublicForward]
//   - [StoryViewPublicRepost]
//
// Example:
//
//	g, err := tg.DecodeStoryView(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.StoryView: // storyView#b0bdeac5
//	case *tg.StoryViewPublicForward: // storyViewPublicForward#9083670b
//	case *tg.StoryViewPublicRepost: // storyViewPublicRepost#bd74cf49
//	default: panic(v)
//	}
type StoryViewClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StoryViewClass

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

	// Blocked field of StoryView.
	GetBlocked() (value bool)
	// BlockedMyStoriesFrom field of StoryView.
	GetBlockedMyStoriesFrom() (value bool)
}

// DecodeStoryView implements binary de-serialization for StoryViewClass.
func DecodeStoryView(buf *bin.Buffer) (StoryViewClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StoryViewTypeID:
		// Decoding storyView#b0bdeac5.
		v := StoryView{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryViewClass: %w", err)
		}
		return &v, nil
	case StoryViewPublicForwardTypeID:
		// Decoding storyViewPublicForward#9083670b.
		v := StoryViewPublicForward{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryViewClass: %w", err)
		}
		return &v, nil
	case StoryViewPublicRepostTypeID:
		// Decoding storyViewPublicRepost#bd74cf49.
		v := StoryViewPublicRepost{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryViewClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoryViewClass: %w", bin.NewUnexpectedID(id))
	}
}

// StoryView boxes the StoryViewClass providing a helper.
type StoryViewBox struct {
	StoryView StoryViewClass
}

// Decode implements bin.Decoder for StoryViewBox.
func (b *StoryViewBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoryViewBox to nil")
	}
	v, err := DecodeStoryView(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StoryView = v
	return nil
}

// Encode implements bin.Encode for StoryViewBox.
func (b *StoryViewBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StoryView == nil {
		return fmt.Errorf("unable to encode StoryViewClass as nil")
	}
	return b.StoryView.Encode(buf)
}
