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

// GetAutoDownloadSettingsPresetsRequest represents TL type `getAutoDownloadSettingsPresets#996a4737`.
type GetAutoDownloadSettingsPresetsRequest struct {
}

// GetAutoDownloadSettingsPresetsRequestTypeID is TL type id of GetAutoDownloadSettingsPresetsRequest.
const GetAutoDownloadSettingsPresetsRequestTypeID = 0x996a4737

// Ensuring interfaces in compile-time for GetAutoDownloadSettingsPresetsRequest.
var (
	_ bin.Encoder     = &GetAutoDownloadSettingsPresetsRequest{}
	_ bin.Decoder     = &GetAutoDownloadSettingsPresetsRequest{}
	_ bin.BareEncoder = &GetAutoDownloadSettingsPresetsRequest{}
	_ bin.BareDecoder = &GetAutoDownloadSettingsPresetsRequest{}
)

func (g *GetAutoDownloadSettingsPresetsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetAutoDownloadSettingsPresetsRequest) String() string {
	if g == nil {
		return "GetAutoDownloadSettingsPresetsRequest(nil)"
	}
	type Alias GetAutoDownloadSettingsPresetsRequest
	return fmt.Sprintf("GetAutoDownloadSettingsPresetsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetAutoDownloadSettingsPresetsRequest) TypeID() uint32 {
	return GetAutoDownloadSettingsPresetsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetAutoDownloadSettingsPresetsRequest) TypeName() string {
	return "getAutoDownloadSettingsPresets"
}

// TypeInfo returns info about TL type.
func (g *GetAutoDownloadSettingsPresetsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getAutoDownloadSettingsPresets",
		ID:   GetAutoDownloadSettingsPresetsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetAutoDownloadSettingsPresetsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getAutoDownloadSettingsPresets#996a4737 as nil")
	}
	b.PutID(GetAutoDownloadSettingsPresetsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetAutoDownloadSettingsPresetsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getAutoDownloadSettingsPresets#996a4737 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetAutoDownloadSettingsPresetsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getAutoDownloadSettingsPresets#996a4737 to nil")
	}
	if err := b.ConsumeID(GetAutoDownloadSettingsPresetsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getAutoDownloadSettingsPresets#996a4737: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetAutoDownloadSettingsPresetsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getAutoDownloadSettingsPresets#996a4737 to nil")
	}
	return nil
}

// GetAutoDownloadSettingsPresets invokes method getAutoDownloadSettingsPresets#996a4737 returning error if any.
func (c *Client) GetAutoDownloadSettingsPresets(ctx context.Context) (*AutoDownloadSettingsPresets, error) {
	var result AutoDownloadSettingsPresets

	request := &GetAutoDownloadSettingsPresetsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
