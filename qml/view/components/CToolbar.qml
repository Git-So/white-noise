import QtQuick 2.7

Item {
    visible: true
    anchors.fill: parent

    // 选中ID
    property int activeID: 0


    Row {
        height: 60

        anchors.fill: parent
        anchors.margins: 10
        spacing: 3

        // 中间选项
        Rectangle {
            anchors.horizontalCenter: parent.horizontalCenter

            Row {
                height: 60

                anchors.horizontalCenter: parent.horizontalCenter
                spacing: 20

                Text {
                    text: "推荐"
                    color: activeID != 0 ? "#ddd" : "#fff"

                    anchors.verticalCenter: parent.verticalCenter

                }

                Text {
                    text: "自选"
                    color: activeID != 1 ? "#ddd" : "#fff"

                    anchors.verticalCenter: parent.verticalCenter
                }
            }
        }

        // 右边选项
        Rectangle {
            anchors.top: parent.top
            anchors.right: parent.right

            // 关于
            Image {
                width: 60
                height: 60
                source: "qrc:/qml/drawable/action_mode_immersion_more_dark_n.png"

                anchors.top: parent.top
                anchors.right: parent.right

            }
        }

    }

}
