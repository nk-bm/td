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

// MessagesReportReactionRequest represents TL type `messages.reportReaction#3f64c076`.
type MessagesReportReactionRequest struct {
	// Peer field of MessagesReportReactionRequest.
	Peer InputPeerClass
	// ID field of MessagesReportReactionRequest.
	ID int
	// ReactionPeer field of MessagesReportReactionRequest.
	ReactionPeer InputPeerClass
}

// MessagesReportReactionRequestTypeID is TL type id of MessagesReportReactionRequest.
const MessagesReportReactionRequestTypeID = 0x3f64c076

// Ensuring interfaces in compile-time for MessagesReportReactionRequest.
var (
	_ bin.Encoder     = &MessagesReportReactionRequest{}
	_ bin.Decoder     = &MessagesReportReactionRequest{}
	_ bin.BareEncoder = &MessagesReportReactionRequest{}
	_ bin.BareDecoder = &MessagesReportReactionRequest{}
)

func (r *MessagesReportReactionRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.ID == 0) {
		return false
	}
	if !(r.ReactionPeer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReportReactionRequest) String() string {
	if r == nil {
		return "MessagesReportReactionRequest(nil)"
	}
	type Alias MessagesReportReactionRequest
	return fmt.Sprintf("MessagesReportReactionRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReportReactionRequest) TypeID() uint32 {
	return MessagesReportReactionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReportReactionRequest) TypeName() string {
	return "messages.reportReaction"
}

// TypeInfo returns info about TL type.
func (r *MessagesReportReactionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.reportReaction",
		ID:   MessagesReportReactionRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "ReactionPeer",
			SchemaName: "reaction_peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReportReactionRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reportReaction#3f64c076 as nil")
	}
	b.PutID(MessagesReportReactionRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReportReactionRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reportReaction#3f64c076 as nil")
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.reportReaction#3f64c076: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.reportReaction#3f64c076: field peer: %w", err)
	}
	b.PutInt(r.ID)
	if r.ReactionPeer == nil {
		return fmt.Errorf("unable to encode messages.reportReaction#3f64c076: field reaction_peer is nil")
	}
	if err := r.ReactionPeer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.reportReaction#3f64c076: field reaction_peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReportReactionRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reportReaction#3f64c076 to nil")
	}
	if err := b.ConsumeID(MessagesReportReactionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.reportReaction#3f64c076: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReportReactionRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reportReaction#3f64c076 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.reportReaction#3f64c076: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.reportReaction#3f64c076: field id: %w", err)
		}
		r.ID = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.reportReaction#3f64c076: field reaction_peer: %w", err)
		}
		r.ReactionPeer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *MessagesReportReactionRequest) GetPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// GetID returns value of ID field.
func (r *MessagesReportReactionRequest) GetID() (value int) {
	if r == nil {
		return
	}
	return r.ID
}

// GetReactionPeer returns value of ReactionPeer field.
func (r *MessagesReportReactionRequest) GetReactionPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.ReactionPeer
}

// MessagesReportReaction invokes method messages.reportReaction#3f64c076 returning error if any.
func (c *Client) MessagesReportReaction(ctx context.Context, request *MessagesReportReactionRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
