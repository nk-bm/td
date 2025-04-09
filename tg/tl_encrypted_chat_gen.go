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

// EncryptedChatEmpty represents TL type `encryptedChatEmpty#ab7ec0a0`.
type EncryptedChatEmpty struct {
	// ID field of EncryptedChatEmpty.
	ID int
}

// EncryptedChatEmptyTypeID is TL type id of EncryptedChatEmpty.
const EncryptedChatEmptyTypeID = 0xab7ec0a0

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChatEmpty) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChatEmpty.
var (
	_ bin.Encoder     = &EncryptedChatEmpty{}
	_ bin.Decoder     = &EncryptedChatEmpty{}
	_ bin.BareEncoder = &EncryptedChatEmpty{}
	_ bin.BareDecoder = &EncryptedChatEmpty{}

	_ EncryptedChatClass = &EncryptedChatEmpty{}
)

func (e *EncryptedChatEmpty) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedChatEmpty) String() string {
	if e == nil {
		return "EncryptedChatEmpty(nil)"
	}
	type Alias EncryptedChatEmpty
	return fmt.Sprintf("EncryptedChatEmpty%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedChatEmpty) TypeID() uint32 {
	return EncryptedChatEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedChatEmpty) TypeName() string {
	return "encryptedChatEmpty"
}

// TypeInfo returns info about TL type.
func (e *EncryptedChatEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedChatEmpty",
		ID:   EncryptedChatEmptyTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EncryptedChatEmpty) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatEmpty#ab7ec0a0 as nil")
	}
	b.PutID(EncryptedChatEmptyTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EncryptedChatEmpty) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatEmpty#ab7ec0a0 as nil")
	}
	b.PutInt(e.ID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedChatEmpty) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatEmpty#ab7ec0a0 to nil")
	}
	if err := b.ConsumeID(EncryptedChatEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChatEmpty#ab7ec0a0: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EncryptedChatEmpty) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatEmpty#ab7ec0a0 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatEmpty#ab7ec0a0: field id: %w", err)
		}
		e.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (e *EncryptedChatEmpty) GetID() (value int) {
	if e == nil {
		return
	}
	return e.ID
}

// EncryptedChatWaiting represents TL type `encryptedChatWaiting#66b25953`.
type EncryptedChatWaiting struct {
	// ID field of EncryptedChatWaiting.
	ID int
	// AccessHash field of EncryptedChatWaiting.
	AccessHash int64
	// Date field of EncryptedChatWaiting.
	Date int
	// AdminID field of EncryptedChatWaiting.
	AdminID int64
	// ParticipantID field of EncryptedChatWaiting.
	ParticipantID int64
}

// EncryptedChatWaitingTypeID is TL type id of EncryptedChatWaiting.
const EncryptedChatWaitingTypeID = 0x66b25953

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChatWaiting) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChatWaiting.
var (
	_ bin.Encoder     = &EncryptedChatWaiting{}
	_ bin.Decoder     = &EncryptedChatWaiting{}
	_ bin.BareEncoder = &EncryptedChatWaiting{}
	_ bin.BareDecoder = &EncryptedChatWaiting{}

	_ EncryptedChatClass = &EncryptedChatWaiting{}
)

func (e *EncryptedChatWaiting) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ID == 0) {
		return false
	}
	if !(e.AccessHash == 0) {
		return false
	}
	if !(e.Date == 0) {
		return false
	}
	if !(e.AdminID == 0) {
		return false
	}
	if !(e.ParticipantID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedChatWaiting) String() string {
	if e == nil {
		return "EncryptedChatWaiting(nil)"
	}
	type Alias EncryptedChatWaiting
	return fmt.Sprintf("EncryptedChatWaiting%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedChatWaiting) TypeID() uint32 {
	return EncryptedChatWaitingTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedChatWaiting) TypeName() string {
	return "encryptedChatWaiting"
}

// TypeInfo returns info about TL type.
func (e *EncryptedChatWaiting) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedChatWaiting",
		ID:   EncryptedChatWaitingTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "AdminID",
			SchemaName: "admin_id",
		},
		{
			Name:       "ParticipantID",
			SchemaName: "participant_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EncryptedChatWaiting) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatWaiting#66b25953 as nil")
	}
	b.PutID(EncryptedChatWaitingTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EncryptedChatWaiting) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatWaiting#66b25953 as nil")
	}
	b.PutInt(e.ID)
	b.PutLong(e.AccessHash)
	b.PutInt(e.Date)
	b.PutLong(e.AdminID)
	b.PutLong(e.ParticipantID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedChatWaiting) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatWaiting#66b25953 to nil")
	}
	if err := b.ConsumeID(EncryptedChatWaitingTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChatWaiting#66b25953: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EncryptedChatWaiting) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatWaiting#66b25953 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#66b25953: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#66b25953: field access_hash: %w", err)
		}
		e.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#66b25953: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#66b25953: field admin_id: %w", err)
		}
		e.AdminID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#66b25953: field participant_id: %w", err)
		}
		e.ParticipantID = value
	}
	return nil
}

// GetID returns value of ID field.
func (e *EncryptedChatWaiting) GetID() (value int) {
	if e == nil {
		return
	}
	return e.ID
}

// GetAccessHash returns value of AccessHash field.
func (e *EncryptedChatWaiting) GetAccessHash() (value int64) {
	if e == nil {
		return
	}
	return e.AccessHash
}

// GetDate returns value of Date field.
func (e *EncryptedChatWaiting) GetDate() (value int) {
	if e == nil {
		return
	}
	return e.Date
}

// GetAdminID returns value of AdminID field.
func (e *EncryptedChatWaiting) GetAdminID() (value int64) {
	if e == nil {
		return
	}
	return e.AdminID
}

// GetParticipantID returns value of ParticipantID field.
func (e *EncryptedChatWaiting) GetParticipantID() (value int64) {
	if e == nil {
		return
	}
	return e.ParticipantID
}

// EncryptedChatRequested represents TL type `encryptedChatRequested#48f1d94c`.
type EncryptedChatRequested struct {
	// Flags field of EncryptedChatRequested.
	Flags bin.Fields
	// FolderID field of EncryptedChatRequested.
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// ID field of EncryptedChatRequested.
	ID int
	// AccessHash field of EncryptedChatRequested.
	AccessHash int64
	// Date field of EncryptedChatRequested.
	Date int
	// AdminID field of EncryptedChatRequested.
	AdminID int64
	// ParticipantID field of EncryptedChatRequested.
	ParticipantID int64
	// GA field of EncryptedChatRequested.
	GA []byte
}

// EncryptedChatRequestedTypeID is TL type id of EncryptedChatRequested.
const EncryptedChatRequestedTypeID = 0x48f1d94c

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChatRequested) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChatRequested.
var (
	_ bin.Encoder     = &EncryptedChatRequested{}
	_ bin.Decoder     = &EncryptedChatRequested{}
	_ bin.BareEncoder = &EncryptedChatRequested{}
	_ bin.BareDecoder = &EncryptedChatRequested{}

	_ EncryptedChatClass = &EncryptedChatRequested{}
)

func (e *EncryptedChatRequested) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.FolderID == 0) {
		return false
	}
	if !(e.ID == 0) {
		return false
	}
	if !(e.AccessHash == 0) {
		return false
	}
	if !(e.Date == 0) {
		return false
	}
	if !(e.AdminID == 0) {
		return false
	}
	if !(e.ParticipantID == 0) {
		return false
	}
	if !(e.GA == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedChatRequested) String() string {
	if e == nil {
		return "EncryptedChatRequested(nil)"
	}
	type Alias EncryptedChatRequested
	return fmt.Sprintf("EncryptedChatRequested%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedChatRequested) TypeID() uint32 {
	return EncryptedChatRequestedTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedChatRequested) TypeName() string {
	return "encryptedChatRequested"
}

// TypeInfo returns info about TL type.
func (e *EncryptedChatRequested) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedChatRequested",
		ID:   EncryptedChatRequestedTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FolderID",
			SchemaName: "folder_id",
			Null:       !e.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "AdminID",
			SchemaName: "admin_id",
		},
		{
			Name:       "ParticipantID",
			SchemaName: "participant_id",
		},
		{
			Name:       "GA",
			SchemaName: "g_a",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (e *EncryptedChatRequested) SetFlags() {
	if !(e.FolderID == 0) {
		e.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (e *EncryptedChatRequested) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatRequested#48f1d94c as nil")
	}
	b.PutID(EncryptedChatRequestedTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EncryptedChatRequested) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatRequested#48f1d94c as nil")
	}
	e.SetFlags()
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode encryptedChatRequested#48f1d94c: field flags: %w", err)
	}
	if e.Flags.Has(0) {
		b.PutInt(e.FolderID)
	}
	b.PutInt(e.ID)
	b.PutLong(e.AccessHash)
	b.PutInt(e.Date)
	b.PutLong(e.AdminID)
	b.PutLong(e.ParticipantID)
	b.PutBytes(e.GA)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedChatRequested) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatRequested#48f1d94c to nil")
	}
	if err := b.ConsumeID(EncryptedChatRequestedTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChatRequested#48f1d94c: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EncryptedChatRequested) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatRequested#48f1d94c to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#48f1d94c: field flags: %w", err)
		}
	}
	if e.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#48f1d94c: field folder_id: %w", err)
		}
		e.FolderID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#48f1d94c: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#48f1d94c: field access_hash: %w", err)
		}
		e.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#48f1d94c: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#48f1d94c: field admin_id: %w", err)
		}
		e.AdminID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#48f1d94c: field participant_id: %w", err)
		}
		e.ParticipantID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#48f1d94c: field g_a: %w", err)
		}
		e.GA = value
	}
	return nil
}

// SetFolderID sets value of FolderID conditional field.
func (e *EncryptedChatRequested) SetFolderID(value int) {
	e.Flags.Set(0)
	e.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (e *EncryptedChatRequested) GetFolderID() (value int, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(0) {
		return value, false
	}
	return e.FolderID, true
}

// GetID returns value of ID field.
func (e *EncryptedChatRequested) GetID() (value int) {
	if e == nil {
		return
	}
	return e.ID
}

// GetAccessHash returns value of AccessHash field.
func (e *EncryptedChatRequested) GetAccessHash() (value int64) {
	if e == nil {
		return
	}
	return e.AccessHash
}

// GetDate returns value of Date field.
func (e *EncryptedChatRequested) GetDate() (value int) {
	if e == nil {
		return
	}
	return e.Date
}

// GetAdminID returns value of AdminID field.
func (e *EncryptedChatRequested) GetAdminID() (value int64) {
	if e == nil {
		return
	}
	return e.AdminID
}

// GetParticipantID returns value of ParticipantID field.
func (e *EncryptedChatRequested) GetParticipantID() (value int64) {
	if e == nil {
		return
	}
	return e.ParticipantID
}

// GetGA returns value of GA field.
func (e *EncryptedChatRequested) GetGA() (value []byte) {
	if e == nil {
		return
	}
	return e.GA
}

// EncryptedChat represents TL type `encryptedChat#61f0d4c7`.
type EncryptedChat struct {
	// ID field of EncryptedChat.
	ID int
	// AccessHash field of EncryptedChat.
	AccessHash int64
	// Date field of EncryptedChat.
	Date int
	// AdminID field of EncryptedChat.
	AdminID int64
	// ParticipantID field of EncryptedChat.
	ParticipantID int64
	// GAOrB field of EncryptedChat.
	GAOrB []byte
	// KeyFingerprint field of EncryptedChat.
	KeyFingerprint int64
}

// EncryptedChatTypeID is TL type id of EncryptedChat.
const EncryptedChatTypeID = 0x61f0d4c7

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChat) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChat.
var (
	_ bin.Encoder     = &EncryptedChat{}
	_ bin.Decoder     = &EncryptedChat{}
	_ bin.BareEncoder = &EncryptedChat{}
	_ bin.BareDecoder = &EncryptedChat{}

	_ EncryptedChatClass = &EncryptedChat{}
)

func (e *EncryptedChat) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ID == 0) {
		return false
	}
	if !(e.AccessHash == 0) {
		return false
	}
	if !(e.Date == 0) {
		return false
	}
	if !(e.AdminID == 0) {
		return false
	}
	if !(e.ParticipantID == 0) {
		return false
	}
	if !(e.GAOrB == nil) {
		return false
	}
	if !(e.KeyFingerprint == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedChat) String() string {
	if e == nil {
		return "EncryptedChat(nil)"
	}
	type Alias EncryptedChat
	return fmt.Sprintf("EncryptedChat%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedChat) TypeID() uint32 {
	return EncryptedChatTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedChat) TypeName() string {
	return "encryptedChat"
}

// TypeInfo returns info about TL type.
func (e *EncryptedChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedChat",
		ID:   EncryptedChatTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "AdminID",
			SchemaName: "admin_id",
		},
		{
			Name:       "ParticipantID",
			SchemaName: "participant_id",
		},
		{
			Name:       "GAOrB",
			SchemaName: "g_a_or_b",
		},
		{
			Name:       "KeyFingerprint",
			SchemaName: "key_fingerprint",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EncryptedChat) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChat#61f0d4c7 as nil")
	}
	b.PutID(EncryptedChatTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EncryptedChat) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChat#61f0d4c7 as nil")
	}
	b.PutInt(e.ID)
	b.PutLong(e.AccessHash)
	b.PutInt(e.Date)
	b.PutLong(e.AdminID)
	b.PutLong(e.ParticipantID)
	b.PutBytes(e.GAOrB)
	b.PutLong(e.KeyFingerprint)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedChat) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChat#61f0d4c7 to nil")
	}
	if err := b.ConsumeID(EncryptedChatTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChat#61f0d4c7: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EncryptedChat) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChat#61f0d4c7 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#61f0d4c7: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#61f0d4c7: field access_hash: %w", err)
		}
		e.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#61f0d4c7: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#61f0d4c7: field admin_id: %w", err)
		}
		e.AdminID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#61f0d4c7: field participant_id: %w", err)
		}
		e.ParticipantID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#61f0d4c7: field g_a_or_b: %w", err)
		}
		e.GAOrB = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#61f0d4c7: field key_fingerprint: %w", err)
		}
		e.KeyFingerprint = value
	}
	return nil
}

// GetID returns value of ID field.
func (e *EncryptedChat) GetID() (value int) {
	if e == nil {
		return
	}
	return e.ID
}

// GetAccessHash returns value of AccessHash field.
func (e *EncryptedChat) GetAccessHash() (value int64) {
	if e == nil {
		return
	}
	return e.AccessHash
}

// GetDate returns value of Date field.
func (e *EncryptedChat) GetDate() (value int) {
	if e == nil {
		return
	}
	return e.Date
}

// GetAdminID returns value of AdminID field.
func (e *EncryptedChat) GetAdminID() (value int64) {
	if e == nil {
		return
	}
	return e.AdminID
}

// GetParticipantID returns value of ParticipantID field.
func (e *EncryptedChat) GetParticipantID() (value int64) {
	if e == nil {
		return
	}
	return e.ParticipantID
}

// GetGAOrB returns value of GAOrB field.
func (e *EncryptedChat) GetGAOrB() (value []byte) {
	if e == nil {
		return
	}
	return e.GAOrB
}

// GetKeyFingerprint returns value of KeyFingerprint field.
func (e *EncryptedChat) GetKeyFingerprint() (value int64) {
	if e == nil {
		return
	}
	return e.KeyFingerprint
}

// EncryptedChatDiscarded represents TL type `encryptedChatDiscarded#1e1c7c45`.
type EncryptedChatDiscarded struct {
	// Flags field of EncryptedChatDiscarded.
	Flags bin.Fields
	// HistoryDeleted field of EncryptedChatDiscarded.
	HistoryDeleted bool
	// ID field of EncryptedChatDiscarded.
	ID int
}

// EncryptedChatDiscardedTypeID is TL type id of EncryptedChatDiscarded.
const EncryptedChatDiscardedTypeID = 0x1e1c7c45

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChatDiscarded) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChatDiscarded.
var (
	_ bin.Encoder     = &EncryptedChatDiscarded{}
	_ bin.Decoder     = &EncryptedChatDiscarded{}
	_ bin.BareEncoder = &EncryptedChatDiscarded{}
	_ bin.BareDecoder = &EncryptedChatDiscarded{}

	_ EncryptedChatClass = &EncryptedChatDiscarded{}
)

func (e *EncryptedChatDiscarded) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.HistoryDeleted == false) {
		return false
	}
	if !(e.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedChatDiscarded) String() string {
	if e == nil {
		return "EncryptedChatDiscarded(nil)"
	}
	type Alias EncryptedChatDiscarded
	return fmt.Sprintf("EncryptedChatDiscarded%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedChatDiscarded) TypeID() uint32 {
	return EncryptedChatDiscardedTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedChatDiscarded) TypeName() string {
	return "encryptedChatDiscarded"
}

// TypeInfo returns info about TL type.
func (e *EncryptedChatDiscarded) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedChatDiscarded",
		ID:   EncryptedChatDiscardedTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "HistoryDeleted",
			SchemaName: "history_deleted",
			Null:       !e.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (e *EncryptedChatDiscarded) SetFlags() {
	if !(e.HistoryDeleted == false) {
		e.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (e *EncryptedChatDiscarded) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatDiscarded#1e1c7c45 as nil")
	}
	b.PutID(EncryptedChatDiscardedTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EncryptedChatDiscarded) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatDiscarded#1e1c7c45 as nil")
	}
	e.SetFlags()
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode encryptedChatDiscarded#1e1c7c45: field flags: %w", err)
	}
	b.PutInt(e.ID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedChatDiscarded) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatDiscarded#1e1c7c45 to nil")
	}
	if err := b.ConsumeID(EncryptedChatDiscardedTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChatDiscarded#1e1c7c45: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EncryptedChatDiscarded) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatDiscarded#1e1c7c45 to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode encryptedChatDiscarded#1e1c7c45: field flags: %w", err)
		}
	}
	e.HistoryDeleted = e.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatDiscarded#1e1c7c45: field id: %w", err)
		}
		e.ID = value
	}
	return nil
}

// SetHistoryDeleted sets value of HistoryDeleted conditional field.
func (e *EncryptedChatDiscarded) SetHistoryDeleted(value bool) {
	if value {
		e.Flags.Set(0)
		e.HistoryDeleted = true
	} else {
		e.Flags.Unset(0)
		e.HistoryDeleted = false
	}
}

// GetHistoryDeleted returns value of HistoryDeleted conditional field.
func (e *EncryptedChatDiscarded) GetHistoryDeleted() (value bool) {
	if e == nil {
		return
	}
	return e.Flags.Has(0)
}

// GetID returns value of ID field.
func (e *EncryptedChatDiscarded) GetID() (value int) {
	if e == nil {
		return
	}
	return e.ID
}

// EncryptedChatClassName is schema name of EncryptedChatClass.
const EncryptedChatClassName = "EncryptedChat"

// EncryptedChatClass represents EncryptedChat generic type.
//
// Constructors:
//   - [EncryptedChatEmpty]
//   - [EncryptedChatWaiting]
//   - [EncryptedChatRequested]
//   - [EncryptedChat]
//   - [EncryptedChatDiscarded]
//
// Example:
//
//	g, err := tg.DecodeEncryptedChat(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.EncryptedChatEmpty: // encryptedChatEmpty#ab7ec0a0
//	case *tg.EncryptedChatWaiting: // encryptedChatWaiting#66b25953
//	case *tg.EncryptedChatRequested: // encryptedChatRequested#48f1d94c
//	case *tg.EncryptedChat: // encryptedChat#61f0d4c7
//	case *tg.EncryptedChatDiscarded: // encryptedChatDiscarded#1e1c7c45
//	default: panic(v)
//	}
type EncryptedChatClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() EncryptedChatClass

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

	// ID field of EncryptedChatEmpty.
	GetID() (value int)
}

// DecodeEncryptedChat implements binary de-serialization for EncryptedChatClass.
func DecodeEncryptedChat(buf *bin.Buffer) (EncryptedChatClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EncryptedChatEmptyTypeID:
		// Decoding encryptedChatEmpty#ab7ec0a0.
		v := EncryptedChatEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	case EncryptedChatWaitingTypeID:
		// Decoding encryptedChatWaiting#66b25953.
		v := EncryptedChatWaiting{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	case EncryptedChatRequestedTypeID:
		// Decoding encryptedChatRequested#48f1d94c.
		v := EncryptedChatRequested{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	case EncryptedChatTypeID:
		// Decoding encryptedChat#61f0d4c7.
		v := EncryptedChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	case EncryptedChatDiscardedTypeID:
		// Decoding encryptedChatDiscarded#1e1c7c45.
		v := EncryptedChatDiscarded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", bin.NewUnexpectedID(id))
	}
}

// EncryptedChat boxes the EncryptedChatClass providing a helper.
type EncryptedChatBox struct {
	EncryptedChat EncryptedChatClass
}

// Decode implements bin.Decoder for EncryptedChatBox.
func (b *EncryptedChatBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode EncryptedChatBox to nil")
	}
	v, err := DecodeEncryptedChat(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EncryptedChat = v
	return nil
}

// Encode implements bin.Encode for EncryptedChatBox.
func (b *EncryptedChatBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EncryptedChat == nil {
		return fmt.Errorf("unable to encode EncryptedChatClass as nil")
	}
	return b.EncryptedChat.Encode(buf)
}
