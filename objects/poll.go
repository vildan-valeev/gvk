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

// Poll struct.
type Poll struct {
	ID         int            `json:"id"`       // Идентификатор опроса для получения информации о нем через метод polls.getById.
	OwnerID    int            `json:"owner_id"` // Poll owner's ID
	Created    int            `json:"created"`  // Date when poll has been created in Unixtime
	Question   string         `json:"question"` // Poll question
	Votes      int            `json:"votes"`    // Votes number
	Answers    []PollsAnswer  `json:"answers"`
	Anonymous  bool           `json:"anonymous"` // Information whether the pole is anonymous
	Multiple   bool           `json:"multiple"`
	AnswerIDs  []int          `json:"answer_ids"`
	EndDate    int            `json:"end_date"`
	Closed     bool           `json:"closed"`
	IsBoard    bool           `json:"is_board"`
	CanEdit    bool           `json:"can_edit"`
	CanVote    bool           `json:"can_vote"`
	CanReport  bool           `json:"can_report"`
	CanShare   bool           `json:"can_share"`
	AuthorID   int            `json:"author_id"`
	Photo      Photo          `json:"photo"`
	Background PollBackground `json:"background"`
	Friends    []PollFriend   `json:"friends"`
}

// PollsAnswer struct.
type PollsAnswer struct {
	ID    int     `json:"id"`
	Rate  float64 `json:"rate"`
	Text  string  `json:"text"`
	Votes int     `json:"votes"`
}

// PollBackground struct.
type PollBackground struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Angle  int    `json:"angle"`
	Color  string `json:"color"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	//Images []
	Points []struct {
		Position float64 `json:"position"`
		Color    string  `json:"color"`
	} `json:"points"`
}

// PollFriend struct.
type PollFriend struct {
	ID int `json:"id"`
}
