package gvk

import (
	"encoding/json"
	"fmt"
)

type Response interface {
	Base() APIError
}

//---------------------------------------------------------------

type APIResponseMessagesSend struct {
	Error    APIError `json:"error,omitempty"`
	Response int      `json:"response,omitempty"`
}

func (a APIResponseMessagesSend) Base() APIError {
	return a.Error
}

//---------------------------------------------------------------

type APIResponseMessagesEdit struct {
	Error    APIError `json:"error,omitempty"`
	Response int      `json:"response,omitempty"`
}

func (a APIResponseMessagesEdit) Base() APIError {
	return a.Error
}

// ------------------------------------------------------------------

type APIResponseMessageEventAnswer struct {
	Error    APIError `json:"error,omitempty"`
	Response int      `json:"response,omitempty"`
}

func (a APIResponseMessageEventAnswer) Base() APIError {
	return a.Error
}

// ------------------------------------------------------------

type APIResponseUsersGet struct {
	Error    APIError `json:"error,omitempty"`
	Response []User   `json:"response,omitempty"`
}

type User struct {
	ID              int64       `json:"id"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	ScreenName      string      `json:"screen_name"` // aka username
	Deactivated     Deactivated `json:"deactivated"`
	IsClosed        bool        `json:"is_closed"`
	CanAccessClosed bool        `json:"can_access_closed"`
}

type Deactivated string

const DeactivatedBanned Deactivated = "banned"
const DeactivatedDeleted Deactivated = "deleted"

func (a APIResponseUsersGet) Base() APIError {
	return a.Error
}

//---------------------------------------------------------------

type APIResponseGetLongPollServer struct {
	Response ResponseGetLongPollServer `json:"response,omitempty"`
	Error    APIError                  `json:"error,omitempty"`
}

type ResponseGetLongPollServer struct {
	Key    string `json:"key,omitempty"`
	Server string `json:"server,omitempty"`
	Ts     string `json:"ts,omitempty"`
}

func (a APIResponseGetLongPollServer) Base() APIError {
	return a.Error
}

// ---------------------------------------------------------------

type APIResponseSetLongPollSettings struct {
	Response ResponseSetLongPollSettings `json:"response,omitempty"`
	Error    APIError                    `json:"error,omitempty"`
}

type ResponseSetLongPollSettings struct {
}

func (a APIResponseSetLongPollSettings) Base() APIError {
	return a.Error
}

//---------------------------------------------------------------

type APIResponseGroupsIsMember struct {
	Response int64    `json:"response,omitempty"`
	Error    APIError `json:"error,omitempty"`
}

func (a APIResponseGroupsIsMember) Base() APIError {
	return a.Error
}

// ------ https://dev.vk.com/ru/method/wall.post#Результат -------------

type APIResponseWallPost struct {
	Error    APIError         `json:"error,omitempty"`
	Response ResponseWallPost `json:"response,omitempty"`
}

type ResponseWallPost struct {
	PostID int64 `json:"post_id"`
}

func (a APIResponseWallPost) Base() APIError {
	return a.Error
}

// -------------------------

type APIResponseUpdate struct {
	Ts string `json:"ts,omitempty"`
	//Updates []*Update `json:"updates,omitempty"`
	Updates RawUpdates `json:"updates,omitempty"`
	Failed  int64      `json:"failed,omitempty"` // ошибка
	Error   APIError   `json:"error,omitempty"`
}

func (a APIResponseUpdate) Base() APIError {
	return a.Error
}

type TypeSwitch struct {
	Type    EventType `json:"type"`
	GroupID int       `json:"group_id"`
	EventID string    `json:"event_id"`
	V       string    `json:"v"`
}

type RawUpdates []json.RawMessage

type Updates []*Update

type Update struct {
	TypeSwitch

	Object Object `json:"object"`
}

type Object struct {
	*MessageNew
	*MessageReply
	*MessageEvent
	*MessageEdit
	*WallPostNew
	// TODO: добавить остальные объекты Events
}

//---------------------------------------------------------------

// ChatID returns the ID of the chat the update is coming from.
func (u *Update) ChatID() int64 {
	switch {
	case u.Object.MessageNew != nil:
		return u.Object.MessageNew.Message.FromID
	case u.Object.MessageEvent != nil:
		return u.Object.MessageEvent.UserID
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

func (a RawUpdates) UnmarshalCustom() (Updates, error) {
	//if len(a) > 0 {
	//	fmt.Println("COUNT UPDATES", len(a))
	//}

	var updates Updates
	for _, row := range a {
		update := Update{}

		if err := json.Unmarshal(row, &update); err != nil {
			return updates, err
		}
		updates = append(updates, &update)
	}

	return updates, nil
}

func (u *Update) UnmarshalJSON(data []byte) error {
	var temp struct {
		TypeSwitch
		Object json.RawMessage `json:"object"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	u.TypeSwitch = temp.TypeSwitch

	switch u.Type {
	case EventMessageNew:
		u.Object.MessageNew = &MessageNew{}
		return json.Unmarshal(temp.Object, u.Object.MessageNew)
	case EventMessageReply:
		u.Object.MessageReply = &MessageReply{}
		return json.Unmarshal(temp.Object, u.Object.MessageReply)
	case EventMessageEdit:
		u.Object.MessageEdit = &MessageEdit{}
		return json.Unmarshal(temp.Object, u.Object.MessageEdit)
	case EventMessageEvent:
		u.Object.MessageEvent = &MessageEvent{}
		return json.Unmarshal(temp.Object, u.Object.MessageEvent)
	case EventWallPostNew:
		u.Object.WallPostNew = &WallPostNew{}
		return json.Unmarshal(temp.Object, u.Object.WallPostNew)
	default:
		return fmt.Errorf("unrecognized type value %q", u.Type)
	}

}

// EventType type.
type EventType string

// EventType list.
const (
	EventConfirmation                            = "confirmation"
	EventMessageNew                    EventType = "message_new"
	EventMessageReply                  EventType = "message_reply"
	EventMessageEdit                             = "message_edit"
	EventMessageAllow                            = "message_allow"
	EventMessageDeny                             = "message_deny"
	EventMessageTypingState                      = "message_typing_state"
	EventMessageEvent                  EventType = "message_event"
	EventPhotoNew                                = "photo_new"
	EventPhotoCommentNew                         = "photo_comment_new"
	EventPhotoCommentEdit                        = "photo_comment_edit"
	EventPhotoCommentRestore                     = "photo_comment_restore"
	EventPhotoCommentDelete                      = "photo_comment_delete"
	EventAudioNew                                = "audio_new"
	EventVideoNew                                = "video_new"
	EventVideoCommentNew                         = "video_comment_new"
	EventVideoCommentEdit                        = "video_comment_edit"
	EventVideoCommentRestore                     = "video_comment_restore"
	EventVideoCommentDelete                      = "video_comment_delete"
	EventWallPostNew                             = "wall_post_new"
	EventWallRepost                              = "wall_repost"
	EventWallReplyNew                            = "wall_reply_new"
	EventWallReplyEdit                           = "wall_reply_edit"
	EventWallReplyRestore                        = "wall_reply_restore"
	EventWallReplyDelete                         = "wall_reply_delete"
	EventBoardPostNew                            = "board_post_new"
	EventBoardPostEdit                           = "board_post_edit"
	EventBoardPostRestore                        = "board_post_restore"
	EventBoardPostDelete                         = "board_post_delete"
	EventMarketCommentNew                        = "market_comment_new"
	EventMarketCommentEdit                       = "market_comment_edit"
	EventMarketCommentRestore                    = "market_comment_restore"
	EventMarketCommentDelete                     = "market_comment_delete"
	EventMarketOrderNew                          = "market_order_new"
	EventMarketOrderEdit                         = "market_order_edit"
	EventGroupLeave                              = "group_leave"
	EventGroupJoin                               = "group_join"
	EventUserBlock                               = "user_block"
	EventUserUnblock                             = "user_unblock"
	EventPollVoteNew                             = "poll_vote_new"
	EventGroupOfficersEdit                       = "group_officers_edit"
	EventGroupChangeSettings                     = "group_change_settings"
	EventGroupChangePhoto                        = "group_change_photo"
	EventVkpayTransaction                        = "vkpay_transaction"
	EventLeadFormsNew                            = "lead_forms_new"
	EventAppPayload                              = "app_payload"
	EventMessageRead                             = "message_read"
	EventLikeAdd                                 = "like_add"
	EventLikeRemove                              = "like_remove"
	EventDonutSubscriptionCreate                 = "donut_subscription_create"
	EventDonutSubscriptionProlonged              = "donut_subscription_prolonged"
	EventDonutSubscriptionExpired                = "donut_subscription_expired"
	EventDonutSubscriptionCancelled              = "donut_subscription_cancelled"
	EventDonutSubscriptionPriceChanged           = "donut_subscription_price_changed"
	EventDonutMoneyWithdraw                      = "donut_money_withdraw"
	EventDonutMoneyWithdrawError                 = "donut_money_withdraw_error"
)
