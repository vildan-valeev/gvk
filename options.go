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
	DontParseLinks  int64    `json:"dont_parse_links"`
	DisableMentions int64    `json:"disable_mentions"`
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
