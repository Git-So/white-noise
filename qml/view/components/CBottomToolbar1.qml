import QtQuick 2.7

Rectangle {
    color: "transparent"

    Rectangle {
        width: parent.width * 0.88
        height: parent.height * 0.68
        color: "#282c34"
        radius: 25

        anchors.horizontalCenter: parent.horizontalCenter

        Row {
            width: parent.width * 0.85
            height: parent.height
            spacing: 10

            anchors.centerIn: parent

            //  图标
            Image {
                width: 13
                height: width * 1.5

                // source: "qrc:/qml/drawable/icon_music.png"
                source: "../../drawable/icon_music.png"

                anchors.verticalCenter: parent.verticalCenter
            }

            // 提示
            Text {
                text: "选择喜欢的音乐"
                color: "#fff"

                anchors.leftMargin: 20
                anchors.verticalCenter: parent.verticalCenter
            }
        }
    }
}
