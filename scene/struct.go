package scene

// Info 场景信息
type Info struct {
	Title        string   `json:"title"`
	Desp         string   `json:"desp"`
	EngTitle     string   `json:"engTitle"`
	EngDesp      string   `json:"engDesp"`
	ImagePath    string   `json:"imagePath"`
	CropPicSmall *CropPic `json:"cropPicSmall"`
	CropPicLarge *CropPic `json:"cropPicLarge"`
}

// CropPic 裁剪图片
type CropPic struct {
	CenterLine float32 `json:"centerLine"`
	Top        float32 `json:"top"`
	Bottom     float32 `json:"bottom"`
	Left       float32 `json:"left"`
	Right      float32 `json:"right"`
}

// Audio 音频信息
type Audio struct {
	Name          string   `json:"audioName"`
	Duration      int      `json:"duration"`
	IsLine        bool     `json:"isLineAudio"`
	IsPoint       bool     `json:"isPointAudio"`
	IsRawRes      bool     `json:"isRawRes"`
	StartPlayTime int      `json:"startPlayTime"`
	TotalEndTime  int      `json:"totalEndTime"`
	Frequency     int      `json:"frequency"`
	Names         []string `json:"names"`
	Volume        float32  `json:"volume"`
	Pitch         float32  `json:"pitch"`
	Speed         float32  `json:"speed"`
}
