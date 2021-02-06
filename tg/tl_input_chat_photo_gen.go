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

// InputChatPhotoEmpty represents TL type `inputChatPhotoEmpty#1ca48f57`.
// Empty constructor, remove group photo.
//
// See https://core.telegram.org/constructor/inputChatPhotoEmpty for reference.
type InputChatPhotoEmpty struct {
}

// InputChatPhotoEmptyTypeID is TL type id of InputChatPhotoEmpty.
const InputChatPhotoEmptyTypeID = 0x1ca48f57

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
	var sb strings.Builder
	sb.WriteString("InputChatPhotoEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputChatPhotoEmpty) TypeID() uint32 {
	return InputChatPhotoEmptyTypeID
}

// Encode implements bin.Encoder.
func (i *InputChatPhotoEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhotoEmpty#1ca48f57 as nil")
	}
	b.PutID(InputChatPhotoEmptyTypeID)
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
	return nil
}

// construct implements constructor of InputChatPhotoClass.
func (i InputChatPhotoEmpty) construct() InputChatPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputChatPhotoEmpty.
var (
	_ bin.Encoder = &InputChatPhotoEmpty{}
	_ bin.Decoder = &InputChatPhotoEmpty{}

	_ InputChatPhotoClass = &InputChatPhotoEmpty{}
)

// InputChatUploadedPhoto represents TL type `inputChatUploadedPhoto#c642724e`.
// New photo to be set as group profile photo.
//
// See https://core.telegram.org/constructor/inputChatUploadedPhoto for reference.
type InputChatUploadedPhoto struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// File saved in parts using the method upload.saveFilePart¹
	//
	// Links:
	//  1) https://core.telegram.org/method/upload.saveFilePart
	//
	// Use SetFile and GetFile helpers.
	File InputFileClass
	// Square video for animated profile picture
	//
	// Use SetVideo and GetVideo helpers.
	Video InputFileClass
	// Timestamp that should be shown as static preview to the user (seconds)
	//
	// Use SetVideoStartTs and GetVideoStartTs helpers.
	VideoStartTs float64
}

// InputChatUploadedPhotoTypeID is TL type id of InputChatUploadedPhoto.
const InputChatUploadedPhotoTypeID = 0xc642724e

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

	return true
}

// String implements fmt.Stringer.
func (i *InputChatUploadedPhoto) String() string {
	if i == nil {
		return "InputChatUploadedPhoto(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputChatUploadedPhoto")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(i.Flags))
	sb.WriteString(",\n")
	if i.Flags.Has(0) {
		sb.WriteString("\tFile: ")
		sb.WriteString(fmt.Sprint(i.File))
		sb.WriteString(",\n")
	}
	if i.Flags.Has(1) {
		sb.WriteString("\tVideo: ")
		sb.WriteString(fmt.Sprint(i.Video))
		sb.WriteString(",\n")
	}
	if i.Flags.Has(2) {
		sb.WriteString("\tVideoStartTs: ")
		sb.WriteString(fmt.Sprint(i.VideoStartTs))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputChatUploadedPhoto) TypeID() uint32 {
	return InputChatUploadedPhotoTypeID
}

// Encode implements bin.Encoder.
func (i *InputChatUploadedPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatUploadedPhoto#c642724e as nil")
	}
	b.PutID(InputChatUploadedPhotoTypeID)
	if !(i.File == nil) {
		i.Flags.Set(0)
	}
	if !(i.Video == nil) {
		i.Flags.Set(1)
	}
	if !(i.VideoStartTs == 0) {
		i.Flags.Set(2)
	}
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputChatUploadedPhoto#c642724e: field flags: %w", err)
	}
	if i.Flags.Has(0) {
		if i.File == nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#c642724e: field file is nil")
		}
		if err := i.File.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#c642724e: field file: %w", err)
		}
	}
	if i.Flags.Has(1) {
		if i.Video == nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#c642724e: field video is nil")
		}
		if err := i.Video.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputChatUploadedPhoto#c642724e: field video: %w", err)
		}
	}
	if i.Flags.Has(2) {
		b.PutDouble(i.VideoStartTs)
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
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.VideoStartTs, true
}

// Decode implements bin.Decoder.
func (i *InputChatUploadedPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatUploadedPhoto#c642724e to nil")
	}
	if err := b.ConsumeID(InputChatUploadedPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChatUploadedPhoto#c642724e: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#c642724e: field flags: %w", err)
		}
	}
	if i.Flags.Has(0) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#c642724e: field file: %w", err)
		}
		i.File = value
	}
	if i.Flags.Has(1) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#c642724e: field video: %w", err)
		}
		i.Video = value
	}
	if i.Flags.Has(2) {
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode inputChatUploadedPhoto#c642724e: field video_start_ts: %w", err)
		}
		i.VideoStartTs = value
	}
	return nil
}

// construct implements constructor of InputChatPhotoClass.
func (i InputChatUploadedPhoto) construct() InputChatPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputChatUploadedPhoto.
var (
	_ bin.Encoder = &InputChatUploadedPhoto{}
	_ bin.Decoder = &InputChatUploadedPhoto{}

	_ InputChatPhotoClass = &InputChatUploadedPhoto{}
)

// InputChatPhoto represents TL type `inputChatPhoto#8953ad37`.
// Existing photo to be set as a chat profile photo.
//
// See https://core.telegram.org/constructor/inputChatPhoto for reference.
type InputChatPhoto struct {
	// Existing photo
	ID InputPhotoClass
}

// InputChatPhotoTypeID is TL type id of InputChatPhoto.
const InputChatPhotoTypeID = 0x8953ad37

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
	var sb strings.Builder
	sb.WriteString("InputChatPhoto")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputChatPhoto) TypeID() uint32 {
	return InputChatPhotoTypeID
}

// Encode implements bin.Encoder.
func (i *InputChatPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChatPhoto#8953ad37 as nil")
	}
	b.PutID(InputChatPhotoTypeID)
	if i.ID == nil {
		return fmt.Errorf("unable to encode inputChatPhoto#8953ad37: field id is nil")
	}
	if err := i.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputChatPhoto#8953ad37: field id: %w", err)
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputChatPhoto) GetID() (value InputPhotoClass) {
	return i.ID
}

// Decode implements bin.Decoder.
func (i *InputChatPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChatPhoto#8953ad37 to nil")
	}
	if err := b.ConsumeID(InputChatPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChatPhoto#8953ad37: %w", err)
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

// construct implements constructor of InputChatPhotoClass.
func (i InputChatPhoto) construct() InputChatPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputChatPhoto.
var (
	_ bin.Encoder = &InputChatPhoto{}
	_ bin.Decoder = &InputChatPhoto{}

	_ InputChatPhotoClass = &InputChatPhoto{}
)

// InputChatPhotoClass represents InputChatPhoto generic type.
//
// See https://core.telegram.org/type/InputChatPhoto for reference.
//
// Example:
//  g, err := DecodeInputChatPhoto(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputChatPhotoEmpty: // inputChatPhotoEmpty#1ca48f57
//  case *InputChatUploadedPhoto: // inputChatUploadedPhoto#c642724e
//  case *InputChatPhoto: // inputChatPhoto#8953ad37
//  default: panic(v)
//  }
type InputChatPhotoClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputChatPhotoClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
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
		// Decoding inputChatUploadedPhoto#c642724e.
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
