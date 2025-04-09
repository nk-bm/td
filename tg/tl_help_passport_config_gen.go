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

// HelpPassportConfigNotModified represents TL type `help.passportConfigNotModified#bfb9f457`.
type HelpPassportConfigNotModified struct {
}

// HelpPassportConfigNotModifiedTypeID is TL type id of HelpPassportConfigNotModified.
const HelpPassportConfigNotModifiedTypeID = 0xbfb9f457

// construct implements constructor of HelpPassportConfigClass.
func (p HelpPassportConfigNotModified) construct() HelpPassportConfigClass { return &p }

// Ensuring interfaces in compile-time for HelpPassportConfigNotModified.
var (
	_ bin.Encoder     = &HelpPassportConfigNotModified{}
	_ bin.Decoder     = &HelpPassportConfigNotModified{}
	_ bin.BareEncoder = &HelpPassportConfigNotModified{}
	_ bin.BareDecoder = &HelpPassportConfigNotModified{}

	_ HelpPassportConfigClass = &HelpPassportConfigNotModified{}
)

func (p *HelpPassportConfigNotModified) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *HelpPassportConfigNotModified) String() string {
	if p == nil {
		return "HelpPassportConfigNotModified(nil)"
	}
	type Alias HelpPassportConfigNotModified
	return fmt.Sprintf("HelpPassportConfigNotModified%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpPassportConfigNotModified) TypeID() uint32 {
	return HelpPassportConfigNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpPassportConfigNotModified) TypeName() string {
	return "help.passportConfigNotModified"
}

// TypeInfo returns info about TL type.
func (p *HelpPassportConfigNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.passportConfigNotModified",
		ID:   HelpPassportConfigNotModifiedTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *HelpPassportConfigNotModified) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.passportConfigNotModified#bfb9f457 as nil")
	}
	b.PutID(HelpPassportConfigNotModifiedTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *HelpPassportConfigNotModified) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.passportConfigNotModified#bfb9f457 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *HelpPassportConfigNotModified) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.passportConfigNotModified#bfb9f457 to nil")
	}
	if err := b.ConsumeID(HelpPassportConfigNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode help.passportConfigNotModified#bfb9f457: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *HelpPassportConfigNotModified) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.passportConfigNotModified#bfb9f457 to nil")
	}
	return nil
}

// HelpPassportConfig represents TL type `help.passportConfig#a098d6af`.
type HelpPassportConfig struct {
	// Hash field of HelpPassportConfig.
	Hash int
	// CountriesLangs field of HelpPassportConfig.
	CountriesLangs DataJSON
}

// HelpPassportConfigTypeID is TL type id of HelpPassportConfig.
const HelpPassportConfigTypeID = 0xa098d6af

// construct implements constructor of HelpPassportConfigClass.
func (p HelpPassportConfig) construct() HelpPassportConfigClass { return &p }

// Ensuring interfaces in compile-time for HelpPassportConfig.
var (
	_ bin.Encoder     = &HelpPassportConfig{}
	_ bin.Decoder     = &HelpPassportConfig{}
	_ bin.BareEncoder = &HelpPassportConfig{}
	_ bin.BareDecoder = &HelpPassportConfig{}

	_ HelpPassportConfigClass = &HelpPassportConfig{}
)

func (p *HelpPassportConfig) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Hash == 0) {
		return false
	}
	if !(p.CountriesLangs.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *HelpPassportConfig) String() string {
	if p == nil {
		return "HelpPassportConfig(nil)"
	}
	type Alias HelpPassportConfig
	return fmt.Sprintf("HelpPassportConfig%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpPassportConfig) TypeID() uint32 {
	return HelpPassportConfigTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpPassportConfig) TypeName() string {
	return "help.passportConfig"
}

// TypeInfo returns info about TL type.
func (p *HelpPassportConfig) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.passportConfig",
		ID:   HelpPassportConfigTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "CountriesLangs",
			SchemaName: "countries_langs",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *HelpPassportConfig) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.passportConfig#a098d6af as nil")
	}
	b.PutID(HelpPassportConfigTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *HelpPassportConfig) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.passportConfig#a098d6af as nil")
	}
	b.PutInt(p.Hash)
	if err := p.CountriesLangs.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.passportConfig#a098d6af: field countries_langs: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *HelpPassportConfig) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.passportConfig#a098d6af to nil")
	}
	if err := b.ConsumeID(HelpPassportConfigTypeID); err != nil {
		return fmt.Errorf("unable to decode help.passportConfig#a098d6af: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *HelpPassportConfig) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.passportConfig#a098d6af to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.passportConfig#a098d6af: field hash: %w", err)
		}
		p.Hash = value
	}
	{
		if err := p.CountriesLangs.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.passportConfig#a098d6af: field countries_langs: %w", err)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (p *HelpPassportConfig) GetHash() (value int) {
	if p == nil {
		return
	}
	return p.Hash
}

// GetCountriesLangs returns value of CountriesLangs field.
func (p *HelpPassportConfig) GetCountriesLangs() (value DataJSON) {
	if p == nil {
		return
	}
	return p.CountriesLangs
}

// HelpPassportConfigClassName is schema name of HelpPassportConfigClass.
const HelpPassportConfigClassName = "help.PassportConfig"

// HelpPassportConfigClass represents help.PassportConfig generic type.
//
// Constructors:
//   - [HelpPassportConfigNotModified]
//   - [HelpPassportConfig]
//
// Example:
//
//	g, err := tg.DecodeHelpPassportConfig(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.HelpPassportConfigNotModified: // help.passportConfigNotModified#bfb9f457
//	case *tg.HelpPassportConfig: // help.passportConfig#a098d6af
//	default: panic(v)
//	}
type HelpPassportConfigClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() HelpPassportConfigClass

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
}

// DecodeHelpPassportConfig implements binary de-serialization for HelpPassportConfigClass.
func DecodeHelpPassportConfig(buf *bin.Buffer) (HelpPassportConfigClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpPassportConfigNotModifiedTypeID:
		// Decoding help.passportConfigNotModified#bfb9f457.
		v := HelpPassportConfigNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpPassportConfigClass: %w", err)
		}
		return &v, nil
	case HelpPassportConfigTypeID:
		// Decoding help.passportConfig#a098d6af.
		v := HelpPassportConfig{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpPassportConfigClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpPassportConfigClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpPassportConfig boxes the HelpPassportConfigClass providing a helper.
type HelpPassportConfigBox struct {
	PassportConfig HelpPassportConfigClass
}

// Decode implements bin.Decoder for HelpPassportConfigBox.
func (b *HelpPassportConfigBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpPassportConfigBox to nil")
	}
	v, err := DecodeHelpPassportConfig(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PassportConfig = v
	return nil
}

// Encode implements bin.Encode for HelpPassportConfigBox.
func (b *HelpPassportConfigBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PassportConfig == nil {
		return fmt.Errorf("unable to encode HelpPassportConfigClass as nil")
	}
	return b.PassportConfig.Encode(buf)
}
