package audio

import (
	"fmt"

	"github.com/Git-So/white-noise/logger"
	"github.com/hajimehoshi/oto"
)

var (
	otoObj *oto.Context

	iconPath = "config/audio/"
)

func init() {
	// 默认启用播放
	Enable()
}

// getFileName 获取播放文件路径
func getFileName(name string) string {
	return fmt.Sprintf("%v%v.mp3", iconPath, name)
}

// raiseVolume 音量增益
func raiseVolume(buf []byte, volume float32) (data []byte) {
	data = buf
	var minData int16 = -0x8000
	var maxData int16 = 0x7FFF

	bufLen := len(buf)

	if bufLen < 2 {
		return
	}

	var key int
	// fmt.Println(bufLen)
	for key < bufLen {
		value := int16(buf[key])&0xFF | int16(buf[key+1])<<8
		result := float32(value) * volume
		if result > float32(maxData) {
			value = maxData
		} else if result < float32(minData) {
			value = minData
		} else {
			value = int16(result)
		}
		data[key] = byte(value & 0xFF)
		data[key+1] = byte(value >> 8 & 0xFF)
		key += 2
	}

	return
}

// Disable 关闭全部播放
func Disable() {
	if otoObj == nil {
		return
	}
	err := otoObj.Close()
	if err != nil {
		logger.Error("关闭播放实例失败:", err)
	}
	otoObj = nil
}

// Enable 打开播放实例
func Enable() {
	// 打开全局播放
	PlayState = true

	if otoObj != nil {
		return
	}
	// 初始化播放
	var err error
	otoObj, err = oto.NewContext(48000, 2, 2, 8196)
	if err != nil {
		logger.Error("初始化播放出错: ", err)
	}

	logger.Info("初始化播放成功")

	if playAudioObj != nil {
		playAudioObj.update(true)
	}

	// 播放队伍音频
	for key, val := range playerList {
		if val.state == PLAYER_ENABLE {
			playerList[key].enablePlayer()
		}
	}
}
