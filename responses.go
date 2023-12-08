package gvk

type Response interface {
	Base() APIError
}

type ResponseMessagesSend struct {
	APIError     // ошибка
	Response int `json:"response,omitempty"`
}

func (a ResponseMessagesSend) Base() APIError {
	return a.APIError
}

type ResponseGetLongPollServer struct {
	Key    string `json:"key,omitempty"`
	Server string `json:"server,omitempty"`
	Ts     string `json:"ts,omitempty"`
	APIError
}

func (a ResponseGetLongPollServer) Base() APIError {
	return a.APIError
}

type ResponseSetLongPollSettings struct {
	APIError
}

func (a ResponseSetLongPollSettings) Base() APIError {
	return a.APIError
}

type ResponseUpdate struct {
	Ts      string    `json:"ts,omitempty"`
	Updates []*Update `json:"updates,omitempty"`
	Failed  int64     `json:"failed,omitempty"` // ошибка
}

func (a ResponseUpdate) Base() APIError {
	// TODO переделать
	return APIError{Code: a.Failed}
}

type Update struct {
	Type    EventType `json:"type"`
	GroupID int       `json:"group_id"`
	EventID string    `json:"event_id"`
	V       string    `json:"v"`

	MessageNew *MessageNew `json:"object"`
	//TODO: добавить остальные объекты
}

// ChatID returns the ID of the chat the update is coming from.
func (u Update) ChatID() int64 {
	switch {
	case u.MessageNew != nil:
		return u.MessageNew.Message.FromID
	//case u.EditedMessage != nil:
	//	return u.EditedMessage.Chat.ID
	//case u.ChannelPost != nil:
	//	return u.ChannelPost.Chat.ID
	//case u.EditedChannelPost != nil:
	//	return u.EditedChannelPost.Chat.ID
	//case u.InlineQuery != nil:
	//	return u.InlineQuery.From.ID
	//case u.ChosenInlineResult != nil:
	//	return u.ChosenInlineResult.From.ID
	//case u.CallbackQuery != nil:
	//	return u.CallbackQuery.Message.Chat.ID
	//case u.ShippingQuery != nil:
	//	return u.ShippingQuery.From.ID
	//case u.PreCheckoutQuery != nil:
	//	return u.PreCheckoutQuery.From.ID
	//case u.MyChatMember != nil:
	//	return u.MyChatMember.Chat.ID
	//case u.ChatMember != nil:
	//	return u.ChatMember.Chat.ID
	//case u.ChatJoinRequest != nil:
	//	return u.ChatJoinRequest.Chat.ID
	default:
		return 0
	}
}

// EventType type.
type EventType string

// EventType list.
const (
	EventConfirmation                  = "confirmation"
	EventMessageNew                    = "message_new"
	EventMessageReply                  = "message_reply"
	EventMessageEdit                   = "message_edit"
	EventMessageAllow                  = "message_allow"
	EventMessageDeny                   = "message_deny"
	EventMessageTypingState            = "message_typing_state"
	EventMessageEvent                  = "message_event"
	EventPhotoNew                      = "photo_new"
	EventPhotoCommentNew               = "photo_comment_new"
	EventPhotoCommentEdit              = "photo_comment_edit"
	EventPhotoCommentRestore           = "photo_comment_restore"
	EventPhotoCommentDelete            = "photo_comment_delete"
	EventAudioNew                      = "audio_new"
	EventVideoNew                      = "video_new"
	EventVideoCommentNew               = "video_comment_new"
	EventVideoCommentEdit              = "video_comment_edit"
	EventVideoCommentRestore           = "video_comment_restore"
	EventVideoCommentDelete            = "video_comment_delete"
	EventWallPostNew                   = "wall_post_new"
	EventWallRepost                    = "wall_repost"
	EventWallReplyNew                  = "wall_reply_new"
	EventWallReplyEdit                 = "wall_reply_edit"
	EventWallReplyRestore              = "wall_reply_restore"
	EventWallReplyDelete               = "wall_reply_delete"
	EventBoardPostNew                  = "board_post_new"
	EventBoardPostEdit                 = "board_post_edit"
	EventBoardPostRestore              = "board_post_restore"
	EventBoardPostDelete               = "board_post_delete"
	EventMarketCommentNew              = "market_comment_new"
	EventMarketCommentEdit             = "market_comment_edit"
	EventMarketCommentRestore          = "market_comment_restore"
	EventMarketCommentDelete           = "market_comment_delete"
	EventMarketOrderNew                = "market_order_new"
	EventMarketOrderEdit               = "market_order_edit"
	EventGroupLeave                    = "group_leave"
	EventGroupJoin                     = "group_join"
	EventUserBlock                     = "user_block"
	EventUserUnblock                   = "user_unblock"
	EventPollVoteNew                   = "poll_vote_new"
	EventGroupOfficersEdit             = "group_officers_edit"
	EventGroupChangeSettings           = "group_change_settings"
	EventGroupChangePhoto              = "group_change_photo"
	EventVkpayTransaction              = "vkpay_transaction"
	EventLeadFormsNew                  = "lead_forms_new"
	EventAppPayload                    = "app_payload"
	EventMessageRead                   = "message_read"
	EventLikeAdd                       = "like_add"
	EventLikeRemove                    = "like_remove"
	EventDonutSubscriptionCreate       = "donut_subscription_create"
	EventDonutSubscriptionProlonged    = "donut_subscription_prolonged"
	EventDonutSubscriptionExpired      = "donut_subscription_expired"
	EventDonutSubscriptionCancelled    = "donut_subscription_cancelled"
	EventDonutSubscriptionPriceChanged = "donut_subscription_price_changed"
	EventDonutMoneyWithdraw            = "donut_money_withdraw"
	EventDonutMoneyWithdrawError       = "donut_money_withdraw_error"
)
