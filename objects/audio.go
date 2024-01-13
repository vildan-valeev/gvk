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
