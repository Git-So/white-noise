package audio

import (
	"github.com/Git-So/white-noise/logger"
	"github.com/therecipe/qt/core"
)

const (
	SCENE = iota
	ICON
)

var (
	// PlayState 播放状态
	PlayState = true
	// PlayType 播放类型
	PlayType = SCENE
)

// Play 播放信息
type Play struct {
	core.QObject

	_ bool `property:"state"`

	_ func() `constructor:"init"`

	_ func()                  `signal:"stop,auto"`
	_ func()                  `signal:"start,auto"`
	_ func(tp *core.QVariant) `signal:"setType,auto"`
}

// init Play
func (p *Play) init() {
	p.SetState(PlayState)
}

// stop
func (p *Play) stop() {
	logger.Debug("播放停止")
	p.update(false)
	Close()
}

// start
func (p *Play) start() {
	logger.Debug("播放开始")
	p.update(true)
}

// update
func (p *Play) update(stat bool) {
	logger.Debug("更新播放状态")
	PlayState = stat
	p.SetState(PlayState)
}

// type
func (p *Play) setType(tp *core.QVariant) {
	stat := true
	PlayType = tp.ToInt(&stat)
}
