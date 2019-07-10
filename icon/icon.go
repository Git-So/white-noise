package icon

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Git-So/white-noise/audio"
	"github.com/Git-So/white-noise/logger"
)

// Load 加载文件信息
func Load(file string) (infoList []Info) {
	// 读取文件
	data, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Error("加载自定义音频配置文件信息出错:", err)
		return
	}

	// 解析数据
	err = json.Unmarshal(data, &infoList)
	if err != nil {
		logger.Error("解析自定义音频配置文件文件出错:", err)
		return
	}

	return
}

// player 播放解析
func player() {
	// 获取播放队列
	stat := true
	for _, val := range selectedStack {
		index := val["index"].ToInt(&stat)
		iconInfo := iconList[index]
		for _, item := range iconInfo.AudioUrls {
			audioInfo := audio.Info{
				Name:   item.URL,
				Volume: val["volume"].ToFloat(&stat),
				State:  true,
			}
			audioInfo.AddPlayer()
		}
	}
}
