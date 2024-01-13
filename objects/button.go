package objects

type Button struct {
	Title  string `json:"title"`
	Action Action `json:"action"`
}

type Action struct {
	Type string `json:"type"`
	URL  string `json:"URL"`
}
