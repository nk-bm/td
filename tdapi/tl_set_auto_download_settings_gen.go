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

// SetAutoDownloadSettingsRequest represents TL type `setAutoDownloadSettings#eaeb64f4`.
type SetAutoDownloadSettingsRequest struct {
	// New user auto-download settings
	Settings AutoDownloadSettings
	// Type of the network for which the new settings are relevant
	Type NetworkTypeClass
}

// SetAutoDownloadSettingsRequestTypeID is TL type id of SetAutoDownloadSettingsRequest.
const SetAutoDownloadSettingsRequestTypeID = 0xeaeb64f4

// Ensuring interfaces in compile-time for SetAutoDownloadSettingsRequest.
var (
	_ bin.Encoder     = &SetAutoDownloadSettingsRequest{}
	_ bin.Decoder     = &SetAutoDownloadSettingsRequest{}
	_ bin.BareEncoder = &SetAutoDownloadSettingsRequest{}
	_ bin.BareDecoder = &SetAutoDownloadSettingsRequest{}
)

func (s *SetAutoDownloadSettingsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Settings.Zero()) {
		return false
	}
	if !(s.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetAutoDownloadSettingsRequest) String() string {
	if s == nil {
		return "SetAutoDownloadSettingsRequest(nil)"
	}
	type Alias SetAutoDownloadSettingsRequest
	return fmt.Sprintf("SetAutoDownloadSettingsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetAutoDownloadSettingsRequest) TypeID() uint32 {
	return SetAutoDownloadSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetAutoDownloadSettingsRequest) TypeName() string {
	return "setAutoDownloadSettings"
}

// TypeInfo returns info about TL type.
func (s *SetAutoDownloadSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setAutoDownloadSettings",
		ID:   SetAutoDownloadSettingsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetAutoDownloadSettingsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setAutoDownloadSettings#eaeb64f4 as nil")
	}
	b.PutID(SetAutoDownloadSettingsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetAutoDownloadSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setAutoDownloadSettings#eaeb64f4 as nil")
	}
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setAutoDownloadSettings#eaeb64f4: field settings: %w", err)
	}
	if s.Type == nil {
		return fmt.Errorf("unable to encode setAutoDownloadSettings#eaeb64f4: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setAutoDownloadSettings#eaeb64f4: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetAutoDownloadSettingsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setAutoDownloadSettings#eaeb64f4 to nil")
	}
	if err := b.ConsumeID(SetAutoDownloadSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setAutoDownloadSettings#eaeb64f4: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetAutoDownloadSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setAutoDownloadSettings#eaeb64f4 to nil")
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setAutoDownloadSettings#eaeb64f4: field settings: %w", err)
		}
	}
	{
		value, err := DecodeNetworkType(b)
		if err != nil {
			return fmt.Errorf("unable to decode setAutoDownloadSettings#eaeb64f4: field type: %w", err)
		}
		s.Type = value
	}
	return nil
}

// GetSettings returns value of Settings field.
func (s *SetAutoDownloadSettingsRequest) GetSettings() (value AutoDownloadSettings) {
	return s.Settings
}

// GetType returns value of Type field.
func (s *SetAutoDownloadSettingsRequest) GetType() (value NetworkTypeClass) {
	return s.Type
}

// SetAutoDownloadSettings invokes method setAutoDownloadSettings#eaeb64f4 returning error if any.
func (c *Client) SetAutoDownloadSettings(ctx context.Context, request *SetAutoDownloadSettingsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
