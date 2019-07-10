package audio

import (
	"fmt"

	"github.com/Git-So/white-noise/logger"
	"github.com/bobertlo/go-mpg123/mpg123"
)

var (
	iconPath = "config/audio/"
)

// Info 声音详情
type Info struct {
	Name   string
	Volume float32
	State  bool
}

// RaiseVolume 音量增益
func RaiseVolume(buf []byte, volume float32) (data []byte) {
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
		result := float32(value) / volume
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

// Close 关闭全部播放
func Close() {
	if otoObj == nil {
		return
	}
	err := otoObj.Close()
	if err != nil {
		logger.Error("关闭播放实例失败:", err)
	}
}

// AddPlayer 添加播放
func (i *Info) AddPlayer() {
	// 播放状态
	i.State = true
	// logger.Debug("添加到", i.Name, "播放器")
	logger.Debug("添加播放")
	// 循环播放
	go func() {
		// 播放
		if i.State {
			// 播放
			i.player()
		}
	}()
}

// player 播放音频
func (i *Info) player() {
	// 音频编码
	decoder, err := mpg123.NewDecoder("")
	if err != nil {
		panic(err)
	}

	fileName := getFileName(i.Name)
	err = decoder.Open(fileName)
	if err != nil {
		logger.Debug("无法打开文件", fileName, ":", err)
	}
	defer decoder.Close()

	// get audio format information
	rate, chans, _ := decoder.GetFormat()
	logger.Info("Encoding: Signed 16bit")
	logger.Info("Sample Rate:", rate)
	logger.Info("Channels:", chans)

	// make sure output format does not change
	decoder.FormatNone()
	decoder.Format(rate, chans, mpg123.ENC_SIGNED_16)

	buf := make([]byte, 2048*16)
	p := otoObj.NewPlayer()
	defer p.Close()

	for {
		// 停止播放
		if !i.State {
			break
		}

		len, err := decoder.Read(buf)
		if err != nil {
			// decoder.Delete()
			break
		}
		p.Write(RaiseVolume(buf[0:len], i.Volume))
	}
}

// ClosePlayer 关闭播放
func (i *Info) ClosePlayer() {
	i.State = false
}

// getFileName 获取播放文件路径
func getFileName(name string) string {
	return fmt.Sprintf("%v%v.mp3", iconPath, name)
}
