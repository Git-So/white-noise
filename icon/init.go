package icon

import (
	"fmt"

	"github.com/Git-So/white-noise/logger"
)

var (
	iconPath    = "config/icons/"
	sceneConfig = "icons"
)

func init() {
	file := fmt.Sprintf("%s%s.json", iconPath, sceneConfig)
	iconList = Load(file)
	logger.Debug(fmt.Sprintf("%v", iconList))
}

func init() {
	IconListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "IconListModel")
}

func init() {
	SelectedStackModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "SelectedStackModel")
}
