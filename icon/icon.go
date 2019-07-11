package icon

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Git-So/white-noise/audio"
	"github.com/Git-So/white-noise/logger"
	"github.com/therecipe/qt/core"
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

// addPlayer 添加音频播放
func addPlayer(idx int) {
	stackInfo := getStackInfo(idx)
	if stackInfo == nil {
		logger.Error("获取不到栈数据信息")
	}

	// p := audio.GetPlayer(info["name"].ToString())

	// 获取音频信息
	iconInfo := iconList[idx]
	stat := true
	for _, val := range iconInfo.AudioUrls {
		audio.AddPlayer(
			val.URL,
			stackInfo["volume"].ToFloat(&stat),
			val.Duration,
			audio.PLAYER_ENABLE,
		)
	}

}

// removePlayer 移除音频播放
func removePlayer(idx int) {
	stackInfo := getStackInfo(idx)
	if stackInfo == nil {
		logger.Error("获取不到栈数据信息")
	}

	// 获取音频信息
	iconInfo := iconList[idx]
	// stat := true
	for _, val := range iconInfo.AudioUrls {
		// 获取播放实例
		p := audio.GetPlayer(val.URL)
		// 停止播放
		p.SetState(audio.PLAYER_DISABLE)
	}
}

// updateVolume 更新播放音量
func updateVolume(idx int, volume float32) {
	stackInfo := getStackInfo(idx)
	if stackInfo == nil {
		logger.Error("获取不到栈数据信息")
	}

	// 获取音频信息
	iconInfo := iconList[idx]
	// stat := true
	for _, val := range iconInfo.AudioUrls {
		// 获取播放实例
		p := audio.GetPlayer(val.URL)
		// 停止播放
		p.SetVolume(volume)
	}
}

// getStackInfo
func getStackInfo(idx int) (info map[string]*core.QVariant) {

	// 获取栈中对应数据
	// 移除前获取
	// 添加后获取
	for _, val := range selectedStack {
		if _, ok := val["index"]; ok {
			if idx == val["index"].ToInt(&ok) {
				info = val
			}
		}
	}

	return
}
