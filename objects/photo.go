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
