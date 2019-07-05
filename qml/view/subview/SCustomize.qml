import QtQuick 2.7
import QtQuick.Controls 2.5
import "../components"

Rectangle {
    property int block: 45

    width: 9 * block
    height: 16 * block
    color: "#212121"

    // 栈对象
    property var aboutID: false
    property var stackObj: false

    // 是否选中选项
    property bool isActiveMusic: true

    Column {
        width: parent.width
        height: parent.height
        spacing:5

        property var toolbarHeight: 80
        property var bottomToolbarHeight: 100
        property var gridHeight: parent.height - toolbarHeight - bottomToolbarHeight

        // 工具栏
        Rectangle {
            width: parent.width
            height: parent.toolbarHeight
            color: "transparent"

            CToolbar {
                width: parent.width
                height: parent.height

                aboutID: parent.aboutID
                stackObj: parent.stackObj
                activeID: 1
            }
        }

        // 自选列表
        Rectangle {
            width: parent.width
            height: parent.gridHeight
            color: "transparent"
            // color: "red"

            ScrollView {
                width: parent.width - 30
                height: parent.height
                clip: true

                anchors.horizontalCenter: parent.horizontalCenter

                Flow {
                    width: parent.width
                    height: parent.height
                    // anchors.margins: 20
                    spacing: 5
                    // flow: Grid.TopToBottom

                    // anchors.centerIn: parent

                    Repeater {
                        model: 16

                        property variant colorArray: ["#00bde3", "#67c111", "#ea7025"]
                        property var rectangleWidth: 50

                        Rectangle {
                            width: 120
                            height: width + 30
                            color: "transparent"
                            // color: "red"

                            Column {
                                width: parent.width
                                height: parent.height
                                // color: "transparent"
                                // spacing: 5

                                Rectangle {
                                    width: parent.width * 0.8
                                    height: width

                                    property int colorIndex: Math.floor(Math.random()*3)
                                    color: "#282c34"
                                    radius: 100

                                    Rectangle {
                                        visible: Math.random() > 0.5 ? true : false
                                        radius: 100
                                        rotation: 135
                                        gradient: Gradient{
                                                    GradientStop{
                                                        position: 0.0;
                                                        color: "#ffE27DBF";

                                                    }
                                                    GradientStop{
                                                        position: 1.0;
                                                        color: "#ffFCF075";
                                                    }
                                                }

                                        anchors.fill: parent
                                    }



                                    anchors.horizontalCenter: parent.horizontalCenter

                                    //  图标
                                    Image {
                                        width: parent.width * 0.45
                                        height: width

                                        // source: "qrc:/qml/icons/庭院鸟鸣.png"
                                        source: "../../icons/庭院鸟鸣.png"

                                        anchors.centerIn: parent
                                    }
                                }

                                Rectangle {
                                    width: parent.width
                                    height: width - parent.width * 0.8 + 15
                                    color: "transparent"

                                    Text {
                                        text: "自定义"
                                        color: "#fff"

                                        anchors.centerIn: parent
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }


        // 底部工具栏
        Rectangle {
            width: parent.width
            height: parent.bottomToolbarHeight
            color: "transparent"
            // color: "red"

            CBottomToolbar1 {
                visible: !isActiveMusic

                anchors.fill: parent
            }

            CBottomToolbar2 {
                visible: isActiveMusic

                anchors.fill: parent
            }
        }
    }

}
