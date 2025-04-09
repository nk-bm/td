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

// StarGift represents TL type `starGift#49c577cd`.
type StarGift struct {
	// Flags field of StarGift.
	Flags bin.Fields
	// Limited field of StarGift.
	Limited bool
	// SoldOut field of StarGift.
	SoldOut bool
	// Birthday field of StarGift.
	Birthday bool
	// ID field of StarGift.
	ID int64
	// Sticker field of StarGift.
	Sticker DocumentClass
	// Stars field of StarGift.
	Stars int64
	// AvailabilityRemains field of StarGift.
	//
	// Use SetAvailabilityRemains and GetAvailabilityRemains helpers.
	AvailabilityRemains int
	// AvailabilityTotal field of StarGift.
	//
	// Use SetAvailabilityTotal and GetAvailabilityTotal helpers.
	AvailabilityTotal int
	// ConvertStars field of StarGift.
	ConvertStars int64
	// FirstSaleDate field of StarGift.
	//
	// Use SetFirstSaleDate and GetFirstSaleDate helpers.
	FirstSaleDate int
	// LastSaleDate field of StarGift.
	//
	// Use SetLastSaleDate and GetLastSaleDate helpers.
	LastSaleDate int
}

// StarGiftTypeID is TL type id of StarGift.
const StarGiftTypeID = 0x49c577cd

// Ensuring interfaces in compile-time for StarGift.
var (
	_ bin.Encoder     = &StarGift{}
	_ bin.Decoder     = &StarGift{}
	_ bin.BareEncoder = &StarGift{}
	_ bin.BareDecoder = &StarGift{}
)

func (s *StarGift) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Limited == false) {
		return false
	}
	if !(s.SoldOut == false) {
		return false
	}
	if !(s.Birthday == false) {
		return false
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.Sticker == nil) {
		return false
	}
	if !(s.Stars == 0) {
		return false
	}
	if !(s.AvailabilityRemains == 0) {
		return false
	}
	if !(s.AvailabilityTotal == 0) {
		return false
	}
	if !(s.ConvertStars == 0) {
		return false
	}
	if !(s.FirstSaleDate == 0) {
		return false
	}
	if !(s.LastSaleDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarGift) String() string {
	if s == nil {
		return "StarGift(nil)"
	}
	type Alias StarGift
	return fmt.Sprintf("StarGift%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarGift) TypeID() uint32 {
	return StarGiftTypeID
}

// TypeName returns name of type in TL schema.
func (*StarGift) TypeName() string {
	return "starGift"
}

// TypeInfo returns info about TL type.
func (s *StarGift) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starGift",
		ID:   StarGiftTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Limited",
			SchemaName: "limited",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "SoldOut",
			SchemaName: "sold_out",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Birthday",
			SchemaName: "birthday",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
		{
			Name:       "Stars",
			SchemaName: "stars",
		},
		{
			Name:       "AvailabilityRemains",
			SchemaName: "availability_remains",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "AvailabilityTotal",
			SchemaName: "availability_total",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "ConvertStars",
			SchemaName: "convert_stars",
		},
		{
			Name:       "FirstSaleDate",
			SchemaName: "first_sale_date",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "LastSaleDate",
			SchemaName: "last_sale_date",
			Null:       !s.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StarGift) SetFlags() {
	if !(s.Limited == false) {
		s.Flags.Set(0)
	}
	if !(s.SoldOut == false) {
		s.Flags.Set(1)
	}
	if !(s.Birthday == false) {
		s.Flags.Set(2)
	}
	if !(s.AvailabilityRemains == 0) {
		s.Flags.Set(0)
	}
	if !(s.AvailabilityTotal == 0) {
		s.Flags.Set(0)
	}
	if !(s.FirstSaleDate == 0) {
		s.Flags.Set(1)
	}
	if !(s.LastSaleDate == 0) {
		s.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (s *StarGift) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starGift#49c577cd as nil")
	}
	b.PutID(StarGiftTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarGift) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starGift#49c577cd as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starGift#49c577cd: field flags: %w", err)
	}
	b.PutLong(s.ID)
	if s.Sticker == nil {
		return fmt.Errorf("unable to encode starGift#49c577cd: field sticker is nil")
	}
	if err := s.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starGift#49c577cd: field sticker: %w", err)
	}
	b.PutLong(s.Stars)
	if s.Flags.Has(0) {
		b.PutInt(s.AvailabilityRemains)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.AvailabilityTotal)
	}
	b.PutLong(s.ConvertStars)
	if s.Flags.Has(1) {
		b.PutInt(s.FirstSaleDate)
	}
	if s.Flags.Has(1) {
		b.PutInt(s.LastSaleDate)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarGift) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starGift#49c577cd to nil")
	}
	if err := b.ConsumeID(StarGiftTypeID); err != nil {
		return fmt.Errorf("unable to decode starGift#49c577cd: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarGift) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starGift#49c577cd to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode starGift#49c577cd: field flags: %w", err)
		}
	}
	s.Limited = s.Flags.Has(0)
	s.SoldOut = s.Flags.Has(1)
	s.Birthday = s.Flags.Has(2)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode starGift#49c577cd: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode starGift#49c577cd: field sticker: %w", err)
		}
		s.Sticker = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode starGift#49c577cd: field stars: %w", err)
		}
		s.Stars = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGift#49c577cd: field availability_remains: %w", err)
		}
		s.AvailabilityRemains = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGift#49c577cd: field availability_total: %w", err)
		}
		s.AvailabilityTotal = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode starGift#49c577cd: field convert_stars: %w", err)
		}
		s.ConvertStars = value
	}
	if s.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGift#49c577cd: field first_sale_date: %w", err)
		}
		s.FirstSaleDate = value
	}
	if s.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGift#49c577cd: field last_sale_date: %w", err)
		}
		s.LastSaleDate = value
	}
	return nil
}

// SetLimited sets value of Limited conditional field.
func (s *StarGift) SetLimited(value bool) {
	if value {
		s.Flags.Set(0)
		s.Limited = true
	} else {
		s.Flags.Unset(0)
		s.Limited = false
	}
}

// GetLimited returns value of Limited conditional field.
func (s *StarGift) GetLimited() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// SetSoldOut sets value of SoldOut conditional field.
func (s *StarGift) SetSoldOut(value bool) {
	if value {
		s.Flags.Set(1)
		s.SoldOut = true
	} else {
		s.Flags.Unset(1)
		s.SoldOut = false
	}
}

// GetSoldOut returns value of SoldOut conditional field.
func (s *StarGift) GetSoldOut() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// SetBirthday sets value of Birthday conditional field.
func (s *StarGift) SetBirthday(value bool) {
	if value {
		s.Flags.Set(2)
		s.Birthday = true
	} else {
		s.Flags.Unset(2)
		s.Birthday = false
	}
}

// GetBirthday returns value of Birthday conditional field.
func (s *StarGift) GetBirthday() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(2)
}

// GetID returns value of ID field.
func (s *StarGift) GetID() (value int64) {
	if s == nil {
		return
	}
	return s.ID
}

// GetSticker returns value of Sticker field.
func (s *StarGift) GetSticker() (value DocumentClass) {
	if s == nil {
		return
	}
	return s.Sticker
}

// GetStars returns value of Stars field.
func (s *StarGift) GetStars() (value int64) {
	if s == nil {
		return
	}
	return s.Stars
}

// SetAvailabilityRemains sets value of AvailabilityRemains conditional field.
func (s *StarGift) SetAvailabilityRemains(value int) {
	s.Flags.Set(0)
	s.AvailabilityRemains = value
}

// GetAvailabilityRemains returns value of AvailabilityRemains conditional field and
// boolean which is true if field was set.
func (s *StarGift) GetAvailabilityRemains() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.AvailabilityRemains, true
}

// SetAvailabilityTotal sets value of AvailabilityTotal conditional field.
func (s *StarGift) SetAvailabilityTotal(value int) {
	s.Flags.Set(0)
	s.AvailabilityTotal = value
}

// GetAvailabilityTotal returns value of AvailabilityTotal conditional field and
// boolean which is true if field was set.
func (s *StarGift) GetAvailabilityTotal() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.AvailabilityTotal, true
}

// GetConvertStars returns value of ConvertStars field.
func (s *StarGift) GetConvertStars() (value int64) {
	if s == nil {
		return
	}
	return s.ConvertStars
}

// SetFirstSaleDate sets value of FirstSaleDate conditional field.
func (s *StarGift) SetFirstSaleDate(value int) {
	s.Flags.Set(1)
	s.FirstSaleDate = value
}

// GetFirstSaleDate returns value of FirstSaleDate conditional field and
// boolean which is true if field was set.
func (s *StarGift) GetFirstSaleDate() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.FirstSaleDate, true
}

// SetLastSaleDate sets value of LastSaleDate conditional field.
func (s *StarGift) SetLastSaleDate(value int) {
	s.Flags.Set(1)
	s.LastSaleDate = value
}

// GetLastSaleDate returns value of LastSaleDate conditional field and
// boolean which is true if field was set.
func (s *StarGift) GetLastSaleDate() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.LastSaleDate, true
}
