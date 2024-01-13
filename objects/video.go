package objects

// Video https://dev.vk.com/ru/reference/objects/video.
type Video struct {
	ID            int          `json:"id"`          // Video ID
	OwnerID       int          `json:"owner_id"`    // Video owner ID
	Title         string       `json:"title"`       // Video title
	Description   string       `json:"description"` // Video description
	Duration      int          `json:"duration"`    // Video duration in seconds
	Image         []VideoImage `json:"image"`
	FirstFrame    []FrameImage `json:"first_frame"`
	Date          int          `json:"date"`        // Date when video has been uploaded in Unixtime
	AddingDate    int          `json:"adding_date"` // Date when the video has been added in Unixtime.
	Views         int          `json:"views"`       // Number of views
	LocalViews    int          `json:"local_views"`
	Comments      int          `json:"comments"` // Number of comments
	Player        string       `json:"player"`   // URL of the page with a player that can be used to play the video in the browser.
	Platform      string       `json:"platform"`
	CanAdd        int          `json:"can_add"`    // Может ли пользователь добавить видеозапись к себе. 0 — не может добавить. 1 — может добавить.
	IsPrivate     int          `json:"is_private"` //  Поле возвращается, если видеозапись приватная (например, была загружена в личное сообщение), всегда содержит 1
	AccessKey     string       `json:"access_key"` // Video access key.
	Processing    int          `json:"processing"` // Returns if the video is processing
	IsFavorite    bool         `json:"is_favorite"`
	CanComment    int          `json:"can_comment"`      // Может ли пользователь комментировать видео.0 — не может комментировать. 1 — может комментировать.
	CanEdit       int          `json:"can_edit"`         // Может ли пользователь редактировать видео. 0 — не может редактировать, 1 — может редактировать.
	CanLike       int          `json:"can_like"`         // Может ли пользователь добавить видео в список <<Мне нравится>>. 0 — не может добавить. 1 — может добавить.
	CanRepost     int          `json:"can_repost"`       // Может ли пользователь сделать репост видео. 0 — не может сделать репост/ 1 — может сделать репост.
	CanSubscribe  int          `json:"can_subscribe"`    // Может ли пользователь подписаться на автора видео. 0 — не может подписаться. 1 — может подписаться.
	CanAddToFaves int          `json:"can_add_to_faves"` // Может ли пользователь подписаться на автора видео. 0 — не может подписаться. 1 — может подписаться.
	CanAttachLink int          `json:"can_attach_link"`  // Может ли пользователь прикрепить кнопку действия к видео. 0 — не может, 1 — может
	Width         int          `json:"width"`            // Video width
	Height        int          `json:"height"`           // Video height
	UserID        int          `json:"user_id"`
	Converting    int          `json:"converting"`    // Конвертируется ли видео. 0 — не конвертируется. 1 — конвертируется.
	Added         int          `json:"added"`         // Добавлено ли видео в альбомы пользователя. 0 — не добавлено. 1 — добавлено.
	IsSubscribed  int          `json:"is_subscribed"` // Подписан ли пользователь на автора видео. 0 — не подписан. 1 — подписан.
	Repeat        int          `json:"repeat"`        // Поле возвращается в том случае, если видео зациклено, всегда содержит 1
	Type          string       `json:"type"`          // Тип видеозаписи. Может принимать значения: video, music_video, movie
	Balance       int          `json:"balance"`       //  Баланс донатов в прямой трансляции.
	LiveStatus    string       `json:"live_status"`   // Статус прямой трансляции. Может принимать значения: waiting, started, finished, failed, upcoming.
	Live          int          `json:"live"`          // Поле возвращается в том случае, если видеозапись является прямой трансляцией, всегда содержит 1. Обратите внимание, в этом случае в поле duration содержится значение 0.
	Upcoming      int          `json:"upcoming"`
	Spectators    int          `json:"spectators"`
	Likes         VideoLikes   `json:"likes"`   // Count of likes
	Reposts       VideoReposts `json:"reposts"` // Count of views
}

// VideoImage struct.
type VideoImage struct {
	Height      float64 `json:"height"`
	URL         string  `json:"url"`
	Width       float64 `json:"width"`
	WithPadding int     `json:"with_padding"`
}

type FrameImage struct {
	Height float64 `json:"height"`
	URL    string  `json:"url"`
	Width  float64 `json:"width"`
}

// VideoLikes https://dev.vk.com/ru/reference/objects/video#likes
type VideoLikes struct {
	UserLikes int `json:"user_likes"` // добавлено ли видео в список «Мне нравится» текущего пользователя. 0 — не добавлено. 1 — добавлено.
	Count     int `json:"count"`
}

// VideoReposts https://dev.vk.com/ru/reference/objects/video#reposts
type VideoReposts struct {
	Count        int `json:"count"`
	WallCount    int `json:"wall_count"`
	MailCount    int `json:"mail_count"`
	UserReposted int `json:"user_reposted"`
}
