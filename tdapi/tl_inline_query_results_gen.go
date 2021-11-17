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

// InlineQueryResults represents TL type `inlineQueryResults#6ecde5be`.
type InlineQueryResults struct {
	// Unique identifier of the inline query
	InlineQueryID Int64
	// The offset for the next request. If empty, there are no more results
	NextOffset string
	// Results of the query
	Results []InlineQueryResultClass
	// If non-empty, this text should be shown on the button, which opens a private chat with
	// the bot and sends the bot a start message with the switch_pm_parameter
	SwitchPmText string
	// Parameter for the bot start message
	SwitchPmParameter string
}

// InlineQueryResultsTypeID is TL type id of InlineQueryResults.
const InlineQueryResultsTypeID = 0x6ecde5be

// Ensuring interfaces in compile-time for InlineQueryResults.
var (
	_ bin.Encoder     = &InlineQueryResults{}
	_ bin.Decoder     = &InlineQueryResults{}
	_ bin.BareEncoder = &InlineQueryResults{}
	_ bin.BareDecoder = &InlineQueryResults{}
)

func (i *InlineQueryResults) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.InlineQueryID.Zero()) {
		return false
	}
	if !(i.NextOffset == "") {
		return false
	}
	if !(i.Results == nil) {
		return false
	}
	if !(i.SwitchPmText == "") {
		return false
	}
	if !(i.SwitchPmParameter == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryResults) String() string {
	if i == nil {
		return "InlineQueryResults(nil)"
	}
	type Alias InlineQueryResults
	return fmt.Sprintf("InlineQueryResults%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineQueryResults) TypeID() uint32 {
	return InlineQueryResultsTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineQueryResults) TypeName() string {
	return "inlineQueryResults"
}

// TypeInfo returns info about TL type.
func (i *InlineQueryResults) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineQueryResults",
		ID:   InlineQueryResultsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InlineQueryID",
			SchemaName: "inline_query_id",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
		},
		{
			Name:       "Results",
			SchemaName: "results",
		},
		{
			Name:       "SwitchPmText",
			SchemaName: "switch_pm_text",
		},
		{
			Name:       "SwitchPmParameter",
			SchemaName: "switch_pm_parameter",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineQueryResults) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryResults#6ecde5be as nil")
	}
	b.PutID(InlineQueryResultsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineQueryResults) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryResults#6ecde5be as nil")
	}
	if err := i.InlineQueryID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inlineQueryResults#6ecde5be: field inline_query_id: %w", err)
	}
	b.PutString(i.NextOffset)
	b.PutInt(len(i.Results))
	for idx, v := range i.Results {
		if v == nil {
			return fmt.Errorf("unable to encode inlineQueryResults#6ecde5be: field results element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare inlineQueryResults#6ecde5be: field results element with index %d: %w", idx, err)
		}
	}
	b.PutString(i.SwitchPmText)
	b.PutString(i.SwitchPmParameter)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryResults) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryResults#6ecde5be to nil")
	}
	if err := b.ConsumeID(InlineQueryResultsTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryResults#6ecde5be: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineQueryResults) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryResults#6ecde5be to nil")
	}
	{
		if err := i.InlineQueryID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inlineQueryResults#6ecde5be: field inline_query_id: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineQueryResults#6ecde5be: field next_offset: %w", err)
		}
		i.NextOffset = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inlineQueryResults#6ecde5be: field results: %w", err)
		}

		if headerLen > 0 {
			i.Results = make([]InlineQueryResultClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInlineQueryResult(b)
			if err != nil {
				return fmt.Errorf("unable to decode inlineQueryResults#6ecde5be: field results: %w", err)
			}
			i.Results = append(i.Results, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineQueryResults#6ecde5be: field switch_pm_text: %w", err)
		}
		i.SwitchPmText = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineQueryResults#6ecde5be: field switch_pm_parameter: %w", err)
		}
		i.SwitchPmParameter = value
	}
	return nil
}

// GetInlineQueryID returns value of InlineQueryID field.
func (i *InlineQueryResults) GetInlineQueryID() (value Int64) {
	return i.InlineQueryID
}

// GetNextOffset returns value of NextOffset field.
func (i *InlineQueryResults) GetNextOffset() (value string) {
	return i.NextOffset
}

// GetResults returns value of Results field.
func (i *InlineQueryResults) GetResults() (value []InlineQueryResultClass) {
	return i.Results
}

// GetSwitchPmText returns value of SwitchPmText field.
func (i *InlineQueryResults) GetSwitchPmText() (value string) {
	return i.SwitchPmText
}

// GetSwitchPmParameter returns value of SwitchPmParameter field.
func (i *InlineQueryResults) GetSwitchPmParameter() (value string) {
	return i.SwitchPmParameter
}
