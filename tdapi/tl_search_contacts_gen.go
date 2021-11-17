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

// SearchContactsRequest represents TL type `searchContacts#95073165`.
type SearchContactsRequest struct {
	// Query to search for; may be empty to return all contacts
	Query string
	// The maximum number of users to be returned
	Limit int32
}

// SearchContactsRequestTypeID is TL type id of SearchContactsRequest.
const SearchContactsRequestTypeID = 0x95073165

// Ensuring interfaces in compile-time for SearchContactsRequest.
var (
	_ bin.Encoder     = &SearchContactsRequest{}
	_ bin.Decoder     = &SearchContactsRequest{}
	_ bin.BareEncoder = &SearchContactsRequest{}
	_ bin.BareDecoder = &SearchContactsRequest{}
)

func (s *SearchContactsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Query == "") {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchContactsRequest) String() string {
	if s == nil {
		return "SearchContactsRequest(nil)"
	}
	type Alias SearchContactsRequest
	return fmt.Sprintf("SearchContactsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchContactsRequest) TypeID() uint32 {
	return SearchContactsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchContactsRequest) TypeName() string {
	return "searchContacts"
}

// TypeInfo returns info about TL type.
func (s *SearchContactsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchContacts",
		ID:   SearchContactsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Query",
			SchemaName: "query",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchContactsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchContacts#95073165 as nil")
	}
	b.PutID(SearchContactsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchContactsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchContacts#95073165 as nil")
	}
	b.PutString(s.Query)
	b.PutInt32(s.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchContactsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchContacts#95073165 to nil")
	}
	if err := b.ConsumeID(SearchContactsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchContacts#95073165: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchContactsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchContacts#95073165 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchContacts#95073165: field query: %w", err)
		}
		s.Query = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchContacts#95073165: field limit: %w", err)
		}
		s.Limit = value
	}
	return nil
}

// GetQuery returns value of Query field.
func (s *SearchContactsRequest) GetQuery() (value string) {
	return s.Query
}

// GetLimit returns value of Limit field.
func (s *SearchContactsRequest) GetLimit() (value int32) {
	return s.Limit
}

// SearchContacts invokes method searchContacts#95073165 returning error if any.
func (c *Client) SearchContacts(ctx context.Context, request *SearchContactsRequest) (*Users, error) {
	var result Users

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
