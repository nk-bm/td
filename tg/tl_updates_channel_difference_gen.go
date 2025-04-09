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

// UpdatesChannelDifferenceEmpty represents TL type `updates.channelDifferenceEmpty#3e11affb`.
type UpdatesChannelDifferenceEmpty struct {
	// Flags field of UpdatesChannelDifferenceEmpty.
	Flags bin.Fields
	// Final field of UpdatesChannelDifferenceEmpty.
	Final bool
	// Pts field of UpdatesChannelDifferenceEmpty.
	Pts int
	// Timeout field of UpdatesChannelDifferenceEmpty.
	//
	// Use SetTimeout and GetTimeout helpers.
	Timeout int
}

// UpdatesChannelDifferenceEmptyTypeID is TL type id of UpdatesChannelDifferenceEmpty.
const UpdatesChannelDifferenceEmptyTypeID = 0x3e11affb

// construct implements constructor of UpdatesChannelDifferenceClass.
func (c UpdatesChannelDifferenceEmpty) construct() UpdatesChannelDifferenceClass { return &c }

// Ensuring interfaces in compile-time for UpdatesChannelDifferenceEmpty.
var (
	_ bin.Encoder     = &UpdatesChannelDifferenceEmpty{}
	_ bin.Decoder     = &UpdatesChannelDifferenceEmpty{}
	_ bin.BareEncoder = &UpdatesChannelDifferenceEmpty{}
	_ bin.BareDecoder = &UpdatesChannelDifferenceEmpty{}

	_ UpdatesChannelDifferenceClass = &UpdatesChannelDifferenceEmpty{}
)

func (c *UpdatesChannelDifferenceEmpty) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Final == false) {
		return false
	}
	if !(c.Pts == 0) {
		return false
	}
	if !(c.Timeout == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UpdatesChannelDifferenceEmpty) String() string {
	if c == nil {
		return "UpdatesChannelDifferenceEmpty(nil)"
	}
	type Alias UpdatesChannelDifferenceEmpty
	return fmt.Sprintf("UpdatesChannelDifferenceEmpty%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpdatesChannelDifferenceEmpty) TypeID() uint32 {
	return UpdatesChannelDifferenceEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*UpdatesChannelDifferenceEmpty) TypeName() string {
	return "updates.channelDifferenceEmpty"
}

// TypeInfo returns info about TL type.
func (c *UpdatesChannelDifferenceEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "updates.channelDifferenceEmpty",
		ID:   UpdatesChannelDifferenceEmptyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Final",
			SchemaName: "final",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Pts",
			SchemaName: "pts",
		},
		{
			Name:       "Timeout",
			SchemaName: "timeout",
			Null:       !c.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *UpdatesChannelDifferenceEmpty) SetFlags() {
	if !(c.Final == false) {
		c.Flags.Set(0)
	}
	if !(c.Timeout == 0) {
		c.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (c *UpdatesChannelDifferenceEmpty) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode updates.channelDifferenceEmpty#3e11affb as nil")
	}
	b.PutID(UpdatesChannelDifferenceEmptyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *UpdatesChannelDifferenceEmpty) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode updates.channelDifferenceEmpty#3e11affb as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.channelDifferenceEmpty#3e11affb: field flags: %w", err)
	}
	b.PutInt(c.Pts)
	if c.Flags.Has(1) {
		b.PutInt(c.Timeout)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *UpdatesChannelDifferenceEmpty) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode updates.channelDifferenceEmpty#3e11affb to nil")
	}
	if err := b.ConsumeID(UpdatesChannelDifferenceEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.channelDifferenceEmpty#3e11affb: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *UpdatesChannelDifferenceEmpty) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode updates.channelDifferenceEmpty#3e11affb to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceEmpty#3e11affb: field flags: %w", err)
		}
	}
	c.Final = c.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceEmpty#3e11affb: field pts: %w", err)
		}
		c.Pts = value
	}
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceEmpty#3e11affb: field timeout: %w", err)
		}
		c.Timeout = value
	}
	return nil
}

// SetFinal sets value of Final conditional field.
func (c *UpdatesChannelDifferenceEmpty) SetFinal(value bool) {
	if value {
		c.Flags.Set(0)
		c.Final = true
	} else {
		c.Flags.Unset(0)
		c.Final = false
	}
}

// GetFinal returns value of Final conditional field.
func (c *UpdatesChannelDifferenceEmpty) GetFinal() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// GetPts returns value of Pts field.
func (c *UpdatesChannelDifferenceEmpty) GetPts() (value int) {
	if c == nil {
		return
	}
	return c.Pts
}

// SetTimeout sets value of Timeout conditional field.
func (c *UpdatesChannelDifferenceEmpty) SetTimeout(value int) {
	c.Flags.Set(1)
	c.Timeout = value
}

// GetTimeout returns value of Timeout conditional field and
// boolean which is true if field was set.
func (c *UpdatesChannelDifferenceEmpty) GetTimeout() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.Timeout, true
}

// UpdatesChannelDifferenceTooLong represents TL type `updates.channelDifferenceTooLong#a4bcc6fe`.
type UpdatesChannelDifferenceTooLong struct {
	// Flags field of UpdatesChannelDifferenceTooLong.
	Flags bin.Fields
	// Final field of UpdatesChannelDifferenceTooLong.
	Final bool
	// Timeout field of UpdatesChannelDifferenceTooLong.
	//
	// Use SetTimeout and GetTimeout helpers.
	Timeout int
	// Dialog field of UpdatesChannelDifferenceTooLong.
	Dialog DialogClass
	// Messages field of UpdatesChannelDifferenceTooLong.
	Messages []MessageClass
	// Chats field of UpdatesChannelDifferenceTooLong.
	Chats []ChatClass
	// Users field of UpdatesChannelDifferenceTooLong.
	Users []UserClass
}

// UpdatesChannelDifferenceTooLongTypeID is TL type id of UpdatesChannelDifferenceTooLong.
const UpdatesChannelDifferenceTooLongTypeID = 0xa4bcc6fe

// construct implements constructor of UpdatesChannelDifferenceClass.
func (c UpdatesChannelDifferenceTooLong) construct() UpdatesChannelDifferenceClass { return &c }

// Ensuring interfaces in compile-time for UpdatesChannelDifferenceTooLong.
var (
	_ bin.Encoder     = &UpdatesChannelDifferenceTooLong{}
	_ bin.Decoder     = &UpdatesChannelDifferenceTooLong{}
	_ bin.BareEncoder = &UpdatesChannelDifferenceTooLong{}
	_ bin.BareDecoder = &UpdatesChannelDifferenceTooLong{}

	_ UpdatesChannelDifferenceClass = &UpdatesChannelDifferenceTooLong{}
)

func (c *UpdatesChannelDifferenceTooLong) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Final == false) {
		return false
	}
	if !(c.Timeout == 0) {
		return false
	}
	if !(c.Dialog == nil) {
		return false
	}
	if !(c.Messages == nil) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UpdatesChannelDifferenceTooLong) String() string {
	if c == nil {
		return "UpdatesChannelDifferenceTooLong(nil)"
	}
	type Alias UpdatesChannelDifferenceTooLong
	return fmt.Sprintf("UpdatesChannelDifferenceTooLong%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpdatesChannelDifferenceTooLong) TypeID() uint32 {
	return UpdatesChannelDifferenceTooLongTypeID
}

// TypeName returns name of type in TL schema.
func (*UpdatesChannelDifferenceTooLong) TypeName() string {
	return "updates.channelDifferenceTooLong"
}

// TypeInfo returns info about TL type.
func (c *UpdatesChannelDifferenceTooLong) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "updates.channelDifferenceTooLong",
		ID:   UpdatesChannelDifferenceTooLongTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Final",
			SchemaName: "final",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Timeout",
			SchemaName: "timeout",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "Dialog",
			SchemaName: "dialog",
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
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
func (c *UpdatesChannelDifferenceTooLong) SetFlags() {
	if !(c.Final == false) {
		c.Flags.Set(0)
	}
	if !(c.Timeout == 0) {
		c.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (c *UpdatesChannelDifferenceTooLong) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode updates.channelDifferenceTooLong#a4bcc6fe as nil")
	}
	b.PutID(UpdatesChannelDifferenceTooLongTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *UpdatesChannelDifferenceTooLong) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode updates.channelDifferenceTooLong#a4bcc6fe as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field flags: %w", err)
	}
	if c.Flags.Has(1) {
		b.PutInt(c.Timeout)
	}
	if c.Dialog == nil {
		return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field dialog is nil")
	}
	if err := c.Dialog.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field dialog: %w", err)
	}
	b.PutVectorHeader(len(c.Messages))
	for idx, v := range c.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *UpdatesChannelDifferenceTooLong) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode updates.channelDifferenceTooLong#a4bcc6fe to nil")
	}
	if err := b.ConsumeID(UpdatesChannelDifferenceTooLongTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *UpdatesChannelDifferenceTooLong) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode updates.channelDifferenceTooLong#a4bcc6fe to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field flags: %w", err)
		}
	}
	c.Final = c.Flags.Has(0)
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field timeout: %w", err)
		}
		c.Timeout = value
	}
	{
		value, err := DecodeDialog(b)
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field dialog: %w", err)
		}
		c.Dialog = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field messages: %w", err)
		}

		if headerLen > 0 {
			c.Messages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field messages: %w", err)
			}
			c.Messages = append(c.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field chats: %w", err)
		}

		if headerLen > 0 {
			c.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field users: %w", err)
		}

		if headerLen > 0 {
			c.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// SetFinal sets value of Final conditional field.
func (c *UpdatesChannelDifferenceTooLong) SetFinal(value bool) {
	if value {
		c.Flags.Set(0)
		c.Final = true
	} else {
		c.Flags.Unset(0)
		c.Final = false
	}
}

// GetFinal returns value of Final conditional field.
func (c *UpdatesChannelDifferenceTooLong) GetFinal() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// SetTimeout sets value of Timeout conditional field.
func (c *UpdatesChannelDifferenceTooLong) SetTimeout(value int) {
	c.Flags.Set(1)
	c.Timeout = value
}

// GetTimeout returns value of Timeout conditional field and
// boolean which is true if field was set.
func (c *UpdatesChannelDifferenceTooLong) GetTimeout() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.Timeout, true
}

// GetDialog returns value of Dialog field.
func (c *UpdatesChannelDifferenceTooLong) GetDialog() (value DialogClass) {
	if c == nil {
		return
	}
	return c.Dialog
}

// GetMessages returns value of Messages field.
func (c *UpdatesChannelDifferenceTooLong) GetMessages() (value []MessageClass) {
	if c == nil {
		return
	}
	return c.Messages
}

// GetChats returns value of Chats field.
func (c *UpdatesChannelDifferenceTooLong) GetChats() (value []ChatClass) {
	if c == nil {
		return
	}
	return c.Chats
}

// GetUsers returns value of Users field.
func (c *UpdatesChannelDifferenceTooLong) GetUsers() (value []UserClass) {
	if c == nil {
		return
	}
	return c.Users
}

// UpdatesChannelDifference represents TL type `updates.channelDifference#2064674e`.
type UpdatesChannelDifference struct {
	// Flags field of UpdatesChannelDifference.
	Flags bin.Fields
	// Final field of UpdatesChannelDifference.
	Final bool
	// Pts field of UpdatesChannelDifference.
	Pts int
	// Timeout field of UpdatesChannelDifference.
	//
	// Use SetTimeout and GetTimeout helpers.
	Timeout int
	// NewMessages field of UpdatesChannelDifference.
	NewMessages []MessageClass
	// OtherUpdates field of UpdatesChannelDifference.
	OtherUpdates []UpdateClass
	// Chats field of UpdatesChannelDifference.
	Chats []ChatClass
	// Users field of UpdatesChannelDifference.
	Users []UserClass
}

// UpdatesChannelDifferenceTypeID is TL type id of UpdatesChannelDifference.
const UpdatesChannelDifferenceTypeID = 0x2064674e

// construct implements constructor of UpdatesChannelDifferenceClass.
func (c UpdatesChannelDifference) construct() UpdatesChannelDifferenceClass { return &c }

// Ensuring interfaces in compile-time for UpdatesChannelDifference.
var (
	_ bin.Encoder     = &UpdatesChannelDifference{}
	_ bin.Decoder     = &UpdatesChannelDifference{}
	_ bin.BareEncoder = &UpdatesChannelDifference{}
	_ bin.BareDecoder = &UpdatesChannelDifference{}

	_ UpdatesChannelDifferenceClass = &UpdatesChannelDifference{}
)

func (c *UpdatesChannelDifference) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Final == false) {
		return false
	}
	if !(c.Pts == 0) {
		return false
	}
	if !(c.Timeout == 0) {
		return false
	}
	if !(c.NewMessages == nil) {
		return false
	}
	if !(c.OtherUpdates == nil) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UpdatesChannelDifference) String() string {
	if c == nil {
		return "UpdatesChannelDifference(nil)"
	}
	type Alias UpdatesChannelDifference
	return fmt.Sprintf("UpdatesChannelDifference%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpdatesChannelDifference) TypeID() uint32 {
	return UpdatesChannelDifferenceTypeID
}

// TypeName returns name of type in TL schema.
func (*UpdatesChannelDifference) TypeName() string {
	return "updates.channelDifference"
}

// TypeInfo returns info about TL type.
func (c *UpdatesChannelDifference) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "updates.channelDifference",
		ID:   UpdatesChannelDifferenceTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Final",
			SchemaName: "final",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Pts",
			SchemaName: "pts",
		},
		{
			Name:       "Timeout",
			SchemaName: "timeout",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "NewMessages",
			SchemaName: "new_messages",
		},
		{
			Name:       "OtherUpdates",
			SchemaName: "other_updates",
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
func (c *UpdatesChannelDifference) SetFlags() {
	if !(c.Final == false) {
		c.Flags.Set(0)
	}
	if !(c.Timeout == 0) {
		c.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (c *UpdatesChannelDifference) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode updates.channelDifference#2064674e as nil")
	}
	b.PutID(UpdatesChannelDifferenceTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *UpdatesChannelDifference) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode updates.channelDifference#2064674e as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field flags: %w", err)
	}
	b.PutInt(c.Pts)
	if c.Flags.Has(1) {
		b.PutInt(c.Timeout)
	}
	b.PutVectorHeader(len(c.NewMessages))
	for idx, v := range c.NewMessages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field new_messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field new_messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.OtherUpdates))
	for idx, v := range c.OtherUpdates {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field other_updates element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field other_updates element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *UpdatesChannelDifference) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode updates.channelDifference#2064674e to nil")
	}
	if err := b.ConsumeID(UpdatesChannelDifferenceTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.channelDifference#2064674e: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *UpdatesChannelDifference) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode updates.channelDifference#2064674e to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field flags: %w", err)
		}
	}
	c.Final = c.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field pts: %w", err)
		}
		c.Pts = value
	}
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field timeout: %w", err)
		}
		c.Timeout = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field new_messages: %w", err)
		}

		if headerLen > 0 {
			c.NewMessages = make([]MessageClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field new_messages: %w", err)
			}
			c.NewMessages = append(c.NewMessages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field other_updates: %w", err)
		}

		if headerLen > 0 {
			c.OtherUpdates = make([]UpdateClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUpdate(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field other_updates: %w", err)
			}
			c.OtherUpdates = append(c.OtherUpdates, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field chats: %w", err)
		}

		if headerLen > 0 {
			c.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field users: %w", err)
		}

		if headerLen > 0 {
			c.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// SetFinal sets value of Final conditional field.
func (c *UpdatesChannelDifference) SetFinal(value bool) {
	if value {
		c.Flags.Set(0)
		c.Final = true
	} else {
		c.Flags.Unset(0)
		c.Final = false
	}
}

// GetFinal returns value of Final conditional field.
func (c *UpdatesChannelDifference) GetFinal() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// GetPts returns value of Pts field.
func (c *UpdatesChannelDifference) GetPts() (value int) {
	if c == nil {
		return
	}
	return c.Pts
}

// SetTimeout sets value of Timeout conditional field.
func (c *UpdatesChannelDifference) SetTimeout(value int) {
	c.Flags.Set(1)
	c.Timeout = value
}

// GetTimeout returns value of Timeout conditional field and
// boolean which is true if field was set.
func (c *UpdatesChannelDifference) GetTimeout() (value int, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.Timeout, true
}

// GetNewMessages returns value of NewMessages field.
func (c *UpdatesChannelDifference) GetNewMessages() (value []MessageClass) {
	if c == nil {
		return
	}
	return c.NewMessages
}

// GetOtherUpdates returns value of OtherUpdates field.
func (c *UpdatesChannelDifference) GetOtherUpdates() (value []UpdateClass) {
	if c == nil {
		return
	}
	return c.OtherUpdates
}

// GetChats returns value of Chats field.
func (c *UpdatesChannelDifference) GetChats() (value []ChatClass) {
	if c == nil {
		return
	}
	return c.Chats
}

// GetUsers returns value of Users field.
func (c *UpdatesChannelDifference) GetUsers() (value []UserClass) {
	if c == nil {
		return
	}
	return c.Users
}

// UpdatesChannelDifferenceClassName is schema name of UpdatesChannelDifferenceClass.
const UpdatesChannelDifferenceClassName = "updates.ChannelDifference"

// UpdatesChannelDifferenceClass represents updates.ChannelDifference generic type.
//
// Constructors:
//   - [UpdatesChannelDifferenceEmpty]
//   - [UpdatesChannelDifferenceTooLong]
//   - [UpdatesChannelDifference]
//
// Example:
//
//	g, err := tg.DecodeUpdatesChannelDifference(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.UpdatesChannelDifferenceEmpty: // updates.channelDifferenceEmpty#3e11affb
//	case *tg.UpdatesChannelDifferenceTooLong: // updates.channelDifferenceTooLong#a4bcc6fe
//	case *tg.UpdatesChannelDifference: // updates.channelDifference#2064674e
//	default: panic(v)
//	}
type UpdatesChannelDifferenceClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() UpdatesChannelDifferenceClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Final field of UpdatesChannelDifferenceEmpty.
	GetFinal() (value bool)
	// Timeout field of UpdatesChannelDifferenceEmpty.
	GetTimeout() (value int, ok bool)
}

// DecodeUpdatesChannelDifference implements binary de-serialization for UpdatesChannelDifferenceClass.
func DecodeUpdatesChannelDifference(buf *bin.Buffer) (UpdatesChannelDifferenceClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UpdatesChannelDifferenceEmptyTypeID:
		// Decoding updates.channelDifferenceEmpty#3e11affb.
		v := UpdatesChannelDifferenceEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesChannelDifferenceClass: %w", err)
		}
		return &v, nil
	case UpdatesChannelDifferenceTooLongTypeID:
		// Decoding updates.channelDifferenceTooLong#a4bcc6fe.
		v := UpdatesChannelDifferenceTooLong{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesChannelDifferenceClass: %w", err)
		}
		return &v, nil
	case UpdatesChannelDifferenceTypeID:
		// Decoding updates.channelDifference#2064674e.
		v := UpdatesChannelDifference{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesChannelDifferenceClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UpdatesChannelDifferenceClass: %w", bin.NewUnexpectedID(id))
	}
}

// UpdatesChannelDifference boxes the UpdatesChannelDifferenceClass providing a helper.
type UpdatesChannelDifferenceBox struct {
	ChannelDifference UpdatesChannelDifferenceClass
}

// Decode implements bin.Decoder for UpdatesChannelDifferenceBox.
func (b *UpdatesChannelDifferenceBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UpdatesChannelDifferenceBox to nil")
	}
	v, err := DecodeUpdatesChannelDifference(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChannelDifference = v
	return nil
}

// Encode implements bin.Encode for UpdatesChannelDifferenceBox.
func (b *UpdatesChannelDifferenceBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChannelDifference == nil {
		return fmt.Errorf("unable to encode UpdatesChannelDifferenceClass as nil")
	}
	return b.ChannelDifference.Encode(buf)
}
