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

// PeerNotifySettings represents TL type `peerNotifySettings#99622c0c`.
type PeerNotifySettings struct {
	// Flags field of PeerNotifySettings.
	Flags bin.Fields
	// ShowPreviews field of PeerNotifySettings.
	//
	// Use SetShowPreviews and GetShowPreviews helpers.
	ShowPreviews bool
	// Silent field of PeerNotifySettings.
	//
	// Use SetSilent and GetSilent helpers.
	Silent bool
	// MuteUntil field of PeerNotifySettings.
	//
	// Use SetMuteUntil and GetMuteUntil helpers.
	MuteUntil int
	// IosSound field of PeerNotifySettings.
	//
	// Use SetIosSound and GetIosSound helpers.
	IosSound NotificationSoundClass
	// AndroidSound field of PeerNotifySettings.
	//
	// Use SetAndroidSound and GetAndroidSound helpers.
	AndroidSound NotificationSoundClass
	// OtherSound field of PeerNotifySettings.
	//
	// Use SetOtherSound and GetOtherSound helpers.
	OtherSound NotificationSoundClass
	// StoriesMuted field of PeerNotifySettings.
	//
	// Use SetStoriesMuted and GetStoriesMuted helpers.
	StoriesMuted bool
	// StoriesHideSender field of PeerNotifySettings.
	//
	// Use SetStoriesHideSender and GetStoriesHideSender helpers.
	StoriesHideSender bool
	// StoriesIosSound field of PeerNotifySettings.
	//
	// Use SetStoriesIosSound and GetStoriesIosSound helpers.
	StoriesIosSound NotificationSoundClass
	// StoriesAndroidSound field of PeerNotifySettings.
	//
	// Use SetStoriesAndroidSound and GetStoriesAndroidSound helpers.
	StoriesAndroidSound NotificationSoundClass
	// StoriesOtherSound field of PeerNotifySettings.
	//
	// Use SetStoriesOtherSound and GetStoriesOtherSound helpers.
	StoriesOtherSound NotificationSoundClass
}

// PeerNotifySettingsTypeID is TL type id of PeerNotifySettings.
const PeerNotifySettingsTypeID = 0x99622c0c

// Ensuring interfaces in compile-time for PeerNotifySettings.
var (
	_ bin.Encoder     = &PeerNotifySettings{}
	_ bin.Decoder     = &PeerNotifySettings{}
	_ bin.BareEncoder = &PeerNotifySettings{}
	_ bin.BareDecoder = &PeerNotifySettings{}
)

func (p *PeerNotifySettings) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.ShowPreviews == false) {
		return false
	}
	if !(p.Silent == false) {
		return false
	}
	if !(p.MuteUntil == 0) {
		return false
	}
	if !(p.IosSound == nil) {
		return false
	}
	if !(p.AndroidSound == nil) {
		return false
	}
	if !(p.OtherSound == nil) {
		return false
	}
	if !(p.StoriesMuted == false) {
		return false
	}
	if !(p.StoriesHideSender == false) {
		return false
	}
	if !(p.StoriesIosSound == nil) {
		return false
	}
	if !(p.StoriesAndroidSound == nil) {
		return false
	}
	if !(p.StoriesOtherSound == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PeerNotifySettings) String() string {
	if p == nil {
		return "PeerNotifySettings(nil)"
	}
	type Alias PeerNotifySettings
	return fmt.Sprintf("PeerNotifySettings%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PeerNotifySettings) TypeID() uint32 {
	return PeerNotifySettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*PeerNotifySettings) TypeName() string {
	return "peerNotifySettings"
}

// TypeInfo returns info about TL type.
func (p *PeerNotifySettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "peerNotifySettings",
		ID:   PeerNotifySettingsTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShowPreviews",
			SchemaName: "show_previews",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !p.Flags.Has(1),
		},
		{
			Name:       "MuteUntil",
			SchemaName: "mute_until",
			Null:       !p.Flags.Has(2),
		},
		{
			Name:       "IosSound",
			SchemaName: "ios_sound",
			Null:       !p.Flags.Has(3),
		},
		{
			Name:       "AndroidSound",
			SchemaName: "android_sound",
			Null:       !p.Flags.Has(4),
		},
		{
			Name:       "OtherSound",
			SchemaName: "other_sound",
			Null:       !p.Flags.Has(5),
		},
		{
			Name:       "StoriesMuted",
			SchemaName: "stories_muted",
			Null:       !p.Flags.Has(6),
		},
		{
			Name:       "StoriesHideSender",
			SchemaName: "stories_hide_sender",
			Null:       !p.Flags.Has(7),
		},
		{
			Name:       "StoriesIosSound",
			SchemaName: "stories_ios_sound",
			Null:       !p.Flags.Has(8),
		},
		{
			Name:       "StoriesAndroidSound",
			SchemaName: "stories_android_sound",
			Null:       !p.Flags.Has(9),
		},
		{
			Name:       "StoriesOtherSound",
			SchemaName: "stories_other_sound",
			Null:       !p.Flags.Has(10),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (p *PeerNotifySettings) SetFlags() {
	if !(p.ShowPreviews == false) {
		p.Flags.Set(0)
	}
	if !(p.Silent == false) {
		p.Flags.Set(1)
	}
	if !(p.MuteUntil == 0) {
		p.Flags.Set(2)
	}
	if !(p.IosSound == nil) {
		p.Flags.Set(3)
	}
	if !(p.AndroidSound == nil) {
		p.Flags.Set(4)
	}
	if !(p.OtherSound == nil) {
		p.Flags.Set(5)
	}
	if !(p.StoriesMuted == false) {
		p.Flags.Set(6)
	}
	if !(p.StoriesHideSender == false) {
		p.Flags.Set(7)
	}
	if !(p.StoriesIosSound == nil) {
		p.Flags.Set(8)
	}
	if !(p.StoriesAndroidSound == nil) {
		p.Flags.Set(9)
	}
	if !(p.StoriesOtherSound == nil) {
		p.Flags.Set(10)
	}
}

// Encode implements bin.Encoder.
func (p *PeerNotifySettings) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerNotifySettings#99622c0c as nil")
	}
	b.PutID(PeerNotifySettingsTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PeerNotifySettings) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerNotifySettings#99622c0c as nil")
	}
	p.SetFlags()
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field flags: %w", err)
	}
	if p.Flags.Has(0) {
		b.PutBool(p.ShowPreviews)
	}
	if p.Flags.Has(1) {
		b.PutBool(p.Silent)
	}
	if p.Flags.Has(2) {
		b.PutInt(p.MuteUntil)
	}
	if p.Flags.Has(3) {
		if p.IosSound == nil {
			return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field ios_sound is nil")
		}
		if err := p.IosSound.Encode(b); err != nil {
			return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field ios_sound: %w", err)
		}
	}
	if p.Flags.Has(4) {
		if p.AndroidSound == nil {
			return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field android_sound is nil")
		}
		if err := p.AndroidSound.Encode(b); err != nil {
			return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field android_sound: %w", err)
		}
	}
	if p.Flags.Has(5) {
		if p.OtherSound == nil {
			return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field other_sound is nil")
		}
		if err := p.OtherSound.Encode(b); err != nil {
			return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field other_sound: %w", err)
		}
	}
	if p.Flags.Has(6) {
		b.PutBool(p.StoriesMuted)
	}
	if p.Flags.Has(7) {
		b.PutBool(p.StoriesHideSender)
	}
	if p.Flags.Has(8) {
		if p.StoriesIosSound == nil {
			return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field stories_ios_sound is nil")
		}
		if err := p.StoriesIosSound.Encode(b); err != nil {
			return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field stories_ios_sound: %w", err)
		}
	}
	if p.Flags.Has(9) {
		if p.StoriesAndroidSound == nil {
			return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field stories_android_sound is nil")
		}
		if err := p.StoriesAndroidSound.Encode(b); err != nil {
			return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field stories_android_sound: %w", err)
		}
	}
	if p.Flags.Has(10) {
		if p.StoriesOtherSound == nil {
			return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field stories_other_sound is nil")
		}
		if err := p.StoriesOtherSound.Encode(b); err != nil {
			return fmt.Errorf("unable to encode peerNotifySettings#99622c0c: field stories_other_sound: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PeerNotifySettings) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerNotifySettings#99622c0c to nil")
	}
	if err := b.ConsumeID(PeerNotifySettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PeerNotifySettings) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerNotifySettings#99622c0c to nil")
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: field flags: %w", err)
		}
	}
	if p.Flags.Has(0) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: field show_previews: %w", err)
		}
		p.ShowPreviews = value
	}
	if p.Flags.Has(1) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: field silent: %w", err)
		}
		p.Silent = value
	}
	if p.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: field mute_until: %w", err)
		}
		p.MuteUntil = value
	}
	if p.Flags.Has(3) {
		value, err := DecodeNotificationSound(b)
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: field ios_sound: %w", err)
		}
		p.IosSound = value
	}
	if p.Flags.Has(4) {
		value, err := DecodeNotificationSound(b)
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: field android_sound: %w", err)
		}
		p.AndroidSound = value
	}
	if p.Flags.Has(5) {
		value, err := DecodeNotificationSound(b)
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: field other_sound: %w", err)
		}
		p.OtherSound = value
	}
	if p.Flags.Has(6) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: field stories_muted: %w", err)
		}
		p.StoriesMuted = value
	}
	if p.Flags.Has(7) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: field stories_hide_sender: %w", err)
		}
		p.StoriesHideSender = value
	}
	if p.Flags.Has(8) {
		value, err := DecodeNotificationSound(b)
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: field stories_ios_sound: %w", err)
		}
		p.StoriesIosSound = value
	}
	if p.Flags.Has(9) {
		value, err := DecodeNotificationSound(b)
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: field stories_android_sound: %w", err)
		}
		p.StoriesAndroidSound = value
	}
	if p.Flags.Has(10) {
		value, err := DecodeNotificationSound(b)
		if err != nil {
			return fmt.Errorf("unable to decode peerNotifySettings#99622c0c: field stories_other_sound: %w", err)
		}
		p.StoriesOtherSound = value
	}
	return nil
}

// SetShowPreviews sets value of ShowPreviews conditional field.
func (p *PeerNotifySettings) SetShowPreviews(value bool) {
	p.Flags.Set(0)
	p.ShowPreviews = value
}

// GetShowPreviews returns value of ShowPreviews conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetShowPreviews() (value bool, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.ShowPreviews, true
}

// SetSilent sets value of Silent conditional field.
func (p *PeerNotifySettings) SetSilent(value bool) {
	p.Flags.Set(1)
	p.Silent = value
}

// GetSilent returns value of Silent conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetSilent() (value bool, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.Silent, true
}

// SetMuteUntil sets value of MuteUntil conditional field.
func (p *PeerNotifySettings) SetMuteUntil(value int) {
	p.Flags.Set(2)
	p.MuteUntil = value
}

// GetMuteUntil returns value of MuteUntil conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetMuteUntil() (value int, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.MuteUntil, true
}

// SetIosSound sets value of IosSound conditional field.
func (p *PeerNotifySettings) SetIosSound(value NotificationSoundClass) {
	p.Flags.Set(3)
	p.IosSound = value
}

// GetIosSound returns value of IosSound conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetIosSound() (value NotificationSoundClass, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(3) {
		return value, false
	}
	return p.IosSound, true
}

// SetAndroidSound sets value of AndroidSound conditional field.
func (p *PeerNotifySettings) SetAndroidSound(value NotificationSoundClass) {
	p.Flags.Set(4)
	p.AndroidSound = value
}

// GetAndroidSound returns value of AndroidSound conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetAndroidSound() (value NotificationSoundClass, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.AndroidSound, true
}

// SetOtherSound sets value of OtherSound conditional field.
func (p *PeerNotifySettings) SetOtherSound(value NotificationSoundClass) {
	p.Flags.Set(5)
	p.OtherSound = value
}

// GetOtherSound returns value of OtherSound conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetOtherSound() (value NotificationSoundClass, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(5) {
		return value, false
	}
	return p.OtherSound, true
}

// SetStoriesMuted sets value of StoriesMuted conditional field.
func (p *PeerNotifySettings) SetStoriesMuted(value bool) {
	p.Flags.Set(6)
	p.StoriesMuted = value
}

// GetStoriesMuted returns value of StoriesMuted conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetStoriesMuted() (value bool, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(6) {
		return value, false
	}
	return p.StoriesMuted, true
}

// SetStoriesHideSender sets value of StoriesHideSender conditional field.
func (p *PeerNotifySettings) SetStoriesHideSender(value bool) {
	p.Flags.Set(7)
	p.StoriesHideSender = value
}

// GetStoriesHideSender returns value of StoriesHideSender conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetStoriesHideSender() (value bool, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(7) {
		return value, false
	}
	return p.StoriesHideSender, true
}

// SetStoriesIosSound sets value of StoriesIosSound conditional field.
func (p *PeerNotifySettings) SetStoriesIosSound(value NotificationSoundClass) {
	p.Flags.Set(8)
	p.StoriesIosSound = value
}

// GetStoriesIosSound returns value of StoriesIosSound conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetStoriesIosSound() (value NotificationSoundClass, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(8) {
		return value, false
	}
	return p.StoriesIosSound, true
}

// SetStoriesAndroidSound sets value of StoriesAndroidSound conditional field.
func (p *PeerNotifySettings) SetStoriesAndroidSound(value NotificationSoundClass) {
	p.Flags.Set(9)
	p.StoriesAndroidSound = value
}

// GetStoriesAndroidSound returns value of StoriesAndroidSound conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetStoriesAndroidSound() (value NotificationSoundClass, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(9) {
		return value, false
	}
	return p.StoriesAndroidSound, true
}

// SetStoriesOtherSound sets value of StoriesOtherSound conditional field.
func (p *PeerNotifySettings) SetStoriesOtherSound(value NotificationSoundClass) {
	p.Flags.Set(10)
	p.StoriesOtherSound = value
}

// GetStoriesOtherSound returns value of StoriesOtherSound conditional field and
// boolean which is true if field was set.
func (p *PeerNotifySettings) GetStoriesOtherSound() (value NotificationSoundClass, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(10) {
		return value, false
	}
	return p.StoriesOtherSound, true
}
