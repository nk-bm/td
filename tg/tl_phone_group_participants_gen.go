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

// PhoneGroupParticipants represents TL type `phone.groupParticipants#9cfeb92d`.
//
// See https://core.telegram.org/constructor/phone.groupParticipants for reference.
type PhoneGroupParticipants struct {
	// Count field of PhoneGroupParticipants.
	Count int
	// Participants field of PhoneGroupParticipants.
	Participants []GroupCallParticipant
	// NextOffset field of PhoneGroupParticipants.
	NextOffset string
	// Users field of PhoneGroupParticipants.
	Users []UserClass
	// Version field of PhoneGroupParticipants.
	Version int
}

// PhoneGroupParticipantsTypeID is TL type id of PhoneGroupParticipants.
const PhoneGroupParticipantsTypeID = 0x9cfeb92d

func (g *PhoneGroupParticipants) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Count == 0) {
		return false
	}
	if !(g.Participants == nil) {
		return false
	}
	if !(g.NextOffset == "") {
		return false
	}
	if !(g.Users == nil) {
		return false
	}
	if !(g.Version == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhoneGroupParticipants) String() string {
	if g == nil {
		return "PhoneGroupParticipants(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneGroupParticipants")
	sb.WriteString("{\n")
	sb.WriteString("\tCount: ")
	sb.WriteString(fmt.Sprint(g.Count))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range g.Participants {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("\tNextOffset: ")
	sb.WriteString(fmt.Sprint(g.NextOffset))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range g.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("\tVersion: ")
	sb.WriteString(fmt.Sprint(g.Version))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *PhoneGroupParticipants) TypeID() uint32 {
	return PhoneGroupParticipantsTypeID
}

// Encode implements bin.Encoder.
func (g *PhoneGroupParticipants) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.groupParticipants#9cfeb92d as nil")
	}
	b.PutID(PhoneGroupParticipantsTypeID)
	b.PutInt(g.Count)
	b.PutVectorHeader(len(g.Participants))
	for idx, v := range g.Participants {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.groupParticipants#9cfeb92d: field participants element with index %d: %w", idx, err)
		}
	}
	b.PutString(g.NextOffset)
	b.PutVectorHeader(len(g.Users))
	for idx, v := range g.Users {
		if v == nil {
			return fmt.Errorf("unable to encode phone.groupParticipants#9cfeb92d: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.groupParticipants#9cfeb92d: field users element with index %d: %w", idx, err)
		}
	}
	b.PutInt(g.Version)
	return nil
}

// GetCount returns value of Count field.
func (g *PhoneGroupParticipants) GetCount() (value int) {
	return g.Count
}

// GetParticipants returns value of Participants field.
func (g *PhoneGroupParticipants) GetParticipants() (value []GroupCallParticipant) {
	return g.Participants
}

// GetNextOffset returns value of NextOffset field.
func (g *PhoneGroupParticipants) GetNextOffset() (value string) {
	return g.NextOffset
}

// GetUsers returns value of Users field.
func (g *PhoneGroupParticipants) GetUsers() (value []UserClass) {
	return g.Users
}

// GetVersion returns value of Version field.
func (g *PhoneGroupParticipants) GetVersion() (value int) {
	return g.Version
}

// Decode implements bin.Decoder.
func (g *PhoneGroupParticipants) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.groupParticipants#9cfeb92d to nil")
	}
	if err := b.ConsumeID(PhoneGroupParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.groupParticipants#9cfeb92d: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupParticipants#9cfeb92d: field count: %w", err)
		}
		g.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupParticipants#9cfeb92d: field participants: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value GroupCallParticipant
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode phone.groupParticipants#9cfeb92d: field participants: %w", err)
			}
			g.Participants = append(g.Participants, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupParticipants#9cfeb92d: field next_offset: %w", err)
		}
		g.NextOffset = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupParticipants#9cfeb92d: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode phone.groupParticipants#9cfeb92d: field users: %w", err)
			}
			g.Users = append(g.Users, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.groupParticipants#9cfeb92d: field version: %w", err)
		}
		g.Version = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneGroupParticipants.
var (
	_ bin.Encoder = &PhoneGroupParticipants{}
	_ bin.Decoder = &PhoneGroupParticipants{}
)
