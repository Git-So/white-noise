package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"

	"github.com/Git-So/white-noise/icon"
	"github.com/Git-So/white-noise/scene"
	"github.com/Git-So/white-noise/toast"
)

func main() {
	// core
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// gui
	gui.NewQGuiApplication(len(os.Args), os.Args)

	// style
	quickcontrols2.QQuickStyle_SetStyle("Material")

	engine := qml.NewQQmlApplicationEngine(nil)
	engine.RootContext().SetContextProperty("sceneObj", scene.NewObj(nil))
	engine.RootContext().SetContextProperty("iconStack", icon.NewIconStack(nil))
	engine.RootContext().SetContextProperty("toastObj", toast.NewObj(nil))
	engine.Load(core.NewQUrl3("qrc:/qml/view/App.qml", 0))

	gui.QGuiApplication_Exec()
}
