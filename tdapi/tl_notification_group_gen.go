// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
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
)

// NotificationGroup represents TL type `notificationGroup#d02a41ba`.
type NotificationGroup struct {
	// Unique persistent auto-incremented from 1 identifier of the notification group
	ID int32
	// Type of the group
	Type NotificationGroupTypeClass
	// Identifier of a chat to which all notifications in the group belong
	ChatID int64
	// Total number of active notifications in the group
	TotalCount int32
	// The list of active notifications
	Notifications []Notification
}

// NotificationGroupTypeID is TL type id of NotificationGroup.
const NotificationGroupTypeID = 0xd02a41ba

// Ensuring interfaces in compile-time for NotificationGroup.
var (
	_ bin.Encoder     = &NotificationGroup{}
	_ bin.Decoder     = &NotificationGroup{}
	_ bin.BareEncoder = &NotificationGroup{}
	_ bin.BareDecoder = &NotificationGroup{}
)

func (n *NotificationGroup) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.ID == 0) {
		return false
	}
	if !(n.Type == nil) {
		return false
	}
	if !(n.ChatID == 0) {
		return false
	}
	if !(n.TotalCount == 0) {
		return false
	}
	if !(n.Notifications == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotificationGroup) String() string {
	if n == nil {
		return "NotificationGroup(nil)"
	}
	type Alias NotificationGroup
	return fmt.Sprintf("NotificationGroup%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NotificationGroup) TypeID() uint32 {
	return NotificationGroupTypeID
}

// TypeName returns name of type in TL schema.
func (*NotificationGroup) TypeName() string {
	return "notificationGroup"
}

// TypeInfo returns info about TL type.
func (n *NotificationGroup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "notificationGroup",
		ID:   NotificationGroupTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "Notifications",
			SchemaName: "notifications",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NotificationGroup) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationGroup#d02a41ba as nil")
	}
	b.PutID(NotificationGroupTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NotificationGroup) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationGroup#d02a41ba as nil")
	}
	b.PutInt32(n.ID)
	if n.Type == nil {
		return fmt.Errorf("unable to encode notificationGroup#d02a41ba: field type is nil")
	}
	if err := n.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode notificationGroup#d02a41ba: field type: %w", err)
	}
	b.PutLong(n.ChatID)
	b.PutInt32(n.TotalCount)
	b.PutInt(len(n.Notifications))
	for idx, v := range n.Notifications {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare notificationGroup#d02a41ba: field notifications element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NotificationGroup) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationGroup#d02a41ba to nil")
	}
	if err := b.ConsumeID(NotificationGroupTypeID); err != nil {
		return fmt.Errorf("unable to decode notificationGroup#d02a41ba: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NotificationGroup) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationGroup#d02a41ba to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode notificationGroup#d02a41ba: field id: %w", err)
		}
		n.ID = value
	}
	{
		value, err := DecodeNotificationGroupType(b)
		if err != nil {
			return fmt.Errorf("unable to decode notificationGroup#d02a41ba: field type: %w", err)
		}
		n.Type = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode notificationGroup#d02a41ba: field chat_id: %w", err)
		}
		n.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode notificationGroup#d02a41ba: field total_count: %w", err)
		}
		n.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode notificationGroup#d02a41ba: field notifications: %w", err)
		}

		if headerLen > 0 {
			n.Notifications = make([]Notification, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Notification
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare notificationGroup#d02a41ba: field notifications: %w", err)
			}
			n.Notifications = append(n.Notifications, value)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (n *NotificationGroup) GetID() (value int32) {
	return n.ID
}

// GetType returns value of Type field.
func (n *NotificationGroup) GetType() (value NotificationGroupTypeClass) {
	return n.Type
}

// GetChatID returns value of ChatID field.
func (n *NotificationGroup) GetChatID() (value int64) {
	return n.ChatID
}

// GetTotalCount returns value of TotalCount field.
func (n *NotificationGroup) GetTotalCount() (value int32) {
	return n.TotalCount
}

// GetNotifications returns value of Notifications field.
func (n *NotificationGroup) GetNotifications() (value []Notification) {
	return n.Notifications
}
