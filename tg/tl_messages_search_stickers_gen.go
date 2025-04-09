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

// MessagesSearchStickersRequest represents TL type `messages.searchStickers#29b1c66a`.
type MessagesSearchStickersRequest struct {
	// Flags field of MessagesSearchStickersRequest.
	Flags bin.Fields
	// Emojis field of MessagesSearchStickersRequest.
	Emojis bool
	// Q field of MessagesSearchStickersRequest.
	Q string
	// Emoticon field of MessagesSearchStickersRequest.
	Emoticon string
	// LangCode field of MessagesSearchStickersRequest.
	LangCode []string
	// Offset field of MessagesSearchStickersRequest.
	Offset int
	// Limit field of MessagesSearchStickersRequest.
	Limit int
	// Hash field of MessagesSearchStickersRequest.
	Hash int64
}

// MessagesSearchStickersRequestTypeID is TL type id of MessagesSearchStickersRequest.
const MessagesSearchStickersRequestTypeID = 0x29b1c66a

// Ensuring interfaces in compile-time for MessagesSearchStickersRequest.
var (
	_ bin.Encoder     = &MessagesSearchStickersRequest{}
	_ bin.Decoder     = &MessagesSearchStickersRequest{}
	_ bin.BareEncoder = &MessagesSearchStickersRequest{}
	_ bin.BareDecoder = &MessagesSearchStickersRequest{}
)

func (s *MessagesSearchStickersRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Emojis == false) {
		return false
	}
	if !(s.Q == "") {
		return false
	}
	if !(s.Emoticon == "") {
		return false
	}
	if !(s.LangCode == nil) {
		return false
	}
	if !(s.Offset == 0) {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}
	if !(s.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSearchStickersRequest) String() string {
	if s == nil {
		return "MessagesSearchStickersRequest(nil)"
	}
	type Alias MessagesSearchStickersRequest
	return fmt.Sprintf("MessagesSearchStickersRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSearchStickersRequest) TypeID() uint32 {
	return MessagesSearchStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSearchStickersRequest) TypeName() string {
	return "messages.searchStickers"
}

// TypeInfo returns info about TL type.
func (s *MessagesSearchStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.searchStickers",
		ID:   MessagesSearchStickersRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Emojis",
			SchemaName: "emojis",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Q",
			SchemaName: "q",
		},
		{
			Name:       "Emoticon",
			SchemaName: "emoticon",
		},
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSearchStickersRequest) SetFlags() {
	if !(s.Emojis == false) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSearchStickersRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.searchStickers#29b1c66a as nil")
	}
	b.PutID(MessagesSearchStickersRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSearchStickersRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.searchStickers#29b1c66a as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.searchStickers#29b1c66a: field flags: %w", err)
	}
	b.PutString(s.Q)
	b.PutString(s.Emoticon)
	b.PutVectorHeader(len(s.LangCode))
	for _, v := range s.LangCode {
		b.PutString(v)
	}
	b.PutInt(s.Offset)
	b.PutInt(s.Limit)
	b.PutLong(s.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSearchStickersRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.searchStickers#29b1c66a to nil")
	}
	if err := b.ConsumeID(MessagesSearchStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.searchStickers#29b1c66a: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSearchStickersRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.searchStickers#29b1c66a to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.searchStickers#29b1c66a: field flags: %w", err)
		}
	}
	s.Emojis = s.Flags.Has(0)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchStickers#29b1c66a: field q: %w", err)
		}
		s.Q = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchStickers#29b1c66a: field emoticon: %w", err)
		}
		s.Emoticon = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchStickers#29b1c66a: field lang_code: %w", err)
		}

		if headerLen > 0 {
			s.LangCode = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode messages.searchStickers#29b1c66a: field lang_code: %w", err)
			}
			s.LangCode = append(s.LangCode, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchStickers#29b1c66a: field offset: %w", err)
		}
		s.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchStickers#29b1c66a: field limit: %w", err)
		}
		s.Limit = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchStickers#29b1c66a: field hash: %w", err)
		}
		s.Hash = value
	}
	return nil
}

// SetEmojis sets value of Emojis conditional field.
func (s *MessagesSearchStickersRequest) SetEmojis(value bool) {
	if value {
		s.Flags.Set(0)
		s.Emojis = true
	} else {
		s.Flags.Unset(0)
		s.Emojis = false
	}
}

// GetEmojis returns value of Emojis conditional field.
func (s *MessagesSearchStickersRequest) GetEmojis() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// GetQ returns value of Q field.
func (s *MessagesSearchStickersRequest) GetQ() (value string) {
	if s == nil {
		return
	}
	return s.Q
}

// GetEmoticon returns value of Emoticon field.
func (s *MessagesSearchStickersRequest) GetEmoticon() (value string) {
	if s == nil {
		return
	}
	return s.Emoticon
}

// GetLangCode returns value of LangCode field.
func (s *MessagesSearchStickersRequest) GetLangCode() (value []string) {
	if s == nil {
		return
	}
	return s.LangCode
}

// GetOffset returns value of Offset field.
func (s *MessagesSearchStickersRequest) GetOffset() (value int) {
	if s == nil {
		return
	}
	return s.Offset
}

// GetLimit returns value of Limit field.
func (s *MessagesSearchStickersRequest) GetLimit() (value int) {
	if s == nil {
		return
	}
	return s.Limit
}

// GetHash returns value of Hash field.
func (s *MessagesSearchStickersRequest) GetHash() (value int64) {
	if s == nil {
		return
	}
	return s.Hash
}

// MessagesSearchStickers invokes method messages.searchStickers#29b1c66a returning error if any.
func (c *Client) MessagesSearchStickers(ctx context.Context, request *MessagesSearchStickersRequest) (MessagesFoundStickersClass, error) {
	var result MessagesFoundStickersBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.FoundStickers, nil
}
