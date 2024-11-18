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

func IsPtsUpdate(u UpdateClass) (pts, ptsCount int, ok bool) {
	switch u := u.(type) {
	case *UpdateNewMessage:
		return u.Pts, u.PtsCount, true
	case *UpdateDeleteMessages:
		return u.Pts, u.PtsCount, true
	case *UpdateReadHistoryInbox:
		return u.Pts, u.PtsCount, true
	case *UpdateReadHistoryOutbox:
		return u.Pts, u.PtsCount, true
	case *UpdateWebPage:
		return u.Pts, u.PtsCount, true
	case *UpdateReadMessagesContents:
		return u.Pts, u.PtsCount, true
	case *UpdateEditMessage:
		return u.Pts, u.PtsCount, true
	case *UpdateFolderPeers:
		return u.Pts, u.PtsCount, true
	case *UpdatePinnedMessages:
		return u.Pts, u.PtsCount, true
	}

	return
}

func IsQtsUpdate(u UpdateClass) (qts int, ok bool) {
	switch u := u.(type) {
	case *UpdateNewEncryptedMessage:
		return u.Qts, true
	case *UpdateMessagePollVote:
		return u.Qts, true
	case *UpdateChatParticipant:
		return u.Qts, true
	case *UpdateChannelParticipant:
		return u.Qts, true
	case *UpdateBotStopped:
		return u.Qts, true
	case *UpdateBotChatInviteRequester:
		return u.Qts, true
	case *UpdateBotChatBoost:
		return u.Qts, true
	case *UpdateBotMessageReaction:
		return u.Qts, true
	case *UpdateBotMessageReactions:
		return u.Qts, true
	case *UpdateBotBusinessConnect:
		return u.Qts, true
	case *UpdateBotNewBusinessMessage:
		return u.Qts, true
	case *UpdateBotEditBusinessMessage:
		return u.Qts, true
	case *UpdateBotDeleteBusinessMessage:
		return u.Qts, true
	case *UpdateBotPurchasedPaidMedia:
		return u.Qts, true
	case *UpdateBotSubscriptionExpire:
		return u.Qts, true
	}

	return
}

func IsChannelPtsUpdate(u UpdateClass) (channelID int64, pts, ptsCount int, ok bool, err error) {
	switch u := u.(type) {
	case *UpdateChannelTooLong:
		return u.ChannelID, u.Pts, 0, true, nil
	case *UpdateNewChannelMessage:
		channelID, err = extractChannelID(u.Message)
		return channelID, u.Pts, u.PtsCount, true, err
	case *UpdateReadChannelInbox:
		return u.ChannelID, u.Pts, 0, true, nil
	case *UpdateDeleteChannelMessages:
		return u.ChannelID, u.Pts, u.PtsCount, true, nil
	case *UpdateEditChannelMessage:
		channelID, err = extractChannelID(u.Message)
		return channelID, u.Pts, u.PtsCount, true, err
	case *UpdateChannelWebPage:
		return u.ChannelID, u.Pts, u.PtsCount, true, nil
	case *UpdatePinnedChannelMessages:
		return u.ChannelID, u.Pts, u.PtsCount, true, nil
	}

	return
}

func extractChannelID(msg MessageClass) (int64, error) {
	switch msg := msg.(type) {
	case *MessageEmpty:
		peer, ok := msg.GetPeerID()
		if !ok {
			return 0, errors.New("MessageEmpty have no peerID field")
		}
		if c, ok := peer.(*PeerChannel); ok {
			return c.ChannelID, nil
		}

		return 0, errors.New("unexpected peer type")
	case *Message:
		peer := msg.PeerID

		if c, ok := peer.(*PeerChannel); ok {
			return c.ChannelID, nil
		}

		return 0, errors.New("unexpected peer type")
	case *MessageService:
		peer := msg.PeerID

		if c, ok := peer.(*PeerChannel); ok {
			return c.ChannelID, nil
		}

		return 0, errors.New("unexpected peer type")
	default:
		return 0, errors.New("unexpected MessageClass type")
	}
}
