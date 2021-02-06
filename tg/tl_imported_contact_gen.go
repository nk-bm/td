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

// ImportedContact represents TL type `importedContact#d0028438`.
// Successfully imported contact.
//
// See https://core.telegram.org/constructor/importedContact for reference.
type ImportedContact struct {
	// User identifier
	UserID int
	// The contact's client identifier (passed to one of the InputContact¹ constructors)
	//
	// Links:
	//  1) https://core.telegram.org/type/InputContact
	ClientID int64
}

// ImportedContactTypeID is TL type id of ImportedContact.
const ImportedContactTypeID = 0xd0028438

func (i *ImportedContact) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.UserID == 0) {
		return false
	}
	if !(i.ClientID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *ImportedContact) String() string {
	if i == nil {
		return "ImportedContact(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ImportedContact")
	sb.WriteString("{\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(i.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tClientID: ")
	sb.WriteString(fmt.Sprint(i.ClientID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *ImportedContact) TypeID() uint32 {
	return ImportedContactTypeID
}

// Encode implements bin.Encoder.
func (i *ImportedContact) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode importedContact#d0028438 as nil")
	}
	b.PutID(ImportedContactTypeID)
	b.PutInt(i.UserID)
	b.PutLong(i.ClientID)
	return nil
}

// GetUserID returns value of UserID field.
func (i *ImportedContact) GetUserID() (value int) {
	return i.UserID
}

// GetClientID returns value of ClientID field.
func (i *ImportedContact) GetClientID() (value int64) {
	return i.ClientID
}

// Decode implements bin.Decoder.
func (i *ImportedContact) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode importedContact#d0028438 to nil")
	}
	if err := b.ConsumeID(ImportedContactTypeID); err != nil {
		return fmt.Errorf("unable to decode importedContact#d0028438: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode importedContact#d0028438: field user_id: %w", err)
		}
		i.UserID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode importedContact#d0028438: field client_id: %w", err)
		}
		i.ClientID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ImportedContact.
var (
	_ bin.Encoder = &ImportedContact{}
	_ bin.Decoder = &ImportedContact{}
)
