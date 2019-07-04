import QtQuick 2.7
import QtQuick.Controls 2.5
import "../components"

Rectangle {
    // width: parent.width
    // height: parent.height

    property int block: 45

    width: 9 * block
    height: 16 * block

    color: "#212121"

    // 栈对象
    property var stackObj: false

    Column {
        width: parent.width
        height: parent.height
        spacing:5

        // 高度配置
        property int toolbarHeight: 80
        property int appInfoHeight: 260
        property int aboutListHeight: parent.height - toolbarHeight - appInfoHeight - 20

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
                        source: "qrc:/qml/drawable/action_bar_back_normal_dark.png"
                        // source: "../../drawable/action_bar_back_normal_dark.png"

                        anchors.verticalCenter: parent.verticalCenter

                        MouseArea {
                            anchors.fill: parent

                            onClicked: {
                                if (stackObj) {
                                    stackObj.pop()
                                }
                            }
                        }
                    }
                }
            }
        }

        // 应用信息
        Rectangle {
            width: parent.width
            height: parent.appInfoHeight
            color: "transparent"

            Column {
                width: parent.width

                // 应用图标
                Image {
                    width: 100
                    height: 100

                    source: "qrc:/qml/drawable/setting_icon.png"
                    // source: "../../drawable/setting_icon.png"

                    anchors.horizontalCenter: parent.horizontalCenter

                }

                // 应用名称
                Rectangle {
                    width: parent.width
                    height: 40
                    color: "transparent"

                    Text {
                        text: qsTr("白噪音")
                        color: "#fff"
                        font {
                            pixelSize: 20
                            bold: true
                        }

                        anchors.centerIn: parent
                    }
                }

                // 应用版本
                Rectangle {
                    width: parent.width
                    height: 30
                    color: "transparent"

                    Text {
                        text: qsTr("版本号: 0.1")
                        color: "#aaa"

                        anchors.centerIn: parent
                    }
                }

                // 检查新版本按钮
                Rectangle {
                    width: parent.width
                    height: 80
                    color: "transparent"

                    Rectangle {
                        width: 240
                        height: 50
                        color: "#14d685"
                        radius: 10

                        anchors.centerIn: parent

                        Text {
                            text: qsTr("检查新版本")
                            color: "#fff"

                            anchors.centerIn: parent
                        }

                        MouseArea {
                            anchors.fill: parent

                            onClicked: {
                                Qt.openUrlExternally("https://github.com/Git-So/white-noise/releases")
                            }
                        }
                    }
                }
            }
        }

        // 关于列表
        Rectangle {
            width: parent.width
            height: parent.aboutListHeight
            color: "transparent"

            ScrollView {
                width: parent.width - 10
                height: parent.height
                clip: true

                anchors.horizontalCenter: parent.horizontalCenter

                Column {
                    width: parent.width
                    height: parent.height

                    // 应用相关
                    Rectangle {
                        width: 1
                        height: 15
                        color: "transparent"
                    }
                    Rectangle {
                        width: parent.width
                        height: 120
                        color: "#282c34"
                        radius: 15

                        Column {
                            width: parent.width
                            height: parent.height

                            CListItem1 {
                                name: "项目主页"
                                value: "https://github.com/Git-So/white-noise"
                            }

                            CListItem1 {
                                name:"反馈"
                                value: "https://github.com/Git-So/white-noise/issues/new"
                            }
                        }
                    }

                    // 开发者相关
                    Rectangle {
                        width: 1
                        height: 15
                        color: "transparent"
                    }
                    Rectangle {
                        width: parent.width
                        height: 320
                        color: "#282c34"
                        radius: 15


                        Column {
                            width: parent.width
                            height: parent.height

                            CListItem2 {
                                name: "Github"
                                info: "https://github.com/Git-So"
                                value: "https://github.com/Git-So"
                            }

                            CListItem2 {
                                name: "邮箱"
                                info: "me@sooo.site"
                                value: "mailto:me@sooo.site"
                            }

                            CListItem2 {
                                name: "博客"
                                info: "http://sooo.site"
                                value: "http://sooo.site"
                            }

                            CListItem2 {
                                name: "微博"
                                info: "http://weibo.com/u/3743700645"
                                value: "http://weibo.com/u/3743700645"
                            }
                        }
                    }

                }
            }
        }
    }
}
