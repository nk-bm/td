// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
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
)

// ResendAuthenticationCodeRequest represents TL type `resendAuthenticationCode#cf759719`.
type ResendAuthenticationCodeRequest struct {
}

// ResendAuthenticationCodeRequestTypeID is TL type id of ResendAuthenticationCodeRequest.
const ResendAuthenticationCodeRequestTypeID = 0xcf759719

// Ensuring interfaces in compile-time for ResendAuthenticationCodeRequest.
var (
	_ bin.Encoder     = &ResendAuthenticationCodeRequest{}
	_ bin.Decoder     = &ResendAuthenticationCodeRequest{}
	_ bin.BareEncoder = &ResendAuthenticationCodeRequest{}
	_ bin.BareDecoder = &ResendAuthenticationCodeRequest{}
)

func (r *ResendAuthenticationCodeRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResendAuthenticationCodeRequest) String() string {
	if r == nil {
		return "ResendAuthenticationCodeRequest(nil)"
	}
	type Alias ResendAuthenticationCodeRequest
	return fmt.Sprintf("ResendAuthenticationCodeRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResendAuthenticationCodeRequest) TypeID() uint32 {
	return ResendAuthenticationCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ResendAuthenticationCodeRequest) TypeName() string {
	return "resendAuthenticationCode"
}

// TypeInfo returns info about TL type.
func (r *ResendAuthenticationCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "resendAuthenticationCode",
		ID:   ResendAuthenticationCodeRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResendAuthenticationCodeRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resendAuthenticationCode#cf759719 as nil")
	}
	b.PutID(ResendAuthenticationCodeRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ResendAuthenticationCodeRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resendAuthenticationCode#cf759719 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ResendAuthenticationCodeRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resendAuthenticationCode#cf759719 to nil")
	}
	if err := b.ConsumeID(ResendAuthenticationCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode resendAuthenticationCode#cf759719: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ResendAuthenticationCodeRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resendAuthenticationCode#cf759719 to nil")
	}
	return nil
}

// ResendAuthenticationCode invokes method resendAuthenticationCode#cf759719 returning error if any.
func (c *Client) ResendAuthenticationCode(ctx context.Context) error {
	var ok Ok

	request := &ResendAuthenticationCodeRequest{}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
