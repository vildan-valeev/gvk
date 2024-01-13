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
