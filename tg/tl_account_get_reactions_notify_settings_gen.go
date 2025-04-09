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

// AccountGetReactionsNotifySettingsRequest represents TL type `account.getReactionsNotifySettings#6dd654c`.
type AccountGetReactionsNotifySettingsRequest struct {
}

// AccountGetReactionsNotifySettingsRequestTypeID is TL type id of AccountGetReactionsNotifySettingsRequest.
const AccountGetReactionsNotifySettingsRequestTypeID = 0x6dd654c

// Ensuring interfaces in compile-time for AccountGetReactionsNotifySettingsRequest.
var (
	_ bin.Encoder     = &AccountGetReactionsNotifySettingsRequest{}
	_ bin.Decoder     = &AccountGetReactionsNotifySettingsRequest{}
	_ bin.BareEncoder = &AccountGetReactionsNotifySettingsRequest{}
	_ bin.BareDecoder = &AccountGetReactionsNotifySettingsRequest{}
)

func (g *AccountGetReactionsNotifySettingsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetReactionsNotifySettingsRequest) String() string {
	if g == nil {
		return "AccountGetReactionsNotifySettingsRequest(nil)"
	}
	type Alias AccountGetReactionsNotifySettingsRequest
	return fmt.Sprintf("AccountGetReactionsNotifySettingsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetReactionsNotifySettingsRequest) TypeID() uint32 {
	return AccountGetReactionsNotifySettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetReactionsNotifySettingsRequest) TypeName() string {
	return "account.getReactionsNotifySettings"
}

// TypeInfo returns info about TL type.
func (g *AccountGetReactionsNotifySettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getReactionsNotifySettings",
		ID:   AccountGetReactionsNotifySettingsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetReactionsNotifySettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getReactionsNotifySettings#6dd654c as nil")
	}
	b.PutID(AccountGetReactionsNotifySettingsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetReactionsNotifySettingsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getReactionsNotifySettings#6dd654c as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetReactionsNotifySettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getReactionsNotifySettings#6dd654c to nil")
	}
	if err := b.ConsumeID(AccountGetReactionsNotifySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getReactionsNotifySettings#6dd654c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetReactionsNotifySettingsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getReactionsNotifySettings#6dd654c to nil")
	}
	return nil
}

// AccountGetReactionsNotifySettings invokes method account.getReactionsNotifySettings#6dd654c returning error if any.
func (c *Client) AccountGetReactionsNotifySettings(ctx context.Context) (*ReactionsNotifySettings, error) {
	var result ReactionsNotifySettings

	request := &AccountGetReactionsNotifySettingsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
