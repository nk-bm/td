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

// ReportResultChooseOption represents TL type `reportResultChooseOption#f0e4e0b6`.
// The user must choose one of the following options, and then messages.report¹ must be
// re-invoked, passing the option's option identifier to messages.report².option.
//
// Links:
//  1. https://core.telegram.org/method/messages.report
//  2. https://core.telegram.org/method/messages.report
//
// See https://core.telegram.org/constructor/reportResultChooseOption for reference.
type ReportResultChooseOption struct {
	// Title of the option popup
	Title string
	// Available options, rendered as menu entries.
	Options []MessageReportOption
}

// ReportResultChooseOptionTypeID is TL type id of ReportResultChooseOption.
const ReportResultChooseOptionTypeID = 0xf0e4e0b6

// construct implements constructor of ReportResultClass.
func (r ReportResultChooseOption) construct() ReportResultClass { return &r }

// Ensuring interfaces in compile-time for ReportResultChooseOption.
var (
	_ bin.Encoder     = &ReportResultChooseOption{}
	_ bin.Decoder     = &ReportResultChooseOption{}
	_ bin.BareEncoder = &ReportResultChooseOption{}
	_ bin.BareDecoder = &ReportResultChooseOption{}

	_ ReportResultClass = &ReportResultChooseOption{}
)

func (r *ReportResultChooseOption) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Title == "") {
		return false
	}
	if !(r.Options == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportResultChooseOption) String() string {
	if r == nil {
		return "ReportResultChooseOption(nil)"
	}
	type Alias ReportResultChooseOption
	return fmt.Sprintf("ReportResultChooseOption%+v", Alias(*r))
}

// FillFrom fills ReportResultChooseOption from given interface.
func (r *ReportResultChooseOption) FillFrom(from interface {
	GetTitle() (value string)
	GetOptions() (value []MessageReportOption)
}) {
	r.Title = from.GetTitle()
	r.Options = from.GetOptions()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportResultChooseOption) TypeID() uint32 {
	return ReportResultChooseOptionTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportResultChooseOption) TypeName() string {
	return "reportResultChooseOption"
}

// TypeInfo returns info about TL type.
func (r *ReportResultChooseOption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportResultChooseOption",
		ID:   ReportResultChooseOptionTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Options",
			SchemaName: "options",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportResultChooseOption) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportResultChooseOption#f0e4e0b6 as nil")
	}
	b.PutID(ReportResultChooseOptionTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportResultChooseOption) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportResultChooseOption#f0e4e0b6 as nil")
	}
	b.PutString(r.Title)
	b.PutVectorHeader(len(r.Options))
	for idx, v := range r.Options {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode reportResultChooseOption#f0e4e0b6: field options element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportResultChooseOption) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportResultChooseOption#f0e4e0b6 to nil")
	}
	if err := b.ConsumeID(ReportResultChooseOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode reportResultChooseOption#f0e4e0b6: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportResultChooseOption) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportResultChooseOption#f0e4e0b6 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reportResultChooseOption#f0e4e0b6: field title: %w", err)
		}
		r.Title = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode reportResultChooseOption#f0e4e0b6: field options: %w", err)
		}

		if headerLen > 0 {
			r.Options = make([]MessageReportOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessageReportOption
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode reportResultChooseOption#f0e4e0b6: field options: %w", err)
			}
			r.Options = append(r.Options, value)
		}
	}
	return nil
}

// GetTitle returns value of Title field.
func (r *ReportResultChooseOption) GetTitle() (value string) {
	if r == nil {
		return
	}
	return r.Title
}

// GetOptions returns value of Options field.
func (r *ReportResultChooseOption) GetOptions() (value []MessageReportOption) {
	if r == nil {
		return
	}
	return r.Options
}

// ReportResultAddComment represents TL type `reportResultAddComment#6f09ac31`.
// The user should enter an additional comment for the moderators, and then messages
// report¹ must be re-invoked, passing the comment to messages.report².message.
//
// Links:
//  1. https://core.telegram.org/method/messages.report
//  2. https://core.telegram.org/method/messages.report
//
// See https://core.telegram.org/constructor/reportResultAddComment for reference.
type ReportResultAddComment struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this step can be skipped by the user, passing an empty message to messages
	// report¹, or if a non-empty message is mandatory.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.report
	Optional bool
	// The messages.report¹ method must be re-invoked, passing this option to option
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.report
	Option []byte
}

// ReportResultAddCommentTypeID is TL type id of ReportResultAddComment.
const ReportResultAddCommentTypeID = 0x6f09ac31

// construct implements constructor of ReportResultClass.
func (r ReportResultAddComment) construct() ReportResultClass { return &r }

// Ensuring interfaces in compile-time for ReportResultAddComment.
var (
	_ bin.Encoder     = &ReportResultAddComment{}
	_ bin.Decoder     = &ReportResultAddComment{}
	_ bin.BareEncoder = &ReportResultAddComment{}
	_ bin.BareDecoder = &ReportResultAddComment{}

	_ ReportResultClass = &ReportResultAddComment{}
)

func (r *ReportResultAddComment) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Optional == false) {
		return false
	}
	if !(r.Option == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportResultAddComment) String() string {
	if r == nil {
		return "ReportResultAddComment(nil)"
	}
	type Alias ReportResultAddComment
	return fmt.Sprintf("ReportResultAddComment%+v", Alias(*r))
}

// FillFrom fills ReportResultAddComment from given interface.
func (r *ReportResultAddComment) FillFrom(from interface {
	GetOptional() (value bool)
	GetOption() (value []byte)
}) {
	r.Optional = from.GetOptional()
	r.Option = from.GetOption()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportResultAddComment) TypeID() uint32 {
	return ReportResultAddCommentTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportResultAddComment) TypeName() string {
	return "reportResultAddComment"
}

// TypeInfo returns info about TL type.
func (r *ReportResultAddComment) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportResultAddComment",
		ID:   ReportResultAddCommentTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Optional",
			SchemaName: "optional",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "Option",
			SchemaName: "option",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *ReportResultAddComment) SetFlags() {
	if !(r.Optional == false) {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *ReportResultAddComment) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportResultAddComment#6f09ac31 as nil")
	}
	b.PutID(ReportResultAddCommentTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportResultAddComment) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportResultAddComment#6f09ac31 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reportResultAddComment#6f09ac31: field flags: %w", err)
	}
	b.PutBytes(r.Option)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportResultAddComment) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportResultAddComment#6f09ac31 to nil")
	}
	if err := b.ConsumeID(ReportResultAddCommentTypeID); err != nil {
		return fmt.Errorf("unable to decode reportResultAddComment#6f09ac31: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportResultAddComment) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportResultAddComment#6f09ac31 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode reportResultAddComment#6f09ac31: field flags: %w", err)
		}
	}
	r.Optional = r.Flags.Has(0)
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode reportResultAddComment#6f09ac31: field option: %w", err)
		}
		r.Option = value
	}
	return nil
}

// SetOptional sets value of Optional conditional field.
func (r *ReportResultAddComment) SetOptional(value bool) {
	if value {
		r.Flags.Set(0)
		r.Optional = true
	} else {
		r.Flags.Unset(0)
		r.Optional = false
	}
}

// GetOptional returns value of Optional conditional field.
func (r *ReportResultAddComment) GetOptional() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(0)
}

// GetOption returns value of Option field.
func (r *ReportResultAddComment) GetOption() (value []byte) {
	if r == nil {
		return
	}
	return r.Option
}

// ReportResultReported represents TL type `reportResultReported#8db33c4b`.
// The report was sent successfully, no further actions are required.
//
// See https://core.telegram.org/constructor/reportResultReported for reference.
type ReportResultReported struct {
}

// ReportResultReportedTypeID is TL type id of ReportResultReported.
const ReportResultReportedTypeID = 0x8db33c4b

// construct implements constructor of ReportResultClass.
func (r ReportResultReported) construct() ReportResultClass { return &r }

// Ensuring interfaces in compile-time for ReportResultReported.
var (
	_ bin.Encoder     = &ReportResultReported{}
	_ bin.Decoder     = &ReportResultReported{}
	_ bin.BareEncoder = &ReportResultReported{}
	_ bin.BareDecoder = &ReportResultReported{}

	_ ReportResultClass = &ReportResultReported{}
)

func (r *ReportResultReported) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportResultReported) String() string {
	if r == nil {
		return "ReportResultReported(nil)"
	}
	type Alias ReportResultReported
	return fmt.Sprintf("ReportResultReported%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportResultReported) TypeID() uint32 {
	return ReportResultReportedTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportResultReported) TypeName() string {
	return "reportResultReported"
}

// TypeInfo returns info about TL type.
func (r *ReportResultReported) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportResultReported",
		ID:   ReportResultReportedTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportResultReported) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportResultReported#8db33c4b as nil")
	}
	b.PutID(ReportResultReportedTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportResultReported) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportResultReported#8db33c4b as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportResultReported) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportResultReported#8db33c4b to nil")
	}
	if err := b.ConsumeID(ReportResultReportedTypeID); err != nil {
		return fmt.Errorf("unable to decode reportResultReported#8db33c4b: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportResultReported) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportResultReported#8db33c4b to nil")
	}
	return nil
}

// ReportResultClassName is schema name of ReportResultClass.
const ReportResultClassName = "ReportResult"

// ReportResultClass represents ReportResult generic type.
//
// See https://core.telegram.org/type/ReportResult for reference.
//
// Constructors:
//   - [ReportResultChooseOption]
//   - [ReportResultAddComment]
//   - [ReportResultReported]
//
// Example:
//
//	g, err := tg.DecodeReportResult(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.ReportResultChooseOption: // reportResultChooseOption#f0e4e0b6
//	case *tg.ReportResultAddComment: // reportResultAddComment#6f09ac31
//	case *tg.ReportResultReported: // reportResultReported#8db33c4b
//	default: panic(v)
//	}
type ReportResultClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ReportResultClass

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

// DecodeReportResult implements binary de-serialization for ReportResultClass.
func DecodeReportResult(buf *bin.Buffer) (ReportResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ReportResultChooseOptionTypeID:
		// Decoding reportResultChooseOption#f0e4e0b6.
		v := ReportResultChooseOption{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportResultClass: %w", err)
		}
		return &v, nil
	case ReportResultAddCommentTypeID:
		// Decoding reportResultAddComment#6f09ac31.
		v := ReportResultAddComment{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportResultClass: %w", err)
		}
		return &v, nil
	case ReportResultReportedTypeID:
		// Decoding reportResultReported#8db33c4b.
		v := ReportResultReported{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReportResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// ReportResult boxes the ReportResultClass providing a helper.
type ReportResultBox struct {
	ReportResult ReportResultClass
}

// Decode implements bin.Decoder for ReportResultBox.
func (b *ReportResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReportResultBox to nil")
	}
	v, err := DecodeReportResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReportResult = v
	return nil
}

// Encode implements bin.Encode for ReportResultBox.
func (b *ReportResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ReportResult == nil {
		return fmt.Errorf("unable to encode ReportResultClass as nil")
	}
	return b.ReportResult.Encode(buf)
}
