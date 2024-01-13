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
