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

// AnswerPreCheckoutQueryRequest represents TL type `answerPreCheckoutQuery#a76163eb`.
type AnswerPreCheckoutQueryRequest struct {
	// Identifier of the pre-checkout query
	PreCheckoutQueryID Int64
	// An error message, empty on success
	ErrorMessage string
}

// AnswerPreCheckoutQueryRequestTypeID is TL type id of AnswerPreCheckoutQueryRequest.
const AnswerPreCheckoutQueryRequestTypeID = 0xa76163eb

// Ensuring interfaces in compile-time for AnswerPreCheckoutQueryRequest.
var (
	_ bin.Encoder     = &AnswerPreCheckoutQueryRequest{}
	_ bin.Decoder     = &AnswerPreCheckoutQueryRequest{}
	_ bin.BareEncoder = &AnswerPreCheckoutQueryRequest{}
	_ bin.BareDecoder = &AnswerPreCheckoutQueryRequest{}
)

func (a *AnswerPreCheckoutQueryRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.PreCheckoutQueryID.Zero()) {
		return false
	}
	if !(a.ErrorMessage == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AnswerPreCheckoutQueryRequest) String() string {
	if a == nil {
		return "AnswerPreCheckoutQueryRequest(nil)"
	}
	type Alias AnswerPreCheckoutQueryRequest
	return fmt.Sprintf("AnswerPreCheckoutQueryRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AnswerPreCheckoutQueryRequest) TypeID() uint32 {
	return AnswerPreCheckoutQueryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AnswerPreCheckoutQueryRequest) TypeName() string {
	return "answerPreCheckoutQuery"
}

// TypeInfo returns info about TL type.
func (a *AnswerPreCheckoutQueryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "answerPreCheckoutQuery",
		ID:   AnswerPreCheckoutQueryRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PreCheckoutQueryID",
			SchemaName: "pre_checkout_query_id",
		},
		{
			Name:       "ErrorMessage",
			SchemaName: "error_message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AnswerPreCheckoutQueryRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode answerPreCheckoutQuery#a76163eb as nil")
	}
	b.PutID(AnswerPreCheckoutQueryRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AnswerPreCheckoutQueryRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode answerPreCheckoutQuery#a76163eb as nil")
	}
	if err := a.PreCheckoutQueryID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode answerPreCheckoutQuery#a76163eb: field pre_checkout_query_id: %w", err)
	}
	b.PutString(a.ErrorMessage)
	return nil
}

// Decode implements bin.Decoder.
func (a *AnswerPreCheckoutQueryRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode answerPreCheckoutQuery#a76163eb to nil")
	}
	if err := b.ConsumeID(AnswerPreCheckoutQueryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode answerPreCheckoutQuery#a76163eb: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AnswerPreCheckoutQueryRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode answerPreCheckoutQuery#a76163eb to nil")
	}
	{
		if err := a.PreCheckoutQueryID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode answerPreCheckoutQuery#a76163eb: field pre_checkout_query_id: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode answerPreCheckoutQuery#a76163eb: field error_message: %w", err)
		}
		a.ErrorMessage = value
	}
	return nil
}

// GetPreCheckoutQueryID returns value of PreCheckoutQueryID field.
func (a *AnswerPreCheckoutQueryRequest) GetPreCheckoutQueryID() (value Int64) {
	return a.PreCheckoutQueryID
}

// GetErrorMessage returns value of ErrorMessage field.
func (a *AnswerPreCheckoutQueryRequest) GetErrorMessage() (value string) {
	return a.ErrorMessage
}

// AnswerPreCheckoutQuery invokes method answerPreCheckoutQuery#a76163eb returning error if any.
func (c *Client) AnswerPreCheckoutQuery(ctx context.Context, request *AnswerPreCheckoutQueryRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
