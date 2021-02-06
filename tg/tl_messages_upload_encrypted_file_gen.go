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

// MessagesUploadEncryptedFileRequest represents TL type `messages.uploadEncryptedFile#5057c497`.
// Upload encrypted file and associate it to a secret chat
//
// See https://core.telegram.org/method/messages.uploadEncryptedFile for reference.
type MessagesUploadEncryptedFileRequest struct {
	// The secret chat to associate the file to
	Peer InputEncryptedChat
	// The file
	File InputEncryptedFileClass
}

// MessagesUploadEncryptedFileRequestTypeID is TL type id of MessagesUploadEncryptedFileRequest.
const MessagesUploadEncryptedFileRequestTypeID = 0x5057c497

func (u *MessagesUploadEncryptedFileRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Peer.Zero()) {
		return false
	}
	if !(u.File == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *MessagesUploadEncryptedFileRequest) String() string {
	if u == nil {
		return "MessagesUploadEncryptedFileRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesUploadEncryptedFileRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(u.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tFile: ")
	sb.WriteString(fmt.Sprint(u.File))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (u *MessagesUploadEncryptedFileRequest) TypeID() uint32 {
	return MessagesUploadEncryptedFileRequestTypeID
}

// Encode implements bin.Encoder.
func (u *MessagesUploadEncryptedFileRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.uploadEncryptedFile#5057c497 as nil")
	}
	b.PutID(MessagesUploadEncryptedFileRequestTypeID)
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.uploadEncryptedFile#5057c497: field peer: %w", err)
	}
	if u.File == nil {
		return fmt.Errorf("unable to encode messages.uploadEncryptedFile#5057c497: field file is nil")
	}
	if err := u.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.uploadEncryptedFile#5057c497: field file: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (u *MessagesUploadEncryptedFileRequest) GetPeer() (value InputEncryptedChat) {
	return u.Peer
}

// GetFile returns value of File field.
func (u *MessagesUploadEncryptedFileRequest) GetFile() (value InputEncryptedFileClass) {
	return u.File
}

// Decode implements bin.Decoder.
func (u *MessagesUploadEncryptedFileRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.uploadEncryptedFile#5057c497 to nil")
	}
	if err := b.ConsumeID(MessagesUploadEncryptedFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.uploadEncryptedFile#5057c497: %w", err)
	}
	{
		if err := u.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.uploadEncryptedFile#5057c497: field peer: %w", err)
		}
	}
	{
		value, err := DecodeInputEncryptedFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.uploadEncryptedFile#5057c497: field file: %w", err)
		}
		u.File = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesUploadEncryptedFileRequest.
var (
	_ bin.Encoder = &MessagesUploadEncryptedFileRequest{}
	_ bin.Decoder = &MessagesUploadEncryptedFileRequest{}
)

// MessagesUploadEncryptedFile invokes method messages.uploadEncryptedFile#5057c497 returning error if any.
// Upload encrypted file and associate it to a secret chat
//
// See https://core.telegram.org/method/messages.uploadEncryptedFile for reference.
func (c *Client) MessagesUploadEncryptedFile(ctx context.Context, request *MessagesUploadEncryptedFileRequest) (EncryptedFileClass, error) {
	var result EncryptedFileBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.EncryptedFile, nil
}
