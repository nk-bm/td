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

// AccountThemesNotModified represents TL type `account.themesNotModified#f41eb622`.
type AccountThemesNotModified struct {
}

// AccountThemesNotModifiedTypeID is TL type id of AccountThemesNotModified.
const AccountThemesNotModifiedTypeID = 0xf41eb622

// construct implements constructor of AccountThemesClass.
func (t AccountThemesNotModified) construct() AccountThemesClass { return &t }

// Ensuring interfaces in compile-time for AccountThemesNotModified.
var (
	_ bin.Encoder     = &AccountThemesNotModified{}
	_ bin.Decoder     = &AccountThemesNotModified{}
	_ bin.BareEncoder = &AccountThemesNotModified{}
	_ bin.BareDecoder = &AccountThemesNotModified{}

	_ AccountThemesClass = &AccountThemesNotModified{}
)

func (t *AccountThemesNotModified) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *AccountThemesNotModified) String() string {
	if t == nil {
		return "AccountThemesNotModified(nil)"
	}
	type Alias AccountThemesNotModified
	return fmt.Sprintf("AccountThemesNotModified%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountThemesNotModified) TypeID() uint32 {
	return AccountThemesNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountThemesNotModified) TypeName() string {
	return "account.themesNotModified"
}

// TypeInfo returns info about TL type.
func (t *AccountThemesNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.themesNotModified",
		ID:   AccountThemesNotModifiedTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *AccountThemesNotModified) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.themesNotModified#f41eb622 as nil")
	}
	b.PutID(AccountThemesNotModifiedTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *AccountThemesNotModified) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.themesNotModified#f41eb622 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *AccountThemesNotModified) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.themesNotModified#f41eb622 to nil")
	}
	if err := b.ConsumeID(AccountThemesNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode account.themesNotModified#f41eb622: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *AccountThemesNotModified) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.themesNotModified#f41eb622 to nil")
	}
	return nil
}

// AccountThemes represents TL type `account.themes#9a3d8c6d`.
type AccountThemes struct {
	// Hash field of AccountThemes.
	Hash int64
	// Themes field of AccountThemes.
	Themes []Theme
}

// AccountThemesTypeID is TL type id of AccountThemes.
const AccountThemesTypeID = 0x9a3d8c6d

// construct implements constructor of AccountThemesClass.
func (t AccountThemes) construct() AccountThemesClass { return &t }

// Ensuring interfaces in compile-time for AccountThemes.
var (
	_ bin.Encoder     = &AccountThemes{}
	_ bin.Decoder     = &AccountThemes{}
	_ bin.BareEncoder = &AccountThemes{}
	_ bin.BareDecoder = &AccountThemes{}

	_ AccountThemesClass = &AccountThemes{}
)

func (t *AccountThemes) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Hash == 0) {
		return false
	}
	if !(t.Themes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *AccountThemes) String() string {
	if t == nil {
		return "AccountThemes(nil)"
	}
	type Alias AccountThemes
	return fmt.Sprintf("AccountThemes%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountThemes) TypeID() uint32 {
	return AccountThemesTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountThemes) TypeName() string {
	return "account.themes"
}

// TypeInfo returns info about TL type.
func (t *AccountThemes) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.themes",
		ID:   AccountThemesTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Themes",
			SchemaName: "themes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *AccountThemes) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.themes#9a3d8c6d as nil")
	}
	b.PutID(AccountThemesTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *AccountThemes) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.themes#9a3d8c6d as nil")
	}
	b.PutLong(t.Hash)
	b.PutVectorHeader(len(t.Themes))
	for idx, v := range t.Themes {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.themes#9a3d8c6d: field themes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *AccountThemes) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.themes#9a3d8c6d to nil")
	}
	if err := b.ConsumeID(AccountThemesTypeID); err != nil {
		return fmt.Errorf("unable to decode account.themes#9a3d8c6d: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *AccountThemes) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.themes#9a3d8c6d to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.themes#9a3d8c6d: field hash: %w", err)
		}
		t.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.themes#9a3d8c6d: field themes: %w", err)
		}

		if headerLen > 0 {
			t.Themes = make([]Theme, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Theme
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode account.themes#9a3d8c6d: field themes: %w", err)
			}
			t.Themes = append(t.Themes, value)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (t *AccountThemes) GetHash() (value int64) {
	if t == nil {
		return
	}
	return t.Hash
}

// GetThemes returns value of Themes field.
func (t *AccountThemes) GetThemes() (value []Theme) {
	if t == nil {
		return
	}
	return t.Themes
}

// AccountThemesClassName is schema name of AccountThemesClass.
const AccountThemesClassName = "account.Themes"

// AccountThemesClass represents account.Themes generic type.
//
// Constructors:
//   - [AccountThemesNotModified]
//   - [AccountThemes]
//
// Example:
//
//	g, err := tg.DecodeAccountThemes(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.AccountThemesNotModified: // account.themesNotModified#f41eb622
//	case *tg.AccountThemes: // account.themes#9a3d8c6d
//	default: panic(v)
//	}
type AccountThemesClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AccountThemesClass

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

// DecodeAccountThemes implements binary de-serialization for AccountThemesClass.
func DecodeAccountThemes(buf *bin.Buffer) (AccountThemesClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AccountThemesNotModifiedTypeID:
		// Decoding account.themesNotModified#f41eb622.
		v := AccountThemesNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountThemesClass: %w", err)
		}
		return &v, nil
	case AccountThemesTypeID:
		// Decoding account.themes#9a3d8c6d.
		v := AccountThemes{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountThemesClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AccountThemesClass: %w", bin.NewUnexpectedID(id))
	}
}

// AccountThemes boxes the AccountThemesClass providing a helper.
type AccountThemesBox struct {
	Themes AccountThemesClass
}

// Decode implements bin.Decoder for AccountThemesBox.
func (b *AccountThemesBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AccountThemesBox to nil")
	}
	v, err := DecodeAccountThemes(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Themes = v
	return nil
}

// Encode implements bin.Encode for AccountThemesBox.
func (b *AccountThemesBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Themes == nil {
		return fmt.Errorf("unable to encode AccountThemesClass as nil")
	}
	return b.Themes.Encode(buf)
}
