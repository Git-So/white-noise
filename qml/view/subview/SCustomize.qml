import QtQuick 2.7
import QtQuick.Controls 2.5
import QtGraphicalEffects 1.13
import "../components"

Rectangle {
    id: customizeRoot
    property int block: 45

    width: 9 * block
    height: 16 * block
    color: "#212121"

    // 栈对象
    property var aboutID: false
    property var muiscInfoID: false
    property var stackObj: false

    // 是否选中选项
    property bool isActiveMusic: true

    Column {
        width: parent.width
        height: parent.height
        spacing:5

        property var toolbarHeight: 100
        property var bottomToolbarHeight: 100
        property var gridHeight: parent.height - toolbarHeight - bottomToolbarHeight

        // 工具栏
        Rectangle {
            width: parent.width
            height: parent.toolbarHeight
            color: "transparent"

            CToolbar {
                aboutID: customizeRoot.aboutID
                stackObj: customizeRoot.stackObj
                activeID: 1

                anchors.fill: parent
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
                        model: [
                            {
                                name: "自定义",
                                icon: "../../icons/庭院鸟鸣.png",
                                startColor: "#ffE27DBF",
                                endColor: "#ffFCF075"
                            }
                        ]

                        CIconItem {
                            width: 120
                            height: width + 30

                            name: modelData.name
                            icon: modelData.icon
                            startColor: modelData.startColor
                            endColor: modelData.endColor
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

                stackObj: customizeRoot.stackObj
                muiscInfoID: customizeRoot.muiscInfoID
            }
        }
    }

}
