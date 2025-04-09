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

// ChannelsUpdateUsernameRequest represents TL type `channels.updateUsername#3514b3de`.
type ChannelsUpdateUsernameRequest struct {
	// Channel field of ChannelsUpdateUsernameRequest.
	Channel InputChannelClass
	// Username field of ChannelsUpdateUsernameRequest.
	Username string
}

// ChannelsUpdateUsernameRequestTypeID is TL type id of ChannelsUpdateUsernameRequest.
const ChannelsUpdateUsernameRequestTypeID = 0x3514b3de

// Ensuring interfaces in compile-time for ChannelsUpdateUsernameRequest.
var (
	_ bin.Encoder     = &ChannelsUpdateUsernameRequest{}
	_ bin.Decoder     = &ChannelsUpdateUsernameRequest{}
	_ bin.BareEncoder = &ChannelsUpdateUsernameRequest{}
	_ bin.BareDecoder = &ChannelsUpdateUsernameRequest{}
)

func (u *ChannelsUpdateUsernameRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Channel == nil) {
		return false
	}
	if !(u.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *ChannelsUpdateUsernameRequest) String() string {
	if u == nil {
		return "ChannelsUpdateUsernameRequest(nil)"
	}
	type Alias ChannelsUpdateUsernameRequest
	return fmt.Sprintf("ChannelsUpdateUsernameRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsUpdateUsernameRequest) TypeID() uint32 {
	return ChannelsUpdateUsernameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsUpdateUsernameRequest) TypeName() string {
	return "channels.updateUsername"
}

// TypeInfo returns info about TL type.
func (u *ChannelsUpdateUsernameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.updateUsername",
		ID:   ChannelsUpdateUsernameRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Username",
			SchemaName: "username",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *ChannelsUpdateUsernameRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode channels.updateUsername#3514b3de as nil")
	}
	b.PutID(ChannelsUpdateUsernameRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *ChannelsUpdateUsernameRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode channels.updateUsername#3514b3de as nil")
	}
	if u.Channel == nil {
		return fmt.Errorf("unable to encode channels.updateUsername#3514b3de: field channel is nil")
	}
	if err := u.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.updateUsername#3514b3de: field channel: %w", err)
	}
	b.PutString(u.Username)
	return nil
}

// Decode implements bin.Decoder.
func (u *ChannelsUpdateUsernameRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode channels.updateUsername#3514b3de to nil")
	}
	if err := b.ConsumeID(ChannelsUpdateUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.updateUsername#3514b3de: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *ChannelsUpdateUsernameRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode channels.updateUsername#3514b3de to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.updateUsername#3514b3de: field channel: %w", err)
		}
		u.Channel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.updateUsername#3514b3de: field username: %w", err)
		}
		u.Username = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (u *ChannelsUpdateUsernameRequest) GetChannel() (value InputChannelClass) {
	if u == nil {
		return
	}
	return u.Channel
}

// GetUsername returns value of Username field.
func (u *ChannelsUpdateUsernameRequest) GetUsername() (value string) {
	if u == nil {
		return
	}
	return u.Username
}

// ChannelsUpdateUsername invokes method channels.updateUsername#3514b3de returning error if any.
func (c *Client) ChannelsUpdateUsername(ctx context.Context, request *ChannelsUpdateUsernameRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
