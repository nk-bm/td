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

// BotsToggleUserEmojiStatusPermissionRequest represents TL type `bots.toggleUserEmojiStatusPermission#6de6392`.
type BotsToggleUserEmojiStatusPermissionRequest struct {
	// Bot field of BotsToggleUserEmojiStatusPermissionRequest.
	Bot InputUserClass
	// Enabled field of BotsToggleUserEmojiStatusPermissionRequest.
	Enabled bool
}

// BotsToggleUserEmojiStatusPermissionRequestTypeID is TL type id of BotsToggleUserEmojiStatusPermissionRequest.
const BotsToggleUserEmojiStatusPermissionRequestTypeID = 0x6de6392

// Ensuring interfaces in compile-time for BotsToggleUserEmojiStatusPermissionRequest.
var (
	_ bin.Encoder     = &BotsToggleUserEmojiStatusPermissionRequest{}
	_ bin.Decoder     = &BotsToggleUserEmojiStatusPermissionRequest{}
	_ bin.BareEncoder = &BotsToggleUserEmojiStatusPermissionRequest{}
	_ bin.BareDecoder = &BotsToggleUserEmojiStatusPermissionRequest{}
)

func (t *BotsToggleUserEmojiStatusPermissionRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Bot == nil) {
		return false
	}
	if !(t.Enabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *BotsToggleUserEmojiStatusPermissionRequest) String() string {
	if t == nil {
		return "BotsToggleUserEmojiStatusPermissionRequest(nil)"
	}
	type Alias BotsToggleUserEmojiStatusPermissionRequest
	return fmt.Sprintf("BotsToggleUserEmojiStatusPermissionRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsToggleUserEmojiStatusPermissionRequest) TypeID() uint32 {
	return BotsToggleUserEmojiStatusPermissionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsToggleUserEmojiStatusPermissionRequest) TypeName() string {
	return "bots.toggleUserEmojiStatusPermission"
}

// TypeInfo returns info about TL type.
func (t *BotsToggleUserEmojiStatusPermissionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.toggleUserEmojiStatusPermission",
		ID:   BotsToggleUserEmojiStatusPermissionRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "Enabled",
			SchemaName: "enabled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *BotsToggleUserEmojiStatusPermissionRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode bots.toggleUserEmojiStatusPermission#6de6392 as nil")
	}
	b.PutID(BotsToggleUserEmojiStatusPermissionRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *BotsToggleUserEmojiStatusPermissionRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode bots.toggleUserEmojiStatusPermission#6de6392 as nil")
	}
	if t.Bot == nil {
		return fmt.Errorf("unable to encode bots.toggleUserEmojiStatusPermission#6de6392: field bot is nil")
	}
	if err := t.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.toggleUserEmojiStatusPermission#6de6392: field bot: %w", err)
	}
	b.PutBool(t.Enabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *BotsToggleUserEmojiStatusPermissionRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode bots.toggleUserEmojiStatusPermission#6de6392 to nil")
	}
	if err := b.ConsumeID(BotsToggleUserEmojiStatusPermissionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.toggleUserEmojiStatusPermission#6de6392: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *BotsToggleUserEmojiStatusPermissionRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode bots.toggleUserEmojiStatusPermission#6de6392 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.toggleUserEmojiStatusPermission#6de6392: field bot: %w", err)
		}
		t.Bot = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode bots.toggleUserEmojiStatusPermission#6de6392: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// GetBot returns value of Bot field.
func (t *BotsToggleUserEmojiStatusPermissionRequest) GetBot() (value InputUserClass) {
	if t == nil {
		return
	}
	return t.Bot
}

// GetEnabled returns value of Enabled field.
func (t *BotsToggleUserEmojiStatusPermissionRequest) GetEnabled() (value bool) {
	if t == nil {
		return
	}
	return t.Enabled
}

// BotsToggleUserEmojiStatusPermission invokes method bots.toggleUserEmojiStatusPermission#6de6392 returning error if any.
func (c *Client) BotsToggleUserEmojiStatusPermission(ctx context.Context, request *BotsToggleUserEmojiStatusPermissionRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
