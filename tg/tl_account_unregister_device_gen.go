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

// AccountUnregisterDeviceRequest represents TL type `account.unregisterDevice#6a0d3206`.
type AccountUnregisterDeviceRequest struct {
	// TokenType field of AccountUnregisterDeviceRequest.
	TokenType int
	// Token field of AccountUnregisterDeviceRequest.
	Token string
	// OtherUIDs field of AccountUnregisterDeviceRequest.
	OtherUIDs []int64
}

// AccountUnregisterDeviceRequestTypeID is TL type id of AccountUnregisterDeviceRequest.
const AccountUnregisterDeviceRequestTypeID = 0x6a0d3206

// Ensuring interfaces in compile-time for AccountUnregisterDeviceRequest.
var (
	_ bin.Encoder     = &AccountUnregisterDeviceRequest{}
	_ bin.Decoder     = &AccountUnregisterDeviceRequest{}
	_ bin.BareEncoder = &AccountUnregisterDeviceRequest{}
	_ bin.BareDecoder = &AccountUnregisterDeviceRequest{}
)

func (u *AccountUnregisterDeviceRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.TokenType == 0) {
		return false
	}
	if !(u.Token == "") {
		return false
	}
	if !(u.OtherUIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUnregisterDeviceRequest) String() string {
	if u == nil {
		return "AccountUnregisterDeviceRequest(nil)"
	}
	type Alias AccountUnregisterDeviceRequest
	return fmt.Sprintf("AccountUnregisterDeviceRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUnregisterDeviceRequest) TypeID() uint32 {
	return AccountUnregisterDeviceRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUnregisterDeviceRequest) TypeName() string {
	return "account.unregisterDevice"
}

// TypeInfo returns info about TL type.
func (u *AccountUnregisterDeviceRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.unregisterDevice",
		ID:   AccountUnregisterDeviceRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TokenType",
			SchemaName: "token_type",
		},
		{
			Name:       "Token",
			SchemaName: "token",
		},
		{
			Name:       "OtherUIDs",
			SchemaName: "other_uids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUnregisterDeviceRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.unregisterDevice#6a0d3206 as nil")
	}
	b.PutID(AccountUnregisterDeviceRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUnregisterDeviceRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.unregisterDevice#6a0d3206 as nil")
	}
	b.PutInt(u.TokenType)
	b.PutString(u.Token)
	b.PutVectorHeader(len(u.OtherUIDs))
	for _, v := range u.OtherUIDs {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUnregisterDeviceRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.unregisterDevice#6a0d3206 to nil")
	}
	if err := b.ConsumeID(AccountUnregisterDeviceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.unregisterDevice#6a0d3206: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUnregisterDeviceRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.unregisterDevice#6a0d3206 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.unregisterDevice#6a0d3206: field token_type: %w", err)
		}
		u.TokenType = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.unregisterDevice#6a0d3206: field token: %w", err)
		}
		u.Token = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.unregisterDevice#6a0d3206: field other_uids: %w", err)
		}

		if headerLen > 0 {
			u.OtherUIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode account.unregisterDevice#6a0d3206: field other_uids: %w", err)
			}
			u.OtherUIDs = append(u.OtherUIDs, value)
		}
	}
	return nil
}

// GetTokenType returns value of TokenType field.
func (u *AccountUnregisterDeviceRequest) GetTokenType() (value int) {
	if u == nil {
		return
	}
	return u.TokenType
}

// GetToken returns value of Token field.
func (u *AccountUnregisterDeviceRequest) GetToken() (value string) {
	if u == nil {
		return
	}
	return u.Token
}

// GetOtherUIDs returns value of OtherUIDs field.
func (u *AccountUnregisterDeviceRequest) GetOtherUIDs() (value []int64) {
	if u == nil {
		return
	}
	return u.OtherUIDs
}

// AccountUnregisterDevice invokes method account.unregisterDevice#6a0d3206 returning error if any.
func (c *Client) AccountUnregisterDevice(ctx context.Context, request *AccountUnregisterDeviceRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
