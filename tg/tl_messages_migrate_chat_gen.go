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

// MessagesMigrateChatRequest represents TL type `messages.migrateChat#15a3b8e3`.
// Turn a legacy group into a supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/messages.migrateChat for reference.
type MessagesMigrateChatRequest struct {
	// Legacy group to migrate
	ChatID int
}

// MessagesMigrateChatRequestTypeID is TL type id of MessagesMigrateChatRequest.
const MessagesMigrateChatRequestTypeID = 0x15a3b8e3

func (m *MessagesMigrateChatRequest) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagesMigrateChatRequest) String() string {
	if m == nil {
		return "MessagesMigrateChatRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesMigrateChatRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChatID: ")
	sb.WriteString(fmt.Sprint(m.ChatID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MessagesMigrateChatRequest) TypeID() uint32 {
	return MessagesMigrateChatRequestTypeID
}

// Encode implements bin.Encoder.
func (m *MessagesMigrateChatRequest) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.migrateChat#15a3b8e3 as nil")
	}
	b.PutID(MessagesMigrateChatRequestTypeID)
	b.PutInt(m.ChatID)
	return nil
}

// GetChatID returns value of ChatID field.
func (m *MessagesMigrateChatRequest) GetChatID() (value int) {
	return m.ChatID
}

// Decode implements bin.Decoder.
func (m *MessagesMigrateChatRequest) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.migrateChat#15a3b8e3 to nil")
	}
	if err := b.ConsumeID(MessagesMigrateChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.migrateChat#15a3b8e3: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.migrateChat#15a3b8e3: field chat_id: %w", err)
		}
		m.ChatID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesMigrateChatRequest.
var (
	_ bin.Encoder = &MessagesMigrateChatRequest{}
	_ bin.Decoder = &MessagesMigrateChatRequest{}
)

// MessagesMigrateChat invokes method messages.migrateChat#15a3b8e3 returning error if any.
// Turn a legacy group into a supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  403 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.migrateChat for reference.
func (c *Client) MessagesMigrateChat(ctx context.Context, chatid int) (UpdatesClass, error) {
	var result UpdatesBox

	request := &MessagesMigrateChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
