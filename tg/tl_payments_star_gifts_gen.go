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

// PaymentsStarGiftsNotModified represents TL type `payments.starGiftsNotModified#a388a368`.
// The list of available gifts »¹ hasn't changed.
//
// Links:
//  1. https://core.telegram.org/api/gifts
//
// See https://core.telegram.org/constructor/payments.starGiftsNotModified for reference.
type PaymentsStarGiftsNotModified struct {
}

// PaymentsStarGiftsNotModifiedTypeID is TL type id of PaymentsStarGiftsNotModified.
const PaymentsStarGiftsNotModifiedTypeID = 0xa388a368

// construct implements constructor of PaymentsStarGiftsClass.
func (s PaymentsStarGiftsNotModified) construct() PaymentsStarGiftsClass { return &s }

// Ensuring interfaces in compile-time for PaymentsStarGiftsNotModified.
var (
	_ bin.Encoder     = &PaymentsStarGiftsNotModified{}
	_ bin.Decoder     = &PaymentsStarGiftsNotModified{}
	_ bin.BareEncoder = &PaymentsStarGiftsNotModified{}
	_ bin.BareDecoder = &PaymentsStarGiftsNotModified{}

	_ PaymentsStarGiftsClass = &PaymentsStarGiftsNotModified{}
)

func (s *PaymentsStarGiftsNotModified) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *PaymentsStarGiftsNotModified) String() string {
	if s == nil {
		return "PaymentsStarGiftsNotModified(nil)"
	}
	type Alias PaymentsStarGiftsNotModified
	return fmt.Sprintf("PaymentsStarGiftsNotModified%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsStarGiftsNotModified) TypeID() uint32 {
	return PaymentsStarGiftsNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsStarGiftsNotModified) TypeName() string {
	return "payments.starGiftsNotModified"
}

// TypeInfo returns info about TL type.
func (s *PaymentsStarGiftsNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.starGiftsNotModified",
		ID:   PaymentsStarGiftsNotModifiedTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *PaymentsStarGiftsNotModified) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.starGiftsNotModified#a388a368 as nil")
	}
	b.PutID(PaymentsStarGiftsNotModifiedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *PaymentsStarGiftsNotModified) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.starGiftsNotModified#a388a368 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *PaymentsStarGiftsNotModified) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.starGiftsNotModified#a388a368 to nil")
	}
	if err := b.ConsumeID(PaymentsStarGiftsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.starGiftsNotModified#a388a368: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *PaymentsStarGiftsNotModified) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.starGiftsNotModified#a388a368 to nil")
	}
	return nil
}

// PaymentsStarGifts represents TL type `payments.starGifts#901689ea`.
// Available gifts »¹.
//
// Links:
//  1. https://core.telegram.org/api/gifts
//
// See https://core.telegram.org/constructor/payments.starGifts for reference.
type PaymentsStarGifts struct {
	// Hash used for caching, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// List of available gifts.
	Gifts []StarGift
}

// PaymentsStarGiftsTypeID is TL type id of PaymentsStarGifts.
const PaymentsStarGiftsTypeID = 0x901689ea

// construct implements constructor of PaymentsStarGiftsClass.
func (s PaymentsStarGifts) construct() PaymentsStarGiftsClass { return &s }

// Ensuring interfaces in compile-time for PaymentsStarGifts.
var (
	_ bin.Encoder     = &PaymentsStarGifts{}
	_ bin.Decoder     = &PaymentsStarGifts{}
	_ bin.BareEncoder = &PaymentsStarGifts{}
	_ bin.BareDecoder = &PaymentsStarGifts{}

	_ PaymentsStarGiftsClass = &PaymentsStarGifts{}
)

func (s *PaymentsStarGifts) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Hash == 0) {
		return false
	}
	if !(s.Gifts == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PaymentsStarGifts) String() string {
	if s == nil {
		return "PaymentsStarGifts(nil)"
	}
	type Alias PaymentsStarGifts
	return fmt.Sprintf("PaymentsStarGifts%+v", Alias(*s))
}

// FillFrom fills PaymentsStarGifts from given interface.
func (s *PaymentsStarGifts) FillFrom(from interface {
	GetHash() (value int)
	GetGifts() (value []StarGift)
}) {
	s.Hash = from.GetHash()
	s.Gifts = from.GetGifts()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsStarGifts) TypeID() uint32 {
	return PaymentsStarGiftsTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsStarGifts) TypeName() string {
	return "payments.starGifts"
}

// TypeInfo returns info about TL type.
func (s *PaymentsStarGifts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.starGifts",
		ID:   PaymentsStarGiftsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Gifts",
			SchemaName: "gifts",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *PaymentsStarGifts) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.starGifts#901689ea as nil")
	}
	b.PutID(PaymentsStarGiftsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *PaymentsStarGifts) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.starGifts#901689ea as nil")
	}
	b.PutInt(s.Hash)
	b.PutVectorHeader(len(s.Gifts))
	for idx, v := range s.Gifts {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.starGifts#901689ea: field gifts element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *PaymentsStarGifts) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.starGifts#901689ea to nil")
	}
	if err := b.ConsumeID(PaymentsStarGiftsTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.starGifts#901689ea: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *PaymentsStarGifts) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.starGifts#901689ea to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.starGifts#901689ea: field hash: %w", err)
		}
		s.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.starGifts#901689ea: field gifts: %w", err)
		}

		if headerLen > 0 {
			s.Gifts = make([]StarGift, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StarGift
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode payments.starGifts#901689ea: field gifts: %w", err)
			}
			s.Gifts = append(s.Gifts, value)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (s *PaymentsStarGifts) GetHash() (value int) {
	if s == nil {
		return
	}
	return s.Hash
}

// GetGifts returns value of Gifts field.
func (s *PaymentsStarGifts) GetGifts() (value []StarGift) {
	if s == nil {
		return
	}
	return s.Gifts
}

// PaymentsStarGiftsClassName is schema name of PaymentsStarGiftsClass.
const PaymentsStarGiftsClassName = "payments.StarGifts"

// PaymentsStarGiftsClass represents payments.StarGifts generic type.
//
// See https://core.telegram.org/type/payments.StarGifts for reference.
//
// Constructors:
//   - [PaymentsStarGiftsNotModified]
//   - [PaymentsStarGifts]
//
// Example:
//
//	g, err := tg.DecodePaymentsStarGifts(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.PaymentsStarGiftsNotModified: // payments.starGiftsNotModified#a388a368
//	case *tg.PaymentsStarGifts: // payments.starGifts#901689ea
//	default: panic(v)
//	}
type PaymentsStarGiftsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PaymentsStarGiftsClass

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

	// AsModified tries to map PaymentsStarGiftsClass to PaymentsStarGifts.
	AsModified() (*PaymentsStarGifts, bool)
}

// AsModified tries to map PaymentsStarGiftsNotModified to PaymentsStarGifts.
func (s *PaymentsStarGiftsNotModified) AsModified() (*PaymentsStarGifts, bool) {
	return nil, false
}

// AsModified tries to map PaymentsStarGifts to PaymentsStarGifts.
func (s *PaymentsStarGifts) AsModified() (*PaymentsStarGifts, bool) {
	return s, true
}

// DecodePaymentsStarGifts implements binary de-serialization for PaymentsStarGiftsClass.
func DecodePaymentsStarGifts(buf *bin.Buffer) (PaymentsStarGiftsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PaymentsStarGiftsNotModifiedTypeID:
		// Decoding payments.starGiftsNotModified#a388a368.
		v := PaymentsStarGiftsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PaymentsStarGiftsClass: %w", err)
		}
		return &v, nil
	case PaymentsStarGiftsTypeID:
		// Decoding payments.starGifts#901689ea.
		v := PaymentsStarGifts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PaymentsStarGiftsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PaymentsStarGiftsClass: %w", bin.NewUnexpectedID(id))
	}
}

// PaymentsStarGifts boxes the PaymentsStarGiftsClass providing a helper.
type PaymentsStarGiftsBox struct {
	StarGifts PaymentsStarGiftsClass
}

// Decode implements bin.Decoder for PaymentsStarGiftsBox.
func (b *PaymentsStarGiftsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PaymentsStarGiftsBox to nil")
	}
	v, err := DecodePaymentsStarGifts(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StarGifts = v
	return nil
}

// Encode implements bin.Encode for PaymentsStarGiftsBox.
func (b *PaymentsStarGiftsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StarGifts == nil {
		return fmt.Errorf("unable to encode PaymentsStarGiftsClass as nil")
	}
	return b.StarGifts.Encode(buf)
}
