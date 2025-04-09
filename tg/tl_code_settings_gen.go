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

// CodeSettings represents TL type `codeSettings#ad253d78`.
type CodeSettings struct {
	// Flags field of CodeSettings.
	Flags bin.Fields
	// AllowFlashcall field of CodeSettings.
	AllowFlashcall bool
	// CurrentNumber field of CodeSettings.
	CurrentNumber bool
	// AllowAppHash field of CodeSettings.
	AllowAppHash bool
	// AllowMissedCall field of CodeSettings.
	AllowMissedCall bool
	// AllowFirebase field of CodeSettings.
	AllowFirebase bool
	// UnknownNumber field of CodeSettings.
	UnknownNumber bool
	// LogoutTokens field of CodeSettings.
	//
	// Use SetLogoutTokens and GetLogoutTokens helpers.
	LogoutTokens [][]byte
	// Token field of CodeSettings.
	//
	// Use SetToken and GetToken helpers.
	Token string
	// AppSandbox field of CodeSettings.
	//
	// Use SetAppSandbox and GetAppSandbox helpers.
	AppSandbox bool
}

// CodeSettingsTypeID is TL type id of CodeSettings.
const CodeSettingsTypeID = 0xad253d78

// Ensuring interfaces in compile-time for CodeSettings.
var (
	_ bin.Encoder     = &CodeSettings{}
	_ bin.Decoder     = &CodeSettings{}
	_ bin.BareEncoder = &CodeSettings{}
	_ bin.BareDecoder = &CodeSettings{}
)

func (c *CodeSettings) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.AllowFlashcall == false) {
		return false
	}
	if !(c.CurrentNumber == false) {
		return false
	}
	if !(c.AllowAppHash == false) {
		return false
	}
	if !(c.AllowMissedCall == false) {
		return false
	}
	if !(c.AllowFirebase == false) {
		return false
	}
	if !(c.UnknownNumber == false) {
		return false
	}
	if !(c.LogoutTokens == nil) {
		return false
	}
	if !(c.Token == "") {
		return false
	}
	if !(c.AppSandbox == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CodeSettings) String() string {
	if c == nil {
		return "CodeSettings(nil)"
	}
	type Alias CodeSettings
	return fmt.Sprintf("CodeSettings%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CodeSettings) TypeID() uint32 {
	return CodeSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*CodeSettings) TypeName() string {
	return "codeSettings"
}

// TypeInfo returns info about TL type.
func (c *CodeSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "codeSettings",
		ID:   CodeSettingsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AllowFlashcall",
			SchemaName: "allow_flashcall",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "CurrentNumber",
			SchemaName: "current_number",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "AllowAppHash",
			SchemaName: "allow_app_hash",
			Null:       !c.Flags.Has(4),
		},
		{
			Name:       "AllowMissedCall",
			SchemaName: "allow_missed_call",
			Null:       !c.Flags.Has(5),
		},
		{
			Name:       "AllowFirebase",
			SchemaName: "allow_firebase",
			Null:       !c.Flags.Has(7),
		},
		{
			Name:       "UnknownNumber",
			SchemaName: "unknown_number",
			Null:       !c.Flags.Has(9),
		},
		{
			Name:       "LogoutTokens",
			SchemaName: "logout_tokens",
			Null:       !c.Flags.Has(6),
		},
		{
			Name:       "Token",
			SchemaName: "token",
			Null:       !c.Flags.Has(8),
		},
		{
			Name:       "AppSandbox",
			SchemaName: "app_sandbox",
			Null:       !c.Flags.Has(8),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *CodeSettings) SetFlags() {
	if !(c.AllowFlashcall == false) {
		c.Flags.Set(0)
	}
	if !(c.CurrentNumber == false) {
		c.Flags.Set(1)
	}
	if !(c.AllowAppHash == false) {
		c.Flags.Set(4)
	}
	if !(c.AllowMissedCall == false) {
		c.Flags.Set(5)
	}
	if !(c.AllowFirebase == false) {
		c.Flags.Set(7)
	}
	if !(c.UnknownNumber == false) {
		c.Flags.Set(9)
	}
	if !(c.LogoutTokens == nil) {
		c.Flags.Set(6)
	}
	if !(c.Token == "") {
		c.Flags.Set(8)
	}
	if !(c.AppSandbox == false) {
		c.Flags.Set(8)
	}
}

// Encode implements bin.Encoder.
func (c *CodeSettings) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode codeSettings#ad253d78 as nil")
	}
	b.PutID(CodeSettingsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CodeSettings) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode codeSettings#ad253d78 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode codeSettings#ad253d78: field flags: %w", err)
	}
	if c.Flags.Has(6) {
		b.PutVectorHeader(len(c.LogoutTokens))
		for _, v := range c.LogoutTokens {
			b.PutBytes(v)
		}
	}
	if c.Flags.Has(8) {
		b.PutString(c.Token)
	}
	if c.Flags.Has(8) {
		b.PutBool(c.AppSandbox)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CodeSettings) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode codeSettings#ad253d78 to nil")
	}
	if err := b.ConsumeID(CodeSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode codeSettings#ad253d78: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CodeSettings) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode codeSettings#ad253d78 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode codeSettings#ad253d78: field flags: %w", err)
		}
	}
	c.AllowFlashcall = c.Flags.Has(0)
	c.CurrentNumber = c.Flags.Has(1)
	c.AllowAppHash = c.Flags.Has(4)
	c.AllowMissedCall = c.Flags.Has(5)
	c.AllowFirebase = c.Flags.Has(7)
	c.UnknownNumber = c.Flags.Has(9)
	if c.Flags.Has(6) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode codeSettings#ad253d78: field logout_tokens: %w", err)
		}

		if headerLen > 0 {
			c.LogoutTokens = make([][]byte, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode codeSettings#ad253d78: field logout_tokens: %w", err)
			}
			c.LogoutTokens = append(c.LogoutTokens, value)
		}
	}
	if c.Flags.Has(8) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode codeSettings#ad253d78: field token: %w", err)
		}
		c.Token = value
	}
	if c.Flags.Has(8) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode codeSettings#ad253d78: field app_sandbox: %w", err)
		}
		c.AppSandbox = value
	}
	return nil
}

// SetAllowFlashcall sets value of AllowFlashcall conditional field.
func (c *CodeSettings) SetAllowFlashcall(value bool) {
	if value {
		c.Flags.Set(0)
		c.AllowFlashcall = true
	} else {
		c.Flags.Unset(0)
		c.AllowFlashcall = false
	}
}

// GetAllowFlashcall returns value of AllowFlashcall conditional field.
func (c *CodeSettings) GetAllowFlashcall() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// SetCurrentNumber sets value of CurrentNumber conditional field.
func (c *CodeSettings) SetCurrentNumber(value bool) {
	if value {
		c.Flags.Set(1)
		c.CurrentNumber = true
	} else {
		c.Flags.Unset(1)
		c.CurrentNumber = false
	}
}

// GetCurrentNumber returns value of CurrentNumber conditional field.
func (c *CodeSettings) GetCurrentNumber() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(1)
}

// SetAllowAppHash sets value of AllowAppHash conditional field.
func (c *CodeSettings) SetAllowAppHash(value bool) {
	if value {
		c.Flags.Set(4)
		c.AllowAppHash = true
	} else {
		c.Flags.Unset(4)
		c.AllowAppHash = false
	}
}

// GetAllowAppHash returns value of AllowAppHash conditional field.
func (c *CodeSettings) GetAllowAppHash() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(4)
}

// SetAllowMissedCall sets value of AllowMissedCall conditional field.
func (c *CodeSettings) SetAllowMissedCall(value bool) {
	if value {
		c.Flags.Set(5)
		c.AllowMissedCall = true
	} else {
		c.Flags.Unset(5)
		c.AllowMissedCall = false
	}
}

// GetAllowMissedCall returns value of AllowMissedCall conditional field.
func (c *CodeSettings) GetAllowMissedCall() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(5)
}

// SetAllowFirebase sets value of AllowFirebase conditional field.
func (c *CodeSettings) SetAllowFirebase(value bool) {
	if value {
		c.Flags.Set(7)
		c.AllowFirebase = true
	} else {
		c.Flags.Unset(7)
		c.AllowFirebase = false
	}
}

// GetAllowFirebase returns value of AllowFirebase conditional field.
func (c *CodeSettings) GetAllowFirebase() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(7)
}

// SetUnknownNumber sets value of UnknownNumber conditional field.
func (c *CodeSettings) SetUnknownNumber(value bool) {
	if value {
		c.Flags.Set(9)
		c.UnknownNumber = true
	} else {
		c.Flags.Unset(9)
		c.UnknownNumber = false
	}
}

// GetUnknownNumber returns value of UnknownNumber conditional field.
func (c *CodeSettings) GetUnknownNumber() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(9)
}

// SetLogoutTokens sets value of LogoutTokens conditional field.
func (c *CodeSettings) SetLogoutTokens(value [][]byte) {
	c.Flags.Set(6)
	c.LogoutTokens = value
}

// GetLogoutTokens returns value of LogoutTokens conditional field and
// boolean which is true if field was set.
func (c *CodeSettings) GetLogoutTokens() (value [][]byte, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(6) {
		return value, false
	}
	return c.LogoutTokens, true
}

// SetToken sets value of Token conditional field.
func (c *CodeSettings) SetToken(value string) {
	c.Flags.Set(8)
	c.Token = value
}

// GetToken returns value of Token conditional field and
// boolean which is true if field was set.
func (c *CodeSettings) GetToken() (value string, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(8) {
		return value, false
	}
	return c.Token, true
}

// SetAppSandbox sets value of AppSandbox conditional field.
func (c *CodeSettings) SetAppSandbox(value bool) {
	c.Flags.Set(8)
	c.AppSandbox = value
}

// GetAppSandbox returns value of AppSandbox conditional field and
// boolean which is true if field was set.
func (c *CodeSettings) GetAppSandbox() (value bool, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(8) {
		return value, false
	}
	return c.AppSandbox, true
}
