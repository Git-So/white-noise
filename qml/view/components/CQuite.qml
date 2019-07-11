import QtQuick 2.7

// 禁音
Rectangle {
    property var quite: false

    width: parent.width
    height: parent.height / 1.5
    y: 100
    color: "transparent"
    // color: "red"
    // anchors.fill: parent

    Image {
        width: 80
        height: 80
        visible: !quite

        // source: "qrc:/qml/drawable/ic_play.png"
        source: "../../drawable/ic_play.png"

        anchors.centerIn: parent
    }

    // 切换播放状态
    MouseArea {
        anchors.fill: parent

        onClicked: {
            if (playObj.state) {
                playObj.stop()
            } else {
                playObj.start()
            }
        }
    }
}
