package objects

// Doc https://dev.vk.com/ru/reference/objects/doc
type Doc struct {
	ID      int        `json:"id"`       // Document ID
	OwnerID int        `json:"owner_id"` // Document owner ID
	Title   string     `json:"title"`    // Document title
	Size    int        `json:"size"`     // File size in bites
	Ext     string     `json:"ext"`      // File extension
	URL     string     `json:"url"`      // File URL
	Date    int        `json:"date"`     // Date when file has been uploaded in Unixtime
	Type    int        `json:"type"`     // Document type
	Preview DocPreview `json:"preview"`
}

// DocPreview https://dev.vk.com/ru/reference/objects/doc#preview.
type DocPreview struct {
	Photo        DocPreviewPhoto        `json:"photo"`
	Graffiti     DocPreviewGraffiti     `json:"graffiti"`
	AudioMessage DocPreviewAudioMessage `json:"audio_message"`
}

type DocPreviewPhoto struct {
	Sizes []DocPreviewPhotoSizes `json:"sizes"`
}

// DocPreviewPhotoSizes https://dev.vk.com/ru/reference/objects/photo-sizes.
type DocPreviewPhotoSizes struct {
	Height int    `json:"height"` // Height in px
	URL    string `json:"url"`    // URL of the image
	Type   string `json:"type"`
	Width  int    `json:"width"` // Width in px
}

type DocPreviewGraffiti struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// DocPreviewAudioMessage struct.
type DocPreviewAudioMessage struct {
	Duration int    `json:"duration"`
	Waveform []int  `json:"waveform"`
	LinkOgg  string `json:"link_ogg"`
	LinkMp3  string `json:"link_mp3"`
}
