import QtQuick 2.7

Rectangle {
    id: iconItemRoot
    color: "transparent"
    // color: "red"

    property string name: "自定义"
    property string icon: "庭院鸟鸣.png"
    property string startColor: "#ffE27DBF"
    property string endColor: "#ffFCF075"
    property bool isActive: false

    Column {
        width: parent.width
        height: parent.height
        // color: "transparent"
        // spacing: 5

        Rectangle {
            width: parent.width * 0.8
            height: width

            color: "#282c34"
            radius: 100

            anchors.horizontalCenter: parent.horizontalCenter



            // 渐变背景
            Rectangle {
                visible: iconItemRoot.isActive
                radius: 100
                rotation: 135
                gradient: Gradient{
                            GradientStop{
                                position: 0.0;
                                color: iconItemRoot.startColor;

                            }
                            GradientStop{
                                position: 1.0;
                                color: iconItemRoot.endColor;
                            }
                        }

                anchors.fill: parent
            }

            //  图标
            Image {
                width: parent.width * 0.45
                height: width

                // source: "qrc:/qml/icons/庭院鸟鸣.png"
                source: "qrc:/qml/icons/" + iconItemRoot.icon

                anchors.centerIn: parent
            }

            // 切换选中状态
            // MouseArea {
            //     anchors.fill: parent

            //     onClicked: {
            //         parent.isActive = !parent.isActive
            //     }
            // }
        }

        Rectangle {
            width: parent.width
            height: width - parent.width * 0.8 + 15
            color: "transparent"

            Text {
                text: iconItemRoot.name
                color: "#fff"

                anchors.centerIn: parent
            }
        }
    }
}
