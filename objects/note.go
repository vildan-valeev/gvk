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

// Note https://dev.vk.com/ru/reference/objects/note.
type Note struct {
	ID           int    `json:"id"`            // Note ID
	OwnerID      int    `json:"owner_id"`      // Note owner's ID
	Title        string `json:"title"`         // Note title
	Date         int    `json:"date"`          // Date when the note has been created in Unixtime
	Comments     int    `json:"comments"`      // Comments number
	ReadComments int    `json:"read_comments"` // Read comments number
	ViewURL      string `json:"view_url"`      // URL of the page with note preview
	PrivacyView  string `json:"privacy_view"`
	CanComment   int    `json:"can_comment"`
	TextWiki     string `json:"text_wiki"`
}
