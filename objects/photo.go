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

/*
https://dev.vk.com/ru/reference/objects/photo
*/

type Photo struct {
	ID      int          `json:"id"`       // Photo ID
	AlbumID int          `json:"album_id"` // Album ID
	OwnerID int          `json:"owner_id"` // Photo owner's ID
	UserID  int          `json:"user_id"`  // ID of the user who have uploaded the photo
	Text    string       `json:"text"`     // Photo caption
	Date    int          `json:"date"`     // Date when uploaded
	Sizes   []PhotoSizes `json:"sizes"`
	Width   int          `json:"width"`  // Original photo width
	Height  int          `json:"height"` // Original photo height
}

type PhotoSizes struct {
	Height float64 `json:"height"`
	URL    string  `json:"url"`
	Width  float64 `json:"width"`
	Type   string  `json:"type"`
}
