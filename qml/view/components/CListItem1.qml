import QtQuick 2.7

Rectangle {
    height: 60

    color: "transparent"

    anchors.left: parent.left
    anchors.right: parent.right
    anchors.leftMargin: 20
    anchors.rightMargin: 10

    property string name: "关于名称"
    property string value: ""

    Text {
        text: name
        color: "#fff"

        anchors.verticalCenter: parent.verticalCenter
    }

    Image {
        width: 45
        height: 45
        source: "qrc:/qml/drawable/action_bar_back_normal_dark.png"
        // source: "../../drawable/action_bar_back_normal_dark.png"

        rotation: 180

        anchors.verticalCenter: parent.verticalCenter
        anchors.right: parent.right
    }

    MouseArea {
        anchors.fill: parent

        onClicked: {
            if (value.length > 0) {
                Qt.openUrlExternally(value)
            }
        }
    }
}
