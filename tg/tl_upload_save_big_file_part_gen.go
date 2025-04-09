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

// UploadSaveBigFilePartRequest represents TL type `upload.saveBigFilePart#de7b673d`.
type UploadSaveBigFilePartRequest struct {
	// FileID field of UploadSaveBigFilePartRequest.
	FileID int64
	// FilePart field of UploadSaveBigFilePartRequest.
	FilePart int
	// FileTotalParts field of UploadSaveBigFilePartRequest.
	FileTotalParts int
	// Bytes field of UploadSaveBigFilePartRequest.
	Bytes []byte
}

// UploadSaveBigFilePartRequestTypeID is TL type id of UploadSaveBigFilePartRequest.
const UploadSaveBigFilePartRequestTypeID = 0xde7b673d

// Ensuring interfaces in compile-time for UploadSaveBigFilePartRequest.
var (
	_ bin.Encoder     = &UploadSaveBigFilePartRequest{}
	_ bin.Decoder     = &UploadSaveBigFilePartRequest{}
	_ bin.BareEncoder = &UploadSaveBigFilePartRequest{}
	_ bin.BareDecoder = &UploadSaveBigFilePartRequest{}
)

func (s *UploadSaveBigFilePartRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.FileID == 0) {
		return false
	}
	if !(s.FilePart == 0) {
		return false
	}
	if !(s.FileTotalParts == 0) {
		return false
	}
	if !(s.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *UploadSaveBigFilePartRequest) String() string {
	if s == nil {
		return "UploadSaveBigFilePartRequest(nil)"
	}
	type Alias UploadSaveBigFilePartRequest
	return fmt.Sprintf("UploadSaveBigFilePartRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadSaveBigFilePartRequest) TypeID() uint32 {
	return UploadSaveBigFilePartRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadSaveBigFilePartRequest) TypeName() string {
	return "upload.saveBigFilePart"
}

// TypeInfo returns info about TL type.
func (s *UploadSaveBigFilePartRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.saveBigFilePart",
		ID:   UploadSaveBigFilePartRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileID",
			SchemaName: "file_id",
		},
		{
			Name:       "FilePart",
			SchemaName: "file_part",
		},
		{
			Name:       "FileTotalParts",
			SchemaName: "file_total_parts",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *UploadSaveBigFilePartRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode upload.saveBigFilePart#de7b673d as nil")
	}
	b.PutID(UploadSaveBigFilePartRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *UploadSaveBigFilePartRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode upload.saveBigFilePart#de7b673d as nil")
	}
	b.PutLong(s.FileID)
	b.PutInt(s.FilePart)
	b.PutInt(s.FileTotalParts)
	b.PutBytes(s.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (s *UploadSaveBigFilePartRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode upload.saveBigFilePart#de7b673d to nil")
	}
	if err := b.ConsumeID(UploadSaveBigFilePartRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.saveBigFilePart#de7b673d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *UploadSaveBigFilePartRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode upload.saveBigFilePart#de7b673d to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveBigFilePart#de7b673d: field file_id: %w", err)
		}
		s.FileID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveBigFilePart#de7b673d: field file_part: %w", err)
		}
		s.FilePart = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveBigFilePart#de7b673d: field file_total_parts: %w", err)
		}
		s.FileTotalParts = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveBigFilePart#de7b673d: field bytes: %w", err)
		}
		s.Bytes = value
	}
	return nil
}

// GetFileID returns value of FileID field.
func (s *UploadSaveBigFilePartRequest) GetFileID() (value int64) {
	if s == nil {
		return
	}
	return s.FileID
}

// GetFilePart returns value of FilePart field.
func (s *UploadSaveBigFilePartRequest) GetFilePart() (value int) {
	if s == nil {
		return
	}
	return s.FilePart
}

// GetFileTotalParts returns value of FileTotalParts field.
func (s *UploadSaveBigFilePartRequest) GetFileTotalParts() (value int) {
	if s == nil {
		return
	}
	return s.FileTotalParts
}

// GetBytes returns value of Bytes field.
func (s *UploadSaveBigFilePartRequest) GetBytes() (value []byte) {
	if s == nil {
		return
	}
	return s.Bytes
}

// UploadSaveBigFilePart invokes method upload.saveBigFilePart#de7b673d returning error if any.
func (c *Client) UploadSaveBigFilePart(ctx context.Context, request *UploadSaveBigFilePartRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
