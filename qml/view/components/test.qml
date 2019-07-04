import QtQuick 2.7


Rectangle {
    width: 100
    height: 100
    color: "transparent"

    Text {
        text: qsTr("推荐")
        color: activeID != 0 ? "#ddd" : "#fff"

        anchors.verticalCenter: parent.verticalCenter

        MouseArea {
        anchors.fill: parent

        onClicked: {

            console.log("121: ")

        }
    }
    }


}
