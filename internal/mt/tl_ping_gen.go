// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// PingRequest represents TL type `ping#7abe77ec`.
type PingRequest struct {
	// PingID field of PingRequest.
	PingID int64
}

// PingRequestTypeID is TL type id of PingRequest.
const PingRequestTypeID = 0x7abe77ec

func (p *PingRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.PingID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PingRequest) String() string {
	if p == nil {
		return "PingRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PingRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPingID: ")
	sb.WriteString(fmt.Sprint(p.PingID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PingRequest) TypeID() uint32 {
	return PingRequestTypeID
}

// Encode implements bin.Encoder.
func (p *PingRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode ping#7abe77ec as nil")
	}
	b.PutID(PingRequestTypeID)
	b.PutLong(p.PingID)
	return nil
}

// GetPingID returns value of PingID field.
func (p *PingRequest) GetPingID() (value int64) {
	return p.PingID
}

// Decode implements bin.Decoder.
func (p *PingRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode ping#7abe77ec to nil")
	}
	if err := b.ConsumeID(PingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode ping#7abe77ec: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode ping#7abe77ec: field ping_id: %w", err)
		}
		p.PingID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PingRequest.
var (
	_ bin.Encoder = &PingRequest{}
	_ bin.Decoder = &PingRequest{}
)

// Ping invokes method ping#7abe77ec returning error if any.
func (c *Client) Ping(ctx context.Context, pingid int64) (*Pong, error) {
	var result Pong

	request := &PingRequest{
		PingID: pingid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
