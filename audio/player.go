package audio

import (
	"time"

	"github.com/Git-So/white-noise/logger"
	"github.com/bobertlo/go-mpg123/mpg123"
	"github.com/hajimehoshi/oto"
)

const (
	// PLAYER_ENABLE 播放器启用
	PLAYER_ENABLE = true
	// PLAYER_DISABLE 播放器停用
	PLAYER_DISABLE = false
)

// 播放器列表
var playerList = make(map[string]*player, 100)

// player 播放器信息
type player struct {
	name     string
	duration int // 我怀疑持续时间就是间接时间，暂时当间接时间处理
	volume   float32
	state    bool
}

// AddPlayer 添加播放音频到播放器
func AddPlayer(name string, volume float32, duration int, state bool) {
	// 不存在添加到列表
	if _, ok := playerList[name]; !ok {
		playerList[name] = &player{
			name:     name,
			duration: 0,
			volume:   volume,
			state:    state,
		}
	}

	// 启用播放
	Enable()

	// 开始播放
	playerList[name].enablePlayer()
}

// GetPlayer 获取播放器实例
func GetPlayer(name string) (info *player) {
	if _, ok := playerList[name]; ok {
		info = playerList[name]
	}
	return
}

// SetVolume 设置播放器音量
func (p *player) SetVolume(volume float32) {
	logger.Debug(p.name, "音量设置为", volume)
	p.volume = volume
}

// SetState 设置播放状态
func (p *player) SetState(state bool) {
	logger.Debug(p.name, "播放状态改变为", state)
	p.state = state
}

// enablePlayer 启用播放状态
func (p *player) enablePlayer() {
	logger.Debug(p.name, "启用播放,音量", p.volume)

	// 启用播放
	p.state = PLAYER_ENABLE

	// 按一定时间循环播放
	go func() {
		// 播放器有效状态播放
		if p.state == PLAYER_ENABLE && PlayState {
			// 等待播放
			time.Sleep(time.Duration(p.duration) * time.Millisecond)

			// 添加播放实例
			play := otoObj.NewPlayer()
			defer play.Close()
			p.player(play)
		}
	}()
}

// player 播放实例
func (p *player) player(op *oto.Player) {
	// 音频解码实例
	decoder, err := mpg123.NewDecoder("")
	if err != nil {
		logger.Error("初始化音频解码失败:", err)
	}

	// 获取解码文件
	fileName := getFileName(p.name)
	err = decoder.Open(fileName)
	if err != nil {
		logger.Error("无法打开文件", fileName, ":", err)
	}
	defer decoder.Close()

	// 获取音频格式信息
	rate, chans, _ := decoder.GetFormat()
	logger.Info("Encoding: Signed 16bit")
	logger.Info("Sample Rate:", rate)
	logger.Info("Channels:", chans)

	// make sure output format does not change
	decoder.FormatNone()
	decoder.Format(rate, chans, mpg123.ENC_SIGNED_16)

	// 暂存区
	buf := make([]byte, 2048*16)

	// 解析播放音频
	for {
		// 状态改变停止播放
		if p.state == PLAYER_DISABLE || !PlayState {
			// decoder.Delete()
			break
		}

		// 读取流
		len, err := decoder.Read(buf)
		if err != nil {
			break
		}

		// 写入流
		op.Write(raiseVolume(buf[0:len], p.volume))
	}
}
