package scene

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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
func LoadAudio(sceneName string) (audioList []*Audio, err error) {
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
