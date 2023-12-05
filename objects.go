package gvk

// ClientInfo struct.
type ClientInfo struct {
	ButtonActions  []string `json:"button_actions"`
	Keyboard       bool     `json:"keyboard"`
	InlineKeyboard bool     `json:"inline_keyboard"`
	Carousel       bool     `json:"carousel"`
	LangID         int      `json:"lang_id"`
}

// BaseMessageGeo struct.
type BaseMessageGeo struct {
	Coordinates BaseGeoCoordinates `json:"coordinates"`
	Place       BasePlace          `json:"place"`
	Showmap     int                `json:"showmap"`
	Type        string             `json:"type"`
}

// BaseGeoCoordinates struct.
type BaseGeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// BasePlace struct.
type BasePlace struct {
	Address        string             `json:"address"`
	Checkins       int                `json:"checkins"`
	City           interface{}        `json:"city"` // BUG(VK): https://github.com/VKCOM/vk-api-schema/issues/143
	Country        interface{}        `json:"country"`
	Created        int                `json:"created"`
	ID             int                `json:"id"`
	Icon           string             `json:"icon"`
	Latitude       float64            `json:"latitude"`
	Longitude      float64            `json:"longitude"`
	Title          string             `json:"title"`
	Type           string             `json:"type"`
	IsDeleted      bool               `json:"is_deleted"`
	TotalCheckins  int                `json:"total_checkins"`
	Updated        int                `json:"updated"`
	CategoryObject BaseCategoryObject `json:"category_object"`
}

// MessagesMessage struct.
type MessagesMessage struct {
	// Only for messages from community. Contains user ID of community admin,
	// who sent this message.
	AdminAuthorID int                         `json:"admin_author_id"`
	Action        MessagesMessageAction       `json:"action"`
	Attachments   []MessagesMessageAttachment `json:"attachments"`

	// Unique auto-incremented number for all messages with this peer.
	ConversationMessageID int `json:"conversation_message_id"`

	// Date when the message has been sent in Unixtime.
	Date int `json:"date"`

	// Message author's ID.
	FromID int64 `json:"from_id"`

	// Forwarded messages.
	FwdMessages  []MessagesMessage `json:"fwd_Messages"`
	ReplyMessage *MessagesMessage  `json:"reply_message"`
	Geo          BaseMessageGeo    `json:"geo"`
	PinnedAt     int               `json:"pinned_at,omitempty"`
	ID           int               `json:"id"`        // Message ID
	Deleted      bool              `json:"deleted"`   // Is it an deleted message
	Important    bool              `json:"important"` // Is it an important message
	IsHidden     bool              `json:"is_hidden"`
	IsCropped    bool              `json:"is_cropped"`
	IsSilent     bool              `json:"is_silent"`
	Out          bool              `json:"out"` // Information whether the message is outcoming
	WasListened  bool              `json:"was_listened,omitempty"`
	Keyboard     MessagesKeyboard  `json:"keyboard"`
	Template     MessagesTemplate  `json:"template"`
	Payload      string            `json:"payload"`
	PeerID       int               `json:"peer_id"` // Peer ID

	// ID used for sending messages. It returned only for outgoing messages.
	RandomID     int    `json:"random_id"`
	Ref          string `json:"ref"`
	RefSource    string `json:"ref_source"`
	Text         string `json:"text"`          // Message text
	UpdateTime   int    `json:"update_time"`   // Date when the message has been updated in Unixtime
	MembersCount int    `json:"members_count"` // Members number
	ExpireTTL    int    `json:"expire_ttl"`
	MessageTag   string `json:"message_tag"` // for https://notify.mail.ru/
}

type MessageReplyObject MessagesMessage

// MessagesMessageAction struct.
type MessagesMessageAction struct {
	ConversationMessageID int `json:"conversation_message_id"` // Message ID

	// Email address for chat_invite_user or chat_kick_user actions.
	Email    string                     `json:"email"`
	MemberID int                        `json:"member_id"` // User or email peer ID
	Message  string                     `json:"message"`   // Message body of related message
	Photo    MessagesMessageActionPhoto `json:"photo"`

	// New chat title for chat_create and chat_title_update actions.
	Text string `json:"text"`
	Type string `json:"type"`
}

// MessagesMessageAttachment struct.
type MessagesMessageAttachment struct {
	Audio             AudioAudio        `json:"audio"`
	Doc               DocsDoc           `json:"doc"`
	Gift              GiftsLayout       `json:"gift"`
	Link              BaseLink          `json:"link"`
	Market            MarketMarketItem  `json:"market"`
	MarketMarketAlbum MarketMarketAlbum `json:"market_market_album"`
	Photo             PhotosPhoto       `json:"photo"`
	Sticker           BaseSticker       `json:"sticker"`
	Type              string            `json:"type"`
	Video             VideoVideo        `json:"video"`
	Wall              WallWallpost      `json:"wall"`
	WallReply         WallWallComment   `json:"wall_reply"`
	AudioMessage      DocsDoc           `json:"audio_message"`
	Graffiti          DocsDoc           `json:"graffiti"`
	Poll              PollsPoll         `json:"poll"`
	Call              MessageCall       `json:"call"`
	Story             StoriesStory      `json:"story"`
	Podcast           PodcastsEpisode   `json:"podcast"`
}
