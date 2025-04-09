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

// ExportedStoryLink represents TL type `exportedStoryLink#3fc9053b`.
type ExportedStoryLink struct {
	// Link field of ExportedStoryLink.
	Link string
}

// ExportedStoryLinkTypeID is TL type id of ExportedStoryLink.
const ExportedStoryLinkTypeID = 0x3fc9053b

// Ensuring interfaces in compile-time for ExportedStoryLink.
var (
	_ bin.Encoder     = &ExportedStoryLink{}
	_ bin.Decoder     = &ExportedStoryLink{}
	_ bin.BareEncoder = &ExportedStoryLink{}
	_ bin.BareDecoder = &ExportedStoryLink{}
)

func (e *ExportedStoryLink) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Link == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ExportedStoryLink) String() string {
	if e == nil {
		return "ExportedStoryLink(nil)"
	}
	type Alias ExportedStoryLink
	return fmt.Sprintf("ExportedStoryLink%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ExportedStoryLink) TypeID() uint32 {
	return ExportedStoryLinkTypeID
}

// TypeName returns name of type in TL schema.
func (*ExportedStoryLink) TypeName() string {
	return "exportedStoryLink"
}

// TypeInfo returns info about TL type.
func (e *ExportedStoryLink) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "exportedStoryLink",
		ID:   ExportedStoryLinkTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Link",
			SchemaName: "link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *ExportedStoryLink) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode exportedStoryLink#3fc9053b as nil")
	}
	b.PutID(ExportedStoryLinkTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *ExportedStoryLink) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode exportedStoryLink#3fc9053b as nil")
	}
	b.PutString(e.Link)
	return nil
}

// Decode implements bin.Decoder.
func (e *ExportedStoryLink) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode exportedStoryLink#3fc9053b to nil")
	}
	if err := b.ConsumeID(ExportedStoryLinkTypeID); err != nil {
		return fmt.Errorf("unable to decode exportedStoryLink#3fc9053b: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *ExportedStoryLink) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode exportedStoryLink#3fc9053b to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode exportedStoryLink#3fc9053b: field link: %w", err)
		}
		e.Link = value
	}
	return nil
}

// GetLink returns value of Link field.
func (e *ExportedStoryLink) GetLink() (value string) {
	if e == nil {
		return
	}
	return e.Link
}
