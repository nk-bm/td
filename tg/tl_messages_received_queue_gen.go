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

// MessagesReceivedQueueRequest represents TL type `messages.receivedQueue#55a5bb66`.
// Confirms receipt of messages in a secret chat by client, cancels push notifications.
//
// See https://core.telegram.org/method/messages.receivedQueue for reference.
type MessagesReceivedQueueRequest struct {
	// Maximum qts value available at the client
	MaxQts int
}

// MessagesReceivedQueueRequestTypeID is TL type id of MessagesReceivedQueueRequest.
const MessagesReceivedQueueRequestTypeID = 0x55a5bb66

func (r *MessagesReceivedQueueRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.MaxQts == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReceivedQueueRequest) String() string {
	if r == nil {
		return "MessagesReceivedQueueRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesReceivedQueueRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tMaxQts: ")
	sb.WriteString(fmt.Sprint(r.MaxQts))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *MessagesReceivedQueueRequest) TypeID() uint32 {
	return MessagesReceivedQueueRequestTypeID
}

// Encode implements bin.Encoder.
func (r *MessagesReceivedQueueRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.receivedQueue#55a5bb66 as nil")
	}
	b.PutID(MessagesReceivedQueueRequestTypeID)
	b.PutInt(r.MaxQts)
	return nil
}

// GetMaxQts returns value of MaxQts field.
func (r *MessagesReceivedQueueRequest) GetMaxQts() (value int) {
	return r.MaxQts
}

// Decode implements bin.Decoder.
func (r *MessagesReceivedQueueRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.receivedQueue#55a5bb66 to nil")
	}
	if err := b.ConsumeID(MessagesReceivedQueueRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.receivedQueue#55a5bb66: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.receivedQueue#55a5bb66: field max_qts: %w", err)
		}
		r.MaxQts = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesReceivedQueueRequest.
var (
	_ bin.Encoder = &MessagesReceivedQueueRequest{}
	_ bin.Decoder = &MessagesReceivedQueueRequest{}
)

// MessagesReceivedQueue invokes method messages.receivedQueue#55a5bb66 returning error if any.
// Confirms receipt of messages in a secret chat by client, cancels push notifications.
//
// Possible errors:
//  400 MSG_WAIT_FAILED: A waiting call returned an error
//
// See https://core.telegram.org/method/messages.receivedQueue for reference.
func (c *Client) MessagesReceivedQueue(ctx context.Context, maxqts int) ([]int64, error) {
	var result LongVector

	request := &MessagesReceivedQueueRequest{
		MaxQts: maxqts,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}
