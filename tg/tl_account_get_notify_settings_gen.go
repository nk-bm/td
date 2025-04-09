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

// AccountGetNotifySettingsRequest represents TL type `account.getNotifySettings#12b3ad31`.
type AccountGetNotifySettingsRequest struct {
	// Peer field of AccountGetNotifySettingsRequest.
	Peer InputNotifyPeerClass
}

// AccountGetNotifySettingsRequestTypeID is TL type id of AccountGetNotifySettingsRequest.
const AccountGetNotifySettingsRequestTypeID = 0x12b3ad31

// Ensuring interfaces in compile-time for AccountGetNotifySettingsRequest.
var (
	_ bin.Encoder     = &AccountGetNotifySettingsRequest{}
	_ bin.Decoder     = &AccountGetNotifySettingsRequest{}
	_ bin.BareEncoder = &AccountGetNotifySettingsRequest{}
	_ bin.BareDecoder = &AccountGetNotifySettingsRequest{}
)

func (g *AccountGetNotifySettingsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetNotifySettingsRequest) String() string {
	if g == nil {
		return "AccountGetNotifySettingsRequest(nil)"
	}
	type Alias AccountGetNotifySettingsRequest
	return fmt.Sprintf("AccountGetNotifySettingsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetNotifySettingsRequest) TypeID() uint32 {
	return AccountGetNotifySettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetNotifySettingsRequest) TypeName() string {
	return "account.getNotifySettings"
}

// TypeInfo returns info about TL type.
func (g *AccountGetNotifySettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getNotifySettings",
		ID:   AccountGetNotifySettingsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetNotifySettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getNotifySettings#12b3ad31 as nil")
	}
	b.PutID(AccountGetNotifySettingsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetNotifySettingsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getNotifySettings#12b3ad31 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode account.getNotifySettings#12b3ad31: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getNotifySettings#12b3ad31: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetNotifySettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getNotifySettings#12b3ad31 to nil")
	}
	if err := b.ConsumeID(AccountGetNotifySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getNotifySettings#12b3ad31: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetNotifySettingsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getNotifySettings#12b3ad31 to nil")
	}
	{
		value, err := DecodeInputNotifyPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getNotifySettings#12b3ad31: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *AccountGetNotifySettingsRequest) GetPeer() (value InputNotifyPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// AccountGetNotifySettings invokes method account.getNotifySettings#12b3ad31 returning error if any.
func (c *Client) AccountGetNotifySettings(ctx context.Context, peer InputNotifyPeerClass) (*PeerNotifySettings, error) {
	var result PeerNotifySettings

	request := &AccountGetNotifySettingsRequest{
		Peer: peer,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
