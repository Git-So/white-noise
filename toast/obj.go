package toast

import (
	"github.com/Git-So/white-noise/logger"
	"github.com/therecipe/qt/core"
)

var toastObj *Obj

// Obj struct
type Obj struct {
	core.QObject

	_ func() `constructor:"init"`

	_ bool   `property:"visible"`
	_ string `property:"msg"`
}

// init Obj
func (o *Obj) init() {
	toastObj = o
	o.SetVisible(false)
	o.SetMsg("true")
}

// Msg 发送消息
func Msg(msg string) {
	logger.Debug("发送 toast ：", msg)
	if toastObj == nil {
		logger.Debug("toast 对象不存在 ")
		return
	}
	toastObj.SetVisible(false)
	toastObj.SetVisible(true)
	toastObj.SetMsg(msg)
	// toastObj.show(msg, 1000)
}
