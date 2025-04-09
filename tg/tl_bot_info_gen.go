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

// BotInfo represents TL type `botInfo#36607333`.
type BotInfo struct {
	// Flags field of BotInfo.
	Flags bin.Fields
	// HasPreviewMedias field of BotInfo.
	HasPreviewMedias bool
	// UserID field of BotInfo.
	//
	// Use SetUserID and GetUserID helpers.
	UserID int64
	// Description field of BotInfo.
	//
	// Use SetDescription and GetDescription helpers.
	Description string
	// DescriptionPhoto field of BotInfo.
	//
	// Use SetDescriptionPhoto and GetDescriptionPhoto helpers.
	DescriptionPhoto PhotoClass
	// DescriptionDocument field of BotInfo.
	//
	// Use SetDescriptionDocument and GetDescriptionDocument helpers.
	DescriptionDocument DocumentClass
	// Commands field of BotInfo.
	//
	// Use SetCommands and GetCommands helpers.
	Commands []BotCommand
	// MenuButton field of BotInfo.
	//
	// Use SetMenuButton and GetMenuButton helpers.
	MenuButton BotMenuButtonClass
	// PrivacyPolicyURL field of BotInfo.
	//
	// Use SetPrivacyPolicyURL and GetPrivacyPolicyURL helpers.
	PrivacyPolicyURL string
	// AppSettings field of BotInfo.
	//
	// Use SetAppSettings and GetAppSettings helpers.
	AppSettings BotAppSettings
}

// BotInfoTypeID is TL type id of BotInfo.
const BotInfoTypeID = 0x36607333

// Ensuring interfaces in compile-time for BotInfo.
var (
	_ bin.Encoder     = &BotInfo{}
	_ bin.Decoder     = &BotInfo{}
	_ bin.BareEncoder = &BotInfo{}
	_ bin.BareDecoder = &BotInfo{}
)

func (b *BotInfo) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Flags.Zero()) {
		return false
	}
	if !(b.HasPreviewMedias == false) {
		return false
	}
	if !(b.UserID == 0) {
		return false
	}
	if !(b.Description == "") {
		return false
	}
	if !(b.DescriptionPhoto == nil) {
		return false
	}
	if !(b.DescriptionDocument == nil) {
		return false
	}
	if !(b.Commands == nil) {
		return false
	}
	if !(b.MenuButton == nil) {
		return false
	}
	if !(b.PrivacyPolicyURL == "") {
		return false
	}
	if !(b.AppSettings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotInfo) String() string {
	if b == nil {
		return "BotInfo(nil)"
	}
	type Alias BotInfo
	return fmt.Sprintf("BotInfo%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotInfo) TypeID() uint32 {
	return BotInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*BotInfo) TypeName() string {
	return "botInfo"
}

// TypeInfo returns info about TL type.
func (b *BotInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botInfo",
		ID:   BotInfoTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "HasPreviewMedias",
			SchemaName: "has_preview_medias",
			Null:       !b.Flags.Has(6),
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
			Null:       !b.Flags.Has(0),
		},
		{
			Name:       "Description",
			SchemaName: "description",
			Null:       !b.Flags.Has(1),
		},
		{
			Name:       "DescriptionPhoto",
			SchemaName: "description_photo",
			Null:       !b.Flags.Has(4),
		},
		{
			Name:       "DescriptionDocument",
			SchemaName: "description_document",
			Null:       !b.Flags.Has(5),
		},
		{
			Name:       "Commands",
			SchemaName: "commands",
			Null:       !b.Flags.Has(2),
		},
		{
			Name:       "MenuButton",
			SchemaName: "menu_button",
			Null:       !b.Flags.Has(3),
		},
		{
			Name:       "PrivacyPolicyURL",
			SchemaName: "privacy_policy_url",
			Null:       !b.Flags.Has(7),
		},
		{
			Name:       "AppSettings",
			SchemaName: "app_settings",
			Null:       !b.Flags.Has(8),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (b *BotInfo) SetFlags() {
	if !(b.HasPreviewMedias == false) {
		b.Flags.Set(6)
	}
	if !(b.UserID == 0) {
		b.Flags.Set(0)
	}
	if !(b.Description == "") {
		b.Flags.Set(1)
	}
	if !(b.DescriptionPhoto == nil) {
		b.Flags.Set(4)
	}
	if !(b.DescriptionDocument == nil) {
		b.Flags.Set(5)
	}
	if !(b.Commands == nil) {
		b.Flags.Set(2)
	}
	if !(b.MenuButton == nil) {
		b.Flags.Set(3)
	}
	if !(b.PrivacyPolicyURL == "") {
		b.Flags.Set(7)
	}
	if !(b.AppSettings.Zero()) {
		b.Flags.Set(8)
	}
}

// Encode implements bin.Encoder.
func (b *BotInfo) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInfo#36607333 as nil")
	}
	buf.PutID(BotInfoTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotInfo) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInfo#36607333 as nil")
	}
	b.SetFlags()
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInfo#36607333: field flags: %w", err)
	}
	if b.Flags.Has(0) {
		buf.PutLong(b.UserID)
	}
	if b.Flags.Has(1) {
		buf.PutString(b.Description)
	}
	if b.Flags.Has(4) {
		if b.DescriptionPhoto == nil {
			return fmt.Errorf("unable to encode botInfo#36607333: field description_photo is nil")
		}
		if err := b.DescriptionPhoto.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInfo#36607333: field description_photo: %w", err)
		}
	}
	if b.Flags.Has(5) {
		if b.DescriptionDocument == nil {
			return fmt.Errorf("unable to encode botInfo#36607333: field description_document is nil")
		}
		if err := b.DescriptionDocument.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInfo#36607333: field description_document: %w", err)
		}
	}
	if b.Flags.Has(2) {
		buf.PutVectorHeader(len(b.Commands))
		for idx, v := range b.Commands {
			if err := v.Encode(buf); err != nil {
				return fmt.Errorf("unable to encode botInfo#36607333: field commands element with index %d: %w", idx, err)
			}
		}
	}
	if b.Flags.Has(3) {
		if b.MenuButton == nil {
			return fmt.Errorf("unable to encode botInfo#36607333: field menu_button is nil")
		}
		if err := b.MenuButton.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInfo#36607333: field menu_button: %w", err)
		}
	}
	if b.Flags.Has(7) {
		buf.PutString(b.PrivacyPolicyURL)
	}
	if b.Flags.Has(8) {
		if err := b.AppSettings.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInfo#36607333: field app_settings: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotInfo) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInfo#36607333 to nil")
	}
	if err := buf.ConsumeID(BotInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode botInfo#36607333: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotInfo) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInfo#36607333 to nil")
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInfo#36607333: field flags: %w", err)
		}
	}
	b.HasPreviewMedias = b.Flags.Has(6)
	if b.Flags.Has(0) {
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#36607333: field user_id: %w", err)
		}
		b.UserID = value
	}
	if b.Flags.Has(1) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#36607333: field description: %w", err)
		}
		b.Description = value
	}
	if b.Flags.Has(4) {
		value, err := DecodePhoto(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#36607333: field description_photo: %w", err)
		}
		b.DescriptionPhoto = value
	}
	if b.Flags.Has(5) {
		value, err := DecodeDocument(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#36607333: field description_document: %w", err)
		}
		b.DescriptionDocument = value
	}
	if b.Flags.Has(2) {
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#36607333: field commands: %w", err)
		}

		if headerLen > 0 {
			b.Commands = make([]BotCommand, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BotCommand
			if err := value.Decode(buf); err != nil {
				return fmt.Errorf("unable to decode botInfo#36607333: field commands: %w", err)
			}
			b.Commands = append(b.Commands, value)
		}
	}
	if b.Flags.Has(3) {
		value, err := DecodeBotMenuButton(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#36607333: field menu_button: %w", err)
		}
		b.MenuButton = value
	}
	if b.Flags.Has(7) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#36607333: field privacy_policy_url: %w", err)
		}
		b.PrivacyPolicyURL = value
	}
	if b.Flags.Has(8) {
		if err := b.AppSettings.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInfo#36607333: field app_settings: %w", err)
		}
	}
	return nil
}

// SetHasPreviewMedias sets value of HasPreviewMedias conditional field.
func (b *BotInfo) SetHasPreviewMedias(value bool) {
	if value {
		b.Flags.Set(6)
		b.HasPreviewMedias = true
	} else {
		b.Flags.Unset(6)
		b.HasPreviewMedias = false
	}
}

// GetHasPreviewMedias returns value of HasPreviewMedias conditional field.
func (b *BotInfo) GetHasPreviewMedias() (value bool) {
	if b == nil {
		return
	}
	return b.Flags.Has(6)
}

// SetUserID sets value of UserID conditional field.
func (b *BotInfo) SetUserID(value int64) {
	b.Flags.Set(0)
	b.UserID = value
}

// GetUserID returns value of UserID conditional field and
// boolean which is true if field was set.
func (b *BotInfo) GetUserID() (value int64, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(0) {
		return value, false
	}
	return b.UserID, true
}

// SetDescription sets value of Description conditional field.
func (b *BotInfo) SetDescription(value string) {
	b.Flags.Set(1)
	b.Description = value
}

// GetDescription returns value of Description conditional field and
// boolean which is true if field was set.
func (b *BotInfo) GetDescription() (value string, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(1) {
		return value, false
	}
	return b.Description, true
}

// SetDescriptionPhoto sets value of DescriptionPhoto conditional field.
func (b *BotInfo) SetDescriptionPhoto(value PhotoClass) {
	b.Flags.Set(4)
	b.DescriptionPhoto = value
}

// GetDescriptionPhoto returns value of DescriptionPhoto conditional field and
// boolean which is true if field was set.
func (b *BotInfo) GetDescriptionPhoto() (value PhotoClass, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(4) {
		return value, false
	}
	return b.DescriptionPhoto, true
}

// SetDescriptionDocument sets value of DescriptionDocument conditional field.
func (b *BotInfo) SetDescriptionDocument(value DocumentClass) {
	b.Flags.Set(5)
	b.DescriptionDocument = value
}

// GetDescriptionDocument returns value of DescriptionDocument conditional field and
// boolean which is true if field was set.
func (b *BotInfo) GetDescriptionDocument() (value DocumentClass, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(5) {
		return value, false
	}
	return b.DescriptionDocument, true
}

// SetCommands sets value of Commands conditional field.
func (b *BotInfo) SetCommands(value []BotCommand) {
	b.Flags.Set(2)
	b.Commands = value
}

// GetCommands returns value of Commands conditional field and
// boolean which is true if field was set.
func (b *BotInfo) GetCommands() (value []BotCommand, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.Commands, true
}

// SetMenuButton sets value of MenuButton conditional field.
func (b *BotInfo) SetMenuButton(value BotMenuButtonClass) {
	b.Flags.Set(3)
	b.MenuButton = value
}

// GetMenuButton returns value of MenuButton conditional field and
// boolean which is true if field was set.
func (b *BotInfo) GetMenuButton() (value BotMenuButtonClass, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(3) {
		return value, false
	}
	return b.MenuButton, true
}

// SetPrivacyPolicyURL sets value of PrivacyPolicyURL conditional field.
func (b *BotInfo) SetPrivacyPolicyURL(value string) {
	b.Flags.Set(7)
	b.PrivacyPolicyURL = value
}

// GetPrivacyPolicyURL returns value of PrivacyPolicyURL conditional field and
// boolean which is true if field was set.
func (b *BotInfo) GetPrivacyPolicyURL() (value string, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(7) {
		return value, false
	}
	return b.PrivacyPolicyURL, true
}

// SetAppSettings sets value of AppSettings conditional field.
func (b *BotInfo) SetAppSettings(value BotAppSettings) {
	b.Flags.Set(8)
	b.AppSettings = value
}

// GetAppSettings returns value of AppSettings conditional field and
// boolean which is true if field was set.
func (b *BotInfo) GetAppSettings() (value BotAppSettings, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(8) {
		return value, false
	}
	return b.AppSettings, true
}
