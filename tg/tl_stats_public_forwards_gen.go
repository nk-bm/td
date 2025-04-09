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

// StatsPublicForwards represents TL type `stats.publicForwards#93037e20`.
type StatsPublicForwards struct {
	// Flags field of StatsPublicForwards.
	Flags bin.Fields
	// Count field of StatsPublicForwards.
	Count int
	// Forwards field of StatsPublicForwards.
	Forwards []PublicForwardClass
	// NextOffset field of StatsPublicForwards.
	//
	// Use SetNextOffset and GetNextOffset helpers.
	NextOffset string
	// Chats field of StatsPublicForwards.
	Chats []ChatClass
	// Users field of StatsPublicForwards.
	Users []UserClass
}

// StatsPublicForwardsTypeID is TL type id of StatsPublicForwards.
const StatsPublicForwardsTypeID = 0x93037e20

// Ensuring interfaces in compile-time for StatsPublicForwards.
var (
	_ bin.Encoder     = &StatsPublicForwards{}
	_ bin.Decoder     = &StatsPublicForwards{}
	_ bin.BareEncoder = &StatsPublicForwards{}
	_ bin.BareDecoder = &StatsPublicForwards{}
)

func (p *StatsPublicForwards) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Count == 0) {
		return false
	}
	if !(p.Forwards == nil) {
		return false
	}
	if !(p.NextOffset == "") {
		return false
	}
	if !(p.Chats == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *StatsPublicForwards) String() string {
	if p == nil {
		return "StatsPublicForwards(nil)"
	}
	type Alias StatsPublicForwards
	return fmt.Sprintf("StatsPublicForwards%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsPublicForwards) TypeID() uint32 {
	return StatsPublicForwardsTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsPublicForwards) TypeName() string {
	return "stats.publicForwards"
}

// TypeInfo returns info about TL type.
func (p *StatsPublicForwards) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.publicForwards",
		ID:   StatsPublicForwardsTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Forwards",
			SchemaName: "forwards",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (p *StatsPublicForwards) SetFlags() {
	if !(p.NextOffset == "") {
		p.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (p *StatsPublicForwards) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode stats.publicForwards#93037e20 as nil")
	}
	b.PutID(StatsPublicForwardsTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *StatsPublicForwards) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode stats.publicForwards#93037e20 as nil")
	}
	p.SetFlags()
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.publicForwards#93037e20: field flags: %w", err)
	}
	b.PutInt(p.Count)
	b.PutVectorHeader(len(p.Forwards))
	for idx, v := range p.Forwards {
		if v == nil {
			return fmt.Errorf("unable to encode stats.publicForwards#93037e20: field forwards element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.publicForwards#93037e20: field forwards element with index %d: %w", idx, err)
		}
	}
	if p.Flags.Has(0) {
		b.PutString(p.NextOffset)
	}
	b.PutVectorHeader(len(p.Chats))
	for idx, v := range p.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode stats.publicForwards#93037e20: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.publicForwards#93037e20: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode stats.publicForwards#93037e20: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stats.publicForwards#93037e20: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *StatsPublicForwards) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode stats.publicForwards#93037e20 to nil")
	}
	if err := b.ConsumeID(StatsPublicForwardsTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.publicForwards#93037e20: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *StatsPublicForwards) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode stats.publicForwards#93037e20 to nil")
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.publicForwards#93037e20: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.publicForwards#93037e20: field count: %w", err)
		}
		p.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.publicForwards#93037e20: field forwards: %w", err)
		}

		if headerLen > 0 {
			p.Forwards = make([]PublicForwardClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePublicForward(b)
			if err != nil {
				return fmt.Errorf("unable to decode stats.publicForwards#93037e20: field forwards: %w", err)
			}
			p.Forwards = append(p.Forwards, value)
		}
	}
	if p.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stats.publicForwards#93037e20: field next_offset: %w", err)
		}
		p.NextOffset = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.publicForwards#93037e20: field chats: %w", err)
		}

		if headerLen > 0 {
			p.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode stats.publicForwards#93037e20: field chats: %w", err)
			}
			p.Chats = append(p.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.publicForwards#93037e20: field users: %w", err)
		}

		if headerLen > 0 {
			p.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode stats.publicForwards#93037e20: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (p *StatsPublicForwards) GetCount() (value int) {
	if p == nil {
		return
	}
	return p.Count
}

// GetForwards returns value of Forwards field.
func (p *StatsPublicForwards) GetForwards() (value []PublicForwardClass) {
	if p == nil {
		return
	}
	return p.Forwards
}

// SetNextOffset sets value of NextOffset conditional field.
func (p *StatsPublicForwards) SetNextOffset(value string) {
	p.Flags.Set(0)
	p.NextOffset = value
}

// GetNextOffset returns value of NextOffset conditional field and
// boolean which is true if field was set.
func (p *StatsPublicForwards) GetNextOffset() (value string, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.NextOffset, true
}

// GetChats returns value of Chats field.
func (p *StatsPublicForwards) GetChats() (value []ChatClass) {
	if p == nil {
		return
	}
	return p.Chats
}

// GetUsers returns value of Users field.
func (p *StatsPublicForwards) GetUsers() (value []UserClass) {
	if p == nil {
		return
	}
	return p.Users
}
