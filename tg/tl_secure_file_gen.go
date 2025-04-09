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

// SecureFileEmpty represents TL type `secureFileEmpty#64199744`.
type SecureFileEmpty struct {
}

// SecureFileEmptyTypeID is TL type id of SecureFileEmpty.
const SecureFileEmptyTypeID = 0x64199744

// construct implements constructor of SecureFileClass.
func (s SecureFileEmpty) construct() SecureFileClass { return &s }

// Ensuring interfaces in compile-time for SecureFileEmpty.
var (
	_ bin.Encoder     = &SecureFileEmpty{}
	_ bin.Decoder     = &SecureFileEmpty{}
	_ bin.BareEncoder = &SecureFileEmpty{}
	_ bin.BareDecoder = &SecureFileEmpty{}

	_ SecureFileClass = &SecureFileEmpty{}
)

func (s *SecureFileEmpty) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureFileEmpty) String() string {
	if s == nil {
		return "SecureFileEmpty(nil)"
	}
	type Alias SecureFileEmpty
	return fmt.Sprintf("SecureFileEmpty%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SecureFileEmpty) TypeID() uint32 {
	return SecureFileEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*SecureFileEmpty) TypeName() string {
	return "secureFileEmpty"
}

// TypeInfo returns info about TL type.
func (s *SecureFileEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "secureFileEmpty",
		ID:   SecureFileEmptyTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *SecureFileEmpty) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureFileEmpty#64199744 as nil")
	}
	b.PutID(SecureFileEmptyTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SecureFileEmpty) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureFileEmpty#64199744 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureFileEmpty) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureFileEmpty#64199744 to nil")
	}
	if err := b.ConsumeID(SecureFileEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode secureFileEmpty#64199744: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SecureFileEmpty) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureFileEmpty#64199744 to nil")
	}
	return nil
}

// SecureFile represents TL type `secureFile#7d09c27e`.
type SecureFile struct {
	// ID field of SecureFile.
	ID int64
	// AccessHash field of SecureFile.
	AccessHash int64
	// Size field of SecureFile.
	Size int64
	// DCID field of SecureFile.
	DCID int
	// Date field of SecureFile.
	Date int
	// FileHash field of SecureFile.
	FileHash []byte
	// Secret field of SecureFile.
	Secret []byte
}

// SecureFileTypeID is TL type id of SecureFile.
const SecureFileTypeID = 0x7d09c27e

// construct implements constructor of SecureFileClass.
func (s SecureFile) construct() SecureFileClass { return &s }

// Ensuring interfaces in compile-time for SecureFile.
var (
	_ bin.Encoder     = &SecureFile{}
	_ bin.Decoder     = &SecureFile{}
	_ bin.BareEncoder = &SecureFile{}
	_ bin.BareDecoder = &SecureFile{}

	_ SecureFileClass = &SecureFile{}
)

func (s *SecureFile) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.AccessHash == 0) {
		return false
	}
	if !(s.Size == 0) {
		return false
	}
	if !(s.DCID == 0) {
		return false
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.FileHash == nil) {
		return false
	}
	if !(s.Secret == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureFile) String() string {
	if s == nil {
		return "SecureFile(nil)"
	}
	type Alias SecureFile
	return fmt.Sprintf("SecureFile%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SecureFile) TypeID() uint32 {
	return SecureFileTypeID
}

// TypeName returns name of type in TL schema.
func (*SecureFile) TypeName() string {
	return "secureFile"
}

// TypeInfo returns info about TL type.
func (s *SecureFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "secureFile",
		ID:   SecureFileTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "FileHash",
			SchemaName: "file_hash",
		},
		{
			Name:       "Secret",
			SchemaName: "secret",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SecureFile) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureFile#7d09c27e as nil")
	}
	b.PutID(SecureFileTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SecureFile) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureFile#7d09c27e as nil")
	}
	b.PutLong(s.ID)
	b.PutLong(s.AccessHash)
	b.PutLong(s.Size)
	b.PutInt(s.DCID)
	b.PutInt(s.Date)
	b.PutBytes(s.FileHash)
	b.PutBytes(s.Secret)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureFile) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureFile#7d09c27e to nil")
	}
	if err := b.ConsumeID(SecureFileTypeID); err != nil {
		return fmt.Errorf("unable to decode secureFile#7d09c27e: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SecureFile) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureFile#7d09c27e to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#7d09c27e: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#7d09c27e: field access_hash: %w", err)
		}
		s.AccessHash = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#7d09c27e: field size: %w", err)
		}
		s.Size = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#7d09c27e: field dc_id: %w", err)
		}
		s.DCID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#7d09c27e: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#7d09c27e: field file_hash: %w", err)
		}
		s.FileHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureFile#7d09c27e: field secret: %w", err)
		}
		s.Secret = value
	}
	return nil
}

// GetID returns value of ID field.
func (s *SecureFile) GetID() (value int64) {
	if s == nil {
		return
	}
	return s.ID
}

// GetAccessHash returns value of AccessHash field.
func (s *SecureFile) GetAccessHash() (value int64) {
	if s == nil {
		return
	}
	return s.AccessHash
}

// GetSize returns value of Size field.
func (s *SecureFile) GetSize() (value int64) {
	if s == nil {
		return
	}
	return s.Size
}

// GetDCID returns value of DCID field.
func (s *SecureFile) GetDCID() (value int) {
	if s == nil {
		return
	}
	return s.DCID
}

// GetDate returns value of Date field.
func (s *SecureFile) GetDate() (value int) {
	if s == nil {
		return
	}
	return s.Date
}

// GetFileHash returns value of FileHash field.
func (s *SecureFile) GetFileHash() (value []byte) {
	if s == nil {
		return
	}
	return s.FileHash
}

// GetSecret returns value of Secret field.
func (s *SecureFile) GetSecret() (value []byte) {
	if s == nil {
		return
	}
	return s.Secret
}

// SecureFileClassName is schema name of SecureFileClass.
const SecureFileClassName = "SecureFile"

// SecureFileClass represents SecureFile generic type.
//
// Constructors:
//   - [SecureFileEmpty]
//   - [SecureFile]
//
// Example:
//
//	g, err := tg.DecodeSecureFile(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.SecureFileEmpty: // secureFileEmpty#64199744
//	case *tg.SecureFile: // secureFile#7d09c27e
//	default: panic(v)
//	}
type SecureFileClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() SecureFileClass

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

// DecodeSecureFile implements binary de-serialization for SecureFileClass.
func DecodeSecureFile(buf *bin.Buffer) (SecureFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SecureFileEmptyTypeID:
		// Decoding secureFileEmpty#64199744.
		v := SecureFileEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureFileClass: %w", err)
		}
		return &v, nil
	case SecureFileTypeID:
		// Decoding secureFile#7d09c27e.
		v := SecureFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SecureFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// SecureFile boxes the SecureFileClass providing a helper.
type SecureFileBox struct {
	SecureFile SecureFileClass
}

// Decode implements bin.Decoder for SecureFileBox.
func (b *SecureFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SecureFileBox to nil")
	}
	v, err := DecodeSecureFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SecureFile = v
	return nil
}

// Encode implements bin.Encode for SecureFileBox.
func (b *SecureFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SecureFile == nil {
		return fmt.Errorf("unable to encode SecureFileClass as nil")
	}
	return b.SecureFile.Encode(buf)
}
