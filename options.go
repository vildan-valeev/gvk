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
	UserID   int64    `query:"user_id"`
	Keyboard Keyboard `query:"keyboard"`
	//RandomID int32 `query:"random_id"`
	//PeerID   int64 `query:"peer_id"`
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
