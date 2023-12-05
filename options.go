package gvk

type UpdateOptions struct {
	AllowedUpdates []UpdateType `query:"allowed_updates"`
	Offset         int          `query:"offset"`
	Limit          int          `query:"limit"`
	v              float32      `query:"v"`
}

type GetLongPollServerOptions struct {
	GroupID int64 `query:"group_id"`
}

type MessagesSendOptions struct {
	UserID   int64 `query:"user_id"`
	RandomID int32 `query:"random_id"`
	PeerID   int64 `query:"peer_id"`
}
