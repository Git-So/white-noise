package scene

import (
	"github.com/Git-So/white-noise/logger"
	"github.com/therecipe/qt/core"
)

// sceneID
var sceneID int
var sceneList []*Info

// Obj 场景对象
type Obj struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"title"`
	_ string `property:"desp"`
	_ string `property:"imagePath"`

	_ func() `signal:"next,auto"`
	_ func() `signal:"play,auto"`
}

// init 初始化对象
func (o *Obj) init() {
	o.update(sceneID)

	// 播放场景音频
	o.play()
}

// update 更新对象
func (o *Obj) update(sceneID int) {
	info := *(sceneList[sceneID])
	o.SetTitle(info.Title)
	o.SetDesp(info.Desp)
	o.SetImagePath(info.ImagePath)
}

// next 下一场景
func (o *Obj) next() {
	// 场景循环切换
	if sceneID == len(sceneList)-1 {
		sceneID = 0
	} else {
		sceneID++
	}
	o.update(sceneID)

	// 播放场景音频
	o.play()
}

// play 播放场景音乐
func (o *Obj) play() {
	logger.Debug("播放场景音乐")
	player()
}
