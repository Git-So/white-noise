import QtQuick 2.7
import "../components"

Rectangle {
    id: homeRoot

    // property string imgUrl: "qrc:/qml/scene/森林.jpg"
    property string imgUrl: "../../scene/森林.jpg"

    property var customizeID: false
    property var aboutID: false
    property var stackObj: false

    property var quite: false

    // 背景
    Image {
        source: homeRoot.imgUrl

        asynchronous: true
        cache: false
        fillMode: Image.PreserveAspectCrop

        anchors.fill: parent
    }

    // 禁音
    CQuite {
        quite: homeRoot.quite
    }

    // 声音详情
    Rectangle {
        width: parent.width
        height: 180
        color: "transparent"
        // color: "#f00"

        anchors.bottom: parent.bottom

        Row {
            width: parent.width - 40
            height: parent.height

            anchors.horizontalCenter: parent.horizontalCenter

            Rectangle {
                width: parent.width / 7 * 5
                height: parent.height
                color: "transparent"
                // color: "#00f"

                anchors.left: parent.left
                anchors.margins: 20

                Rectangle {
                    width: parent.width
                    height: parent.height / 2.1
                    color: "transparent"

                    anchors.verticalCenter: parent.verticalCenter

                    Column {
                        // 标题
                        Text {
                            text: "森林"
                            color: "#fff"
                            font {
                                pixelSize: 35
                                bold: true
                            }
                        }

                        // 介绍
                        Text {
                            text: "把你的秘密藏进森林里"
                            color: "#eee"
                            font.pixelSize: 17
                        }
                    }
                }
            }

            // 下一场景
            Rectangle {
                width: parent.width / 7 * 2
                height: parent.height
                color: "transparent"
                // color: "#0f0"

                anchors.right: parent.right

                Rectangle {
                    width: parent.width / 10 * 8
                    height: parent.height / 4.2
                    radius: 20
                    color: "transparent"
                    border {
                        width: 2
                        color: "#eee"
                    }

                    anchors.verticalCenter: parent.verticalCenter

                    Text {
                        text: "下一场景"
                        color: "#fff"

                        anchors.centerIn: parent
                    }
                }
            }
        }
    }

    // 工具栏
    CToolbar {
        width: parent.width
        height: 100

        customizeID: homeRoot.customizeID
        aboutID: homeRoot.aboutID
        stackObj: homeRoot.stackObj
        activeID: 0
    }
}
