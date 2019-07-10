package audio

import (
	"github.com/Git-So/white-noise/logger"
	"github.com/hajimehoshi/oto"
)

var (
	otoObj *oto.Context
)

func init() {
	if otoObj != nil {
		return
	}
	// 初始化播放
	var err error
	otoObj, err = oto.NewContext(48000, 2, 2, 8196)
	if err != nil {
		logger.Error("初始化播放出错: ", err)
	}
	logger.Debug("初始化播放成功")
}
