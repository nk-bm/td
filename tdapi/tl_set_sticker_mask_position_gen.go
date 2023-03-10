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

// SetStickerMaskPositionRequest represents TL type `setStickerMaskPosition#47a959d0`.
type SetStickerMaskPositionRequest struct {
	// Sticker
	Sticker InputFileClass
	// Position where the mask is placed; pass null to remove mask position
	MaskPosition MaskPosition
}

// SetStickerMaskPositionRequestTypeID is TL type id of SetStickerMaskPositionRequest.
const SetStickerMaskPositionRequestTypeID = 0x47a959d0

// Ensuring interfaces in compile-time for SetStickerMaskPositionRequest.
var (
	_ bin.Encoder     = &SetStickerMaskPositionRequest{}
	_ bin.Decoder     = &SetStickerMaskPositionRequest{}
	_ bin.BareEncoder = &SetStickerMaskPositionRequest{}
	_ bin.BareDecoder = &SetStickerMaskPositionRequest{}
)

func (s *SetStickerMaskPositionRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Sticker == nil) {
		return false
	}
	if !(s.MaskPosition.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetStickerMaskPositionRequest) String() string {
	if s == nil {
		return "SetStickerMaskPositionRequest(nil)"
	}
	type Alias SetStickerMaskPositionRequest
	return fmt.Sprintf("SetStickerMaskPositionRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetStickerMaskPositionRequest) TypeID() uint32 {
	return SetStickerMaskPositionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetStickerMaskPositionRequest) TypeName() string {
	return "setStickerMaskPosition"
}

// TypeInfo returns info about TL type.
func (s *SetStickerMaskPositionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setStickerMaskPosition",
		ID:   SetStickerMaskPositionRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
		{
			Name:       "MaskPosition",
			SchemaName: "mask_position",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetStickerMaskPositionRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setStickerMaskPosition#47a959d0 as nil")
	}
	b.PutID(SetStickerMaskPositionRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetStickerMaskPositionRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setStickerMaskPosition#47a959d0 as nil")
	}
	if s.Sticker == nil {
		return fmt.Errorf("unable to encode setStickerMaskPosition#47a959d0: field sticker is nil")
	}
	if err := s.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setStickerMaskPosition#47a959d0: field sticker: %w", err)
	}
	if err := s.MaskPosition.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setStickerMaskPosition#47a959d0: field mask_position: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetStickerMaskPositionRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setStickerMaskPosition#47a959d0 to nil")
	}
	if err := b.ConsumeID(SetStickerMaskPositionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setStickerMaskPosition#47a959d0: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetStickerMaskPositionRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setStickerMaskPosition#47a959d0 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode setStickerMaskPosition#47a959d0: field sticker: %w", err)
		}
		s.Sticker = value
	}
	{
		if err := s.MaskPosition.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setStickerMaskPosition#47a959d0: field mask_position: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetStickerMaskPositionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setStickerMaskPosition#47a959d0 as nil")
	}
	b.ObjStart()
	b.PutID("setStickerMaskPosition")
	b.Comma()
	b.FieldStart("sticker")
	if s.Sticker == nil {
		return fmt.Errorf("unable to encode setStickerMaskPosition#47a959d0: field sticker is nil")
	}
	if err := s.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setStickerMaskPosition#47a959d0: field sticker: %w", err)
	}
	b.Comma()
	b.FieldStart("mask_position")
	if err := s.MaskPosition.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setStickerMaskPosition#47a959d0: field mask_position: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetStickerMaskPositionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setStickerMaskPosition#47a959d0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setStickerMaskPosition"); err != nil {
				return fmt.Errorf("unable to decode setStickerMaskPosition#47a959d0: %w", err)
			}
		case "sticker":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode setStickerMaskPosition#47a959d0: field sticker: %w", err)
			}
			s.Sticker = value
		case "mask_position":
			if err := s.MaskPosition.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setStickerMaskPosition#47a959d0: field mask_position: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSticker returns value of Sticker field.
func (s *SetStickerMaskPositionRequest) GetSticker() (value InputFileClass) {
	if s == nil {
		return
	}
	return s.Sticker
}

// GetMaskPosition returns value of MaskPosition field.
func (s *SetStickerMaskPositionRequest) GetMaskPosition() (value MaskPosition) {
	if s == nil {
		return
	}
	return s.MaskPosition
}

// SetStickerMaskPosition invokes method setStickerMaskPosition#47a959d0 returning error if any.
func (c *Client) SetStickerMaskPosition(ctx context.Context, request *SetStickerMaskPositionRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
