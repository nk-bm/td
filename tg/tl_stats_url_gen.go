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

// StatsURL represents TL type `statsURL#47a971e0`.
type StatsURL struct {
	// URL field of StatsURL.
	URL string
}

// StatsURLTypeID is TL type id of StatsURL.
const StatsURLTypeID = 0x47a971e0

// Ensuring interfaces in compile-time for StatsURL.
var (
	_ bin.Encoder     = &StatsURL{}
	_ bin.Decoder     = &StatsURL{}
	_ bin.BareEncoder = &StatsURL{}
	_ bin.BareDecoder = &StatsURL{}
)

func (s *StatsURL) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsURL) String() string {
	if s == nil {
		return "StatsURL(nil)"
	}
	type Alias StatsURL
	return fmt.Sprintf("StatsURL%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsURL) TypeID() uint32 {
	return StatsURLTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsURL) TypeName() string {
	return "statsURL"
}

// TypeInfo returns info about TL type.
func (s *StatsURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statsURL",
		ID:   StatsURLTypeID,
	}
	if s == nil {
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
func (s *StatsURL) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsURL#47a971e0 as nil")
	}
	b.PutID(StatsURLTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatsURL) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsURL#47a971e0 as nil")
	}
	b.PutString(s.URL)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsURL) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsURL#47a971e0 to nil")
	}
	if err := b.ConsumeID(StatsURLTypeID); err != nil {
		return fmt.Errorf("unable to decode statsURL#47a971e0: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatsURL) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsURL#47a971e0 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode statsURL#47a971e0: field url: %w", err)
		}
		s.URL = value
	}
	return nil
}

// GetURL returns value of URL field.
func (s *StatsURL) GetURL() (value string) {
	if s == nil {
		return
	}
	return s.URL
}
