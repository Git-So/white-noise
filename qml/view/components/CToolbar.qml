import QtQuick 2.7

Item {
    id: toolbarRoot

    anchors.fill: parent

    // 选中ID
    property int activeID: 0

    // 栈对象
    property var stackObj: false

    // 关于界面ID
    property var aboutID: false

    // 自定义界面ID
    property var customizeID: false


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
                spacing: 40

                Text {
                    text: qsTr("推荐")
                    color: toolbarRoot.activeID != 0 ? "#ddd" : "#fff"

                    anchors.verticalCenter: parent.verticalCenter

                    MouseArea {
                        anchors.fill: parent
                        onClicked: {
                            if (toolbarRoot.stackObj && toolbarRoot.activeID != 0 ) {
                                stackObj.pop()
                            }
                        }
                    }
                }

                Text {
                    z: 100
                    text: qsTr("自选")
                    color: activeID != 1 ? "#ddd" : "#fff"

                    anchors.verticalCenter: parent.verticalCenter

                    MouseArea {
                        anchors.fill: parent

                        onClicked: {

                            if (toolbarRoot.stackObj && toolbarRoot.customizeID && toolbarRoot.activeID != 1 ) {
                                toolbarRoot.stackObj.push(customizeID)
                            }
                        }
                    }
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

                MouseArea {
                    anchors.fill: parent

                    onClicked: {
                        if (toolbarRoot.stackObj && toolbarRoot.aboutID) {
                            toolbarRoot.stackObj.push(aboutID)
                        }
                    }
                }
            }
        }

    }

}
