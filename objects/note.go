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
