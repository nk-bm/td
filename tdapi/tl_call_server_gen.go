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

// CallServer represents TL type `callServer#6f37df97`.
type CallServer struct {
	// Server identifier
	ID Int64
	// Server IPv4 address
	IPAddress string
	// Server IPv6 address
	Ipv6Address string
	// Server port number
	Port int32
	// Server type
	Type CallServerTypeClass
}

// CallServerTypeID is TL type id of CallServer.
const CallServerTypeID = 0x6f37df97

// Ensuring interfaces in compile-time for CallServer.
var (
	_ bin.Encoder     = &CallServer{}
	_ bin.Decoder     = &CallServer{}
	_ bin.BareEncoder = &CallServer{}
	_ bin.BareDecoder = &CallServer{}
)

func (c *CallServer) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ID.Zero()) {
		return false
	}
	if !(c.IPAddress == "") {
		return false
	}
	if !(c.Ipv6Address == "") {
		return false
	}
	if !(c.Port == 0) {
		return false
	}
	if !(c.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CallServer) String() string {
	if c == nil {
		return "CallServer(nil)"
	}
	type Alias CallServer
	return fmt.Sprintf("CallServer%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CallServer) TypeID() uint32 {
	return CallServerTypeID
}

// TypeName returns name of type in TL schema.
func (*CallServer) TypeName() string {
	return "callServer"
}

// TypeInfo returns info about TL type.
func (c *CallServer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "callServer",
		ID:   CallServerTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "IPAddress",
			SchemaName: "ip_address",
		},
		{
			Name:       "Ipv6Address",
			SchemaName: "ipv6_address",
		},
		{
			Name:       "Port",
			SchemaName: "port",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CallServer) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callServer#6f37df97 as nil")
	}
	b.PutID(CallServerTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CallServer) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callServer#6f37df97 as nil")
	}
	if err := c.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode callServer#6f37df97: field id: %w", err)
	}
	b.PutString(c.IPAddress)
	b.PutString(c.Ipv6Address)
	b.PutInt32(c.Port)
	if c.Type == nil {
		return fmt.Errorf("unable to encode callServer#6f37df97: field type is nil")
	}
	if err := c.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode callServer#6f37df97: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CallServer) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callServer#6f37df97 to nil")
	}
	if err := b.ConsumeID(CallServerTypeID); err != nil {
		return fmt.Errorf("unable to decode callServer#6f37df97: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CallServer) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callServer#6f37df97 to nil")
	}
	{
		if err := c.ID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode callServer#6f37df97: field id: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode callServer#6f37df97: field ip_address: %w", err)
		}
		c.IPAddress = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode callServer#6f37df97: field ipv6_address: %w", err)
		}
		c.Ipv6Address = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode callServer#6f37df97: field port: %w", err)
		}
		c.Port = value
	}
	{
		value, err := DecodeCallServerType(b)
		if err != nil {
			return fmt.Errorf("unable to decode callServer#6f37df97: field type: %w", err)
		}
		c.Type = value
	}
	return nil
}

// GetID returns value of ID field.
func (c *CallServer) GetID() (value Int64) {
	return c.ID
}

// GetIPAddress returns value of IPAddress field.
func (c *CallServer) GetIPAddress() (value string) {
	return c.IPAddress
}

// GetIpv6Address returns value of Ipv6Address field.
func (c *CallServer) GetIpv6Address() (value string) {
	return c.Ipv6Address
}

// GetPort returns value of Port field.
func (c *CallServer) GetPort() (value int32) {
	return c.Port
}

// GetType returns value of Type field.
func (c *CallServer) GetType() (value CallServerTypeClass) {
	return c.Type
}
