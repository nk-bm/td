// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// GeoPointEmpty represents TL type `geoPointEmpty#1117dd5f`.
// Empty constructor.
//
// See https://core.telegram.org/constructor/geoPointEmpty for reference.
type GeoPointEmpty struct {
}

// GeoPointEmptyTypeID is TL type id of GeoPointEmpty.
const GeoPointEmptyTypeID = 0x1117dd5f

func (g *GeoPointEmpty) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GeoPointEmpty) String() string {
	if g == nil {
		return "GeoPointEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("GeoPointEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *GeoPointEmpty) TypeID() uint32 {
	return GeoPointEmptyTypeID
}

// Encode implements bin.Encoder.
func (g *GeoPointEmpty) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode geoPointEmpty#1117dd5f as nil")
	}
	b.PutID(GeoPointEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GeoPointEmpty) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode geoPointEmpty#1117dd5f to nil")
	}
	if err := b.ConsumeID(GeoPointEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode geoPointEmpty#1117dd5f: %w", err)
	}
	return nil
}

// construct implements constructor of GeoPointClass.
func (g GeoPointEmpty) construct() GeoPointClass { return &g }

// Ensuring interfaces in compile-time for GeoPointEmpty.
var (
	_ bin.Encoder = &GeoPointEmpty{}
	_ bin.Decoder = &GeoPointEmpty{}

	_ GeoPointClass = &GeoPointEmpty{}
)

// GeoPoint represents TL type `geoPoint#b2a2f663`.
// GeoPoint.
//
// See https://core.telegram.org/constructor/geoPoint for reference.
type GeoPoint struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Longtitude
	Long float64
	// Latitude
	Lat float64
	// Access hash
	AccessHash int64
	// The estimated horizontal accuracy of the location, in meters; as defined by the sender.
	//
	// Use SetAccuracyRadius and GetAccuracyRadius helpers.
	AccuracyRadius int
}

// GeoPointTypeID is TL type id of GeoPoint.
const GeoPointTypeID = 0xb2a2f663

func (g *GeoPoint) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Long == 0) {
		return false
	}
	if !(g.Lat == 0) {
		return false
	}
	if !(g.AccessHash == 0) {
		return false
	}
	if !(g.AccuracyRadius == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GeoPoint) String() string {
	if g == nil {
		return "GeoPoint(nil)"
	}
	var sb strings.Builder
	sb.WriteString("GeoPoint")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(g.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tLong: ")
	sb.WriteString(fmt.Sprint(g.Long))
	sb.WriteString(",\n")
	sb.WriteString("\tLat: ")
	sb.WriteString(fmt.Sprint(g.Lat))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(g.AccessHash))
	sb.WriteString(",\n")
	if g.Flags.Has(0) {
		sb.WriteString("\tAccuracyRadius: ")
		sb.WriteString(fmt.Sprint(g.AccuracyRadius))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *GeoPoint) TypeID() uint32 {
	return GeoPointTypeID
}

// Encode implements bin.Encoder.
func (g *GeoPoint) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode geoPoint#b2a2f663 as nil")
	}
	b.PutID(GeoPointTypeID)
	if !(g.AccuracyRadius == 0) {
		g.Flags.Set(0)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode geoPoint#b2a2f663: field flags: %w", err)
	}
	b.PutDouble(g.Long)
	b.PutDouble(g.Lat)
	b.PutLong(g.AccessHash)
	if g.Flags.Has(0) {
		b.PutInt(g.AccuracyRadius)
	}
	return nil
}

// GetLong returns value of Long field.
func (g *GeoPoint) GetLong() (value float64) {
	return g.Long
}

// GetLat returns value of Lat field.
func (g *GeoPoint) GetLat() (value float64) {
	return g.Lat
}

// GetAccessHash returns value of AccessHash field.
func (g *GeoPoint) GetAccessHash() (value int64) {
	return g.AccessHash
}

// SetAccuracyRadius sets value of AccuracyRadius conditional field.
func (g *GeoPoint) SetAccuracyRadius(value int) {
	g.Flags.Set(0)
	g.AccuracyRadius = value
}

// GetAccuracyRadius returns value of AccuracyRadius conditional field and
// boolean which is true if field was set.
func (g *GeoPoint) GetAccuracyRadius() (value int, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.AccuracyRadius, true
}

// Decode implements bin.Decoder.
func (g *GeoPoint) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode geoPoint#b2a2f663 to nil")
	}
	if err := b.ConsumeID(GeoPointTypeID); err != nil {
		return fmt.Errorf("unable to decode geoPoint#b2a2f663: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode geoPoint#b2a2f663: field flags: %w", err)
		}
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode geoPoint#b2a2f663: field long: %w", err)
		}
		g.Long = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode geoPoint#b2a2f663: field lat: %w", err)
		}
		g.Lat = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode geoPoint#b2a2f663: field access_hash: %w", err)
		}
		g.AccessHash = value
	}
	if g.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode geoPoint#b2a2f663: field accuracy_radius: %w", err)
		}
		g.AccuracyRadius = value
	}
	return nil
}

// construct implements constructor of GeoPointClass.
func (g GeoPoint) construct() GeoPointClass { return &g }

// Ensuring interfaces in compile-time for GeoPoint.
var (
	_ bin.Encoder = &GeoPoint{}
	_ bin.Decoder = &GeoPoint{}

	_ GeoPointClass = &GeoPoint{}
)

// GeoPointClass represents GeoPoint generic type.
//
// See https://core.telegram.org/type/GeoPoint for reference.
//
// Example:
//  g, err := DecodeGeoPoint(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *GeoPointEmpty: // geoPointEmpty#1117dd5f
//  case *GeoPoint: // geoPoint#b2a2f663
//  default: panic(v)
//  }
type GeoPointClass interface {
	bin.Encoder
	bin.Decoder
	construct() GeoPointClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeGeoPoint implements binary de-serialization for GeoPointClass.
func DecodeGeoPoint(buf *bin.Buffer) (GeoPointClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case GeoPointEmptyTypeID:
		// Decoding geoPointEmpty#1117dd5f.
		v := GeoPointEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode GeoPointClass: %w", err)
		}
		return &v, nil
	case GeoPointTypeID:
		// Decoding geoPoint#b2a2f663.
		v := GeoPoint{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode GeoPointClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode GeoPointClass: %w", bin.NewUnexpectedID(id))
	}
}

// GeoPoint boxes the GeoPointClass providing a helper.
type GeoPointBox struct {
	GeoPoint GeoPointClass
}

// Decode implements bin.Decoder for GeoPointBox.
func (b *GeoPointBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode GeoPointBox to nil")
	}
	v, err := DecodeGeoPoint(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.GeoPoint = v
	return nil
}

// Encode implements bin.Encode for GeoPointBox.
func (b *GeoPointBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.GeoPoint == nil {
		return fmt.Errorf("unable to encode GeoPointClass as nil")
	}
	return b.GeoPoint.Encode(buf)
}
