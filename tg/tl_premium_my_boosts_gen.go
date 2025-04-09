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

// PremiumMyBoosts represents TL type `premium.myBoosts#9ae228e2`.
type PremiumMyBoosts struct {
	// MyBoosts field of PremiumMyBoosts.
	MyBoosts []MyBoost
	// Chats field of PremiumMyBoosts.
	Chats []ChatClass
	// Users field of PremiumMyBoosts.
	Users []UserClass
}

// PremiumMyBoostsTypeID is TL type id of PremiumMyBoosts.
const PremiumMyBoostsTypeID = 0x9ae228e2

// Ensuring interfaces in compile-time for PremiumMyBoosts.
var (
	_ bin.Encoder     = &PremiumMyBoosts{}
	_ bin.Decoder     = &PremiumMyBoosts{}
	_ bin.BareEncoder = &PremiumMyBoosts{}
	_ bin.BareDecoder = &PremiumMyBoosts{}
)

func (m *PremiumMyBoosts) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.MyBoosts == nil) {
		return false
	}
	if !(m.Chats == nil) {
		return false
	}
	if !(m.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *PremiumMyBoosts) String() string {
	if m == nil {
		return "PremiumMyBoosts(nil)"
	}
	type Alias PremiumMyBoosts
	return fmt.Sprintf("PremiumMyBoosts%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumMyBoosts) TypeID() uint32 {
	return PremiumMyBoostsTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumMyBoosts) TypeName() string {
	return "premium.myBoosts"
}

// TypeInfo returns info about TL type.
func (m *PremiumMyBoosts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premium.myBoosts",
		ID:   PremiumMyBoostsTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MyBoosts",
			SchemaName: "my_boosts",
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

// Encode implements bin.Encoder.
func (m *PremiumMyBoosts) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode premium.myBoosts#9ae228e2 as nil")
	}
	b.PutID(PremiumMyBoostsTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *PremiumMyBoosts) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode premium.myBoosts#9ae228e2 as nil")
	}
	b.PutVectorHeader(len(m.MyBoosts))
	for idx, v := range m.MyBoosts {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode premium.myBoosts#9ae228e2: field my_boosts element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Chats))
	for idx, v := range m.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode premium.myBoosts#9ae228e2: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode premium.myBoosts#9ae228e2: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Users))
	for idx, v := range m.Users {
		if v == nil {
			return fmt.Errorf("unable to encode premium.myBoosts#9ae228e2: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode premium.myBoosts#9ae228e2: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *PremiumMyBoosts) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode premium.myBoosts#9ae228e2 to nil")
	}
	if err := b.ConsumeID(PremiumMyBoostsTypeID); err != nil {
		return fmt.Errorf("unable to decode premium.myBoosts#9ae228e2: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *PremiumMyBoosts) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode premium.myBoosts#9ae228e2 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode premium.myBoosts#9ae228e2: field my_boosts: %w", err)
		}

		if headerLen > 0 {
			m.MyBoosts = make([]MyBoost, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MyBoost
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode premium.myBoosts#9ae228e2: field my_boosts: %w", err)
			}
			m.MyBoosts = append(m.MyBoosts, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode premium.myBoosts#9ae228e2: field chats: %w", err)
		}

		if headerLen > 0 {
			m.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode premium.myBoosts#9ae228e2: field chats: %w", err)
			}
			m.Chats = append(m.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode premium.myBoosts#9ae228e2: field users: %w", err)
		}

		if headerLen > 0 {
			m.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode premium.myBoosts#9ae228e2: field users: %w", err)
			}
			m.Users = append(m.Users, value)
		}
	}
	return nil
}

// GetMyBoosts returns value of MyBoosts field.
func (m *PremiumMyBoosts) GetMyBoosts() (value []MyBoost) {
	if m == nil {
		return
	}
	return m.MyBoosts
}

// GetChats returns value of Chats field.
func (m *PremiumMyBoosts) GetChats() (value []ChatClass) {
	if m == nil {
		return
	}
	return m.Chats
}

// GetUsers returns value of Users field.
func (m *PremiumMyBoosts) GetUsers() (value []UserClass) {
	if m == nil {
		return
	}
	return m.Users
}
