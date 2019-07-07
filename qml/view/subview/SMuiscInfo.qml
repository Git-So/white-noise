import QtQuick 2.7
import QtGraphicalEffects 1.13

import "../components"

Rectangle {
    id: muiscInfoRoot
    // width: parent.width
    // height: parent.height

    property int block: 45

    width: 9 * block
    height: 16 * block

    gradient: Gradient {
        GradientStop {
            position: 1
            color: "#ffE27DBF"
        }
        GradientStop {
            position: 0.0
            color: "#ffFCF075"
        }
    }

    // 栈对象
    property var stackObj: false

    // 调节状态
    property bool isSet: false

    // 音频数量
    property int musicNum: 2

    Column {
        width: parent.width
        height: parent.height
        spacing:5

        // 高度配置
        property int toolbarHeight: 80
        property int appInfoHeight: 180

        // 操作栏
        Rectangle {
            width: parent.width
            height: parent.toolbarHeight
            color: "transparent"

            Row {
                width: parent.width
                height: parent.height

                Rectangle {
                    width: parent.width
                    height: parent.height
                    color: "transparent"

                    Image {
                        width: 60
                        height: 60
                        x: 10
                        // source: "qrc:/qml/drawable/icon_editpage_finish_n.png"
                        source: "../../drawable/icon_editpage_finish_n.png"

                        anchors.verticalCenter: parent.verticalCenter

                        MouseArea {
                            anchors.fill: parent

                            onClicked: {
                                if (!muiscInfoRoot.isSet && stackObj) {
                                    stackObj.pop()
                                }
                            }
                        }
                    }
                }
            }
        }

        // 简单信息
        Rectangle {
            width: parent.width
            height: parent.appInfoHeight
            color: "transparent"
            // color: "red"
            visible: !muiscInfoRoot.isSet

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
                                    text: "自选"
                                    color: "#fff"
                                    font {
                                        pixelSize: 35
                                        bold: true
                                    }
                                }

                                // 介绍
                                Text {
                                    text: "夏雨/雨打树叶"
                                    color: "#eee"
                                    font.pixelSize: 17
                                }
                            }
                        }
                    }

                    // 音乐配置
                    Rectangle {
                        width: parent.width / 7 * 2
                        height: parent.height
                        color: "transparent"
                        // color: "#0f0"

                        anchors.right: parent.right

                        Image {
                            width: 70
                            height: 70
                            // source: "qrc:/qml/drawable/ic_audio_adjust_n.png"
                            source: "../../drawable/ic_audio_adjust_n.png"

                            anchors.verticalCenter: parent.verticalCenter

                            MouseArea {
                                anchors.fill: parent

                                onClicked: {
                                    muiscInfoRoot.isSet = true
                                }
                            }
                        }
                    }
                }
        }

        // 配置调节
        Rectangle {
            // color: "transparent"
            visible: muiscInfoRoot.isSet

            gradient: Gradient {
                GradientStop {
                    position: 1
                    color: "#ffE27DBF"
                }
                GradientStop {
                    position: 0.0
                    color: "#ffFCF075"
                }
            }

            anchors.fill: parent

            Rectangle {
                width: parent.width
                height: parent.height
                color: "transparent"
                // color: "#0f0"

                anchors.bottom: parent.bottom

                Column {
                    width: parent.width
                    height: parent.height
                    spacing:20

                    anchors.bottom: parent.bottom

                    property var closeBtnHeight: parent.height / 8

                    // 调节列表
                    Rectangle {
                        property var itemBlockheight: parent.height / 6

                        width: parent.width
                        height: itemBlockheight * muiscInfoRoot.musicNum
                        color: "transparent"
                        // color: "#f00"

                        anchors.bottomMargin: parent.closeBtnHeight
                        anchors.bottom: parent.bottom

                        Column {

                            anchors.fill: parent
                            anchors.bottom: parent.bottom

                            ListView {


                                model: [
                                    {
                                        name: "夏雨",
                                        size: 3
                                    },
                                    {
                                        name: "鱼打树叶",
                                        size: 3
                                    }
                                ]

                                delegate: CMuiscSet {
                                    width: parent.width
                                    height: parent.parent.parent.parent.itemBlockheight
                                }

                                anchors.fill: parent
                                anchors.bottom: parent.bottom
                            }
                        }
                    }

                    // 关闭按钮
                    Rectangle {
                        width: parent.width
                        height: parent.closeBtnHeight
                        color: "transparent"
                        // color: "#f00"
                        anchors.bottom: parent.bottom

                        Image {
                            width: 22
                            height: 22
                            // source: "qrc:/qml/drawable/timer_full_screen_shut_down.png"
                            source: "../../drawable/timer_full_screen_shut_down.png"

                            anchors.horizontalCenter: parent.horizontalCenter

                            MouseArea {
                                anchors.fill: parent

                                onClicked: {
                                    muiscInfoRoot.isSet = false
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}