package scene

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Git-So/white-noise/audio"
	"github.com/Git-So/white-noise/icon"
	"github.com/Git-So/white-noise/logger"
)

var (
	scenePath   = "config/scene/"
	sceneConfig = "prebuilt_scene_config"
)

func init() {
	file := fmt.Sprintf("%s%s.json", scenePath, sceneConfig)
	sceneList = Load(file)
}

// Load 加载文件信息
func Load(file string) (infoList []*Info) {
	// 读取文件
	data, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Error("加载场景文件信息出错:", err)
		return
	}

	// 解析数据
	err = json.Unmarshal(data, &infoList)
	if err != nil {
		logger.Error("解析场景文件出错:", err)
		return
	}

	return
}

// LoadAudio 加载音频配置
func LoadAudio(sceneName string) (audioList []*Audio) {
	// 读取文件
	file := getAudioFilePath(sceneName)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Error("加载音频配置文件信息出错:", err)
		return
	}

	// 解析数据
	err = json.Unmarshal(data, &audioList)
	if err != nil {
		logger.Error("解析音频配置文件出错:", err)
		return
	}

	return
}

// getAudioFilePath 获取音频文件
func getAudioFilePath(sceneName string) (filePath string) {
	return fmt.Sprintf("%s%s.json", scenePath, sceneName)
}

// player 播放场景
func player() {
	// 停止当前播放
	audio.CloseAllPlayer()

	// 播放当前场景
	sceneInfo := sceneList[sceneID]

	// 加载场景音频文件
	audioList := LoadAudio(sceneInfo.Title)
	if audioList == nil {
		logger.Error("加载场景音频文件信息出错")
		return
	}

	for _, audioInfo := range audioList {
		// 加载自定义音乐
		if len(audioInfo.Name) > 0 {
			icon.AddPlayer(audioInfo.Name)
		}

		// 加载系列音频
		for _, name := range audioInfo.Names {
			audio.AddPlayer(
				name,
				audioInfo.Volume,
				audioInfo.Duration,
				audio.PLAYER_ENABLE,
			)
		}
	}

}
