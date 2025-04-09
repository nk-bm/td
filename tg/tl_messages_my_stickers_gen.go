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

// MessagesMyStickers represents TL type `messages.myStickers#faff629d`.
type MessagesMyStickers struct {
	// Count field of MessagesMyStickers.
	Count int
	// Sets field of MessagesMyStickers.
	Sets []StickerSetCoveredClass
}

// MessagesMyStickersTypeID is TL type id of MessagesMyStickers.
const MessagesMyStickersTypeID = 0xfaff629d

// Ensuring interfaces in compile-time for MessagesMyStickers.
var (
	_ bin.Encoder     = &MessagesMyStickers{}
	_ bin.Decoder     = &MessagesMyStickers{}
	_ bin.BareEncoder = &MessagesMyStickers{}
	_ bin.BareDecoder = &MessagesMyStickers{}
)

func (m *MessagesMyStickers) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Count == 0) {
		return false
	}
	if !(m.Sets == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagesMyStickers) String() string {
	if m == nil {
		return "MessagesMyStickers(nil)"
	}
	type Alias MessagesMyStickers
	return fmt.Sprintf("MessagesMyStickers%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesMyStickers) TypeID() uint32 {
	return MessagesMyStickersTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesMyStickers) TypeName() string {
	return "messages.myStickers"
}

// TypeInfo returns info about TL type.
func (m *MessagesMyStickers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.myStickers",
		ID:   MessagesMyStickersTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Sets",
			SchemaName: "sets",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessagesMyStickers) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.myStickers#faff629d as nil")
	}
	b.PutID(MessagesMyStickersTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessagesMyStickers) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.myStickers#faff629d as nil")
	}
	b.PutInt(m.Count)
	b.PutVectorHeader(len(m.Sets))
	for idx, v := range m.Sets {
		if v == nil {
			return fmt.Errorf("unable to encode messages.myStickers#faff629d: field sets element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.myStickers#faff629d: field sets element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagesMyStickers) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.myStickers#faff629d to nil")
	}
	if err := b.ConsumeID(MessagesMyStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.myStickers#faff629d: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessagesMyStickers) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.myStickers#faff629d to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.myStickers#faff629d: field count: %w", err)
		}
		m.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.myStickers#faff629d: field sets: %w", err)
		}

		if headerLen > 0 {
			m.Sets = make([]StickerSetCoveredClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStickerSetCovered(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.myStickers#faff629d: field sets: %w", err)
			}
			m.Sets = append(m.Sets, value)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (m *MessagesMyStickers) GetCount() (value int) {
	if m == nil {
		return
	}
	return m.Count
}

// GetSets returns value of Sets field.
func (m *MessagesMyStickers) GetSets() (value []StickerSetCoveredClass) {
	if m == nil {
		return
	}
	return m.Sets
}
