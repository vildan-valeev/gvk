package objects

// BaseSticker https://dev.vk.com/ru/reference/objects/sticker.
type Sticker struct {
	ProductID            int            `json:"product_id"`
	StickerID            int            `json:"sticker_id"`
	Images               []StickerImage `json:"images"`
	ImagesWithBackground []StickerImage `json:"images_with_background"`
	IsAllowed            bool           `json:"is_allowed"`
	AnimationURL         string         `json:"animation_url"`
}

type StickerImage struct {
	Height float64 `json:"height"`
	URL    string  `json:"url"`
	Width  float64 `json:"width"`
}
