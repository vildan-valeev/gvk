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

// MarketAlbum https://dev.vk.com/ru/reference/objects/market-album.
type MarketAlbum struct {
	ID       int    `json:"id"`       // Market album ID
	OwnerID  int    `json:"owner_id"` // Market album owner's ID
	Title    string `json:"title"`    // Market album title
	IsMain   bool   `json:"is_main"`
	IsHidden bool   `json:"is_hidden"`
	Count    int    `json:"count"` // Items number
	Photo    Photo  `json:"photo"`
}
