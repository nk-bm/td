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

// MessageReplyHeader represents TL type `messageReplyHeader#a6d57763`.
// Message replies and thread¹ information
//
// Links:
//  1) https://core.telegram.org/api/threads
//
// See https://core.telegram.org/constructor/messageReplyHeader for reference.
type MessageReplyHeader struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// ID of message to which this message is replying
	ReplyToMsgID int
	// For replies sent in channel discussion threads¹ of which the current user is not a member, the discussion group ID
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetReplyToPeerID and GetReplyToPeerID helpers.
	ReplyToPeerID PeerClass
	// ID of the message that started this message thread¹
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetReplyToTopID and GetReplyToTopID helpers.
	ReplyToTopID int
}

// MessageReplyHeaderTypeID is TL type id of MessageReplyHeader.
const MessageReplyHeaderTypeID = 0xa6d57763

func (m *MessageReplyHeader) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.ReplyToMsgID == 0) {
		return false
	}
	if !(m.ReplyToPeerID == nil) {
		return false
	}
	if !(m.ReplyToTopID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageReplyHeader) String() string {
	if m == nil {
		return "MessageReplyHeader(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessageReplyHeader")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(m.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tReplyToMsgID: ")
	sb.WriteString(fmt.Sprint(m.ReplyToMsgID))
	sb.WriteString(",\n")
	if m.Flags.Has(0) {
		sb.WriteString("\tReplyToPeerID: ")
		sb.WriteString(fmt.Sprint(m.ReplyToPeerID))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(1) {
		sb.WriteString("\tReplyToTopID: ")
		sb.WriteString(fmt.Sprint(m.ReplyToTopID))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MessageReplyHeader) TypeID() uint32 {
	return MessageReplyHeaderTypeID
}

// Encode implements bin.Encoder.
func (m *MessageReplyHeader) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyHeader#a6d57763 as nil")
	}
	b.PutID(MessageReplyHeaderTypeID)
	if !(m.ReplyToPeerID == nil) {
		m.Flags.Set(0)
	}
	if !(m.ReplyToTopID == 0) {
		m.Flags.Set(1)
	}
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageReplyHeader#a6d57763: field flags: %w", err)
	}
	b.PutInt(m.ReplyToMsgID)
	if m.Flags.Has(0) {
		if m.ReplyToPeerID == nil {
			return fmt.Errorf("unable to encode messageReplyHeader#a6d57763: field reply_to_peer_id is nil")
		}
		if err := m.ReplyToPeerID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageReplyHeader#a6d57763: field reply_to_peer_id: %w", err)
		}
	}
	if m.Flags.Has(1) {
		b.PutInt(m.ReplyToTopID)
	}
	return nil
}

// GetReplyToMsgID returns value of ReplyToMsgID field.
func (m *MessageReplyHeader) GetReplyToMsgID() (value int) {
	return m.ReplyToMsgID
}

// SetReplyToPeerID sets value of ReplyToPeerID conditional field.
func (m *MessageReplyHeader) SetReplyToPeerID(value PeerClass) {
	m.Flags.Set(0)
	m.ReplyToPeerID = value
}

// GetReplyToPeerID returns value of ReplyToPeerID conditional field and
// boolean which is true if field was set.
func (m *MessageReplyHeader) GetReplyToPeerID() (value PeerClass, ok bool) {
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.ReplyToPeerID, true
}

// SetReplyToTopID sets value of ReplyToTopID conditional field.
func (m *MessageReplyHeader) SetReplyToTopID(value int) {
	m.Flags.Set(1)
	m.ReplyToTopID = value
}

// GetReplyToTopID returns value of ReplyToTopID conditional field and
// boolean which is true if field was set.
func (m *MessageReplyHeader) GetReplyToTopID() (value int, ok bool) {
	if !m.Flags.Has(1) {
		return value, false
	}
	return m.ReplyToTopID, true
}

// Decode implements bin.Decoder.
func (m *MessageReplyHeader) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyHeader#a6d57763 to nil")
	}
	if err := b.ConsumeID(MessageReplyHeaderTypeID); err != nil {
		return fmt.Errorf("unable to decode messageReplyHeader#a6d57763: %w", err)
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#a6d57763: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#a6d57763: field reply_to_msg_id: %w", err)
		}
		m.ReplyToMsgID = value
	}
	if m.Flags.Has(0) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#a6d57763: field reply_to_peer_id: %w", err)
		}
		m.ReplyToPeerID = value
	}
	if m.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyHeader#a6d57763: field reply_to_top_id: %w", err)
		}
		m.ReplyToTopID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessageReplyHeader.
var (
	_ bin.Encoder = &MessageReplyHeader{}
	_ bin.Decoder = &MessageReplyHeader{}
)
