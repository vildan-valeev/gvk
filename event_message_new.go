package gvk

// ClientInfo struct.
type ClientInfo struct {
	ButtonActions  []string `json:"button_actions"`
	Keyboard       bool     `json:"keyboard"`
	InlineKeyboard bool     `json:"inline_keyboard"`
	Carousel       bool     `json:"carousel"`
	LangID         int      `json:"lang_id"`
}

// Messages struct.
type Message struct {
	ID                    int            `json:"id"` // Message ID
	Date                  int            `json:"date"`
	PeerID                int            `json:"peer_id"` // Peer ID
	FromID                int64          `json:"from_id"`
	Text                  string         `json:"text"` // Message text
	RandomID              int            `json:"random_id"`
	Ref                   string         `json:"ref"`
	RefSource             string         `json:"ref_source"`
	Attachments           []Attachment   `json:"attachments"`
	Important             bool           `json:"important"` // Is it an important message
	Geo                   BaseMessageGeo `json:"geo"`
	Payload               string         `json:"payload"`
	Keyboard              Keyboard       `json:"keyboard"`
	FwdMessages           []Message      `json:"fwd_Messages"`
	ReplyMessage          *Message       `json:"reply_message"`
	Action                MessageAction  `json:"action"`
	AdminAuthorID         int            `json:"admin_author_id"`
	ConversationMessageID int            `json:"conversation_message_id"`
	IsCropped             bool           `json:"is_cropped"`
	MembersCount          int            `json:"members_count"` // Members number
	UpdateTime            int            `json:"update_time"`   // Date when the message has been updated in Unixtime
	WasListened           bool           `json:"was_listened,omitempty"`
	PinnedAt              int            `json:"pinned_at,omitempty"`
	MessageTag            string         `json:"message_tag"` // for https://notify.mail.ru/
	IsMentionedUser       bool           `json:"is_mentioned_user,omitempty"`
}

type MessageAction struct {
	ConversationMessageID int                `json:"conversation_message_id"` // Message ID
	Email                 string             `json:"email"`
	MemberID              int                `json:"member_id"` // User or email peer ID
	Message               string             `json:"message"`   // Message body of related message
	Photo                 MessageActionPhoto `json:"photo"`
	Text                  string             `json:"text"`
	Type                  string             `json:"type"`
}

type MessageActionPhoto struct {
	Photo100 string `json:"photo_100"` // URL of the preview image with 100px in width
	Photo200 string `json:"photo_200"` // URL of the preview image with 200px in width
	Photo50  string `json:"photo_50"`  // URL of the preview image with 50px in width
}

type Attachment struct {
	//Audio             AudioAudio        `json:"audio"`
	//Doc               DocsDoc           `json:"doc"`
	//Gift              GiftsLayout       `json:"gift"`
	//Link              BaseLink          `json:"link"`
	//Market            MarketMarketItem  `json:"market"`
	//MarketMarketAlbum MarketMarketAlbum `json:"market_market_album"`
	//Photo             PhotosPhoto       `json:"photo"`
	//Sticker           BaseSticker       `json:"sticker"`
	Type string `json:"type"`
	//Video             VideoVideo        `json:"video"`
	//Wall              WallWallpost      `json:"wall"`
	//WallReply         WallWallComment   `json:"wall_reply"`
	//AudioMessage      DocsDoc           `json:"audio_message"`
	//Graffiti          DocsDoc           `json:"graffiti"`
	//Poll              PollsPoll         `json:"poll"`
	//Call              MessageCall       `json:"call"`
	//Story             StoriesStory      `json:"story"`
	//Podcast           PodcastsEpisode   `json:"podcast"`
}
type Keyboard struct {
	AuthorID int        `json:"author_id,omitempty"` // Community or bot, which set this keyboard
	Buttons  [][]Button `json:"buttons"`
	OneTime  bool       `json:"one_time,omitempty"` // Should this keyboard disappear on first use
	Inline   bool       `json:"inline,omitempty"`
}

type Button struct {
	Action ButtonAction `json:"action"`
	Color  string       `json:"color,omitempty"` // Button color
}
type ButtonAction struct {
	AppID   int    `json:"app_id,omitempty"`   // Fragment value in app link like vk.com/app{app_id}_-654321#hash
	Hash    string `json:"hash,omitempty"`     // Fragment value in app link like vk.com/app123456_-654321#{hash}
	Label   string `json:"label,omitempty"`    // Label for button
	OwnerID int    `json:"owner_id,omitempty"` // Fragment value in app link like vk.com/app123456_{owner_id}#hash
	Payload string `json:"payload,omitempty"`  // Additional data sent along with message for developer convenience
	Type    string `json:"type"`               // Button type
	Link    string `json:"link,omitempty"`     // Link URL
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
type BaseCategoryObject struct {
	ID    int         `json:"id"`
	Title string      `json:"title"`
	Icons []BaseImage `json:"icons"`
}

// BaseImage struct.
type BaseImage struct {
	Height float64 `json:"height"`
	URL    string  `json:"url"`
	Width  float64 `json:"width"`
	Type   string  `json:"type"`
}
