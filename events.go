/*
GVK
Copyright (C) 2023-2024 The GVK Devs

GVK is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published
by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

GVK is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package gvk

import (
	"encoding/json"
)

/*
https://dev.vk.com/ru/api/community-events/json-schema
*/

// MessageNew struct.
type MessageNew struct {
	Message    Message    `json:"message"`
	ClientInfo ClientInfo `json:"client_info"`
}

// MessageReply struct.
type MessageReply Message

// MessageEdit struct.
type MessageEdit Message

// MessageAllow struct.
type MessageAllow struct {
	UserID int64  `json:"user_id"`
	Key    string `json:"key"`
}

// MessageDeny struct.
type MessageDeny struct {
	UserID int64 `json:"user_id"`
}

// MessageTypingState struct.
type MessageTypingState struct {
	State  string `json:"state"`
	FromID int64  `json:"from_id"`
	ToID   int    `json:"to_id"`
}

// MessageEvent struct.
type MessageEvent struct {
	UserID                int64           `json:"user_id"`
	PeerID                int64           `json:"peer_id"`
	EventID               string          `json:"event_id"`
	Payload               json.RawMessage `json:"payload"`
	ConversationMessageID int64           `json:"conversation_message_id"`
}

//// PhotoNewObject struct.
//type PhotoNew objects.Photo

// // PhotoCommentNewObject struct.
// type PhotoCommentNewObject object.WallWallComment
//
// // PhotoCommentEditObject struct.
// type PhotoCommentEditObject object.WallWallComment
//
// // PhotoCommentRestoreObject struct.
// type PhotoCommentRestoreObject object.WallWallComment
//
// // PhotoCommentDeleteObject struct.
//
//	type PhotoCommentDeleteObject struct {
//		OwnerID   int `json:"owner_id"`
//		ID        int `json:"id"`
//		UserID    int `json:"user_id"`
//		DeleterID int `json:"deleter_id"`
//		PhotoID   int `json:"photo_id"`
//	}
//
// // AudioNewObject struct.
// type AudioNewObject object.AudioAudio
//
// // VideoNewObject struct.
// type VideoNewObject object.VideoVideo
//
// // VideoCommentNewObject struct.
// type VideoCommentNewObject object.WallWallComment
//
// // VideoCommentEditObject struct.
// type VideoCommentEditObject object.WallWallComment
//
// // VideoCommentRestoreObject struct.
// type VideoCommentRestoreObject object.WallWallComment
//
// // VideoCommentDeleteObject struct.
//
//	type VideoCommentDeleteObject struct {
//		OwnerID   int `json:"owner_id"`
//		ID        int `json:"id"`
//		UserID    int `json:"user_id"`
//		DeleterID int `json:"deleter_id"`
//		VideoID   int `json:"video_id"`
//	}

// WallPostNew struct.
type WallPostNew WallPost

//
//// WallRepostObject struct.
//type WallRepostObject object.WallWallpost
//
//// WallReplyNewObject struct.
//type WallReplyNewObject object.WallWallComment
//
//// WallReplyEditObject struct.
//type WallReplyEditObject object.WallWallComment
//
//// WallReplyRestoreObject struct.
//type WallReplyRestoreObject object.WallWallComment
//
//// WallReplyDeleteObject struct.
//type WallReplyDeleteObject struct {
//	OwnerID   int `json:"owner_id"`
//	ID        int `json:"id"`
//	DeleterID int `json:"deleter_id"`
//	PostID    int `json:"post_id"`
//}
//
//// BoardPostNewObject struct.
//type BoardPostNewObject object.BoardTopicComment
//
//// BoardPostEditObject struct.
//type BoardPostEditObject object.BoardTopicComment
//
//// BoardPostRestoreObject struct.
//type BoardPostRestoreObject object.BoardTopicComment
//
//// BoardPostDeleteObject struct.
//type BoardPostDeleteObject struct {
//	TopicOwnerID int `json:"topic_owner_id"`
//	TopicID      int `json:"topic_id"`
//	ID           int `json:"id"`
//}
//
//// MarketCommentNewObject struct.
//type MarketCommentNewObject object.WallWallComment
//
//// MarketCommentEditObject struct.
//type MarketCommentEditObject object.WallWallComment
//
//// MarketCommentRestoreObject struct.
//type MarketCommentRestoreObject object.WallWallComment
//
//// MarketCommentDeleteObject struct.
//type MarketCommentDeleteObject struct {
//	OwnerID   int `json:"owner_id"`
//	ID        int `json:"id"`
//	UserID    int `json:"user_id"`
//	DeleterID int `json:"deleter_id"`
//	ItemID    int `json:"item_id"`
//}
//
//// MarketOrderNewObject struct.
//type MarketOrderNewObject object.MarketOrder
//
//// MarketOrderEditObject struct.
//type MarketOrderEditObject object.MarketOrder
//
//// GroupLeaveObject struct.
//type GroupLeaveObject struct {
//	UserID int                `json:"user_id"`
//	Self   object.BaseBoolInt `json:"self"`
//}
//
//// GroupJoinObject struct.
//type GroupJoinObject struct {
//	UserID   int    `json:"user_id"`
//	JoinType string `json:"join_type"`
//}
//
//// UserBlockObject struct.
//type UserBlockObject struct {
//	AdminID     int    `json:"admin_id"`
//	UserID      int    `json:"user_id"`
//	UnblockDate int    `json:"unblock_date"`
//	Reason      int    `json:"reason"`
//	Comment     string `json:"comment"`
//}
//
//// UserUnblockObject struct.
//type UserUnblockObject struct {
//	AdminID   int `json:"admin_id"`
//	UserID    int `json:"user_id"`
//	ByEndDate int `json:"by_end_date"`
//}
//
//// PollVoteNewObject struct.
////
//// BUG(VK): при голосовании за несколько вариантов, возвращается только один.
//type PollVoteNewObject struct {
//	OwnerID  int `json:"owner_id"`
//	PollID   int `json:"poll_id"`
//	OptionID int `json:"option_id"`
//	UserID   int `json:"user_id"`
//}
//
//// GroupOfficersEditObject struct.
//type GroupOfficersEditObject struct {
//	AdminID  int `json:"admin_id"`
//	UserID   int `json:"user_id"`
//	LevelOld int `json:"level_old"`
//	LevelNew int `json:"level_new"`
//}
//
//// Changes struct.
//type Changes struct {
//	OldValue string `json:"old_value"`
//	NewValue string `json:"new_value"`
//}
//
//// ChangesInt struct.
//type ChangesInt struct {
//	OldValue int `json:"old_value"`
//	NewValue int `json:"new_value"`
//}
//
//// GroupChangeSettingsObject struct.
////
//// BUG(VK): Phone https://vk.com/bugtracker?act=show&id=64240
////
//// BUG(VK): Email https://vk.com/bugtracker?act=show&id=86650
//type GroupChangeSettingsObject struct {
//	UserID  int `json:"user_id"`
//	Changes struct {
//		Title                 Changes    `json:"title"`
//		Description           Changes    `json:"description"`
//		Access                ChangesInt `json:"access"`
//		ScreenName            Changes    `json:"screen_name"`
//		PublicCategory        ChangesInt `json:"public_category"`
//		PublicSubcategory     ChangesInt `json:"public_subcategory"`
//		AgeLimits             ChangesInt `json:"age_limits"`
//		Website               Changes    `json:"website"`
//		StatusDefault         Changes    `json:"status_default"`
//		Wall                  ChangesInt `json:"wall"`                    // на основе ответа
//		Replies               ChangesInt `json:"replies"`                 // на основе ответа
//		Topics                ChangesInt `json:"topics"`                  // на основе ответа
//		Audio                 ChangesInt `json:"audio"`                   // на основе ответа
//		Photos                ChangesInt `json:"photos"`                  // на основе ответа
//		Video                 ChangesInt `json:"video"`                   // на основе ответа
//		Market                ChangesInt `json:"market"`                  // на основе ответа
//		Docs                  ChangesInt `json:"docs"`                    // на основе ответа
//		Messages              ChangesInt `json:"messages"`                // на основе ответа
//		EventGroupID          ChangesInt `json:"event_group_id"`          // на основе ответа
//		Links                 Changes    `json:"links"`                   // на основе ответа
//		Email                 Changes    `json:"email"`                   // на основе ответа
//		EventStartDate        ChangesInt `json:"event_start_date::"`      // на основе ответа
//		EventFinishDate       ChangesInt `json:"event_finish_date:"`      // на основе ответа
//		Subject               Changes    `json:"subject"`                 // на основе ответа
//		MarketWiki            Changes    `json:"market_wiki"`             // на основе ответа
//		DisableMarketComments ChangesInt `json:"disable_market_comments"` // на основе ответа
//		Phone                 ChangesInt `json:"phone"`                   // на основе ответа
//		CountryID             ChangesInt `json:"country_id"`              // на основе ответа
//		CityID                ChangesInt `json:"city_id"`                 // на основе ответа
//	} `json:"Changes"`
//}
//
//// GroupChangePhotoObject struct.
//type GroupChangePhotoObject struct {
//	UserID int                `json:"user_id"`
//	Photo  object.PhotosPhoto `json:"photo"`
//}
//
//// VkpayTransactionObject struct.
//type VkpayTransactionObject struct {
//	FromID      int    `json:"from_id"`
//	Amount      int    `json:"amount"`
//	Description string `json:"description"`
//	Date        int    `json:"date"`
//}
//
//// LeadFormsNewObject struct.
//type LeadFormsNewObject struct {
//	LeadID   int    `json:"lead_id"`
//	GroupID  int    `json:"group_id"`
//	UserID   int    `json:"user_id"`
//	FormID   int    `json:"form_id"`
//	FormName string `json:"form_name"`
//	AdID     int    `json:"ad_id"`
//	Answers  []struct {
//		Key      string `json:"key"`
//		Question string `json:"question"`
//		Answer   string `json:"answer"`
//	} `json:"answers"`
//}
//
//// AppPayloadObject struct.
//type AppPayloadObject struct {
//	UserID  int    `json:"user_id"`
//	AppID   int    `json:"app_id"`
//	Payload string `json:"payload"`
//}
//
//// MessageReadObject struct.
//type MessageReadObject struct {
//	FromID        int `json:"from_id"`
//	PeerID        int `json:"peer_id"`
//	ReadMessageID int `json:"read_message_id"`
//}
//
//// LikeAddObject struct.
//type LikeAddObject struct {
//	LikerID       int    `json:"liker_id"`
//	ObjectType    string `json:"object_type"`
//	ObjectOwnerID int    `json:"object_owner_id"`
//	ObjectID      int    `json:"object_id"`
//	ThreadReplyID int    `json:"thread_reply_id"`
//	PostID        int    `json:"post_id"` // for comment
//}
//
//// LikeRemoveObject struct.
//type LikeRemoveObject struct {
//	LikerID       int    `json:"liker_id"`
//	ObjectType    string `json:"object_type"`
//	ObjectOwnerID int    `json:"object_owner_id"`
//	ObjectID      int    `json:"object_id"`
//	ThreadReplyID int    `json:"thread_reply_id"`
//	PostID        int    `json:"post_id"` // for comment
//}
//
//// DonutSubscriptionCreateObject struct.
//type DonutSubscriptionCreateObject struct {
//	Amount           float64 `json:"amount"`
//	AmountWithoutFee float64 `json:"amount_without_fee"`
//	UserID           float64 `json:"user_id"`
//}
//
//// DonutSubscriptionProlongedObject struct.
//type DonutSubscriptionProlongedObject struct {
//	Amount           float64 `json:"amount"`
//	AmountWithoutFee float64 `json:"amount_without_fee"`
//	UserID           float64 `json:"user_id"`
//}
//
//// DonutSubscriptionExpiredObject struct.
//type DonutSubscriptionExpiredObject struct {
//	UserID float64 `json:"user_id"`
//}
//
//// DonutSubscriptionCancelledObject struct.
//type DonutSubscriptionCancelledObject struct {
//	UserID float64 `json:"user_id"`
//}
//
//// DonutSubscriptionPriceChangedObject struct.
//type DonutSubscriptionPriceChangedObject struct {
//	AmountOld            float64 `json:"amount_old"`
//	AmountNew            float64 `json:"amount_new"`
//	AmountDiff           float64 `json:"amount_diff"`
//	AmountDiffWithoutFee float64 `json:"amount_diff_without_fee"`
//	UserID               float64 `json:"user_id"`
//}
//
//// DonutMoneyWithdrawObject struct.
//type DonutMoneyWithdrawObject struct {
//	Amount           float64 `json:"amount"`
//	AmountWithoutFee float64 `json:"amount_without_fee"`
//}
//
//// DonutMoneyWithdrawErrorObject struct.
//type DonutMoneyWithdrawErrorObject struct {
//	Reason string `json:"reason"`
//}
//
