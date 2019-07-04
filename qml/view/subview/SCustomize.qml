import QtQuick 2.7
import "../components"

Rectangle {
    color: "#212121"

    // 栈对象
    property var aboutID: false
    property var stackObj: false

    // 工具栏
    CToolbar {
        width: parent.width
        height: 100

        aboutID: parent.aboutID
        stackObj: parent.stackObj
        activeID: 1
    }
}
