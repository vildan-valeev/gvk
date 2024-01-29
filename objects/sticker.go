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

// BaseSticker https://dev.vk.com/ru/reference/objects/sticker.
type Sticker struct {
	ProductID            int            `json:"product_id"`
	StickerID            int            `json:"sticker_id"`
	Images               []StickerImage `json:"images"`
	ImagesWithBackground []StickerImage `json:"images_with_background"`
	IsAllowed            bool           `json:"is_allowed"`
	AnimationURL         string         `json:"animation_url"`
}

type StickerImage struct {
	Height float64 `json:"height"`
	URL    string  `json:"url"`
	Width  float64 `json:"width"`
}
