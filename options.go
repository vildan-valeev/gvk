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
	UserID int64 `query:"user_id"`
	//RandomID int32 `query:"random_id"`
	//PeerID   int64 `query:"peer_id"`
}

type SetLongPollSettingsOptions struct {
	GroupID    int64  `query:"group_id"`
	Enable     int64  `query:"enable"`
	APIVersion string `query:"api_version"`
}

type UsersGetOptions struct {
	UserIDS string `json:"user_ids"`
	Fields  string `json:"fields"`
}
