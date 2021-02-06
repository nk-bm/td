// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// HelpDismissSuggestionRequest represents TL type `help.dismissSuggestion#77fa99f`.
// Dismiss a suggestion
//
// See https://core.telegram.org/method/help.dismissSuggestion for reference.
type HelpDismissSuggestionRequest struct {
	// Suggestion
	Suggestion string
}

// HelpDismissSuggestionRequestTypeID is TL type id of HelpDismissSuggestionRequest.
const HelpDismissSuggestionRequestTypeID = 0x77fa99f

func (d *HelpDismissSuggestionRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Suggestion == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *HelpDismissSuggestionRequest) String() string {
	if d == nil {
		return "HelpDismissSuggestionRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpDismissSuggestionRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tSuggestion: ")
	sb.WriteString(fmt.Sprint(d.Suggestion))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *HelpDismissSuggestionRequest) TypeID() uint32 {
	return HelpDismissSuggestionRequestTypeID
}

// Encode implements bin.Encoder.
func (d *HelpDismissSuggestionRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode help.dismissSuggestion#77fa99f as nil")
	}
	b.PutID(HelpDismissSuggestionRequestTypeID)
	b.PutString(d.Suggestion)
	return nil
}

// GetSuggestion returns value of Suggestion field.
func (d *HelpDismissSuggestionRequest) GetSuggestion() (value string) {
	return d.Suggestion
}

// Decode implements bin.Decoder.
func (d *HelpDismissSuggestionRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode help.dismissSuggestion#77fa99f to nil")
	}
	if err := b.ConsumeID(HelpDismissSuggestionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.dismissSuggestion#77fa99f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.dismissSuggestion#77fa99f: field suggestion: %w", err)
		}
		d.Suggestion = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpDismissSuggestionRequest.
var (
	_ bin.Encoder = &HelpDismissSuggestionRequest{}
	_ bin.Decoder = &HelpDismissSuggestionRequest{}
)

// HelpDismissSuggestion invokes method help.dismissSuggestion#77fa99f returning error if any.
// Dismiss a suggestion
//
// See https://core.telegram.org/method/help.dismissSuggestion for reference.
// Can be used by bots.
func (c *Client) HelpDismissSuggestion(ctx context.Context, suggestion string) (bool, error) {
	var result BoolBox

	request := &HelpDismissSuggestionRequest{
		Suggestion: suggestion,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
