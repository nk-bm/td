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

// SponsoredMessage represents TL type `sponsoredMessage#4d93a990`.
type SponsoredMessage struct {
	// Flags field of SponsoredMessage.
	Flags bin.Fields
	// Recommended field of SponsoredMessage.
	Recommended bool
	// CanReport field of SponsoredMessage.
	CanReport bool
	// RandomID field of SponsoredMessage.
	RandomID []byte
	// URL field of SponsoredMessage.
	URL string
	// Title field of SponsoredMessage.
	Title string
	// Message field of SponsoredMessage.
	Message string
	// Entities field of SponsoredMessage.
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// Photo field of SponsoredMessage.
	//
	// Use SetPhoto and GetPhoto helpers.
	Photo PhotoClass
	// Media field of SponsoredMessage.
	//
	// Use SetMedia and GetMedia helpers.
	Media MessageMediaClass
	// Color field of SponsoredMessage.
	//
	// Use SetColor and GetColor helpers.
	Color PeerColor
	// ButtonText field of SponsoredMessage.
	ButtonText string
	// SponsorInfo field of SponsoredMessage.
	//
	// Use SetSponsorInfo and GetSponsorInfo helpers.
	SponsorInfo string
	// AdditionalInfo field of SponsoredMessage.
	//
	// Use SetAdditionalInfo and GetAdditionalInfo helpers.
	AdditionalInfo string
}

// SponsoredMessageTypeID is TL type id of SponsoredMessage.
const SponsoredMessageTypeID = 0x4d93a990

// Ensuring interfaces in compile-time for SponsoredMessage.
var (
	_ bin.Encoder     = &SponsoredMessage{}
	_ bin.Decoder     = &SponsoredMessage{}
	_ bin.BareEncoder = &SponsoredMessage{}
	_ bin.BareDecoder = &SponsoredMessage{}
)

func (s *SponsoredMessage) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Recommended == false) {
		return false
	}
	if !(s.CanReport == false) {
		return false
	}
	if !(s.RandomID == nil) {
		return false
	}
	if !(s.URL == "") {
		return false
	}
	if !(s.Title == "") {
		return false
	}
	if !(s.Message == "") {
		return false
	}
	if !(s.Entities == nil) {
		return false
	}
	if !(s.Photo == nil) {
		return false
	}
	if !(s.Media == nil) {
		return false
	}
	if !(s.Color.Zero()) {
		return false
	}
	if !(s.ButtonText == "") {
		return false
	}
	if !(s.SponsorInfo == "") {
		return false
	}
	if !(s.AdditionalInfo == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SponsoredMessage) String() string {
	if s == nil {
		return "SponsoredMessage(nil)"
	}
	type Alias SponsoredMessage
	return fmt.Sprintf("SponsoredMessage%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SponsoredMessage) TypeID() uint32 {
	return SponsoredMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*SponsoredMessage) TypeName() string {
	return "sponsoredMessage"
}

// TypeInfo returns info about TL type.
func (s *SponsoredMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sponsoredMessage",
		ID:   SponsoredMessageTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Recommended",
			SchemaName: "recommended",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "CanReport",
			SchemaName: "can_report",
			Null:       !s.Flags.Has(12),
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
			Null:       !s.Flags.Has(6),
		},
		{
			Name:       "Media",
			SchemaName: "media",
			Null:       !s.Flags.Has(14),
		},
		{
			Name:       "Color",
			SchemaName: "color",
			Null:       !s.Flags.Has(13),
		},
		{
			Name:       "ButtonText",
			SchemaName: "button_text",
		},
		{
			Name:       "SponsorInfo",
			SchemaName: "sponsor_info",
			Null:       !s.Flags.Has(7),
		},
		{
			Name:       "AdditionalInfo",
			SchemaName: "additional_info",
			Null:       !s.Flags.Has(8),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *SponsoredMessage) SetFlags() {
	if !(s.Recommended == false) {
		s.Flags.Set(5)
	}
	if !(s.CanReport == false) {
		s.Flags.Set(12)
	}
	if !(s.Entities == nil) {
		s.Flags.Set(1)
	}
	if !(s.Photo == nil) {
		s.Flags.Set(6)
	}
	if !(s.Media == nil) {
		s.Flags.Set(14)
	}
	if !(s.Color.Zero()) {
		s.Flags.Set(13)
	}
	if !(s.SponsorInfo == "") {
		s.Flags.Set(7)
	}
	if !(s.AdditionalInfo == "") {
		s.Flags.Set(8)
	}
}

// Encode implements bin.Encoder.
func (s *SponsoredMessage) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessage#4d93a990 as nil")
	}
	b.PutID(SponsoredMessageTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SponsoredMessage) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessage#4d93a990 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#4d93a990: field flags: %w", err)
	}
	b.PutBytes(s.RandomID)
	b.PutString(s.URL)
	b.PutString(s.Title)
	b.PutString(s.Message)
	if s.Flags.Has(1) {
		b.PutVectorHeader(len(s.Entities))
		for idx, v := range s.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode sponsoredMessage#4d93a990: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode sponsoredMessage#4d93a990: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(6) {
		if s.Photo == nil {
			return fmt.Errorf("unable to encode sponsoredMessage#4d93a990: field photo is nil")
		}
		if err := s.Photo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode sponsoredMessage#4d93a990: field photo: %w", err)
		}
	}
	if s.Flags.Has(14) {
		if s.Media == nil {
			return fmt.Errorf("unable to encode sponsoredMessage#4d93a990: field media is nil")
		}
		if err := s.Media.Encode(b); err != nil {
			return fmt.Errorf("unable to encode sponsoredMessage#4d93a990: field media: %w", err)
		}
	}
	if s.Flags.Has(13) {
		if err := s.Color.Encode(b); err != nil {
			return fmt.Errorf("unable to encode sponsoredMessage#4d93a990: field color: %w", err)
		}
	}
	b.PutString(s.ButtonText)
	if s.Flags.Has(7) {
		b.PutString(s.SponsorInfo)
	}
	if s.Flags.Has(8) {
		b.PutString(s.AdditionalInfo)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SponsoredMessage) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessage#4d93a990 to nil")
	}
	if err := b.ConsumeID(SponsoredMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SponsoredMessage) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessage#4d93a990 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field flags: %w", err)
		}
	}
	s.Recommended = s.Flags.Has(5)
	s.CanReport = s.Flags.Has(12)
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field url: %w", err)
		}
		s.URL = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field title: %w", err)
		}
		s.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field message: %w", err)
		}
		s.Message = value
	}
	if s.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field entities: %w", err)
		}

		if headerLen > 0 {
			s.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field entities: %w", err)
			}
			s.Entities = append(s.Entities, value)
		}
	}
	if s.Flags.Has(6) {
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field photo: %w", err)
		}
		s.Photo = value
	}
	if s.Flags.Has(14) {
		value, err := DecodeMessageMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field media: %w", err)
		}
		s.Media = value
	}
	if s.Flags.Has(13) {
		if err := s.Color.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field color: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field button_text: %w", err)
		}
		s.ButtonText = value
	}
	if s.Flags.Has(7) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field sponsor_info: %w", err)
		}
		s.SponsorInfo = value
	}
	if s.Flags.Has(8) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#4d93a990: field additional_info: %w", err)
		}
		s.AdditionalInfo = value
	}
	return nil
}

// SetRecommended sets value of Recommended conditional field.
func (s *SponsoredMessage) SetRecommended(value bool) {
	if value {
		s.Flags.Set(5)
		s.Recommended = true
	} else {
		s.Flags.Unset(5)
		s.Recommended = false
	}
}

// GetRecommended returns value of Recommended conditional field.
func (s *SponsoredMessage) GetRecommended() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(5)
}

// SetCanReport sets value of CanReport conditional field.
func (s *SponsoredMessage) SetCanReport(value bool) {
	if value {
		s.Flags.Set(12)
		s.CanReport = true
	} else {
		s.Flags.Unset(12)
		s.CanReport = false
	}
}

// GetCanReport returns value of CanReport conditional field.
func (s *SponsoredMessage) GetCanReport() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(12)
}

// GetRandomID returns value of RandomID field.
func (s *SponsoredMessage) GetRandomID() (value []byte) {
	if s == nil {
		return
	}
	return s.RandomID
}

// GetURL returns value of URL field.
func (s *SponsoredMessage) GetURL() (value string) {
	if s == nil {
		return
	}
	return s.URL
}

// GetTitle returns value of Title field.
func (s *SponsoredMessage) GetTitle() (value string) {
	if s == nil {
		return
	}
	return s.Title
}

// GetMessage returns value of Message field.
func (s *SponsoredMessage) GetMessage() (value string) {
	if s == nil {
		return
	}
	return s.Message
}

// SetEntities sets value of Entities conditional field.
func (s *SponsoredMessage) SetEntities(value []MessageEntityClass) {
	s.Flags.Set(1)
	s.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetEntities() (value []MessageEntityClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.Entities, true
}

// SetPhoto sets value of Photo conditional field.
func (s *SponsoredMessage) SetPhoto(value PhotoClass) {
	s.Flags.Set(6)
	s.Photo = value
}

// GetPhoto returns value of Photo conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetPhoto() (value PhotoClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(6) {
		return value, false
	}
	return s.Photo, true
}

// SetMedia sets value of Media conditional field.
func (s *SponsoredMessage) SetMedia(value MessageMediaClass) {
	s.Flags.Set(14)
	s.Media = value
}

// GetMedia returns value of Media conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetMedia() (value MessageMediaClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(14) {
		return value, false
	}
	return s.Media, true
}

// SetColor sets value of Color conditional field.
func (s *SponsoredMessage) SetColor(value PeerColor) {
	s.Flags.Set(13)
	s.Color = value
}

// GetColor returns value of Color conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetColor() (value PeerColor, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(13) {
		return value, false
	}
	return s.Color, true
}

// GetButtonText returns value of ButtonText field.
func (s *SponsoredMessage) GetButtonText() (value string) {
	if s == nil {
		return
	}
	return s.ButtonText
}

// SetSponsorInfo sets value of SponsorInfo conditional field.
func (s *SponsoredMessage) SetSponsorInfo(value string) {
	s.Flags.Set(7)
	s.SponsorInfo = value
}

// GetSponsorInfo returns value of SponsorInfo conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetSponsorInfo() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(7) {
		return value, false
	}
	return s.SponsorInfo, true
}

// SetAdditionalInfo sets value of AdditionalInfo conditional field.
func (s *SponsoredMessage) SetAdditionalInfo(value string) {
	s.Flags.Set(8)
	s.AdditionalInfo = value
}

// GetAdditionalInfo returns value of AdditionalInfo conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetAdditionalInfo() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(8) {
		return value, false
	}
	return s.AdditionalInfo, true
}
