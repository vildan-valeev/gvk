package gvk

// EventType type.
type EventType string

// EventType list.
const (
	EventConfirmation                  EventType = "confirmation"
	EventMessageNew                              = "message_new"
	EventMessageReply                            = "message_reply"
	EventMessageEdit                             = "message_edit"
	EventMessageAllow                            = "message_allow"
	EventMessageDeny                             = "message_deny"
	EventMessageTypingState                      = "message_typing_state"
	EventMessageEvent                            = "message_event"
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

type Object struct {
	MessageNew *MessageNewObject
	//messageReply *MessageReplyObject
	//messageEdit  *MessageEditObject
	//messageAllow                  *MessageAllowObject
	//messageDeny                   *MessageDenyObject
	//messageTypingState            *MessageTypingStateObject
	//messageEvent                  *MessageEventObject
	//photoNew                      *PhotoNewObject
	//photoCommentNew               *PhotoCommentNewObject
	//photoCommentEdit              *PhotoCommentEditObject
	//photoCommentRestore           *PhotoCommentRestoreObject
	//photoCommentDelete            *PhotoCommentDeleteObject
	//audioNew                      *AudioNewObject
	//videoNew                      *VideoNewObject
	//videoCommentNew               *VideoCommentNewObject
	//videoCommentEdit              *VideoCommentEditObject
	//videoCommentRestore           *VideoCommentRestoreObject
	//videoCommentDelete            *VideoCommentDeleteObject
	//wallPostNew                   *WallPostNewObject
	//wallRepost                    *WallRepostObject
	//wallReplyNew                  *WallReplyNewObject
	//wallReplyEdit                 *WallReplyEditObject
	//wallReplyRestore              *WallReplyRestoreObject
	//wallReplyDelete               *WallReplyDeleteObject
	//boardPostNew                  *BoardPostNewObject
	//boardPostEdit                 *BoardPostEditObject
	//boardPostRestore              *BoardPostRestoreObject
	//boardPostDelete               *BoardPostDeleteObject
	//marketCommentNew              *MarketCommentNewObject
	//marketCommentEdit             *MarketCommentEditObject
	//marketCommentRestore          *MarketCommentRestoreObject
	//marketCommentDelete           *MarketCommentDeleteObject
	//marketOrderNew                *MarketOrderNewObject
	//marketOrderEdit               *MarketOrderEditObject
	//groupLeave                    *GroupLeaveObject
	//groupJoin                     *GroupJoinObject
	//userBlock                     *UserBlockObject
	//userUnblock                   *UserUnblockObject
	//pollVoteNew                   *PollVoteNewObject
	//groupOfficersEdit             *GroupOfficersEditObject
	//groupChangeSettings           *GroupChangeSettingsObject
	//groupChangePhoto              *GroupChangePhotoObject
	//vkpayTransaction              *VkpayTransactionObject
	//leadFormsNew                  *LeadFormsNewObject
	//appPayload                    *AppPayloadObject
	//messageRead                   *MessageReadObject
	//likeAdd                       *LikeAddObject
	//likeRemove                    *LikeRemoveObject
	//donutSubscriptionCreate       *DonutSubscriptionCreateObject
	//donutSubscriptionProlonged    *DonutSubscriptionProlongedObject
	//donutSubscriptionExpired      *DonutSubscriptionExpiredObject
	//donutSubscriptionCancelled    *DonutSubscriptionCancelledObject
	//donutSubscriptionPriceChanged *DonutSubscriptionPriceChangedObject
	//donutMoneyWithdraw            *DonutMoneyWithdrawObject
	//donutMoneyWithdrawError       *DonutMoneyWithdrawErrorObject
}

// MessageNewObject struct.
type MessageNewObject struct {
	Message    Message    `json:"message"`
	ClientInfo ClientInfo `json:"client_info"`
}
