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

// AccountCreateThemeRequest represents TL type `account.createTheme#8432c21f`.
// Create a theme
//
// See https://core.telegram.org/method/account.createTheme for reference.
type AccountCreateThemeRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Unique theme ID
	Slug string
	// Theme name
	Title string
	// Theme file
	//
	// Use SetDocument and GetDocument helpers.
	Document InputDocumentClass
	// Theme settings
	//
	// Use SetSettings and GetSettings helpers.
	Settings InputThemeSettings
}

// AccountCreateThemeRequestTypeID is TL type id of AccountCreateThemeRequest.
const AccountCreateThemeRequestTypeID = 0x8432c21f

func (c *AccountCreateThemeRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Slug == "") {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.Document == nil) {
		return false
	}
	if !(c.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AccountCreateThemeRequest) String() string {
	if c == nil {
		return "AccountCreateThemeRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountCreateThemeRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tSlug: ")
	sb.WriteString(fmt.Sprint(c.Slug))
	sb.WriteString(",\n")
	sb.WriteString("\tTitle: ")
	sb.WriteString(fmt.Sprint(c.Title))
	sb.WriteString(",\n")
	if c.Flags.Has(2) {
		sb.WriteString("\tDocument: ")
		sb.WriteString(fmt.Sprint(c.Document))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(3) {
		sb.WriteString("\tSettings: ")
		sb.WriteString(fmt.Sprint(c.Settings))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *AccountCreateThemeRequest) TypeID() uint32 {
	return AccountCreateThemeRequestTypeID
}

// Encode implements bin.Encoder.
func (c *AccountCreateThemeRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.createTheme#8432c21f as nil")
	}
	b.PutID(AccountCreateThemeRequestTypeID)
	if !(c.Document == nil) {
		c.Flags.Set(2)
	}
	if !(c.Settings.Zero()) {
		c.Flags.Set(3)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.createTheme#8432c21f: field flags: %w", err)
	}
	b.PutString(c.Slug)
	b.PutString(c.Title)
	if c.Flags.Has(2) {
		if c.Document == nil {
			return fmt.Errorf("unable to encode account.createTheme#8432c21f: field document is nil")
		}
		if err := c.Document.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.createTheme#8432c21f: field document: %w", err)
		}
	}
	if c.Flags.Has(3) {
		if err := c.Settings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.createTheme#8432c21f: field settings: %w", err)
		}
	}
	return nil
}

// GetSlug returns value of Slug field.
func (c *AccountCreateThemeRequest) GetSlug() (value string) {
	return c.Slug
}

// GetTitle returns value of Title field.
func (c *AccountCreateThemeRequest) GetTitle() (value string) {
	return c.Title
}

// SetDocument sets value of Document conditional field.
func (c *AccountCreateThemeRequest) SetDocument(value InputDocumentClass) {
	c.Flags.Set(2)
	c.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (c *AccountCreateThemeRequest) GetDocument() (value InputDocumentClass, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.Document, true
}

// SetSettings sets value of Settings conditional field.
func (c *AccountCreateThemeRequest) SetSettings(value InputThemeSettings) {
	c.Flags.Set(3)
	c.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (c *AccountCreateThemeRequest) GetSettings() (value InputThemeSettings, ok bool) {
	if !c.Flags.Has(3) {
		return value, false
	}
	return c.Settings, true
}

// Decode implements bin.Decoder.
func (c *AccountCreateThemeRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.createTheme#8432c21f to nil")
	}
	if err := b.ConsumeID(AccountCreateThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.createTheme#8432c21f: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.createTheme#8432c21f: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.createTheme#8432c21f: field slug: %w", err)
		}
		c.Slug = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.createTheme#8432c21f: field title: %w", err)
		}
		c.Title = value
	}
	if c.Flags.Has(2) {
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.createTheme#8432c21f: field document: %w", err)
		}
		c.Document = value
	}
	if c.Flags.Has(3) {
		if err := c.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.createTheme#8432c21f: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountCreateThemeRequest.
var (
	_ bin.Encoder = &AccountCreateThemeRequest{}
	_ bin.Decoder = &AccountCreateThemeRequest{}
)

// AccountCreateTheme invokes method account.createTheme#8432c21f returning error if any.
// Create a theme
//
// See https://core.telegram.org/method/account.createTheme for reference.
func (c *Client) AccountCreateTheme(ctx context.Context, request *AccountCreateThemeRequest) (*Theme, error) {
	var result Theme

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
