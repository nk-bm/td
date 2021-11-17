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

// TerminateAllOtherSessionsRequest represents TL type `terminateAllOtherSessions#6fba6113`.
type TerminateAllOtherSessionsRequest struct {
}

// TerminateAllOtherSessionsRequestTypeID is TL type id of TerminateAllOtherSessionsRequest.
const TerminateAllOtherSessionsRequestTypeID = 0x6fba6113

// Ensuring interfaces in compile-time for TerminateAllOtherSessionsRequest.
var (
	_ bin.Encoder     = &TerminateAllOtherSessionsRequest{}
	_ bin.Decoder     = &TerminateAllOtherSessionsRequest{}
	_ bin.BareEncoder = &TerminateAllOtherSessionsRequest{}
	_ bin.BareDecoder = &TerminateAllOtherSessionsRequest{}
)

func (t *TerminateAllOtherSessionsRequest) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TerminateAllOtherSessionsRequest) String() string {
	if t == nil {
		return "TerminateAllOtherSessionsRequest(nil)"
	}
	type Alias TerminateAllOtherSessionsRequest
	return fmt.Sprintf("TerminateAllOtherSessionsRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TerminateAllOtherSessionsRequest) TypeID() uint32 {
	return TerminateAllOtherSessionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TerminateAllOtherSessionsRequest) TypeName() string {
	return "terminateAllOtherSessions"
}

// TypeInfo returns info about TL type.
func (t *TerminateAllOtherSessionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "terminateAllOtherSessions",
		ID:   TerminateAllOtherSessionsRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *TerminateAllOtherSessionsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode terminateAllOtherSessions#6fba6113 as nil")
	}
	b.PutID(TerminateAllOtherSessionsRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TerminateAllOtherSessionsRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode terminateAllOtherSessions#6fba6113 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TerminateAllOtherSessionsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode terminateAllOtherSessions#6fba6113 to nil")
	}
	if err := b.ConsumeID(TerminateAllOtherSessionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode terminateAllOtherSessions#6fba6113: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TerminateAllOtherSessionsRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode terminateAllOtherSessions#6fba6113 to nil")
	}
	return nil
}

// TerminateAllOtherSessions invokes method terminateAllOtherSessions#6fba6113 returning error if any.
func (c *Client) TerminateAllOtherSessions(ctx context.Context) error {
	var ok Ok

	request := &TerminateAllOtherSessionsRequest{}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
