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

// AccountGetPasswordRequest represents TL type `account.getPassword#548a30f5`.
type AccountGetPasswordRequest struct {
}

// AccountGetPasswordRequestTypeID is TL type id of AccountGetPasswordRequest.
const AccountGetPasswordRequestTypeID = 0x548a30f5

// Ensuring interfaces in compile-time for AccountGetPasswordRequest.
var (
	_ bin.Encoder     = &AccountGetPasswordRequest{}
	_ bin.Decoder     = &AccountGetPasswordRequest{}
	_ bin.BareEncoder = &AccountGetPasswordRequest{}
	_ bin.BareDecoder = &AccountGetPasswordRequest{}
)

func (g *AccountGetPasswordRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetPasswordRequest) String() string {
	if g == nil {
		return "AccountGetPasswordRequest(nil)"
	}
	type Alias AccountGetPasswordRequest
	return fmt.Sprintf("AccountGetPasswordRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetPasswordRequest) TypeID() uint32 {
	return AccountGetPasswordRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetPasswordRequest) TypeName() string {
	return "account.getPassword"
}

// TypeInfo returns info about TL type.
func (g *AccountGetPasswordRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getPassword",
		ID:   AccountGetPasswordRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetPasswordRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getPassword#548a30f5 as nil")
	}
	b.PutID(AccountGetPasswordRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetPasswordRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getPassword#548a30f5 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetPasswordRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getPassword#548a30f5 to nil")
	}
	if err := b.ConsumeID(AccountGetPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getPassword#548a30f5: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetPasswordRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getPassword#548a30f5 to nil")
	}
	return nil
}

// AccountGetPassword invokes method account.getPassword#548a30f5 returning error if any.
func (c *Client) AccountGetPassword(ctx context.Context) (*AccountPassword, error) {
	var result AccountPassword

	request := &AccountGetPasswordRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
