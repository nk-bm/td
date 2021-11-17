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

// GetCommandsRequest represents TL type `getCommands#58ba8ff7`.
type GetCommandsRequest struct {
	// The scope to which the commands are relevant
	Scope BotCommandScopeClass
	// A two-letter ISO 639-1 country code or an empty string
	LanguageCode string
}

// GetCommandsRequestTypeID is TL type id of GetCommandsRequest.
const GetCommandsRequestTypeID = 0x58ba8ff7

// Ensuring interfaces in compile-time for GetCommandsRequest.
var (
	_ bin.Encoder     = &GetCommandsRequest{}
	_ bin.Decoder     = &GetCommandsRequest{}
	_ bin.BareEncoder = &GetCommandsRequest{}
	_ bin.BareDecoder = &GetCommandsRequest{}
)

func (g *GetCommandsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Scope == nil) {
		return false
	}
	if !(g.LanguageCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetCommandsRequest) String() string {
	if g == nil {
		return "GetCommandsRequest(nil)"
	}
	type Alias GetCommandsRequest
	return fmt.Sprintf("GetCommandsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetCommandsRequest) TypeID() uint32 {
	return GetCommandsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetCommandsRequest) TypeName() string {
	return "getCommands"
}

// TypeInfo returns info about TL type.
func (g *GetCommandsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getCommands",
		ID:   GetCommandsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Scope",
			SchemaName: "scope",
		},
		{
			Name:       "LanguageCode",
			SchemaName: "language_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetCommandsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCommands#58ba8ff7 as nil")
	}
	b.PutID(GetCommandsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetCommandsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCommands#58ba8ff7 as nil")
	}
	if g.Scope == nil {
		return fmt.Errorf("unable to encode getCommands#58ba8ff7: field scope is nil")
	}
	if err := g.Scope.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getCommands#58ba8ff7: field scope: %w", err)
	}
	b.PutString(g.LanguageCode)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetCommandsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCommands#58ba8ff7 to nil")
	}
	if err := b.ConsumeID(GetCommandsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getCommands#58ba8ff7: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetCommandsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCommands#58ba8ff7 to nil")
	}
	{
		value, err := DecodeBotCommandScope(b)
		if err != nil {
			return fmt.Errorf("unable to decode getCommands#58ba8ff7: field scope: %w", err)
		}
		g.Scope = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getCommands#58ba8ff7: field language_code: %w", err)
		}
		g.LanguageCode = value
	}
	return nil
}

// GetScope returns value of Scope field.
func (g *GetCommandsRequest) GetScope() (value BotCommandScopeClass) {
	return g.Scope
}

// GetLanguageCode returns value of LanguageCode field.
func (g *GetCommandsRequest) GetLanguageCode() (value string) {
	return g.LanguageCode
}

// GetCommands invokes method getCommands#58ba8ff7 returning error if any.
func (c *Client) GetCommands(ctx context.Context, request *GetCommandsRequest) (*BotCommands, error) {
	var result BotCommands

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
