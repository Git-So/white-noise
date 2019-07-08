package scene

import (
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
}

// init 初始化对象
func (o *Obj) init() {
	o.update(sceneID)
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
}
