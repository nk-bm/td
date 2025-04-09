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

// BotsReorderUsernamesRequest represents TL type `bots.reorderUsernames#9709b1c2`.
type BotsReorderUsernamesRequest struct {
	// Bot field of BotsReorderUsernamesRequest.
	Bot InputUserClass
	// Order field of BotsReorderUsernamesRequest.
	Order []string
}

// BotsReorderUsernamesRequestTypeID is TL type id of BotsReorderUsernamesRequest.
const BotsReorderUsernamesRequestTypeID = 0x9709b1c2

// Ensuring interfaces in compile-time for BotsReorderUsernamesRequest.
var (
	_ bin.Encoder     = &BotsReorderUsernamesRequest{}
	_ bin.Decoder     = &BotsReorderUsernamesRequest{}
	_ bin.BareEncoder = &BotsReorderUsernamesRequest{}
	_ bin.BareDecoder = &BotsReorderUsernamesRequest{}
)

func (r *BotsReorderUsernamesRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Bot == nil) {
		return false
	}
	if !(r.Order == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *BotsReorderUsernamesRequest) String() string {
	if r == nil {
		return "BotsReorderUsernamesRequest(nil)"
	}
	type Alias BotsReorderUsernamesRequest
	return fmt.Sprintf("BotsReorderUsernamesRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsReorderUsernamesRequest) TypeID() uint32 {
	return BotsReorderUsernamesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsReorderUsernamesRequest) TypeName() string {
	return "bots.reorderUsernames"
}

// TypeInfo returns info about TL type.
func (r *BotsReorderUsernamesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.reorderUsernames",
		ID:   BotsReorderUsernamesRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "Order",
			SchemaName: "order",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *BotsReorderUsernamesRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode bots.reorderUsernames#9709b1c2 as nil")
	}
	b.PutID(BotsReorderUsernamesRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *BotsReorderUsernamesRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode bots.reorderUsernames#9709b1c2 as nil")
	}
	if r.Bot == nil {
		return fmt.Errorf("unable to encode bots.reorderUsernames#9709b1c2: field bot is nil")
	}
	if err := r.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.reorderUsernames#9709b1c2: field bot: %w", err)
	}
	b.PutVectorHeader(len(r.Order))
	for _, v := range r.Order {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *BotsReorderUsernamesRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode bots.reorderUsernames#9709b1c2 to nil")
	}
	if err := b.ConsumeID(BotsReorderUsernamesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.reorderUsernames#9709b1c2: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *BotsReorderUsernamesRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode bots.reorderUsernames#9709b1c2 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.reorderUsernames#9709b1c2: field bot: %w", err)
		}
		r.Bot = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode bots.reorderUsernames#9709b1c2: field order: %w", err)
		}

		if headerLen > 0 {
			r.Order = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode bots.reorderUsernames#9709b1c2: field order: %w", err)
			}
			r.Order = append(r.Order, value)
		}
	}
	return nil
}

// GetBot returns value of Bot field.
func (r *BotsReorderUsernamesRequest) GetBot() (value InputUserClass) {
	if r == nil {
		return
	}
	return r.Bot
}

// GetOrder returns value of Order field.
func (r *BotsReorderUsernamesRequest) GetOrder() (value []string) {
	if r == nil {
		return
	}
	return r.Order
}

// BotsReorderUsernames invokes method bots.reorderUsernames#9709b1c2 returning error if any.
func (c *Client) BotsReorderUsernames(ctx context.Context, request *BotsReorderUsernamesRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
