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

// MessageInteractionCounters represents TL type `messageInteractionCounters#ad4fc9bd`.
// Message interaction counters
//
// See https://core.telegram.org/constructor/messageInteractionCounters for reference.
type MessageInteractionCounters struct {
	// Message ID
	MsgID int
	// Views
	Views int
	// Number of times this message was forwarded
	Forwards int
}

// MessageInteractionCountersTypeID is TL type id of MessageInteractionCounters.
const MessageInteractionCountersTypeID = 0xad4fc9bd

func (m *MessageInteractionCounters) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.MsgID == 0) {
		return false
	}
	if !(m.Views == 0) {
		return false
	}
	if !(m.Forwards == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageInteractionCounters) String() string {
	if m == nil {
		return "MessageInteractionCounters(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessageInteractionCounters")
	sb.WriteString("{\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(m.MsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tViews: ")
	sb.WriteString(fmt.Sprint(m.Views))
	sb.WriteString(",\n")
	sb.WriteString("\tForwards: ")
	sb.WriteString(fmt.Sprint(m.Forwards))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MessageInteractionCounters) TypeID() uint32 {
	return MessageInteractionCountersTypeID
}

// Encode implements bin.Encoder.
func (m *MessageInteractionCounters) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageInteractionCounters#ad4fc9bd as nil")
	}
	b.PutID(MessageInteractionCountersTypeID)
	b.PutInt(m.MsgID)
	b.PutInt(m.Views)
	b.PutInt(m.Forwards)
	return nil
}

// GetMsgID returns value of MsgID field.
func (m *MessageInteractionCounters) GetMsgID() (value int) {
	return m.MsgID
}

// GetViews returns value of Views field.
func (m *MessageInteractionCounters) GetViews() (value int) {
	return m.Views
}

// GetForwards returns value of Forwards field.
func (m *MessageInteractionCounters) GetForwards() (value int) {
	return m.Forwards
}

// Decode implements bin.Decoder.
func (m *MessageInteractionCounters) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageInteractionCounters#ad4fc9bd to nil")
	}
	if err := b.ConsumeID(MessageInteractionCountersTypeID); err != nil {
		return fmt.Errorf("unable to decode messageInteractionCounters#ad4fc9bd: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageInteractionCounters#ad4fc9bd: field msg_id: %w", err)
		}
		m.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageInteractionCounters#ad4fc9bd: field views: %w", err)
		}
		m.Views = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageInteractionCounters#ad4fc9bd: field forwards: %w", err)
		}
		m.Forwards = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessageInteractionCounters.
var (
	_ bin.Encoder = &MessageInteractionCounters{}
	_ bin.Decoder = &MessageInteractionCounters{}
)
