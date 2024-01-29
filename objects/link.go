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

// Link . https://dev.vk.com/ru/reference/objects/link
type Link struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Caption     string `json:"caption"`
	Description string `json:"description"`
	Photo       Photo  `json:"photo"`
	PreviewPage string `json:"preview_page"`
	PreviewURL  string `json:"preview_url"`
}
