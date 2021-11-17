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

// InlineKeyboardButtonTypeURL represents TL type `inlineKeyboardButtonTypeUrl#4365beac`.
type InlineKeyboardButtonTypeURL struct {
	// HTTP or tg:// URL to open
	URL string
}

// InlineKeyboardButtonTypeURLTypeID is TL type id of InlineKeyboardButtonTypeURL.
const InlineKeyboardButtonTypeURLTypeID = 0x4365beac

// construct implements constructor of InlineKeyboardButtonTypeClass.
func (i InlineKeyboardButtonTypeURL) construct() InlineKeyboardButtonTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineKeyboardButtonTypeURL.
var (
	_ bin.Encoder     = &InlineKeyboardButtonTypeURL{}
	_ bin.Decoder     = &InlineKeyboardButtonTypeURL{}
	_ bin.BareEncoder = &InlineKeyboardButtonTypeURL{}
	_ bin.BareDecoder = &InlineKeyboardButtonTypeURL{}

	_ InlineKeyboardButtonTypeClass = &InlineKeyboardButtonTypeURL{}
)

func (i *InlineKeyboardButtonTypeURL) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineKeyboardButtonTypeURL) String() string {
	if i == nil {
		return "InlineKeyboardButtonTypeURL(nil)"
	}
	type Alias InlineKeyboardButtonTypeURL
	return fmt.Sprintf("InlineKeyboardButtonTypeURL%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineKeyboardButtonTypeURL) TypeID() uint32 {
	return InlineKeyboardButtonTypeURLTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineKeyboardButtonTypeURL) TypeName() string {
	return "inlineKeyboardButtonTypeUrl"
}

// TypeInfo returns info about TL type.
func (i *InlineKeyboardButtonTypeURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineKeyboardButtonTypeUrl",
		ID:   InlineKeyboardButtonTypeURLTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineKeyboardButtonTypeURL) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeUrl#4365beac as nil")
	}
	b.PutID(InlineKeyboardButtonTypeURLTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineKeyboardButtonTypeURL) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeUrl#4365beac as nil")
	}
	b.PutString(i.URL)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineKeyboardButtonTypeURL) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeUrl#4365beac to nil")
	}
	if err := b.ConsumeID(InlineKeyboardButtonTypeURLTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineKeyboardButtonTypeUrl#4365beac: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineKeyboardButtonTypeURL) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeUrl#4365beac to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineKeyboardButtonTypeUrl#4365beac: field url: %w", err)
		}
		i.URL = value
	}
	return nil
}

// GetURL returns value of URL field.
func (i *InlineKeyboardButtonTypeURL) GetURL() (value string) {
	return i.URL
}

// InlineKeyboardButtonTypeLoginURL represents TL type `inlineKeyboardButtonTypeLoginUrl#10c65d93`.
type InlineKeyboardButtonTypeLoginURL struct {
	// An HTTP URL to open
	URL string
	// Unique button identifier
	ID int32
	// If non-empty, new text of the button in forwarded messages
	ForwardText string
}

// InlineKeyboardButtonTypeLoginURLTypeID is TL type id of InlineKeyboardButtonTypeLoginURL.
const InlineKeyboardButtonTypeLoginURLTypeID = 0x10c65d93

// construct implements constructor of InlineKeyboardButtonTypeClass.
func (i InlineKeyboardButtonTypeLoginURL) construct() InlineKeyboardButtonTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineKeyboardButtonTypeLoginURL.
var (
	_ bin.Encoder     = &InlineKeyboardButtonTypeLoginURL{}
	_ bin.Decoder     = &InlineKeyboardButtonTypeLoginURL{}
	_ bin.BareEncoder = &InlineKeyboardButtonTypeLoginURL{}
	_ bin.BareDecoder = &InlineKeyboardButtonTypeLoginURL{}

	_ InlineKeyboardButtonTypeClass = &InlineKeyboardButtonTypeLoginURL{}
)

func (i *InlineKeyboardButtonTypeLoginURL) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.URL == "") {
		return false
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.ForwardText == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineKeyboardButtonTypeLoginURL) String() string {
	if i == nil {
		return "InlineKeyboardButtonTypeLoginURL(nil)"
	}
	type Alias InlineKeyboardButtonTypeLoginURL
	return fmt.Sprintf("InlineKeyboardButtonTypeLoginURL%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineKeyboardButtonTypeLoginURL) TypeID() uint32 {
	return InlineKeyboardButtonTypeLoginURLTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineKeyboardButtonTypeLoginURL) TypeName() string {
	return "inlineKeyboardButtonTypeLoginUrl"
}

// TypeInfo returns info about TL type.
func (i *InlineKeyboardButtonTypeLoginURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineKeyboardButtonTypeLoginUrl",
		ID:   InlineKeyboardButtonTypeLoginURLTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "ForwardText",
			SchemaName: "forward_text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineKeyboardButtonTypeLoginURL) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeLoginUrl#10c65d93 as nil")
	}
	b.PutID(InlineKeyboardButtonTypeLoginURLTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineKeyboardButtonTypeLoginURL) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeLoginUrl#10c65d93 as nil")
	}
	b.PutString(i.URL)
	b.PutInt32(i.ID)
	b.PutString(i.ForwardText)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineKeyboardButtonTypeLoginURL) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeLoginUrl#10c65d93 to nil")
	}
	if err := b.ConsumeID(InlineKeyboardButtonTypeLoginURLTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineKeyboardButtonTypeLoginUrl#10c65d93: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineKeyboardButtonTypeLoginURL) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeLoginUrl#10c65d93 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineKeyboardButtonTypeLoginUrl#10c65d93: field url: %w", err)
		}
		i.URL = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inlineKeyboardButtonTypeLoginUrl#10c65d93: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineKeyboardButtonTypeLoginUrl#10c65d93: field forward_text: %w", err)
		}
		i.ForwardText = value
	}
	return nil
}

// GetURL returns value of URL field.
func (i *InlineKeyboardButtonTypeLoginURL) GetURL() (value string) {
	return i.URL
}

// GetID returns value of ID field.
func (i *InlineKeyboardButtonTypeLoginURL) GetID() (value int32) {
	return i.ID
}

// GetForwardText returns value of ForwardText field.
func (i *InlineKeyboardButtonTypeLoginURL) GetForwardText() (value string) {
	return i.ForwardText
}

// InlineKeyboardButtonTypeCallback represents TL type `inlineKeyboardButtonTypeCallback#bccb7bfd`.
type InlineKeyboardButtonTypeCallback struct {
	// Data to be sent to the bot via a callback query
	Data []byte
}

// InlineKeyboardButtonTypeCallbackTypeID is TL type id of InlineKeyboardButtonTypeCallback.
const InlineKeyboardButtonTypeCallbackTypeID = 0xbccb7bfd

// construct implements constructor of InlineKeyboardButtonTypeClass.
func (i InlineKeyboardButtonTypeCallback) construct() InlineKeyboardButtonTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineKeyboardButtonTypeCallback.
var (
	_ bin.Encoder     = &InlineKeyboardButtonTypeCallback{}
	_ bin.Decoder     = &InlineKeyboardButtonTypeCallback{}
	_ bin.BareEncoder = &InlineKeyboardButtonTypeCallback{}
	_ bin.BareDecoder = &InlineKeyboardButtonTypeCallback{}

	_ InlineKeyboardButtonTypeClass = &InlineKeyboardButtonTypeCallback{}
)

func (i *InlineKeyboardButtonTypeCallback) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Data == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineKeyboardButtonTypeCallback) String() string {
	if i == nil {
		return "InlineKeyboardButtonTypeCallback(nil)"
	}
	type Alias InlineKeyboardButtonTypeCallback
	return fmt.Sprintf("InlineKeyboardButtonTypeCallback%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineKeyboardButtonTypeCallback) TypeID() uint32 {
	return InlineKeyboardButtonTypeCallbackTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineKeyboardButtonTypeCallback) TypeName() string {
	return "inlineKeyboardButtonTypeCallback"
}

// TypeInfo returns info about TL type.
func (i *InlineKeyboardButtonTypeCallback) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineKeyboardButtonTypeCallback",
		ID:   InlineKeyboardButtonTypeCallbackTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineKeyboardButtonTypeCallback) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeCallback#bccb7bfd as nil")
	}
	b.PutID(InlineKeyboardButtonTypeCallbackTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineKeyboardButtonTypeCallback) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeCallback#bccb7bfd as nil")
	}
	b.PutBytes(i.Data)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineKeyboardButtonTypeCallback) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeCallback#bccb7bfd to nil")
	}
	if err := b.ConsumeID(InlineKeyboardButtonTypeCallbackTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineKeyboardButtonTypeCallback#bccb7bfd: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineKeyboardButtonTypeCallback) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeCallback#bccb7bfd to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inlineKeyboardButtonTypeCallback#bccb7bfd: field data: %w", err)
		}
		i.Data = value
	}
	return nil
}

// GetData returns value of Data field.
func (i *InlineKeyboardButtonTypeCallback) GetData() (value []byte) {
	return i.Data
}

// InlineKeyboardButtonTypeCallbackWithPassword represents TL type `inlineKeyboardButtonTypeCallbackWithPassword#361f4248`.
type InlineKeyboardButtonTypeCallbackWithPassword struct {
	// Data to be sent to the bot via a callback query
	Data []byte
}

// InlineKeyboardButtonTypeCallbackWithPasswordTypeID is TL type id of InlineKeyboardButtonTypeCallbackWithPassword.
const InlineKeyboardButtonTypeCallbackWithPasswordTypeID = 0x361f4248

// construct implements constructor of InlineKeyboardButtonTypeClass.
func (i InlineKeyboardButtonTypeCallbackWithPassword) construct() InlineKeyboardButtonTypeClass {
	return &i
}

// Ensuring interfaces in compile-time for InlineKeyboardButtonTypeCallbackWithPassword.
var (
	_ bin.Encoder     = &InlineKeyboardButtonTypeCallbackWithPassword{}
	_ bin.Decoder     = &InlineKeyboardButtonTypeCallbackWithPassword{}
	_ bin.BareEncoder = &InlineKeyboardButtonTypeCallbackWithPassword{}
	_ bin.BareDecoder = &InlineKeyboardButtonTypeCallbackWithPassword{}

	_ InlineKeyboardButtonTypeClass = &InlineKeyboardButtonTypeCallbackWithPassword{}
)

func (i *InlineKeyboardButtonTypeCallbackWithPassword) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Data == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineKeyboardButtonTypeCallbackWithPassword) String() string {
	if i == nil {
		return "InlineKeyboardButtonTypeCallbackWithPassword(nil)"
	}
	type Alias InlineKeyboardButtonTypeCallbackWithPassword
	return fmt.Sprintf("InlineKeyboardButtonTypeCallbackWithPassword%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineKeyboardButtonTypeCallbackWithPassword) TypeID() uint32 {
	return InlineKeyboardButtonTypeCallbackWithPasswordTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineKeyboardButtonTypeCallbackWithPassword) TypeName() string {
	return "inlineKeyboardButtonTypeCallbackWithPassword"
}

// TypeInfo returns info about TL type.
func (i *InlineKeyboardButtonTypeCallbackWithPassword) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineKeyboardButtonTypeCallbackWithPassword",
		ID:   InlineKeyboardButtonTypeCallbackWithPasswordTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineKeyboardButtonTypeCallbackWithPassword) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeCallbackWithPassword#361f4248 as nil")
	}
	b.PutID(InlineKeyboardButtonTypeCallbackWithPasswordTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineKeyboardButtonTypeCallbackWithPassword) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeCallbackWithPassword#361f4248 as nil")
	}
	b.PutBytes(i.Data)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineKeyboardButtonTypeCallbackWithPassword) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeCallbackWithPassword#361f4248 to nil")
	}
	if err := b.ConsumeID(InlineKeyboardButtonTypeCallbackWithPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineKeyboardButtonTypeCallbackWithPassword#361f4248: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineKeyboardButtonTypeCallbackWithPassword) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeCallbackWithPassword#361f4248 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inlineKeyboardButtonTypeCallbackWithPassword#361f4248: field data: %w", err)
		}
		i.Data = value
	}
	return nil
}

// GetData returns value of Data field.
func (i *InlineKeyboardButtonTypeCallbackWithPassword) GetData() (value []byte) {
	return i.Data
}

// InlineKeyboardButtonTypeCallbackGame represents TL type `inlineKeyboardButtonTypeCallbackGame#e9255468`.
type InlineKeyboardButtonTypeCallbackGame struct {
}

// InlineKeyboardButtonTypeCallbackGameTypeID is TL type id of InlineKeyboardButtonTypeCallbackGame.
const InlineKeyboardButtonTypeCallbackGameTypeID = 0xe9255468

// construct implements constructor of InlineKeyboardButtonTypeClass.
func (i InlineKeyboardButtonTypeCallbackGame) construct() InlineKeyboardButtonTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineKeyboardButtonTypeCallbackGame.
var (
	_ bin.Encoder     = &InlineKeyboardButtonTypeCallbackGame{}
	_ bin.Decoder     = &InlineKeyboardButtonTypeCallbackGame{}
	_ bin.BareEncoder = &InlineKeyboardButtonTypeCallbackGame{}
	_ bin.BareDecoder = &InlineKeyboardButtonTypeCallbackGame{}

	_ InlineKeyboardButtonTypeClass = &InlineKeyboardButtonTypeCallbackGame{}
)

func (i *InlineKeyboardButtonTypeCallbackGame) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineKeyboardButtonTypeCallbackGame) String() string {
	if i == nil {
		return "InlineKeyboardButtonTypeCallbackGame(nil)"
	}
	type Alias InlineKeyboardButtonTypeCallbackGame
	return fmt.Sprintf("InlineKeyboardButtonTypeCallbackGame%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineKeyboardButtonTypeCallbackGame) TypeID() uint32 {
	return InlineKeyboardButtonTypeCallbackGameTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineKeyboardButtonTypeCallbackGame) TypeName() string {
	return "inlineKeyboardButtonTypeCallbackGame"
}

// TypeInfo returns info about TL type.
func (i *InlineKeyboardButtonTypeCallbackGame) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineKeyboardButtonTypeCallbackGame",
		ID:   InlineKeyboardButtonTypeCallbackGameTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineKeyboardButtonTypeCallbackGame) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeCallbackGame#e9255468 as nil")
	}
	b.PutID(InlineKeyboardButtonTypeCallbackGameTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineKeyboardButtonTypeCallbackGame) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeCallbackGame#e9255468 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineKeyboardButtonTypeCallbackGame) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeCallbackGame#e9255468 to nil")
	}
	if err := b.ConsumeID(InlineKeyboardButtonTypeCallbackGameTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineKeyboardButtonTypeCallbackGame#e9255468: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineKeyboardButtonTypeCallbackGame) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeCallbackGame#e9255468 to nil")
	}
	return nil
}

// InlineKeyboardButtonTypeSwitchInline represents TL type `inlineKeyboardButtonTypeSwitchInline#86abc4d5`.
type InlineKeyboardButtonTypeSwitchInline struct {
	// Inline query to be sent to the bot
	Query string
	// True, if the inline query should be sent from the current chat
	InCurrentChat bool
}

// InlineKeyboardButtonTypeSwitchInlineTypeID is TL type id of InlineKeyboardButtonTypeSwitchInline.
const InlineKeyboardButtonTypeSwitchInlineTypeID = 0x86abc4d5

// construct implements constructor of InlineKeyboardButtonTypeClass.
func (i InlineKeyboardButtonTypeSwitchInline) construct() InlineKeyboardButtonTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineKeyboardButtonTypeSwitchInline.
var (
	_ bin.Encoder     = &InlineKeyboardButtonTypeSwitchInline{}
	_ bin.Decoder     = &InlineKeyboardButtonTypeSwitchInline{}
	_ bin.BareEncoder = &InlineKeyboardButtonTypeSwitchInline{}
	_ bin.BareDecoder = &InlineKeyboardButtonTypeSwitchInline{}

	_ InlineKeyboardButtonTypeClass = &InlineKeyboardButtonTypeSwitchInline{}
)

func (i *InlineKeyboardButtonTypeSwitchInline) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Query == "") {
		return false
	}
	if !(i.InCurrentChat == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineKeyboardButtonTypeSwitchInline) String() string {
	if i == nil {
		return "InlineKeyboardButtonTypeSwitchInline(nil)"
	}
	type Alias InlineKeyboardButtonTypeSwitchInline
	return fmt.Sprintf("InlineKeyboardButtonTypeSwitchInline%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineKeyboardButtonTypeSwitchInline) TypeID() uint32 {
	return InlineKeyboardButtonTypeSwitchInlineTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineKeyboardButtonTypeSwitchInline) TypeName() string {
	return "inlineKeyboardButtonTypeSwitchInline"
}

// TypeInfo returns info about TL type.
func (i *InlineKeyboardButtonTypeSwitchInline) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineKeyboardButtonTypeSwitchInline",
		ID:   InlineKeyboardButtonTypeSwitchInlineTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Query",
			SchemaName: "query",
		},
		{
			Name:       "InCurrentChat",
			SchemaName: "in_current_chat",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineKeyboardButtonTypeSwitchInline) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeSwitchInline#86abc4d5 as nil")
	}
	b.PutID(InlineKeyboardButtonTypeSwitchInlineTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineKeyboardButtonTypeSwitchInline) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeSwitchInline#86abc4d5 as nil")
	}
	b.PutString(i.Query)
	b.PutBool(i.InCurrentChat)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineKeyboardButtonTypeSwitchInline) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeSwitchInline#86abc4d5 to nil")
	}
	if err := b.ConsumeID(InlineKeyboardButtonTypeSwitchInlineTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineKeyboardButtonTypeSwitchInline#86abc4d5: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineKeyboardButtonTypeSwitchInline) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeSwitchInline#86abc4d5 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineKeyboardButtonTypeSwitchInline#86abc4d5: field query: %w", err)
		}
		i.Query = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode inlineKeyboardButtonTypeSwitchInline#86abc4d5: field in_current_chat: %w", err)
		}
		i.InCurrentChat = value
	}
	return nil
}

// GetQuery returns value of Query field.
func (i *InlineKeyboardButtonTypeSwitchInline) GetQuery() (value string) {
	return i.Query
}

// GetInCurrentChat returns value of InCurrentChat field.
func (i *InlineKeyboardButtonTypeSwitchInline) GetInCurrentChat() (value bool) {
	return i.InCurrentChat
}

// InlineKeyboardButtonTypeBuy represents TL type `inlineKeyboardButtonTypeBuy#511b3c70`.
type InlineKeyboardButtonTypeBuy struct {
}

// InlineKeyboardButtonTypeBuyTypeID is TL type id of InlineKeyboardButtonTypeBuy.
const InlineKeyboardButtonTypeBuyTypeID = 0x511b3c70

// construct implements constructor of InlineKeyboardButtonTypeClass.
func (i InlineKeyboardButtonTypeBuy) construct() InlineKeyboardButtonTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineKeyboardButtonTypeBuy.
var (
	_ bin.Encoder     = &InlineKeyboardButtonTypeBuy{}
	_ bin.Decoder     = &InlineKeyboardButtonTypeBuy{}
	_ bin.BareEncoder = &InlineKeyboardButtonTypeBuy{}
	_ bin.BareDecoder = &InlineKeyboardButtonTypeBuy{}

	_ InlineKeyboardButtonTypeClass = &InlineKeyboardButtonTypeBuy{}
)

func (i *InlineKeyboardButtonTypeBuy) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineKeyboardButtonTypeBuy) String() string {
	if i == nil {
		return "InlineKeyboardButtonTypeBuy(nil)"
	}
	type Alias InlineKeyboardButtonTypeBuy
	return fmt.Sprintf("InlineKeyboardButtonTypeBuy%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineKeyboardButtonTypeBuy) TypeID() uint32 {
	return InlineKeyboardButtonTypeBuyTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineKeyboardButtonTypeBuy) TypeName() string {
	return "inlineKeyboardButtonTypeBuy"
}

// TypeInfo returns info about TL type.
func (i *InlineKeyboardButtonTypeBuy) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineKeyboardButtonTypeBuy",
		ID:   InlineKeyboardButtonTypeBuyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineKeyboardButtonTypeBuy) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeBuy#511b3c70 as nil")
	}
	b.PutID(InlineKeyboardButtonTypeBuyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineKeyboardButtonTypeBuy) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineKeyboardButtonTypeBuy#511b3c70 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineKeyboardButtonTypeBuy) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeBuy#511b3c70 to nil")
	}
	if err := b.ConsumeID(InlineKeyboardButtonTypeBuyTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineKeyboardButtonTypeBuy#511b3c70: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineKeyboardButtonTypeBuy) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineKeyboardButtonTypeBuy#511b3c70 to nil")
	}
	return nil
}

// InlineKeyboardButtonTypeClass represents InlineKeyboardButtonType generic type.
//
// Example:
//  g, err := tdapi.DecodeInlineKeyboardButtonType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.InlineKeyboardButtonTypeURL: // inlineKeyboardButtonTypeUrl#4365beac
//  case *tdapi.InlineKeyboardButtonTypeLoginURL: // inlineKeyboardButtonTypeLoginUrl#10c65d93
//  case *tdapi.InlineKeyboardButtonTypeCallback: // inlineKeyboardButtonTypeCallback#bccb7bfd
//  case *tdapi.InlineKeyboardButtonTypeCallbackWithPassword: // inlineKeyboardButtonTypeCallbackWithPassword#361f4248
//  case *tdapi.InlineKeyboardButtonTypeCallbackGame: // inlineKeyboardButtonTypeCallbackGame#e9255468
//  case *tdapi.InlineKeyboardButtonTypeSwitchInline: // inlineKeyboardButtonTypeSwitchInline#86abc4d5
//  case *tdapi.InlineKeyboardButtonTypeBuy: // inlineKeyboardButtonTypeBuy#511b3c70
//  default: panic(v)
//  }
type InlineKeyboardButtonTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InlineKeyboardButtonTypeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInlineKeyboardButtonType implements binary de-serialization for InlineKeyboardButtonTypeClass.
func DecodeInlineKeyboardButtonType(buf *bin.Buffer) (InlineKeyboardButtonTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InlineKeyboardButtonTypeURLTypeID:
		// Decoding inlineKeyboardButtonTypeUrl#4365beac.
		v := InlineKeyboardButtonTypeURL{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineKeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case InlineKeyboardButtonTypeLoginURLTypeID:
		// Decoding inlineKeyboardButtonTypeLoginUrl#10c65d93.
		v := InlineKeyboardButtonTypeLoginURL{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineKeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case InlineKeyboardButtonTypeCallbackTypeID:
		// Decoding inlineKeyboardButtonTypeCallback#bccb7bfd.
		v := InlineKeyboardButtonTypeCallback{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineKeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case InlineKeyboardButtonTypeCallbackWithPasswordTypeID:
		// Decoding inlineKeyboardButtonTypeCallbackWithPassword#361f4248.
		v := InlineKeyboardButtonTypeCallbackWithPassword{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineKeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case InlineKeyboardButtonTypeCallbackGameTypeID:
		// Decoding inlineKeyboardButtonTypeCallbackGame#e9255468.
		v := InlineKeyboardButtonTypeCallbackGame{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineKeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case InlineKeyboardButtonTypeSwitchInlineTypeID:
		// Decoding inlineKeyboardButtonTypeSwitchInline#86abc4d5.
		v := InlineKeyboardButtonTypeSwitchInline{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineKeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case InlineKeyboardButtonTypeBuyTypeID:
		// Decoding inlineKeyboardButtonTypeBuy#511b3c70.
		v := InlineKeyboardButtonTypeBuy{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineKeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InlineKeyboardButtonTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// InlineKeyboardButtonType boxes the InlineKeyboardButtonTypeClass providing a helper.
type InlineKeyboardButtonTypeBox struct {
	InlineKeyboardButtonType InlineKeyboardButtonTypeClass
}

// Decode implements bin.Decoder for InlineKeyboardButtonTypeBox.
func (b *InlineKeyboardButtonTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InlineKeyboardButtonTypeBox to nil")
	}
	v, err := DecodeInlineKeyboardButtonType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InlineKeyboardButtonType = v
	return nil
}

// Encode implements bin.Encode for InlineKeyboardButtonTypeBox.
func (b *InlineKeyboardButtonTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InlineKeyboardButtonType == nil {
		return fmt.Errorf("unable to encode InlineKeyboardButtonTypeClass as nil")
	}
	return b.InlineKeyboardButtonType.Encode(buf)
}
