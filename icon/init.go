package icon

import "fmt"

func init() {
	file := fmt.Sprintf("%s%s.json", iconPath, sceneConfig)
	iconList = Load(file)
}

func init() {
	IconListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "IconListModel")
}
