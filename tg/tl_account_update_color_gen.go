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

// AccountUpdateColorRequest represents TL type `account.updateColor#7cefa15d`.
type AccountUpdateColorRequest struct {
	// Flags field of AccountUpdateColorRequest.
	Flags bin.Fields
	// ForProfile field of AccountUpdateColorRequest.
	ForProfile bool
	// Color field of AccountUpdateColorRequest.
	//
	// Use SetColor and GetColor helpers.
	Color int
	// BackgroundEmojiID field of AccountUpdateColorRequest.
	//
	// Use SetBackgroundEmojiID and GetBackgroundEmojiID helpers.
	BackgroundEmojiID int64
}

// AccountUpdateColorRequestTypeID is TL type id of AccountUpdateColorRequest.
const AccountUpdateColorRequestTypeID = 0x7cefa15d

// Ensuring interfaces in compile-time for AccountUpdateColorRequest.
var (
	_ bin.Encoder     = &AccountUpdateColorRequest{}
	_ bin.Decoder     = &AccountUpdateColorRequest{}
	_ bin.BareEncoder = &AccountUpdateColorRequest{}
	_ bin.BareDecoder = &AccountUpdateColorRequest{}
)

func (u *AccountUpdateColorRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.ForProfile == false) {
		return false
	}
	if !(u.Color == 0) {
		return false
	}
	if !(u.BackgroundEmojiID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateColorRequest) String() string {
	if u == nil {
		return "AccountUpdateColorRequest(nil)"
	}
	type Alias AccountUpdateColorRequest
	return fmt.Sprintf("AccountUpdateColorRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateColorRequest) TypeID() uint32 {
	return AccountUpdateColorRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateColorRequest) TypeName() string {
	return "account.updateColor"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateColorRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateColor",
		ID:   AccountUpdateColorRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ForProfile",
			SchemaName: "for_profile",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "Color",
			SchemaName: "color",
			Null:       !u.Flags.Has(2),
		},
		{
			Name:       "BackgroundEmojiID",
			SchemaName: "background_emoji_id",
			Null:       !u.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *AccountUpdateColorRequest) SetFlags() {
	if !(u.ForProfile == false) {
		u.Flags.Set(1)
	}
	if !(u.Color == 0) {
		u.Flags.Set(2)
	}
	if !(u.BackgroundEmojiID == 0) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *AccountUpdateColorRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateColor#7cefa15d as nil")
	}
	b.PutID(AccountUpdateColorRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdateColorRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateColor#7cefa15d as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateColor#7cefa15d: field flags: %w", err)
	}
	if u.Flags.Has(2) {
		b.PutInt(u.Color)
	}
	if u.Flags.Has(0) {
		b.PutLong(u.BackgroundEmojiID)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateColorRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateColor#7cefa15d to nil")
	}
	if err := b.ConsumeID(AccountUpdateColorRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateColor#7cefa15d: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdateColorRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateColor#7cefa15d to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateColor#7cefa15d: field flags: %w", err)
		}
	}
	u.ForProfile = u.Flags.Has(1)
	if u.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateColor#7cefa15d: field color: %w", err)
		}
		u.Color = value
	}
	if u.Flags.Has(0) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateColor#7cefa15d: field background_emoji_id: %w", err)
		}
		u.BackgroundEmojiID = value
	}
	return nil
}

// SetForProfile sets value of ForProfile conditional field.
func (u *AccountUpdateColorRequest) SetForProfile(value bool) {
	if value {
		u.Flags.Set(1)
		u.ForProfile = true
	} else {
		u.Flags.Unset(1)
		u.ForProfile = false
	}
}

// GetForProfile returns value of ForProfile conditional field.
func (u *AccountUpdateColorRequest) GetForProfile() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(1)
}

// SetColor sets value of Color conditional field.
func (u *AccountUpdateColorRequest) SetColor(value int) {
	u.Flags.Set(2)
	u.Color = value
}

// GetColor returns value of Color conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateColorRequest) GetColor() (value int, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.Color, true
}

// SetBackgroundEmojiID sets value of BackgroundEmojiID conditional field.
func (u *AccountUpdateColorRequest) SetBackgroundEmojiID(value int64) {
	u.Flags.Set(0)
	u.BackgroundEmojiID = value
}

// GetBackgroundEmojiID returns value of BackgroundEmojiID conditional field and
// boolean which is true if field was set.
func (u *AccountUpdateColorRequest) GetBackgroundEmojiID() (value int64, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.BackgroundEmojiID, true
}

// AccountUpdateColor invokes method account.updateColor#7cefa15d returning error if any.
func (c *Client) AccountUpdateColor(ctx context.Context, request *AccountUpdateColorRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
