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

// LangpackGetStringsRequest represents TL type `langpack.getStrings#efea3803`.
type LangpackGetStringsRequest struct {
	// LangPack field of LangpackGetStringsRequest.
	LangPack string
	// LangCode field of LangpackGetStringsRequest.
	LangCode string
	// Keys field of LangpackGetStringsRequest.
	Keys []string
}

// LangpackGetStringsRequestTypeID is TL type id of LangpackGetStringsRequest.
const LangpackGetStringsRequestTypeID = 0xefea3803

// Ensuring interfaces in compile-time for LangpackGetStringsRequest.
var (
	_ bin.Encoder     = &LangpackGetStringsRequest{}
	_ bin.Decoder     = &LangpackGetStringsRequest{}
	_ bin.BareEncoder = &LangpackGetStringsRequest{}
	_ bin.BareDecoder = &LangpackGetStringsRequest{}
)

func (g *LangpackGetStringsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangPack == "") {
		return false
	}
	if !(g.LangCode == "") {
		return false
	}
	if !(g.Keys == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *LangpackGetStringsRequest) String() string {
	if g == nil {
		return "LangpackGetStringsRequest(nil)"
	}
	type Alias LangpackGetStringsRequest
	return fmt.Sprintf("LangpackGetStringsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LangpackGetStringsRequest) TypeID() uint32 {
	return LangpackGetStringsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*LangpackGetStringsRequest) TypeName() string {
	return "langpack.getStrings"
}

// TypeInfo returns info about TL type.
func (g *LangpackGetStringsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "langpack.getStrings",
		ID:   LangpackGetStringsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangPack",
			SchemaName: "lang_pack",
		},
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
		{
			Name:       "Keys",
			SchemaName: "keys",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *LangpackGetStringsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode langpack.getStrings#efea3803 as nil")
	}
	b.PutID(LangpackGetStringsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *LangpackGetStringsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode langpack.getStrings#efea3803 as nil")
	}
	b.PutString(g.LangPack)
	b.PutString(g.LangCode)
	b.PutVectorHeader(len(g.Keys))
	for _, v := range g.Keys {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *LangpackGetStringsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode langpack.getStrings#efea3803 to nil")
	}
	if err := b.ConsumeID(LangpackGetStringsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode langpack.getStrings#efea3803: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *LangpackGetStringsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode langpack.getStrings#efea3803 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getStrings#efea3803: field lang_pack: %w", err)
		}
		g.LangPack = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getStrings#efea3803: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getStrings#efea3803: field keys: %w", err)
		}

		if headerLen > 0 {
			g.Keys = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode langpack.getStrings#efea3803: field keys: %w", err)
			}
			g.Keys = append(g.Keys, value)
		}
	}
	return nil
}

// GetLangPack returns value of LangPack field.
func (g *LangpackGetStringsRequest) GetLangPack() (value string) {
	if g == nil {
		return
	}
	return g.LangPack
}

// GetLangCode returns value of LangCode field.
func (g *LangpackGetStringsRequest) GetLangCode() (value string) {
	if g == nil {
		return
	}
	return g.LangCode
}

// GetKeys returns value of Keys field.
func (g *LangpackGetStringsRequest) GetKeys() (value []string) {
	if g == nil {
		return
	}
	return g.Keys
}

// LangpackGetStrings invokes method langpack.getStrings#efea3803 returning error if any.
func (c *Client) LangpackGetStrings(ctx context.Context, request *LangpackGetStringsRequest) ([]LangPackStringClass, error) {
	var result LangPackStringClassVector

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []LangPackStringClass(result.Elems), nil
}
