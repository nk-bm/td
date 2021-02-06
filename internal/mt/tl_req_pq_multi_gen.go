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

// ReqPqMultiRequest represents TL type `req_pq_multi#be7e8ef1`.
type ReqPqMultiRequest struct {
	// Nonce field of ReqPqMultiRequest.
	Nonce bin.Int128
}

// ReqPqMultiRequestTypeID is TL type id of ReqPqMultiRequest.
const ReqPqMultiRequestTypeID = 0xbe7e8ef1

func (r *ReqPqMultiRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Nonce == bin.Int128{}) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReqPqMultiRequest) String() string {
	if r == nil {
		return "ReqPqMultiRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ReqPqMultiRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tNonce: ")
	sb.WriteString(fmt.Sprint(r.Nonce))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *ReqPqMultiRequest) TypeID() uint32 {
	return ReqPqMultiRequestTypeID
}

// Encode implements bin.Encoder.
func (r *ReqPqMultiRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode req_pq_multi#be7e8ef1 as nil")
	}
	b.PutID(ReqPqMultiRequestTypeID)
	b.PutInt128(r.Nonce)
	return nil
}

// GetNonce returns value of Nonce field.
func (r *ReqPqMultiRequest) GetNonce() (value bin.Int128) {
	return r.Nonce
}

// Decode implements bin.Decoder.
func (r *ReqPqMultiRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode req_pq_multi#be7e8ef1 to nil")
	}
	if err := b.ConsumeID(ReqPqMultiRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode req_pq_multi#be7e8ef1: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode req_pq_multi#be7e8ef1: field nonce: %w", err)
		}
		r.Nonce = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ReqPqMultiRequest.
var (
	_ bin.Encoder = &ReqPqMultiRequest{}
	_ bin.Decoder = &ReqPqMultiRequest{}
)

// ReqPqMulti invokes method req_pq_multi#be7e8ef1 returning error if any.
func (c *Client) ReqPqMulti(ctx context.Context, nonce bin.Int128) (*ResPQ, error) {
	var result ResPQ

	request := &ReqPqMultiRequest{
		Nonce: nonce,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
