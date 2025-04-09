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

// StoriesActivateStealthModeRequest represents TL type `stories.activateStealthMode#57bbd166`.
type StoriesActivateStealthModeRequest struct {
	// Flags field of StoriesActivateStealthModeRequest.
	Flags bin.Fields
	// Past field of StoriesActivateStealthModeRequest.
	Past bool
	// Future field of StoriesActivateStealthModeRequest.
	Future bool
}

// StoriesActivateStealthModeRequestTypeID is TL type id of StoriesActivateStealthModeRequest.
const StoriesActivateStealthModeRequestTypeID = 0x57bbd166

// Ensuring interfaces in compile-time for StoriesActivateStealthModeRequest.
var (
	_ bin.Encoder     = &StoriesActivateStealthModeRequest{}
	_ bin.Decoder     = &StoriesActivateStealthModeRequest{}
	_ bin.BareEncoder = &StoriesActivateStealthModeRequest{}
	_ bin.BareDecoder = &StoriesActivateStealthModeRequest{}
)

func (a *StoriesActivateStealthModeRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.Past == false) {
		return false
	}
	if !(a.Future == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *StoriesActivateStealthModeRequest) String() string {
	if a == nil {
		return "StoriesActivateStealthModeRequest(nil)"
	}
	type Alias StoriesActivateStealthModeRequest
	return fmt.Sprintf("StoriesActivateStealthModeRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesActivateStealthModeRequest) TypeID() uint32 {
	return StoriesActivateStealthModeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesActivateStealthModeRequest) TypeName() string {
	return "stories.activateStealthMode"
}

// TypeInfo returns info about TL type.
func (a *StoriesActivateStealthModeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.activateStealthMode",
		ID:   StoriesActivateStealthModeRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Past",
			SchemaName: "past",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "Future",
			SchemaName: "future",
			Null:       !a.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *StoriesActivateStealthModeRequest) SetFlags() {
	if !(a.Past == false) {
		a.Flags.Set(0)
	}
	if !(a.Future == false) {
		a.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (a *StoriesActivateStealthModeRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stories.activateStealthMode#57bbd166 as nil")
	}
	b.PutID(StoriesActivateStealthModeRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *StoriesActivateStealthModeRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stories.activateStealthMode#57bbd166 as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.activateStealthMode#57bbd166: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *StoriesActivateStealthModeRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stories.activateStealthMode#57bbd166 to nil")
	}
	if err := b.ConsumeID(StoriesActivateStealthModeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.activateStealthMode#57bbd166: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *StoriesActivateStealthModeRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stories.activateStealthMode#57bbd166 to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.activateStealthMode#57bbd166: field flags: %w", err)
		}
	}
	a.Past = a.Flags.Has(0)
	a.Future = a.Flags.Has(1)
	return nil
}

// SetPast sets value of Past conditional field.
func (a *StoriesActivateStealthModeRequest) SetPast(value bool) {
	if value {
		a.Flags.Set(0)
		a.Past = true
	} else {
		a.Flags.Unset(0)
		a.Past = false
	}
}

// GetPast returns value of Past conditional field.
func (a *StoriesActivateStealthModeRequest) GetPast() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(0)
}

// SetFuture sets value of Future conditional field.
func (a *StoriesActivateStealthModeRequest) SetFuture(value bool) {
	if value {
		a.Flags.Set(1)
		a.Future = true
	} else {
		a.Flags.Unset(1)
		a.Future = false
	}
}

// GetFuture returns value of Future conditional field.
func (a *StoriesActivateStealthModeRequest) GetFuture() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(1)
}

// StoriesActivateStealthMode invokes method stories.activateStealthMode#57bbd166 returning error if any.
func (c *Client) StoriesActivateStealthMode(ctx context.Context, request *StoriesActivateStealthModeRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
