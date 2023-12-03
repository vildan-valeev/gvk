package gvk

import (
	"encoding/json"
	"github.com/SevereCloud/vksdk/v2/events"
)

type GroupEvent struct {
	*events.GroupEvent
	ParsedObj *events.MessageNewObject
	//obj *events.MessageNewObject
	//obj *events.MessageNewObject
	//obj *events.MessageNewObject
	//obj *events.MessageNewObject

}

// Handler group event handler.
func (g *GroupEvent) Handle() error { //nolint:gocyclo
	//ctx = context.WithValue(ctx, internal.GroupIDKey, g.GroupID)
	//ctx = context.WithValue(ctx, internal.EventIDKey, g.EventID)
	//ctx = context.WithValue(ctx, internal.EventVersionKey, g.V)

	switch g.Type {
	case events.EventMessageNew:
		var obj events.MessageNewObject
		if err := json.Unmarshal(g.Object, &obj); err != nil {
			return err
		}
		g.ParsedObj = &obj
	case events.EventMessageReply:
		var obj events.MessageReplyObject
		if err := json.Unmarshal(g.Object, &obj); err != nil {
			return err
		}

	}

	return nil
}

//
//// ChatID returns the ID of the chat the update is coming from.
//func (u Update) ChatID() int64 {
//	switch {
//	case u.Message != nil:
//		return u.Message.Chat.ID
//	case u.EditedMessage != nil:
//		return u.EditedMessage.Chat.ID
//	case u.ChannelPost != nil:
//		return u.ChannelPost.Chat.ID
//	case u.EditedChannelPost != nil:
//		return u.EditedChannelPost.Chat.ID
//	case u.InlineQuery != nil:
//		return u.InlineQuery.From.ID
//	case u.ChosenInlineResult != nil:
//		return u.ChosenInlineResult.From.ID
//	case u.CallbackQuery != nil:
//		return u.CallbackQuery.Message.Chat.ID
//	case u.ShippingQuery != nil:
//		return u.ShippingQuery.From.ID
//	case u.PreCheckoutQuery != nil:
//		return u.PreCheckoutQuery.From.ID
//	case u.MyChatMember != nil:
//		return u.MyChatMember.Chat.ID
//	case u.ChatMember != nil:
//		return u.ChatMember.Chat.ID
//	case u.ChatJoinRequest != nil:
//		return u.ChatJoinRequest.Chat.ID
//	default:
//		return 0
//	}
//}
