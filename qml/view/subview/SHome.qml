import QtQuick 2.7
import "../components"

Rectangle {

    property string imgUrl: "qrc:/qml/scene/森林.jpg"

    // 背景
    Image {
        source: imgUrl

        asynchronous: true
        cache: false
        fillMode: Image.PreserveAspectCrop

        anchors.fill: parent
    }

    // 工具栏
    CToolbar {
        width: parent.width
        height: 100
    }
}
