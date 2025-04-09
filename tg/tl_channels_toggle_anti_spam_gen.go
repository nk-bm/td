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

// ChannelsToggleAntiSpamRequest represents TL type `channels.toggleAntiSpam#68f3e4eb`.
type ChannelsToggleAntiSpamRequest struct {
	// Channel field of ChannelsToggleAntiSpamRequest.
	Channel InputChannelClass
	// Enabled field of ChannelsToggleAntiSpamRequest.
	Enabled bool
}

// ChannelsToggleAntiSpamRequestTypeID is TL type id of ChannelsToggleAntiSpamRequest.
const ChannelsToggleAntiSpamRequestTypeID = 0x68f3e4eb

// Ensuring interfaces in compile-time for ChannelsToggleAntiSpamRequest.
var (
	_ bin.Encoder     = &ChannelsToggleAntiSpamRequest{}
	_ bin.Decoder     = &ChannelsToggleAntiSpamRequest{}
	_ bin.BareEncoder = &ChannelsToggleAntiSpamRequest{}
	_ bin.BareDecoder = &ChannelsToggleAntiSpamRequest{}
)

func (t *ChannelsToggleAntiSpamRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Channel == nil) {
		return false
	}
	if !(t.Enabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ChannelsToggleAntiSpamRequest) String() string {
	if t == nil {
		return "ChannelsToggleAntiSpamRequest(nil)"
	}
	type Alias ChannelsToggleAntiSpamRequest
	return fmt.Sprintf("ChannelsToggleAntiSpamRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsToggleAntiSpamRequest) TypeID() uint32 {
	return ChannelsToggleAntiSpamRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsToggleAntiSpamRequest) TypeName() string {
	return "channels.toggleAntiSpam"
}

// TypeInfo returns info about TL type.
func (t *ChannelsToggleAntiSpamRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.toggleAntiSpam",
		ID:   ChannelsToggleAntiSpamRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Enabled",
			SchemaName: "enabled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ChannelsToggleAntiSpamRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleAntiSpam#68f3e4eb as nil")
	}
	b.PutID(ChannelsToggleAntiSpamRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ChannelsToggleAntiSpamRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleAntiSpam#68f3e4eb as nil")
	}
	if t.Channel == nil {
		return fmt.Errorf("unable to encode channels.toggleAntiSpam#68f3e4eb: field channel is nil")
	}
	if err := t.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.toggleAntiSpam#68f3e4eb: field channel: %w", err)
	}
	b.PutBool(t.Enabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *ChannelsToggleAntiSpamRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleAntiSpam#68f3e4eb to nil")
	}
	if err := b.ConsumeID(ChannelsToggleAntiSpamRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.toggleAntiSpam#68f3e4eb: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ChannelsToggleAntiSpamRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleAntiSpam#68f3e4eb to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleAntiSpam#68f3e4eb: field channel: %w", err)
		}
		t.Channel = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleAntiSpam#68f3e4eb: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (t *ChannelsToggleAntiSpamRequest) GetChannel() (value InputChannelClass) {
	if t == nil {
		return
	}
	return t.Channel
}

// GetEnabled returns value of Enabled field.
func (t *ChannelsToggleAntiSpamRequest) GetEnabled() (value bool) {
	if t == nil {
		return
	}
	return t.Enabled
}

// ChannelsToggleAntiSpam invokes method channels.toggleAntiSpam#68f3e4eb returning error if any.
func (c *Client) ChannelsToggleAntiSpam(ctx context.Context, request *ChannelsToggleAntiSpamRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
