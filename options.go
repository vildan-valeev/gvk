package gvk

type UpdateOptions struct {
	Server string `query:"server"`
	Act    string `query:"act"`
	Key    string `query:"key"`
	Ts     string `query:"ts"`
	Wait   int64  `query:"wait"`
}

type GetLongPollServerOptions struct {
	GroupID int64 `query:"group_id"`
}

type MessagesSendOptions struct {
	Message         string   `query:"message"`
	UserID          int64    `query:"user_id"`
	Keyboard        Keyboard `query:"keyboard"`
	DontParseLinks  int64    `query:"dont_parse_links"` // 1 - true, 0 - false
	DisableMentions int64    `query:"disable_mentions"` // 1 - true, 0 - false
}

type MessagesEditOptions struct {
	MessageID int64  `query:"message_id"`
	Message   string `query:"message"`
	PeerID    int64  `query:"peer_id"`

	Lat                   string   `query:"lat"`  // float64 as string: from -90 to 90
	Long                  string   `query:"long"` // float64 as string: from -90 to 90
	Attachment            string   `query:"attachment"`
	KeepForwardMessages   int64    `query:"keep_forward_messages"`
	KeepSnippets          int64    `query:"keep_snippets"`
	GroupID               int64    `query:"group_id"`
	DontParseLinks        int64    `query:"dont_parse_links"`
	DisableMentions       int64    `query:"disable_mentions"`
	ConversationMessageID int64    `query:"conversation_message_id"`
	Template              string   `query:"template"`
	Keyboard              Keyboard `query:"keyboard"`
}

type SetLongPollSettingsOptions struct {
	GroupID    int64  `query:"group_id"`
	Enable     int64  `query:"enable"`
	APIVersion string `query:"api_version"`
}

type UsersGetOptions struct {
	UserIDS string `query:"user_ids"`
	Fields  string `query:"fields"`
}

type MessageEventAnswerOptions struct {
	EventID   string `query:"event_id"`
	UserID    int64  `query:"user_id"`
	PeerID    int64  `query:"peer_id"`
	EventData string `query:"text"`
}

type GroupsIsMemberOptions struct {
	GroupID  int64  `query:"group_id"`
	UserID   int64  `query:"user_id"`
	UserIDs  string `query:"user_ids"`
	Extended int64  `query:"extended"`
}

type WallPostOptions struct {
	Message           string `query:"message"`
	OwnerID           int64  `query:"owner_id"`
	FriendsOnly       int64  `query:"friends_only"` // 1 - true, 0 - false
	FromGroup         int64  `query:"from_group"`   // 1 - true, 0 - false
	Signed            int64  `query:"signed"`       // 1 - true, 0 - false
	PublishDate       int64  `query:"publish_date"`
	Lat               string `query:"lat"`         // float64 as string: from -90 to 90
	Long              string `query:"long"`        // float64 as string: from -90 to 90
	MarkAsAds         string `query:"mark_as_ads"` // 1 - true, 0 - false
	LinkPhotoID       string `query:"link_photo_id"`
	LinkTitle         string `query:"link_title"`
	CloseComments     int64  `query:"close_comments"`     // 1 - true, 0 - false
	MuteNotifications int64  `query:"mute_notifications"` // 1 - true, 0 - false
	Copyright         string `query:"copyright"`
}

// WallEditOptions https://dev.vk.com/ru/method/wall.edit
type WallEditOptions struct {
	OwnerID     int64  `query:"owner_id"`
	PostID      int64  `query:"post_id"`      // required
	FriendsOnly int64  `query:"friends_only"` // 1 - true, 0 - false
	Message     string `query:"message"`
	//Attachments         string `query:"attachments"` // <type><owner_id>_<media_id>,<type><owner_id>_<media_id>
	Services            string `query:"services"`
	Signed              int64  `query:"signed"` // 1 - true, 0 - false
	PublishDate         int64  `query:"publish_date"`
	Lat                 string `query:"lat"`  // float64 as string: from -90 to 90
	Long                string `query:"long"` // float64 as string: from -90 to 90
	PlaceID             int64  `query:"place_id"`
	MarkAsAds           string `query:"mark_as_ads"`    // 1 - true, 0 - false
	CloseComments       int64  `query:"close_comments"` // 1 - true, 0 - false
	DonutPaidDuration   int64  `query:"donut_paid_duration"`
	PosterBkgID         int64  `query:"poster_bkg_id"`
	PosterBkgIDOwnerID  int64  `query:"poster_bkg_id_owner_id"`
	PosterBkgAccessHash string `query:"poster_bkg_access_hash"`
	Copyright           string `query:"copyright"`
}
