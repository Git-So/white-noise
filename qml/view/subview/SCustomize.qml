import QtQuick 2.7
import QtQuick.Controls 2.5
import QtGraphicalEffects 1.13
import CustomQmlTypes 1.0
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
                    spacing: 5

                    Repeater {
                        id: iconListView
                        model: IconListModel{}

                        CIconItem {
                            width: 120
                            height: width + 30

                            name: display.name
                            icon: display.icon
                            startColor: display.colorParam.split(";")[0]
                            endColor: display.colorParam.split(";")[1]
                            isActive: display.isActive

                            // 切换选中状态
                            MouseArea {
                                anchors.fill: parent

                                onClicked: {
                                    iconListView.model.change(display.index)
                                    if (display.isActive) {
                                        iconStack.push(display)
                                    } else {
                                        iconStack.pop(display.index)
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
                visible: !iconStack.isActive

                anchors.fill: parent
            }

            CBottomToolbar2 {
                visible: iconStack.isActive

                anchors.fill: parent

                stackObj: customizeRoot.stackObj
                muiscInfoID: customizeRoot.muiscInfoID
                startColor: iconStack.colorParam.split(";")[0]
                endColor: iconStack.colorParam.split(";")[1]
                title: iconStack.title
            }
        }
    }

}
