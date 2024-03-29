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

import "github.com/vildan-valeev/gvk/objects"

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
	PeerID                int64          `json:"peer_id"` // Peer ID
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
	Color  ButtonColor  `json:"color,omitempty"` // Button color
}

type ButtonColor string

const (
	ButtonColorPrimary   ButtonColor = "primary"
	ButtonColorSecondary ButtonColor = "secondary"
	ButtonColorNegative  ButtonColor = "negative"
	ButtonColorPositive  ButtonColor = "positive"
)

type ButtonAction struct {
	Type    ButtonType `json:"type"`              // Button type
	Label   string     `json:"label,omitempty"`   // Label for button
	Payload string     `json:"payload,omitempty"` // Additional data sent along with message for developer convenience

	AppID   int    `json:"app_id,omitempty"`   // Fragment value in app link like vk.com/app{app_id}_-654321#hash
	Hash    string `json:"hash,omitempty"`     // Fragment value in app link like vk.com/app123456_-654321#{hash}
	OwnerID int    `json:"owner_id,omitempty"` // Fragment value in app link like vk.com/app123456_{owner_id}#hash
	Link    string `json:"link,omitempty"`     // Link URL
}

type ButtonType string

const (
	ButtonTypeText     ButtonType = "text"
	ButtonTypeOpenLink ButtonType = "open_link"
	ButtonTypeLocation ButtonType = "location"
	ButtonTypeVkPay    ButtonType = "vkpay"
	ButtonTypeOpenApp  ButtonType = "open_app"
	ButtonTypeCallBack ButtonType = "callback"
)

// BaseMessageGeo struct.
type BaseMessageGeo struct {
	Coordinates BaseGeoCoordinates `json:"coordinates"`
	Place       Place              `json:"place"`
	Showmap     int                `json:"showmap"`
	Type        string             `json:"type"`
}

// BaseGeoCoordinates struct.
type BaseGeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// BasePlace struct.
type Place struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Created   int     `json:"created"`
	Icon      string  `json:"icon"`
	Checkins  int     `json:"checkins"`
	Updated   int     `json:"updated"`
	Type      string  `json:"type"`
	Country   int     `json:"country"`
	City      int     `json:"city"`
	Address   string  `json:"address"`
}

type WallPost struct {
	ID           int                  `json:"id"`       // Post ID
	OwnerID      int64                `json:"owner_id"` // Wall owner's ID
	FromID       int64                `json:"from_id"`  // Post author ID
	CreatedBy    int                  `json:"created_by"`
	Date         int                  `json:"date"` // Date of publishing in Unixtime
	Text         string               `json:"text"` // Post text
	ReplyOwnerID int                  `json:"reply_owner_id"`
	ReplyPostID  int                  `json:"reply_post_id"`
	FriendsOnly  int                  `json:"friends_only"` // 1, если запись была создана с опцией «Только для друзей».
	Comments     Comments             `json:"comments"`
	Copyright    Copyright            `json:"copyright"`
	Likes        Likes                `json:"likes"`   // Count of likes
	Reposts      Reposts              `json:"reposts"` // Count of reposts
	Views        Views                `json:"views"`   // Count of views
	PostType     string               `json:"post_type"`
	PostSource   PostSource           `json:"post_source"`
	Attachments  []WallPostAttachment `json:"attachments"`
	Geo          Geo                  `json:"geo"`
	SignerID     int                  `json:"signer_id"` // Post signer ID
	CopyHistory  []WallPost           `json:"copy_history"`
	CanPin       int                  `json:"can_pin"`
	CanDelete    int                  `json:"can_delete"`
	CanEdit      int                  `json:"can_edit"`
	IsPinned     int                  `json:"is_pinned"`
	IsFavorite   bool                 `json:"is_favorite"` // true, если объект добавлен в закладки у текущего пользователя.
	MarkedAsAds  int                  `json:"marked_as_ads"`
	PostponedID  int                  `json:"postponed_id"` // ID from scheduled posts
}

// Comments struct.
type Comments struct {
	Count         int `json:"count"`
	CanPost       int `json:"can_post"` // информация о том, может ли текущий пользователь комментировать запись (1 — может, 0 — не может);
	GroupsCanPost int `json:"groups_can_post"`
	CanClose      int `json:"can_close"`
	CanOpen       int `json:"can_open"`
}

// Copyright information about the source of the post.
type Copyright struct {
	ID   int    `json:"id,omitempty"`
	Link string `json:"link"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// Likes struct.
type Likes struct {
	CanLike    int `json:"can_like"`    // Information whether current user can like the post
	CanPublish int `json:"can_publish"` // Information whether current user can repost
	UserLikes  int `json:"user_likes"`  // Information whether current uer has liked the post
	Count      int `json:"count"`
}

// Reposts struct.
type Reposts struct {
	Count        int `json:"count"`
	UserReposted int `json:"user_reposted"`
}

// Views struct.
type Views struct {
	Count int `json:"count"` // Count
}

type PostSource struct {
	Data     string `json:"data"`     // Additional data
	Platform string `json:"platform"` // Platform name
	Type     string `json:"type"`
	URL      string `json:"url"` // URL to an external site used to publish the post
}

// BaseGeo struct.
type Geo struct {
	Coordinates string `json:"coordinates"`
	Place       Place  `json:"place"`
	Type        string `json:"type"`
}

// WallPostAttachment struct.
type WallPostAttachment struct {
	Type  string        `json:"type"`
	Photo objects.Photo `json:"photo"`
	//Deprecated
	PostedPhoto PostedPhoto   `json:"posted_photo"`
	Video       objects.Video `json:"video"`
	Audio       objects.Audio `json:"audio"`
	Doc         objects.Doc   `json:"doc"`
	//Deprecated
	Graffiti    WallGraffiti        `json:"graffiti"`
	Link        objects.Link        `json:"link"`
	Note        objects.Note        `json:"note"`
	Poll        objects.Poll        `json:"poll"`
	Page        objects.WikiPage    `json:"page"`
	Album       PhotoAlbum          `json:"album"`
	PhotosList  []string            `json:"photos_list"`
	Market      objects.MarketItem  `json:"market"`
	MarketAlbum objects.MarketAlbum `json:"market_market_album"`
	Sticker     objects.Sticker     `json:"sticker"`
	Event       EventAttach         `json:"event"`
	PrettyCards []PrettyCard        `json:"podcast"`
}

type PostedPhoto struct {
	ID       int    `json:"id"`        // Photo ID
	OwnerID  int    `json:"owner_id"`  // Photo owner's ID
	Photo130 string `json:"photo_130"` // URL of the preview image with 130 px in width
	Photo604 string `json:"photo_604"` // URL of the preview image with 604 px in width
}

// WallGraffiti struct.
type WallGraffiti struct {
	ID       int    `json:"id"`        // Graffiti ID
	OwnerID  int    `json:"owner_id"`  // Graffiti owner's ID
	Photo130 string `json:"photo_200"` // URL of the preview image with 130 px in width
	Photo604 string `json:"photo_586"` // URL of the preview image with 604 px in width
}

// PhotoAlbum struct.
type PhotoAlbum struct {
	ID          int           `json:"id"` // Photo album ID
	Thumb       objects.Photo `json:"thumb"`
	OwnerID     int           `json:"owner_id"`    // Album owner's ID
	Title       string        `json:"title"`       // Photo album title
	Description string        `json:"description"` // Photo album description
	Created     int           `json:"created"`     // Date when the album has been created in Unixtime
	Updated     int           `json:"updated"`     // Date when the album has been updated last time in Unixtime
	Size        int           `json:"size"`        // Photos number
}

// EventAttach struct.
type EventAttach struct {
	ID           int    `json:"id"`                      // event ID
	Time         int    `json:"time,omitempty"`          // event start time
	MemberStatus int    `json:"member_status,omitempty"` // Идёт ли текущий пользователь на встречу. Возможные значения: 1 — точно идёт; 2 — возможно пойдёт; 3 — не идёт.
	Address      string `json:"address,omitempty"`       // address of event
	ButtonText   string `json:"button_text"`             // text of attach
	Friends      []int  `json:"friends"`                 // array of friends ids
	IsFavorite   bool   `json:"is_favorite"`             // is favorite
	Text         string `json:"text"`                    // text of attach

}

type PrettyCard struct {
	CardID   string          `json:"card_id"`
	LinkURL  string          `json:"link_url"`
	Title    string          `json:"title"`
	Images   []objects.Photo `json:"images"`
	Button   objects.Button  `json:"button"`
	Price    string          `json:"price"`
	PriceOld string          `json:"price_old"`
}
