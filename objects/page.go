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

// WikiPage struct. https://dev.vk.com/ru/reference/objects/wiki-page
type WikiPage struct {
	ID                       int    `json:"id"`       // Page ID.
	GroupID                  int    `json:"group_id"` // Community ID.
	CreatorID                int    `json:"creator_id"`
	Title                    string `json:"title"`
	CurrentUserCanEdit       int    `json:"current_user_can_edit"`        // 1, если текущий пользователь может редактировать текст вики-страницы, иначе — 0
	CurrentUserCanEditAccess int    `json:"current_user_can_edit_access"` // 1, если текущий пользователь может изменять права доступа на вики-страницу, иначе — 0.
	WhoCanEdit               int    `json:"who_can_edit"`
	WhoCanView               int    `json:"who_can_view"`
	Edited                   int    `json:"edited"`
	Created                  int    `json:"created"`
	EditorID                 int    `json:"editor_id"`
	Views                    int    `json:"views"`
	Parent                   string `json:"parent"`
	Parent2                  string `json:"parent2"`
	Source                   string `json:"source"`
	HTML                     string `json:"html"`
	ViewURL                  string `json:"view_url"`
}
