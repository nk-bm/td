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

// StoryFwdHeader represents TL type `storyFwdHeader#b826e150`.
type StoryFwdHeader struct {
	// Flags field of StoryFwdHeader.
	Flags bin.Fields
	// Modified field of StoryFwdHeader.
	Modified bool
	// From field of StoryFwdHeader.
	//
	// Use SetFrom and GetFrom helpers.
	From PeerClass
	// FromName field of StoryFwdHeader.
	//
	// Use SetFromName and GetFromName helpers.
	FromName string
	// StoryID field of StoryFwdHeader.
	//
	// Use SetStoryID and GetStoryID helpers.
	StoryID int
}

// StoryFwdHeaderTypeID is TL type id of StoryFwdHeader.
const StoryFwdHeaderTypeID = 0xb826e150

// Ensuring interfaces in compile-time for StoryFwdHeader.
var (
	_ bin.Encoder     = &StoryFwdHeader{}
	_ bin.Decoder     = &StoryFwdHeader{}
	_ bin.BareEncoder = &StoryFwdHeader{}
	_ bin.BareDecoder = &StoryFwdHeader{}
)

func (s *StoryFwdHeader) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Modified == false) {
		return false
	}
	if !(s.From == nil) {
		return false
	}
	if !(s.FromName == "") {
		return false
	}
	if !(s.StoryID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryFwdHeader) String() string {
	if s == nil {
		return "StoryFwdHeader(nil)"
	}
	type Alias StoryFwdHeader
	return fmt.Sprintf("StoryFwdHeader%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryFwdHeader) TypeID() uint32 {
	return StoryFwdHeaderTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryFwdHeader) TypeName() string {
	return "storyFwdHeader"
}

// TypeInfo returns info about TL type.
func (s *StoryFwdHeader) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyFwdHeader",
		ID:   StoryFwdHeaderTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Modified",
			SchemaName: "modified",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "From",
			SchemaName: "from",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "FromName",
			SchemaName: "from_name",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "StoryID",
			SchemaName: "story_id",
			Null:       !s.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StoryFwdHeader) SetFlags() {
	if !(s.Modified == false) {
		s.Flags.Set(3)
	}
	if !(s.From == nil) {
		s.Flags.Set(0)
	}
	if !(s.FromName == "") {
		s.Flags.Set(1)
	}
	if !(s.StoryID == 0) {
		s.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (s *StoryFwdHeader) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyFwdHeader#b826e150 as nil")
	}
	b.PutID(StoryFwdHeaderTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryFwdHeader) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyFwdHeader#b826e150 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyFwdHeader#b826e150: field flags: %w", err)
	}
	if s.Flags.Has(0) {
		if s.From == nil {
			return fmt.Errorf("unable to encode storyFwdHeader#b826e150: field from is nil")
		}
		if err := s.From.Encode(b); err != nil {
			return fmt.Errorf("unable to encode storyFwdHeader#b826e150: field from: %w", err)
		}
	}
	if s.Flags.Has(1) {
		b.PutString(s.FromName)
	}
	if s.Flags.Has(2) {
		b.PutInt(s.StoryID)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryFwdHeader) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyFwdHeader#b826e150 to nil")
	}
	if err := b.ConsumeID(StoryFwdHeaderTypeID); err != nil {
		return fmt.Errorf("unable to decode storyFwdHeader#b826e150: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryFwdHeader) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyFwdHeader#b826e150 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyFwdHeader#b826e150: field flags: %w", err)
		}
	}
	s.Modified = s.Flags.Has(3)
	if s.Flags.Has(0) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode storyFwdHeader#b826e150: field from: %w", err)
		}
		s.From = value
	}
	if s.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode storyFwdHeader#b826e150: field from_name: %w", err)
		}
		s.FromName = value
	}
	if s.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyFwdHeader#b826e150: field story_id: %w", err)
		}
		s.StoryID = value
	}
	return nil
}

// SetModified sets value of Modified conditional field.
func (s *StoryFwdHeader) SetModified(value bool) {
	if value {
		s.Flags.Set(3)
		s.Modified = true
	} else {
		s.Flags.Unset(3)
		s.Modified = false
	}
}

// GetModified returns value of Modified conditional field.
func (s *StoryFwdHeader) GetModified() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(3)
}

// SetFrom sets value of From conditional field.
func (s *StoryFwdHeader) SetFrom(value PeerClass) {
	s.Flags.Set(0)
	s.From = value
}

// GetFrom returns value of From conditional field and
// boolean which is true if field was set.
func (s *StoryFwdHeader) GetFrom() (value PeerClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.From, true
}

// SetFromName sets value of FromName conditional field.
func (s *StoryFwdHeader) SetFromName(value string) {
	s.Flags.Set(1)
	s.FromName = value
}

// GetFromName returns value of FromName conditional field and
// boolean which is true if field was set.
func (s *StoryFwdHeader) GetFromName() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.FromName, true
}

// SetStoryID sets value of StoryID conditional field.
func (s *StoryFwdHeader) SetStoryID(value int) {
	s.Flags.Set(2)
	s.StoryID = value
}

// GetStoryID returns value of StoryID conditional field and
// boolean which is true if field was set.
func (s *StoryFwdHeader) GetStoryID() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.StoryID, true
}
