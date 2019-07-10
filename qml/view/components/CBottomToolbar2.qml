import QtQuick 2.7
import QtGraphicalEffects 1.13

Rectangle {
    id: bottomToolbar2Root
    color: "transparent"

    property bool   iconMusicPlay:  false
    property string title:          "鸟鸣"
    property string startColor:     "#ffE27DBF"
    property string endColor:       "#ffFCF075"

    // 栈对象
    property var stackObj: false
    property var muiscInfoID: false

    Rectangle {
        width: parent.width * 0.88
        height: parent.height * 0.68
        color: "transparent"

        anchors.horizontalCenter: parent.horizontalCenter

        Row {
            width: parent.width
            height: parent.height

            // 播放信息
            Rectangle {
                width: parent.width / 10 * 7.5
                height: parent.height
                color: "#282c34"
                radius: 20
                clip: true
                smooth: true

                Rectangle {
                    id: musicInfoBg
                    radius: 20
                    clip: true
                    smooth: true

                    anchors.fill: parent
                }

                RadialGradient {
                    id: musicInfoBgColor
                    visible: false

                    horizontalOffset: parent.width / 2
                    verticalOffset: parent.height / 2
                    verticalRadius: parent.height * 2

                    gradient: Gradient {
                        GradientStop {
                            position: 0.0
                            color: bottomToolbar2Root.startColor
                        }
                        GradientStop {
                            position: 1
                            color: bottomToolbar2Root.endColor
                        }
                    }

                    anchors.fill: parent
                }

                OpacityMask {
                    anchors.fill: parent
                    source: musicInfoBgColor
                    maskSource: musicInfoBg
                }

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
                        text: bottomToolbar2Root.title
                        color: "#fff"

                        anchors.leftMargin: 20
                        anchors.verticalCenter: parent.verticalCenter
                    }
                }

                //  图标
                Image {
                    width: 12
                    height: width * 1.5

                    // source: "qrc:/qml/drawable/icon_music_next.png"
                    source: "../../drawable/icon_music_next.png"

                    anchors.verticalCenter: parent.verticalCenter
                    anchors.right: parent.right
                    anchors.rightMargin: 22

                }

                MouseArea {
                    anchors.fill: parent

                    onClicked: {

                        console.log("muiscInfo",bottomToolbar2Root.stackObj,bottomToolbar2Root.muiscInfoID)

                        if (bottomToolbar2Root.stackObj && bottomToolbar2Root.muiscInfoID) {
                            bottomToolbar2Root.stackObj.push(bottomToolbar2Root.muiscInfoID)
                        }
                    }
                }
            }

            // 播放控制
            Rectangle {
                width: parent.width / 10 * 2.1
                height: parent.height
                color: "#282c34"
                radius: 25
                clip: true

                anchors.right: parent.right

                property int iconSize: 20

                Rectangle {
                    id: musicPlayBg
                    radius: 20
                    clip: true
                    smooth: true

                    anchors.fill: parent
                }

                RadialGradient {
                    id: musicPlayBgColor
                    visible: false
                    horizontalOffset: parent.width / 2
                    verticalOffset: parent.height / 2
                    verticalRadius: parent.height * 2

                    gradient: Gradient {
                        GradientStop {
                            position: 0.0
                            color: bottomToolbar2Root.startColor
                        }
                        GradientStop {
                            position: 1
                            color: bottomToolbar2Root.endColor
                        }
                    }

                    anchors.fill: parent
                }

                OpacityMask {
                    anchors.fill: parent
                    source: musicPlayBgColor
                    maskSource: musicPlayBg
                }

                //  图标
                Image {
                    width: parent.iconSize
                    height: parent.iconSize
                    visible: playObj.state

                    // source: "qrc:/qml/drawable/icon_music_play.png"
                    source: "../../drawable/icon_music_play.png"

                    anchors.centerIn: parent
                }

                //  图标
                Image {
                    width: parent.iconSize
                    height: parent.iconSize
                    visible: !playObj.state

                    // source: "qrc:/qml/drawable/icon_music_pause.png"
                    source: "../../drawable/icon_music_pause.png"

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
        }
    }
}
