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

// AccountInitTakeoutSessionRequest represents TL type `account.initTakeoutSession#8ef3eab0`.
type AccountInitTakeoutSessionRequest struct {
	// Flags field of AccountInitTakeoutSessionRequest.
	Flags bin.Fields
	// Contacts field of AccountInitTakeoutSessionRequest.
	Contacts bool
	// MessageUsers field of AccountInitTakeoutSessionRequest.
	MessageUsers bool
	// MessageChats field of AccountInitTakeoutSessionRequest.
	MessageChats bool
	// MessageMegagroups field of AccountInitTakeoutSessionRequest.
	MessageMegagroups bool
	// MessageChannels field of AccountInitTakeoutSessionRequest.
	MessageChannels bool
	// Files field of AccountInitTakeoutSessionRequest.
	Files bool
	// FileMaxSize field of AccountInitTakeoutSessionRequest.
	//
	// Use SetFileMaxSize and GetFileMaxSize helpers.
	FileMaxSize int64
}

// AccountInitTakeoutSessionRequestTypeID is TL type id of AccountInitTakeoutSessionRequest.
const AccountInitTakeoutSessionRequestTypeID = 0x8ef3eab0

// Ensuring interfaces in compile-time for AccountInitTakeoutSessionRequest.
var (
	_ bin.Encoder     = &AccountInitTakeoutSessionRequest{}
	_ bin.Decoder     = &AccountInitTakeoutSessionRequest{}
	_ bin.BareEncoder = &AccountInitTakeoutSessionRequest{}
	_ bin.BareDecoder = &AccountInitTakeoutSessionRequest{}
)

func (i *AccountInitTakeoutSessionRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Contacts == false) {
		return false
	}
	if !(i.MessageUsers == false) {
		return false
	}
	if !(i.MessageChats == false) {
		return false
	}
	if !(i.MessageMegagroups == false) {
		return false
	}
	if !(i.MessageChannels == false) {
		return false
	}
	if !(i.Files == false) {
		return false
	}
	if !(i.FileMaxSize == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *AccountInitTakeoutSessionRequest) String() string {
	if i == nil {
		return "AccountInitTakeoutSessionRequest(nil)"
	}
	type Alias AccountInitTakeoutSessionRequest
	return fmt.Sprintf("AccountInitTakeoutSessionRequest%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountInitTakeoutSessionRequest) TypeID() uint32 {
	return AccountInitTakeoutSessionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountInitTakeoutSessionRequest) TypeName() string {
	return "account.initTakeoutSession"
}

// TypeInfo returns info about TL type.
func (i *AccountInitTakeoutSessionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.initTakeoutSession",
		ID:   AccountInitTakeoutSessionRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Contacts",
			SchemaName: "contacts",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "MessageUsers",
			SchemaName: "message_users",
			Null:       !i.Flags.Has(1),
		},
		{
			Name:       "MessageChats",
			SchemaName: "message_chats",
			Null:       !i.Flags.Has(2),
		},
		{
			Name:       "MessageMegagroups",
			SchemaName: "message_megagroups",
			Null:       !i.Flags.Has(3),
		},
		{
			Name:       "MessageChannels",
			SchemaName: "message_channels",
			Null:       !i.Flags.Has(4),
		},
		{
			Name:       "Files",
			SchemaName: "files",
			Null:       !i.Flags.Has(5),
		},
		{
			Name:       "FileMaxSize",
			SchemaName: "file_max_size",
			Null:       !i.Flags.Has(5),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (i *AccountInitTakeoutSessionRequest) SetFlags() {
	if !(i.Contacts == false) {
		i.Flags.Set(0)
	}
	if !(i.MessageUsers == false) {
		i.Flags.Set(1)
	}
	if !(i.MessageChats == false) {
		i.Flags.Set(2)
	}
	if !(i.MessageMegagroups == false) {
		i.Flags.Set(3)
	}
	if !(i.MessageChannels == false) {
		i.Flags.Set(4)
	}
	if !(i.Files == false) {
		i.Flags.Set(5)
	}
	if !(i.FileMaxSize == 0) {
		i.Flags.Set(5)
	}
}

// Encode implements bin.Encoder.
func (i *AccountInitTakeoutSessionRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.initTakeoutSession#8ef3eab0 as nil")
	}
	b.PutID(AccountInitTakeoutSessionRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *AccountInitTakeoutSessionRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.initTakeoutSession#8ef3eab0 as nil")
	}
	i.SetFlags()
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.initTakeoutSession#8ef3eab0: field flags: %w", err)
	}
	if i.Flags.Has(5) {
		b.PutLong(i.FileMaxSize)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *AccountInitTakeoutSessionRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.initTakeoutSession#8ef3eab0 to nil")
	}
	if err := b.ConsumeID(AccountInitTakeoutSessionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.initTakeoutSession#8ef3eab0: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *AccountInitTakeoutSessionRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.initTakeoutSession#8ef3eab0 to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.initTakeoutSession#8ef3eab0: field flags: %w", err)
		}
	}
	i.Contacts = i.Flags.Has(0)
	i.MessageUsers = i.Flags.Has(1)
	i.MessageChats = i.Flags.Has(2)
	i.MessageMegagroups = i.Flags.Has(3)
	i.MessageChannels = i.Flags.Has(4)
	i.Files = i.Flags.Has(5)
	if i.Flags.Has(5) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.initTakeoutSession#8ef3eab0: field file_max_size: %w", err)
		}
		i.FileMaxSize = value
	}
	return nil
}

// SetContacts sets value of Contacts conditional field.
func (i *AccountInitTakeoutSessionRequest) SetContacts(value bool) {
	if value {
		i.Flags.Set(0)
		i.Contacts = true
	} else {
		i.Flags.Unset(0)
		i.Contacts = false
	}
}

// GetContacts returns value of Contacts conditional field.
func (i *AccountInitTakeoutSessionRequest) GetContacts() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(0)
}

// SetMessageUsers sets value of MessageUsers conditional field.
func (i *AccountInitTakeoutSessionRequest) SetMessageUsers(value bool) {
	if value {
		i.Flags.Set(1)
		i.MessageUsers = true
	} else {
		i.Flags.Unset(1)
		i.MessageUsers = false
	}
}

// GetMessageUsers returns value of MessageUsers conditional field.
func (i *AccountInitTakeoutSessionRequest) GetMessageUsers() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(1)
}

// SetMessageChats sets value of MessageChats conditional field.
func (i *AccountInitTakeoutSessionRequest) SetMessageChats(value bool) {
	if value {
		i.Flags.Set(2)
		i.MessageChats = true
	} else {
		i.Flags.Unset(2)
		i.MessageChats = false
	}
}

// GetMessageChats returns value of MessageChats conditional field.
func (i *AccountInitTakeoutSessionRequest) GetMessageChats() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(2)
}

// SetMessageMegagroups sets value of MessageMegagroups conditional field.
func (i *AccountInitTakeoutSessionRequest) SetMessageMegagroups(value bool) {
	if value {
		i.Flags.Set(3)
		i.MessageMegagroups = true
	} else {
		i.Flags.Unset(3)
		i.MessageMegagroups = false
	}
}

// GetMessageMegagroups returns value of MessageMegagroups conditional field.
func (i *AccountInitTakeoutSessionRequest) GetMessageMegagroups() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(3)
}

// SetMessageChannels sets value of MessageChannels conditional field.
func (i *AccountInitTakeoutSessionRequest) SetMessageChannels(value bool) {
	if value {
		i.Flags.Set(4)
		i.MessageChannels = true
	} else {
		i.Flags.Unset(4)
		i.MessageChannels = false
	}
}

// GetMessageChannels returns value of MessageChannels conditional field.
func (i *AccountInitTakeoutSessionRequest) GetMessageChannels() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(4)
}

// SetFiles sets value of Files conditional field.
func (i *AccountInitTakeoutSessionRequest) SetFiles(value bool) {
	if value {
		i.Flags.Set(5)
		i.Files = true
	} else {
		i.Flags.Unset(5)
		i.Files = false
	}
}

// GetFiles returns value of Files conditional field.
func (i *AccountInitTakeoutSessionRequest) GetFiles() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(5)
}

// SetFileMaxSize sets value of FileMaxSize conditional field.
func (i *AccountInitTakeoutSessionRequest) SetFileMaxSize(value int64) {
	i.Flags.Set(5)
	i.FileMaxSize = value
}

// GetFileMaxSize returns value of FileMaxSize conditional field and
// boolean which is true if field was set.
func (i *AccountInitTakeoutSessionRequest) GetFileMaxSize() (value int64, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(5) {
		return value, false
	}
	return i.FileMaxSize, true
}

// AccountInitTakeoutSession invokes method account.initTakeoutSession#8ef3eab0 returning error if any.
func (c *Client) AccountInitTakeoutSession(ctx context.Context, request *AccountInitTakeoutSessionRequest) (*AccountTakeout, error) {
	var result AccountTakeout

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
