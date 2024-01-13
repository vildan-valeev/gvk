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
