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

// MessagesEditExportedChatInviteRequest represents TL type `messages.editExportedChatInvite#bdca2f75`.
type MessagesEditExportedChatInviteRequest struct {
	// Flags field of MessagesEditExportedChatInviteRequest.
	Flags bin.Fields
	// Revoked field of MessagesEditExportedChatInviteRequest.
	Revoked bool
	// Peer field of MessagesEditExportedChatInviteRequest.
	Peer InputPeerClass
	// Link field of MessagesEditExportedChatInviteRequest.
	Link string
	// ExpireDate field of MessagesEditExportedChatInviteRequest.
	//
	// Use SetExpireDate and GetExpireDate helpers.
	ExpireDate int
	// UsageLimit field of MessagesEditExportedChatInviteRequest.
	//
	// Use SetUsageLimit and GetUsageLimit helpers.
	UsageLimit int
	// RequestNeeded field of MessagesEditExportedChatInviteRequest.
	//
	// Use SetRequestNeeded and GetRequestNeeded helpers.
	RequestNeeded bool
	// Title field of MessagesEditExportedChatInviteRequest.
	//
	// Use SetTitle and GetTitle helpers.
	Title string
}

// MessagesEditExportedChatInviteRequestTypeID is TL type id of MessagesEditExportedChatInviteRequest.
const MessagesEditExportedChatInviteRequestTypeID = 0xbdca2f75

// Ensuring interfaces in compile-time for MessagesEditExportedChatInviteRequest.
var (
	_ bin.Encoder     = &MessagesEditExportedChatInviteRequest{}
	_ bin.Decoder     = &MessagesEditExportedChatInviteRequest{}
	_ bin.BareEncoder = &MessagesEditExportedChatInviteRequest{}
	_ bin.BareDecoder = &MessagesEditExportedChatInviteRequest{}
)

func (e *MessagesEditExportedChatInviteRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.Revoked == false) {
		return false
	}
	if !(e.Peer == nil) {
		return false
	}
	if !(e.Link == "") {
		return false
	}
	if !(e.ExpireDate == 0) {
		return false
	}
	if !(e.UsageLimit == 0) {
		return false
	}
	if !(e.RequestNeeded == false) {
		return false
	}
	if !(e.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesEditExportedChatInviteRequest) String() string {
	if e == nil {
		return "MessagesEditExportedChatInviteRequest(nil)"
	}
	type Alias MessagesEditExportedChatInviteRequest
	return fmt.Sprintf("MessagesEditExportedChatInviteRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesEditExportedChatInviteRequest) TypeID() uint32 {
	return MessagesEditExportedChatInviteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesEditExportedChatInviteRequest) TypeName() string {
	return "messages.editExportedChatInvite"
}

// TypeInfo returns info about TL type.
func (e *MessagesEditExportedChatInviteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.editExportedChatInvite",
		ID:   MessagesEditExportedChatInviteRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Revoked",
			SchemaName: "revoked",
			Null:       !e.Flags.Has(2),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Link",
			SchemaName: "link",
		},
		{
			Name:       "ExpireDate",
			SchemaName: "expire_date",
			Null:       !e.Flags.Has(0),
		},
		{
			Name:       "UsageLimit",
			SchemaName: "usage_limit",
			Null:       !e.Flags.Has(1),
		},
		{
			Name:       "RequestNeeded",
			SchemaName: "request_needed",
			Null:       !e.Flags.Has(3),
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !e.Flags.Has(4),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (e *MessagesEditExportedChatInviteRequest) SetFlags() {
	if !(e.Revoked == false) {
		e.Flags.Set(2)
	}
	if !(e.ExpireDate == 0) {
		e.Flags.Set(0)
	}
	if !(e.UsageLimit == 0) {
		e.Flags.Set(1)
	}
	if !(e.RequestNeeded == false) {
		e.Flags.Set(3)
	}
	if !(e.Title == "") {
		e.Flags.Set(4)
	}
}

// Encode implements bin.Encoder.
func (e *MessagesEditExportedChatInviteRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editExportedChatInvite#bdca2f75 as nil")
	}
	b.PutID(MessagesEditExportedChatInviteRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *MessagesEditExportedChatInviteRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editExportedChatInvite#bdca2f75 as nil")
	}
	e.SetFlags()
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editExportedChatInvite#bdca2f75: field flags: %w", err)
	}
	if e.Peer == nil {
		return fmt.Errorf("unable to encode messages.editExportedChatInvite#bdca2f75: field peer is nil")
	}
	if err := e.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editExportedChatInvite#bdca2f75: field peer: %w", err)
	}
	b.PutString(e.Link)
	if e.Flags.Has(0) {
		b.PutInt(e.ExpireDate)
	}
	if e.Flags.Has(1) {
		b.PutInt(e.UsageLimit)
	}
	if e.Flags.Has(3) {
		b.PutBool(e.RequestNeeded)
	}
	if e.Flags.Has(4) {
		b.PutString(e.Title)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesEditExportedChatInviteRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editExportedChatInvite#bdca2f75 to nil")
	}
	if err := b.ConsumeID(MessagesEditExportedChatInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editExportedChatInvite#bdca2f75: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *MessagesEditExportedChatInviteRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editExportedChatInvite#bdca2f75 to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.editExportedChatInvite#bdca2f75: field flags: %w", err)
		}
	}
	e.Revoked = e.Flags.Has(2)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editExportedChatInvite#bdca2f75: field peer: %w", err)
		}
		e.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editExportedChatInvite#bdca2f75: field link: %w", err)
		}
		e.Link = value
	}
	if e.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editExportedChatInvite#bdca2f75: field expire_date: %w", err)
		}
		e.ExpireDate = value
	}
	if e.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editExportedChatInvite#bdca2f75: field usage_limit: %w", err)
		}
		e.UsageLimit = value
	}
	if e.Flags.Has(3) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editExportedChatInvite#bdca2f75: field request_needed: %w", err)
		}
		e.RequestNeeded = value
	}
	if e.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editExportedChatInvite#bdca2f75: field title: %w", err)
		}
		e.Title = value
	}
	return nil
}

// SetRevoked sets value of Revoked conditional field.
func (e *MessagesEditExportedChatInviteRequest) SetRevoked(value bool) {
	if value {
		e.Flags.Set(2)
		e.Revoked = true
	} else {
		e.Flags.Unset(2)
		e.Revoked = false
	}
}

// GetRevoked returns value of Revoked conditional field.
func (e *MessagesEditExportedChatInviteRequest) GetRevoked() (value bool) {
	if e == nil {
		return
	}
	return e.Flags.Has(2)
}

// GetPeer returns value of Peer field.
func (e *MessagesEditExportedChatInviteRequest) GetPeer() (value InputPeerClass) {
	if e == nil {
		return
	}
	return e.Peer
}

// GetLink returns value of Link field.
func (e *MessagesEditExportedChatInviteRequest) GetLink() (value string) {
	if e == nil {
		return
	}
	return e.Link
}

// SetExpireDate sets value of ExpireDate conditional field.
func (e *MessagesEditExportedChatInviteRequest) SetExpireDate(value int) {
	e.Flags.Set(0)
	e.ExpireDate = value
}

// GetExpireDate returns value of ExpireDate conditional field and
// boolean which is true if field was set.
func (e *MessagesEditExportedChatInviteRequest) GetExpireDate() (value int, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(0) {
		return value, false
	}
	return e.ExpireDate, true
}

// SetUsageLimit sets value of UsageLimit conditional field.
func (e *MessagesEditExportedChatInviteRequest) SetUsageLimit(value int) {
	e.Flags.Set(1)
	e.UsageLimit = value
}

// GetUsageLimit returns value of UsageLimit conditional field and
// boolean which is true if field was set.
func (e *MessagesEditExportedChatInviteRequest) GetUsageLimit() (value int, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(1) {
		return value, false
	}
	return e.UsageLimit, true
}

// SetRequestNeeded sets value of RequestNeeded conditional field.
func (e *MessagesEditExportedChatInviteRequest) SetRequestNeeded(value bool) {
	e.Flags.Set(3)
	e.RequestNeeded = value
}

// GetRequestNeeded returns value of RequestNeeded conditional field and
// boolean which is true if field was set.
func (e *MessagesEditExportedChatInviteRequest) GetRequestNeeded() (value bool, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(3) {
		return value, false
	}
	return e.RequestNeeded, true
}

// SetTitle sets value of Title conditional field.
func (e *MessagesEditExportedChatInviteRequest) SetTitle(value string) {
	e.Flags.Set(4)
	e.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (e *MessagesEditExportedChatInviteRequest) GetTitle() (value string, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(4) {
		return value, false
	}
	return e.Title, true
}

// MessagesEditExportedChatInvite invokes method messages.editExportedChatInvite#bdca2f75 returning error if any.
func (c *Client) MessagesEditExportedChatInvite(ctx context.Context, request *MessagesEditExportedChatInviteRequest) (MessagesExportedChatInviteClass, error) {
	var result MessagesExportedChatInviteBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ExportedChatInvite, nil
}
