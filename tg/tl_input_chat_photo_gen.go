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

// InputChatPhotoEmpty represents TL type `inputChatPhotoEmpty#1ca48f57`.
type InputChatPhotoEmpty struct {
}

// InputChatPhotoEmptyTypeID is TL type id of InputChatPhotoEmpty.
const InputChatPhotoEmptyTypeID = 0x1ca48f57

// construct implements constructor of InputChatPhotoClass.
func (i InputChatPhotoEmpty) construct() InputChatPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputChatPhotoEmpty.
var (
	_ bin.Encoder     = &InputChatPhotoEmpty{}
	_ bin.Decoder     = &InputChatPhotoEmpty{}
	_ bin.BareEncoder = &InputChatPhotoEmpty{}
	_ bin.BareDecoder = &InputChatPhotoEmpty{}

	_ InputChatPhotoClass = &InputChatPhotoEmpty{}
)

func (i *InputChatPhotoEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChatPhotoEmpty) String() string {
	if i == nil {
		return "InputChatPhotoEmpty(nil)"
	}
	type Alias InputChatPhotoEmpty
	return fmt.Sprintf("InputChatPhotoEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChatPhotoEmpty) TypeID() uint32 {
	return InputChatPhotoEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChatPhotoEmpty) TypeName() string {
	return "inputChatPhotoEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputChatPhotoEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChatPhotoEmpty",
		ID:   InputChatPhotoEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChatPhotoEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoEmpty#1ca48f57 as nil")
	}
	b.PutID(InputChatPhotoEmptyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputChatPhotoEmpty) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoEmpty#1ca48f57 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChatPhotoEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoEmpty#1ca48f57 to nil")
	}
	if err := b.ConsumeID(InputChatPhotoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChatPhotoEmpty#1ca48f57: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputChatPhotoEmpty) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhotoEmpty#1ca48f57 to nil")
	}
	return nil
}

// InputChatUploadedPhoto represents TL type `inputChatUploadedPhoto#bdcdaec0`.
type InputChatUploadedPhoto struct {
	// Flags field of InputChatUploadedPhoto.
	Flags bin.Fields
	// File field of InputChatUploadedPhoto.
	//
	// Use SetFile and GetFile helpers.
	File InputFileClass
	// Video field of InputChatUploadedPhoto.
	//
	// Use SetVideo and GetVideo helpers.
	Video InputFileClass
	// VideoStartTs field of InputChatUploadedPhoto.
	//
	// Use SetVideoStartTs and GetVideoStartTs helpers.
	VideoStartTs float64
	// VideoEmojiMarkup field of InputChatUploadedPhoto.
	//
	// Use SetVideoEmojiMarkup and GetVideoEmojiMarkup helpers.
	VideoEmojiMarkup VideoSizeClass
}

// InputChatUploadedPhotoTypeID is TL type id of InputChatUploadedPhoto.
const InputChatUploadedPhotoTypeID = 0xbdcdaec0

// construct implements constructor of InputChatPhotoClass.
func (i InputChatUploadedPhoto) construct() InputChatPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputChatUploadedPhoto.
var (
	_ bin.Encoder     = &InputChatUploadedPhoto{}
	_ bin.Decoder     = &InputChatUploadedPhoto{}
	_ bin.BareEncoder = &InputChatUploadedPhoto{}
	_ bin.BareDecoder = &InputChatUploadedPhoto{}

	_ InputChatPhotoClass = &InputChatUploadedPhoto{}
)

func (i *InputChatUploadedPhoto) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.File == nil) {
		return false
	}
	if !(i.Video == nil) {
		return false
	}
	if !(i.VideoStartTs == 0) {
		return false
	}
	if !(i.VideoEmojiMarkup == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChatUploadedPhoto) String() string {
	if i == nil {
		return "InputChatUploadedPhoto(nil)"
	}
	type Alias InputChatUploadedPhoto
	return fmt.Sprintf("InputChatUploadedPhoto%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChatUploadedPhoto) TypeID() uint32 {
	return InputChatUploadedPhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChatUploadedPhoto) TypeName() string {
	return "inputChatUploadedPhoto"
}

// TypeInfo returns info about TL type.
func (i *InputChatUploadedPhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChatUploadedPhoto",
		ID:   InputChatUploadedPhotoTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "File",
			SchemaName: "file",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "Video",
			SchemaName: "video",
			Null:       !i.Flags.Has(1),
		},
		{
			Name:       "VideoStartTs",
			SchemaName: "video_start_ts",
			Null:       !i.Flags.Has(2),
		},
		{
			Name:       "VideoEmojiMarkup",
			SchemaName: "video_emoji_markup",
			Null:       !i.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (i *InputChatUploadedPhoto) SetFlags() {
	if !(i.File == nil) {
		i.Flags.Set(0)
	}
	if !(i.Video == nil) {
		i.Flags.Set(1)
	}
	if !(i.VideoStartTs == 0) {
		i.Flags.Set(2)
	}
	if !(i.VideoEmojiMarkup == nil) {
		i.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (i *InputChatUploadedPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatUploadedPhoto#bdcdaec0 as nil")
	}
	b.PutID(InputChatUploadedPhotoTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputChatUploadedPhoto) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatUploadedPhoto#bdcdaec0 as nil")
	}
	i.SetFlags()
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputChatUploadedPhoto#bdcdaec0: field flags: %w", err)
	}
	if i.Flags.Has(0) {
		if i.File == nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#bdcdaec0: field file is nil")
		}
		if err := i.File.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#bdcdaec0: field file: %w", err)
		}
	}
	if i.Flags.Has(1) {
		if i.Video == nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#bdcdaec0: field video is nil")
		}
		if err := i.Video.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#bdcdaec0: field video: %w", err)
		}
	}
	if i.Flags.Has(2) {
		b.PutDouble(i.VideoStartTs)
	}
	if i.Flags.Has(3) {
		if i.VideoEmojiMarkup == nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#bdcdaec0: field video_emoji_markup is nil")
		}
		if err := i.VideoEmojiMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#bdcdaec0: field video_emoji_markup: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChatUploadedPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatUploadedPhoto#bdcdaec0 to nil")
	}
	if err := b.ConsumeID(InputChatUploadedPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChatUploadedPhoto#bdcdaec0: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputChatUploadedPhoto) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatUploadedPhoto#bdcdaec0 to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#bdcdaec0: field flags: %w", err)
		}
	}
	if i.Flags.Has(0) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#bdcdaec0: field file: %w", err)
		}
		i.File = value
	}
	if i.Flags.Has(1) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#bdcdaec0: field video: %w", err)
		}
		i.Video = value
	}
	if i.Flags.Has(2) {
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#bdcdaec0: field video_start_ts: %w", err)
		}
		i.VideoStartTs = value
	}
	if i.Flags.Has(3) {
		value, err := DecodeVideoSize(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#bdcdaec0: field video_emoji_markup: %w", err)
		}
		i.VideoEmojiMarkup = value
	}
	return nil
}

// SetFile sets value of File conditional field.
func (i *InputChatUploadedPhoto) SetFile(value InputFileClass) {
	i.Flags.Set(0)
	i.File = value
}

// GetFile returns value of File conditional field and
// boolean which is true if field was set.
func (i *InputChatUploadedPhoto) GetFile() (value InputFileClass, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.File, true
}

// SetVideo sets value of Video conditional field.
func (i *InputChatUploadedPhoto) SetVideo(value InputFileClass) {
	i.Flags.Set(1)
	i.Video = value
}

// GetVideo returns value of Video conditional field and
// boolean which is true if field was set.
func (i *InputChatUploadedPhoto) GetVideo() (value InputFileClass, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Video, true
}

// SetVideoStartTs sets value of VideoStartTs conditional field.
func (i *InputChatUploadedPhoto) SetVideoStartTs(value float64) {
	i.Flags.Set(2)
	i.VideoStartTs = value
}

// GetVideoStartTs returns value of VideoStartTs conditional field and
// boolean which is true if field was set.
func (i *InputChatUploadedPhoto) GetVideoStartTs() (value float64, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.VideoStartTs, true
}

// SetVideoEmojiMarkup sets value of VideoEmojiMarkup conditional field.
func (i *InputChatUploadedPhoto) SetVideoEmojiMarkup(value VideoSizeClass) {
	i.Flags.Set(3)
	i.VideoEmojiMarkup = value
}

// GetVideoEmojiMarkup returns value of VideoEmojiMarkup conditional field and
// boolean which is true if field was set.
func (i *InputChatUploadedPhoto) GetVideoEmojiMarkup() (value VideoSizeClass, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(3) {
		return value, false
	}
	return i.VideoEmojiMarkup, true
}

// InputChatPhoto represents TL type `inputChatPhoto#8953ad37`.
type InputChatPhoto struct {
	// ID field of InputChatPhoto.
	ID InputPhotoClass
}

// InputChatPhotoTypeID is TL type id of InputChatPhoto.
const InputChatPhotoTypeID = 0x8953ad37

// construct implements constructor of InputChatPhotoClass.
func (i InputChatPhoto) construct() InputChatPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputChatPhoto.
var (
	_ bin.Encoder     = &InputChatPhoto{}
	_ bin.Decoder     = &InputChatPhoto{}
	_ bin.BareEncoder = &InputChatPhoto{}
	_ bin.BareDecoder = &InputChatPhoto{}

	_ InputChatPhotoClass = &InputChatPhoto{}
)

func (i *InputChatPhoto) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChatPhoto) String() string {
	if i == nil {
		return "InputChatPhoto(nil)"
	}
	type Alias InputChatPhoto
	return fmt.Sprintf("InputChatPhoto%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChatPhoto) TypeID() uint32 {
	return InputChatPhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChatPhoto) TypeName() string {
	return "inputChatPhoto"
}

// TypeInfo returns info about TL type.
func (i *InputChatPhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChatPhoto",
		ID:   InputChatPhotoTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChatPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhoto#8953ad37 as nil")
	}
	b.PutID(InputChatPhotoTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputChatPhoto) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhoto#8953ad37 as nil")
	}
	if i.ID == nil {
		return fmt.Errorf("unable to encode inputChatPhoto#8953ad37: field id is nil")
	}
	if err := i.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputChatPhoto#8953ad37: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChatPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhoto#8953ad37 to nil")
	}
	if err := b.ConsumeID(InputChatPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChatPhoto#8953ad37: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputChatPhoto) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhoto#8953ad37 to nil")
	}
	{
		value, err := DecodeInputPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChatPhoto#8953ad37: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputChatPhoto) GetID() (value InputPhotoClass) {
	if i == nil {
		return
	}
	return i.ID
}

// InputChatPhotoClassName is schema name of InputChatPhotoClass.
const InputChatPhotoClassName = "InputChatPhoto"

// InputChatPhotoClass represents InputChatPhoto generic type.
//
// Constructors:
//   - [InputChatPhotoEmpty]
//   - [InputChatUploadedPhoto]
//   - [InputChatPhoto]
//
// Example:
//
//	g, err := tg.DecodeInputChatPhoto(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputChatPhotoEmpty: // inputChatPhotoEmpty#1ca48f57
//	case *tg.InputChatUploadedPhoto: // inputChatUploadedPhoto#bdcdaec0
//	case *tg.InputChatPhoto: // inputChatPhoto#8953ad37
//	default: panic(v)
//	}
type InputChatPhotoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputChatPhotoClass

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

// DecodeInputChatPhoto implements binary de-serialization for InputChatPhotoClass.
func DecodeInputChatPhoto(buf *bin.Buffer) (InputChatPhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputChatPhotoEmptyTypeID:
		// Decoding inputChatPhotoEmpty#1ca48f57.
		v := InputChatPhotoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", err)
		}
		return &v, nil
	case InputChatUploadedPhotoTypeID:
		// Decoding inputChatUploadedPhoto#bdcdaec0.
		v := InputChatUploadedPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", err)
		}
		return &v, nil
	case InputChatPhotoTypeID:
		// Decoding inputChatPhoto#8953ad37.
		v := InputChatPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputChatPhotoClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputChatPhoto boxes the InputChatPhotoClass providing a helper.
type InputChatPhotoBox struct {
	InputChatPhoto InputChatPhotoClass
}

// Decode implements bin.Decoder for InputChatPhotoBox.
func (b *InputChatPhotoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputChatPhotoBox to nil")
	}
	v, err := DecodeInputChatPhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputChatPhoto = v
	return nil
}

// Encode implements bin.Encode for InputChatPhotoBox.
func (b *InputChatPhotoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputChatPhoto == nil {
		return fmt.Errorf("unable to encode InputChatPhotoClass as nil")
	}
	return b.InputChatPhoto.Encode(buf)
}
