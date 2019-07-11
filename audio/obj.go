package audio

import (
	"github.com/Git-So/white-noise/logger"
	"github.com/therecipe/qt/core"
)

var (
	// PlayState 播放状态
	PlayState = true

	playObj *Play
)

// Play 播放信息
type Play struct {
	core.QObject

	_ bool `property:"state"`

	_ func() `constructor:"init"`

	_ func() `signal:"stop,auto"`
	_ func() `signal:"start,auto"`
}

// init Play
func (py *Play) init() {
	py.SetState(PlayState)
}

// stop
func (py *Play) stop() {
	logger.Debug("播放停止")
	py.update(false)
	Disable()
}

// start
func (py *Play) start() {
	logger.Debug("播放开始")
	py.update(true)
	Enable()
}

// update
func (py *Play) update(stat bool) {
	logger.Debug("更新播放状态")
	PlayState = stat
	py.SetState(PlayState)
}
