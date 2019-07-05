import QtQuick 2.7
import QtGraphicalEffects 1.13

Rectangle {
    color: "transparent"


    property bool iconMusicPlay: false
    property string activeMusic: "鸟鸣"

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
                radius: 25
                clip: true

                Rectangle {
                    radius: 25
                    clip: true

                    anchors.fill: parent

                    RadialGradient {
                        anchors.fill: parent
                        horizontalOffset: parent.width / 2
                        verticalOffset: parent.height / 2
                        verticalRadius: parent.height * 2

                        gradient: Gradient {
                            GradientStop {
                                position: 0.0
                                color: "#ffE27DBF"
                            }
                            GradientStop {
                                position: 1
                                color: "#ffFCF075"
                            }
                        }
                    }
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
                        text: activeMusic
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
            }

            // 播放控制
            Rectangle {
                width: parent.width / 10 * 2.1
                height: parent.height
                color: "#282c34"
                radius: 25
                clip: true

                anchors.right: parent.right

                property int iconSize: 25

                Rectangle {
                    width: parent.width
                    height: width
                    radius: 25

                    RadialGradient {
                        anchors.fill: parent
                        horizontalOffset: parent.width / 2
                        verticalOffset: parent.height / 2
                        verticalRadius: parent.height * 2

                        gradient: Gradient {
                            GradientStop {
                                position: 0.0
                                color: "#ffE27DBF"
                            }
                            GradientStop {
                                position: 1
                                color: "#ffFCF075"
                            }
                        }
                    }
                }

                //  图标
                Image {
                    width: parent.iconSize
                    height: parent.iconSize
                    visible: iconMusicPlay

                    // source: "qrc:/qml/drawable/icon_music_play.png"
                    source: "../../drawable/icon_music_play.png"

                    anchors.centerIn: parent
                }

                //  图标
                Image {
                    width: parent.iconSize
                    height: parent.iconSize
                    visible: !iconMusicPlay

                    // source: "qrc:/qml/drawable/icon_music_pause.png"
                    source: "../../drawable/icon_music_pause.png"

                    anchors.centerIn: parent
                }
            }
        }
    }
}
