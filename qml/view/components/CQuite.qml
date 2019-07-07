import QtQuick 2.7

// 禁音
Rectangle {
    property var quite: true

    width: parent.width
    height: parent.height / 1.2
    color: "transparent"
    visible: quite

    // anchors.fill: parent

    Image {
        width: 80
        height: 80

        // source: "qrc:/qml/drawable/ic_play.png"
        source: "../../drawable/ic_play.png"

        anchors.centerIn: parent
    }
}
