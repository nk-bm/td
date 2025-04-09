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

// MessagesHighScores represents TL type `messages.highScores#9a3bfd99`.
type MessagesHighScores struct {
	// Scores field of MessagesHighScores.
	Scores []HighScore
	// Users field of MessagesHighScores.
	Users []UserClass
}

// MessagesHighScoresTypeID is TL type id of MessagesHighScores.
const MessagesHighScoresTypeID = 0x9a3bfd99

// Ensuring interfaces in compile-time for MessagesHighScores.
var (
	_ bin.Encoder     = &MessagesHighScores{}
	_ bin.Decoder     = &MessagesHighScores{}
	_ bin.BareEncoder = &MessagesHighScores{}
	_ bin.BareDecoder = &MessagesHighScores{}
)

func (h *MessagesHighScores) Zero() bool {
	if h == nil {
		return true
	}
	if !(h.Scores == nil) {
		return false
	}
	if !(h.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (h *MessagesHighScores) String() string {
	if h == nil {
		return "MessagesHighScores(nil)"
	}
	type Alias MessagesHighScores
	return fmt.Sprintf("MessagesHighScores%+v", Alias(*h))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesHighScores) TypeID() uint32 {
	return MessagesHighScoresTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesHighScores) TypeName() string {
	return "messages.highScores"
}

// TypeInfo returns info about TL type.
func (h *MessagesHighScores) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.highScores",
		ID:   MessagesHighScoresTypeID,
	}
	if h == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Scores",
			SchemaName: "scores",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (h *MessagesHighScores) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.highScores#9a3bfd99 as nil")
	}
	b.PutID(MessagesHighScoresTypeID)
	return h.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (h *MessagesHighScores) EncodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.highScores#9a3bfd99 as nil")
	}
	b.PutVectorHeader(len(h.Scores))
	for idx, v := range h.Scores {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.highScores#9a3bfd99: field scores element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(h.Users))
	for idx, v := range h.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.highScores#9a3bfd99: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.highScores#9a3bfd99: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (h *MessagesHighScores) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.highScores#9a3bfd99 to nil")
	}
	if err := b.ConsumeID(MessagesHighScoresTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: %w", err)
	}
	return h.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (h *MessagesHighScores) DecodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.highScores#9a3bfd99 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: field scores: %w", err)
		}

		if headerLen > 0 {
			h.Scores = make([]HighScore, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value HighScore
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: field scores: %w", err)
			}
			h.Scores = append(h.Scores, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: field users: %w", err)
		}

		if headerLen > 0 {
			h.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: field users: %w", err)
			}
			h.Users = append(h.Users, value)
		}
	}
	return nil
}

// GetScores returns value of Scores field.
func (h *MessagesHighScores) GetScores() (value []HighScore) {
	if h == nil {
		return
	}
	return h.Scores
}

// GetUsers returns value of Users field.
func (h *MessagesHighScores) GetUsers() (value []UserClass) {
	if h == nil {
		return
	}
	return h.Users
}
