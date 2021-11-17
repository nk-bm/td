// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
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
)

// LogStreamDefault represents TL type `logStreamDefault#52e296bc`.
type LogStreamDefault struct {
}

// LogStreamDefaultTypeID is TL type id of LogStreamDefault.
const LogStreamDefaultTypeID = 0x52e296bc

// construct implements constructor of LogStreamClass.
func (l LogStreamDefault) construct() LogStreamClass { return &l }

// Ensuring interfaces in compile-time for LogStreamDefault.
var (
	_ bin.Encoder     = &LogStreamDefault{}
	_ bin.Decoder     = &LogStreamDefault{}
	_ bin.BareEncoder = &LogStreamDefault{}
	_ bin.BareDecoder = &LogStreamDefault{}

	_ LogStreamClass = &LogStreamDefault{}
)

func (l *LogStreamDefault) Zero() bool {
	if l == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (l *LogStreamDefault) String() string {
	if l == nil {
		return "LogStreamDefault(nil)"
	}
	type Alias LogStreamDefault
	return fmt.Sprintf("LogStreamDefault%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LogStreamDefault) TypeID() uint32 {
	return LogStreamDefaultTypeID
}

// TypeName returns name of type in TL schema.
func (*LogStreamDefault) TypeName() string {
	return "logStreamDefault"
}

// TypeInfo returns info about TL type.
func (l *LogStreamDefault) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "logStreamDefault",
		ID:   LogStreamDefaultTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (l *LogStreamDefault) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode logStreamDefault#52e296bc as nil")
	}
	b.PutID(LogStreamDefaultTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LogStreamDefault) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode logStreamDefault#52e296bc as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *LogStreamDefault) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode logStreamDefault#52e296bc to nil")
	}
	if err := b.ConsumeID(LogStreamDefaultTypeID); err != nil {
		return fmt.Errorf("unable to decode logStreamDefault#52e296bc: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LogStreamDefault) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode logStreamDefault#52e296bc to nil")
	}
	return nil
}

// LogStreamFile represents TL type `logStreamFile#5b528de5`.
type LogStreamFile struct {
	// Path to the file to where the internal TDLib log will be written
	Path string
	// The maximum size of the file to where the internal TDLib log is written before the
	// file will be auto-rotated
	MaxFileSize int64
	// Pass true to additionally redirect stderr to the log file. Ignored on Windows
	RedirectStderr bool
}

// LogStreamFileTypeID is TL type id of LogStreamFile.
const LogStreamFileTypeID = 0x5b528de5

// construct implements constructor of LogStreamClass.
func (l LogStreamFile) construct() LogStreamClass { return &l }

// Ensuring interfaces in compile-time for LogStreamFile.
var (
	_ bin.Encoder     = &LogStreamFile{}
	_ bin.Decoder     = &LogStreamFile{}
	_ bin.BareEncoder = &LogStreamFile{}
	_ bin.BareDecoder = &LogStreamFile{}

	_ LogStreamClass = &LogStreamFile{}
)

func (l *LogStreamFile) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Path == "") {
		return false
	}
	if !(l.MaxFileSize == 0) {
		return false
	}
	if !(l.RedirectStderr == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LogStreamFile) String() string {
	if l == nil {
		return "LogStreamFile(nil)"
	}
	type Alias LogStreamFile
	return fmt.Sprintf("LogStreamFile%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LogStreamFile) TypeID() uint32 {
	return LogStreamFileTypeID
}

// TypeName returns name of type in TL schema.
func (*LogStreamFile) TypeName() string {
	return "logStreamFile"
}

// TypeInfo returns info about TL type.
func (l *LogStreamFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "logStreamFile",
		ID:   LogStreamFileTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Path",
			SchemaName: "path",
		},
		{
			Name:       "MaxFileSize",
			SchemaName: "max_file_size",
		},
		{
			Name:       "RedirectStderr",
			SchemaName: "redirect_stderr",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LogStreamFile) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode logStreamFile#5b528de5 as nil")
	}
	b.PutID(LogStreamFileTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LogStreamFile) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode logStreamFile#5b528de5 as nil")
	}
	b.PutString(l.Path)
	b.PutLong(l.MaxFileSize)
	b.PutBool(l.RedirectStderr)
	return nil
}

// Decode implements bin.Decoder.
func (l *LogStreamFile) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode logStreamFile#5b528de5 to nil")
	}
	if err := b.ConsumeID(LogStreamFileTypeID); err != nil {
		return fmt.Errorf("unable to decode logStreamFile#5b528de5: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LogStreamFile) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode logStreamFile#5b528de5 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode logStreamFile#5b528de5: field path: %w", err)
		}
		l.Path = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode logStreamFile#5b528de5: field max_file_size: %w", err)
		}
		l.MaxFileSize = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode logStreamFile#5b528de5: field redirect_stderr: %w", err)
		}
		l.RedirectStderr = value
	}
	return nil
}

// GetPath returns value of Path field.
func (l *LogStreamFile) GetPath() (value string) {
	return l.Path
}

// GetMaxFileSize returns value of MaxFileSize field.
func (l *LogStreamFile) GetMaxFileSize() (value int64) {
	return l.MaxFileSize
}

// GetRedirectStderr returns value of RedirectStderr field.
func (l *LogStreamFile) GetRedirectStderr() (value bool) {
	return l.RedirectStderr
}

// LogStreamEmpty represents TL type `logStreamEmpty#e233f1cc`.
type LogStreamEmpty struct {
}

// LogStreamEmptyTypeID is TL type id of LogStreamEmpty.
const LogStreamEmptyTypeID = 0xe233f1cc

// construct implements constructor of LogStreamClass.
func (l LogStreamEmpty) construct() LogStreamClass { return &l }

// Ensuring interfaces in compile-time for LogStreamEmpty.
var (
	_ bin.Encoder     = &LogStreamEmpty{}
	_ bin.Decoder     = &LogStreamEmpty{}
	_ bin.BareEncoder = &LogStreamEmpty{}
	_ bin.BareDecoder = &LogStreamEmpty{}

	_ LogStreamClass = &LogStreamEmpty{}
)

func (l *LogStreamEmpty) Zero() bool {
	if l == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (l *LogStreamEmpty) String() string {
	if l == nil {
		return "LogStreamEmpty(nil)"
	}
	type Alias LogStreamEmpty
	return fmt.Sprintf("LogStreamEmpty%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LogStreamEmpty) TypeID() uint32 {
	return LogStreamEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*LogStreamEmpty) TypeName() string {
	return "logStreamEmpty"
}

// TypeInfo returns info about TL type.
func (l *LogStreamEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "logStreamEmpty",
		ID:   LogStreamEmptyTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (l *LogStreamEmpty) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode logStreamEmpty#e233f1cc as nil")
	}
	b.PutID(LogStreamEmptyTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LogStreamEmpty) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode logStreamEmpty#e233f1cc as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *LogStreamEmpty) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode logStreamEmpty#e233f1cc to nil")
	}
	if err := b.ConsumeID(LogStreamEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode logStreamEmpty#e233f1cc: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LogStreamEmpty) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode logStreamEmpty#e233f1cc to nil")
	}
	return nil
}

// LogStreamClass represents LogStream generic type.
//
// Example:
//  g, err := tdapi.DecodeLogStream(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.LogStreamDefault: // logStreamDefault#52e296bc
//  case *tdapi.LogStreamFile: // logStreamFile#5b528de5
//  case *tdapi.LogStreamEmpty: // logStreamEmpty#e233f1cc
//  default: panic(v)
//  }
type LogStreamClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() LogStreamClass

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

// DecodeLogStream implements binary de-serialization for LogStreamClass.
func DecodeLogStream(buf *bin.Buffer) (LogStreamClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case LogStreamDefaultTypeID:
		// Decoding logStreamDefault#52e296bc.
		v := LogStreamDefault{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LogStreamClass: %w", err)
		}
		return &v, nil
	case LogStreamFileTypeID:
		// Decoding logStreamFile#5b528de5.
		v := LogStreamFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LogStreamClass: %w", err)
		}
		return &v, nil
	case LogStreamEmptyTypeID:
		// Decoding logStreamEmpty#e233f1cc.
		v := LogStreamEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LogStreamClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode LogStreamClass: %w", bin.NewUnexpectedID(id))
	}
}

// LogStream boxes the LogStreamClass providing a helper.
type LogStreamBox struct {
	LogStream LogStreamClass
}

// Decode implements bin.Decoder for LogStreamBox.
func (b *LogStreamBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode LogStreamBox to nil")
	}
	v, err := DecodeLogStream(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.LogStream = v
	return nil
}

// Encode implements bin.Encode for LogStreamBox.
func (b *LogStreamBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.LogStream == nil {
		return fmt.Errorf("unable to encode LogStreamClass as nil")
	}
	return b.LogStream.Encode(buf)
}
