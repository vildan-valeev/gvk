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

package objects

// MarketMarketItem struct.
type MarketItem struct {
	ID           int              `json:"id"`          // Item ID
	OwnerID      int              `json:"owner_id"`    // Item owner's ID
	Title        string           `json:"title"`       // Item title
	Description  string           `json:"description"` // Item description
	Price        MarketPrice      `json:"price"`
	Dimensions   MarketDimensions `json:"dimensions"`
	Weight       int              `json:"weight"`
	Category     MarketCategory   `json:"category"`
	ThumbPhoto   string           `json:"thumb_photo"` // URL of the preview image
	Date         int              `json:"date,omitempty"`
	Availability int              `json:"availability"` // Статус доступности товара. Возможные значения: 0 — товар доступен. 1 — товар удален. 2 — товар недоступен.
	IsFavorite   bool             `json:"is_favorite"`
	SKU          string           `json:"sku"`
	RejectInfo   RejectInfo       `json:"reject_info"`

	CanComment  int             `json:"can_comment"` //  Возможность комментировать товар для текущего пользователя (1 — есть, 0 — нет)
	CanRepost   int             `json:"can_repost"`  // Возможность сделать репост товара для текущего пользователя (1 — есть, 0 — нет).
	Likes       MarketItemLikes `json:"likes"`
	URL         string          `json:"url"`          // URL to item
	ButtonTitle string          `json:"button_title"` // Текст на кнопке товара. Возможные значения: Купить, Перейти в магазин, Купить билет
	Photos      []Photo         `json:"photos"`
}

// MarketPrice struct.
type MarketPrice struct {
	Amount    string         `json:"amount"` // Amount
	Currency  MarketCurrency `json:"currency"`
	OldAmount string         `json:"old_amount"`
	Text      string         `json:"text"` // Text
}

// MarketCurrency struct.
type MarketCurrency struct {
	ID   int    `json:"id"`   // Currency ID
	Name string `json:"name"` // Currency sign
}

// MarketDimensions struct.
type MarketDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Length int `json:"length"`
}

// MarketCategory struct.
type MarketCategory struct {
	ID      int           `json:"id"`   // Category ID
	Name    string        `json:"name"` // Category name
	Section MarketSection `json:"section"`
}

// MarketSection struct.
type MarketSection struct {
	ID   int    `json:"id"`   // Section ID
	Name string `json:"name"` // Section name
}

type RejectInfo struct {
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Buttons            []Button `json:"buttons"`
	ModerationStatus   int      `json:"moderation_status"`
	InfoLink           string   `json:"info_link"`
	WhiteToSupportLink string   `json:"white_to_support_link"`
}

type MarketItemLikes struct {
	UserLikes int `json:"user_likes"`
	Count     int `json:"count"`
}
