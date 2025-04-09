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

// SMSJobsLeaveRequest represents TL type `smsjobs.leave#9898ad73`.
type SMSJobsLeaveRequest struct {
}

// SMSJobsLeaveRequestTypeID is TL type id of SMSJobsLeaveRequest.
const SMSJobsLeaveRequestTypeID = 0x9898ad73

// Ensuring interfaces in compile-time for SMSJobsLeaveRequest.
var (
	_ bin.Encoder     = &SMSJobsLeaveRequest{}
	_ bin.Decoder     = &SMSJobsLeaveRequest{}
	_ bin.BareEncoder = &SMSJobsLeaveRequest{}
	_ bin.BareDecoder = &SMSJobsLeaveRequest{}
)

func (l *SMSJobsLeaveRequest) Zero() bool {
	if l == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (l *SMSJobsLeaveRequest) String() string {
	if l == nil {
		return "SMSJobsLeaveRequest(nil)"
	}
	type Alias SMSJobsLeaveRequest
	return fmt.Sprintf("SMSJobsLeaveRequest%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SMSJobsLeaveRequest) TypeID() uint32 {
	return SMSJobsLeaveRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SMSJobsLeaveRequest) TypeName() string {
	return "smsjobs.leave"
}

// TypeInfo returns info about TL type.
func (l *SMSJobsLeaveRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "smsjobs.leave",
		ID:   SMSJobsLeaveRequestTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (l *SMSJobsLeaveRequest) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode smsjobs.leave#9898ad73 as nil")
	}
	b.PutID(SMSJobsLeaveRequestTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *SMSJobsLeaveRequest) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode smsjobs.leave#9898ad73 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *SMSJobsLeaveRequest) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode smsjobs.leave#9898ad73 to nil")
	}
	if err := b.ConsumeID(SMSJobsLeaveRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode smsjobs.leave#9898ad73: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *SMSJobsLeaveRequest) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode smsjobs.leave#9898ad73 to nil")
	}
	return nil
}

// SMSJobsLeave invokes method smsjobs.leave#9898ad73 returning error if any.
func (c *Client) SMSJobsLeave(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &SMSJobsLeaveRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
