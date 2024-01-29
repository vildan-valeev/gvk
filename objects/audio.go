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

// Audio https://dev.vk.com/ru/reference/objects/audio.
type Audio struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	URL      string `json:"url"`
	LyricsID int    `json:"lyrics_id"`
	AlbumID  int    `json:"album_id"`
	GenreID  int    `json:"genre_id"`
	Date     int    `json:"date"`
	NoSearch int    `json:"no_search"`
	IsHq     int    `json:"is_hq"`
}
