import QtQuick 2.7
import "../components"

Rectangle {

    // property string imgUrl: "qrc:/qml/scene/森林.jpg"
    property string imgUrl: "../../scene/森林.jpg"

    property var customizeID: false
    property var aboutID: false
    property var stackObj: false

    property var quite: false

    // 背景
    Image {
        source: imgUrl

        asynchronous: true
        cache: false
        fillMode: Image.PreserveAspectCrop

        anchors.fill: parent
    }

    // 禁音
    Rectangle {
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

    // 声音详情
    Rectangle {
        width: parent.width
        height: 160
        color: "transparent"
        // color: "#f00"

        anchors.bottom: parent.bottom

        Row {
            width: parent.width
            height: parent.height

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
                    height: parent.height / 4
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

        customizeID: parent.customizeID
        aboutID: parent.aboutID
        stackObj: parent.stackObj
        activeID: 0
    }
}
