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

// RestrictionReason represents TL type `restrictionReason#d072acb4`.
// Restriction reason.
// Contains the reason why access to a certain object must be restricted. Clients are supposed to deny access to the channel if the platform field is equal to all or to the current platform (ios, android, wp, etc.). Platforms can be concatenated (ios-android, ios-wp), unknown platforms are to be ignored. The text is the error message that should be shown to the user.
//
// See https://core.telegram.org/constructor/restrictionReason for reference.
type RestrictionReason struct {
	// Platform identifier (ios, android, wp, all, etc.), can be concatenated with a dash as separator (android-ios, ios-wp, etc)
	Platform string
	// Restriction reason (porno, terms, etc.)
	Reason string
	// Error message to be shown to the user
	Text string
}

// RestrictionReasonTypeID is TL type id of RestrictionReason.
const RestrictionReasonTypeID = 0xd072acb4

func (r *RestrictionReason) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Platform == "") {
		return false
	}
	if !(r.Reason == "") {
		return false
	}
	if !(r.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RestrictionReason) String() string {
	if r == nil {
		return "RestrictionReason(nil)"
	}
	var sb strings.Builder
	sb.WriteString("RestrictionReason")
	sb.WriteString("{\n")
	sb.WriteString("\tPlatform: ")
	sb.WriteString(fmt.Sprint(r.Platform))
	sb.WriteString(",\n")
	sb.WriteString("\tReason: ")
	sb.WriteString(fmt.Sprint(r.Reason))
	sb.WriteString(",\n")
	sb.WriteString("\tText: ")
	sb.WriteString(fmt.Sprint(r.Text))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *RestrictionReason) TypeID() uint32 {
	return RestrictionReasonTypeID
}

// Encode implements bin.Encoder.
func (r *RestrictionReason) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode restrictionReason#d072acb4 as nil")
	}
	b.PutID(RestrictionReasonTypeID)
	b.PutString(r.Platform)
	b.PutString(r.Reason)
	b.PutString(r.Text)
	return nil
}

// GetPlatform returns value of Platform field.
func (r *RestrictionReason) GetPlatform() (value string) {
	return r.Platform
}

// GetReason returns value of Reason field.
func (r *RestrictionReason) GetReason() (value string) {
	return r.Reason
}

// GetText returns value of Text field.
func (r *RestrictionReason) GetText() (value string) {
	return r.Text
}

// Decode implements bin.Decoder.
func (r *RestrictionReason) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode restrictionReason#d072acb4 to nil")
	}
	if err := b.ConsumeID(RestrictionReasonTypeID); err != nil {
		return fmt.Errorf("unable to decode restrictionReason#d072acb4: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode restrictionReason#d072acb4: field platform: %w", err)
		}
		r.Platform = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode restrictionReason#d072acb4: field reason: %w", err)
		}
		r.Reason = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode restrictionReason#d072acb4: field text: %w", err)
		}
		r.Text = value
	}
	return nil
}

// Ensuring interfaces in compile-time for RestrictionReason.
var (
	_ bin.Encoder = &RestrictionReason{}
	_ bin.Decoder = &RestrictionReason{}
)
