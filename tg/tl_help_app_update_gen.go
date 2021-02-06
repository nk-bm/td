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

// HelpAppUpdate represents TL type `help.appUpdate#1da7158f`.
// An update is available for the application.
//
// See https://core.telegram.org/constructor/help.appUpdate for reference.
type HelpAppUpdate struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Unskippable, the new info must be shown to the user (with a popup or something else)
	CanNotSkip bool
	// Update ID
	ID int
	// New version name
	Version string
	// Text description of the update
	Text string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	Entities []MessageEntityClass
	// Application binary
	//
	// Use SetDocument and GetDocument helpers.
	Document DocumentClass
	// Application download URL
	//
	// Use SetURL and GetURL helpers.
	URL string
}

// HelpAppUpdateTypeID is TL type id of HelpAppUpdate.
const HelpAppUpdateTypeID = 0x1da7158f

func (a *HelpAppUpdate) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.CanNotSkip == false) {
		return false
	}
	if !(a.ID == 0) {
		return false
	}
	if !(a.Version == "") {
		return false
	}
	if !(a.Text == "") {
		return false
	}
	if !(a.Entities == nil) {
		return false
	}
	if !(a.Document == nil) {
		return false
	}
	if !(a.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *HelpAppUpdate) String() string {
	if a == nil {
		return "HelpAppUpdate(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpAppUpdate")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(a.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(a.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tVersion: ")
	sb.WriteString(fmt.Sprint(a.Version))
	sb.WriteString(",\n")
	sb.WriteString("\tText: ")
	sb.WriteString(fmt.Sprint(a.Text))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range a.Entities {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	if a.Flags.Has(1) {
		sb.WriteString("\tDocument: ")
		sb.WriteString(fmt.Sprint(a.Document))
		sb.WriteString(",\n")
	}
	if a.Flags.Has(2) {
		sb.WriteString("\tURL: ")
		sb.WriteString(fmt.Sprint(a.URL))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (a *HelpAppUpdate) TypeID() uint32 {
	return HelpAppUpdateTypeID
}

// Encode implements bin.Encoder.
func (a *HelpAppUpdate) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode help.appUpdate#1da7158f as nil")
	}
	b.PutID(HelpAppUpdateTypeID)
	if !(a.CanNotSkip == false) {
		a.Flags.Set(0)
	}
	if !(a.Document == nil) {
		a.Flags.Set(1)
	}
	if !(a.URL == "") {
		a.Flags.Set(2)
	}
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.appUpdate#1da7158f: field flags: %w", err)
	}
	b.PutInt(a.ID)
	b.PutString(a.Version)
	b.PutString(a.Text)
	b.PutVectorHeader(len(a.Entities))
	for idx, v := range a.Entities {
		if v == nil {
			return fmt.Errorf("unable to encode help.appUpdate#1da7158f: field entities element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.appUpdate#1da7158f: field entities element with index %d: %w", idx, err)
		}
	}
	if a.Flags.Has(1) {
		if a.Document == nil {
			return fmt.Errorf("unable to encode help.appUpdate#1da7158f: field document is nil")
		}
		if err := a.Document.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.appUpdate#1da7158f: field document: %w", err)
		}
	}
	if a.Flags.Has(2) {
		b.PutString(a.URL)
	}
	return nil
}

// SetCanNotSkip sets value of CanNotSkip conditional field.
func (a *HelpAppUpdate) SetCanNotSkip(value bool) {
	if value {
		a.Flags.Set(0)
		a.CanNotSkip = true
	} else {
		a.Flags.Unset(0)
		a.CanNotSkip = false
	}
}

// GetCanNotSkip returns value of CanNotSkip conditional field.
func (a *HelpAppUpdate) GetCanNotSkip() (value bool) {
	return a.Flags.Has(0)
}

// GetID returns value of ID field.
func (a *HelpAppUpdate) GetID() (value int) {
	return a.ID
}

// GetVersion returns value of Version field.
func (a *HelpAppUpdate) GetVersion() (value string) {
	return a.Version
}

// GetText returns value of Text field.
func (a *HelpAppUpdate) GetText() (value string) {
	return a.Text
}

// GetEntities returns value of Entities field.
func (a *HelpAppUpdate) GetEntities() (value []MessageEntityClass) {
	return a.Entities
}

// SetDocument sets value of Document conditional field.
func (a *HelpAppUpdate) SetDocument(value DocumentClass) {
	a.Flags.Set(1)
	a.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (a *HelpAppUpdate) GetDocument() (value DocumentClass, ok bool) {
	if !a.Flags.Has(1) {
		return value, false
	}
	return a.Document, true
}

// SetURL sets value of URL conditional field.
func (a *HelpAppUpdate) SetURL(value string) {
	a.Flags.Set(2)
	a.URL = value
}

// GetURL returns value of URL conditional field and
// boolean which is true if field was set.
func (a *HelpAppUpdate) GetURL() (value string, ok bool) {
	if !a.Flags.Has(2) {
		return value, false
	}
	return a.URL, true
}

// Decode implements bin.Decoder.
func (a *HelpAppUpdate) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode help.appUpdate#1da7158f to nil")
	}
	if err := b.ConsumeID(HelpAppUpdateTypeID); err != nil {
		return fmt.Errorf("unable to decode help.appUpdate#1da7158f: %w", err)
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field flags: %w", err)
		}
	}
	a.CanNotSkip = a.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field id: %w", err)
		}
		a.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field version: %w", err)
		}
		a.Version = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field text: %w", err)
		}
		a.Text = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field entities: %w", err)
			}
			a.Entities = append(a.Entities, value)
		}
	}
	if a.Flags.Has(1) {
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field document: %w", err)
		}
		a.Document = value
	}
	if a.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field url: %w", err)
		}
		a.URL = value
	}
	return nil
}

// construct implements constructor of HelpAppUpdateClass.
func (a HelpAppUpdate) construct() HelpAppUpdateClass { return &a }

// Ensuring interfaces in compile-time for HelpAppUpdate.
var (
	_ bin.Encoder = &HelpAppUpdate{}
	_ bin.Decoder = &HelpAppUpdate{}

	_ HelpAppUpdateClass = &HelpAppUpdate{}
)

// HelpNoAppUpdate represents TL type `help.noAppUpdate#c45a6536`.
// No updates are available for the application.
//
// See https://core.telegram.org/constructor/help.noAppUpdate for reference.
type HelpNoAppUpdate struct {
}

// HelpNoAppUpdateTypeID is TL type id of HelpNoAppUpdate.
const HelpNoAppUpdateTypeID = 0xc45a6536

func (n *HelpNoAppUpdate) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *HelpNoAppUpdate) String() string {
	if n == nil {
		return "HelpNoAppUpdate(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpNoAppUpdate")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (n *HelpNoAppUpdate) TypeID() uint32 {
	return HelpNoAppUpdateTypeID
}

// Encode implements bin.Encoder.
func (n *HelpNoAppUpdate) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode help.noAppUpdate#c45a6536 as nil")
	}
	b.PutID(HelpNoAppUpdateTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (n *HelpNoAppUpdate) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode help.noAppUpdate#c45a6536 to nil")
	}
	if err := b.ConsumeID(HelpNoAppUpdateTypeID); err != nil {
		return fmt.Errorf("unable to decode help.noAppUpdate#c45a6536: %w", err)
	}
	return nil
}

// construct implements constructor of HelpAppUpdateClass.
func (n HelpNoAppUpdate) construct() HelpAppUpdateClass { return &n }

// Ensuring interfaces in compile-time for HelpNoAppUpdate.
var (
	_ bin.Encoder = &HelpNoAppUpdate{}
	_ bin.Decoder = &HelpNoAppUpdate{}

	_ HelpAppUpdateClass = &HelpNoAppUpdate{}
)

// HelpAppUpdateClass represents help.AppUpdate generic type.
//
// See https://core.telegram.org/type/help.AppUpdate for reference.
//
// Example:
//  g, err := DecodeHelpAppUpdate(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *HelpAppUpdate: // help.appUpdate#1da7158f
//  case *HelpNoAppUpdate: // help.noAppUpdate#c45a6536
//  default: panic(v)
//  }
type HelpAppUpdateClass interface {
	bin.Encoder
	bin.Decoder
	construct() HelpAppUpdateClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeHelpAppUpdate implements binary de-serialization for HelpAppUpdateClass.
func DecodeHelpAppUpdate(buf *bin.Buffer) (HelpAppUpdateClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpAppUpdateTypeID:
		// Decoding help.appUpdate#1da7158f.
		v := HelpAppUpdate{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpAppUpdateClass: %w", err)
		}
		return &v, nil
	case HelpNoAppUpdateTypeID:
		// Decoding help.noAppUpdate#c45a6536.
		v := HelpNoAppUpdate{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpAppUpdateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpAppUpdateClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpAppUpdate boxes the HelpAppUpdateClass providing a helper.
type HelpAppUpdateBox struct {
	AppUpdate HelpAppUpdateClass
}

// Decode implements bin.Decoder for HelpAppUpdateBox.
func (b *HelpAppUpdateBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpAppUpdateBox to nil")
	}
	v, err := DecodeHelpAppUpdate(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.AppUpdate = v
	return nil
}

// Encode implements bin.Encode for HelpAppUpdateBox.
func (b *HelpAppUpdateBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.AppUpdate == nil {
		return fmt.Errorf("unable to encode HelpAppUpdateClass as nil")
	}
	return b.AppUpdate.Encode(buf)
}
