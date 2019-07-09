package icon

import "fmt"

var (
	iconPath    = "config/icons/"
	sceneConfig = "prebuilt_icon_config"
)

func init() {
	file := fmt.Sprintf("%s%s.json", iconPath, sceneConfig)
	iconList = Load(file)
}

func init() {
	IconListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "IconListModel")
}

func init() {
	SelectedStackModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "SelectedStackModel")
}
