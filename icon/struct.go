package icon

// Info 详情
type Info struct {
	Title      string  `json:"title"`
	EngTitle   string  `json:"engTitle"`
	ColorParam string  `json:"colorParam"`
	Icon       string  `json:"icon"`
	Frequency  int     `json:"frequency"`
	Volume     float32 `json:"volume"`
	Type       int     `json:"type"`
	AudioUrls  []Audio `json:"audioUrls"`
	IsActive   bool    `json:"isActive"`
}

// Audio 音频信息
type Audio struct {
	URL      string `json:"url"`
	Duration int    `json:"duration"`
}
