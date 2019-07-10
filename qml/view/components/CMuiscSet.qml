import QtQuick 2.7
import QtQuick.Controls 1.4

Rectangle {
    id: muiscSetRoot
    color: "transparent"
    // color: "red"

    property int index: 0
    property string name: "自定义"
    property var volume: 1.0

    Column {
        width: parent.width - 50
        height: parent.height

        anchors.horizontalCenter: parent.horizontalCenter

        Rectangle {
            width: parent.width
            height: parent.height / 2
            color: "transparent"

            Text {
                text: muiscSetRoot.name
                color: "#fff"

               anchors.verticalCenter: parent.verticalCenter
            }
        }

        Rectangle {
            width: parent.width
            height: parent.height / 2
            color: "transparent"

            Slider {
                width: parent.width * 0.9
                value: muiscSetRoot.volume
                maximumValue: 2.5
                minimumValue: 0.05
                stepSize: 0.01

                anchors.horizontalCenter: parent.horizontalCenter

                onValueChanged: {
                    console.log(muiscSetRoot.volume)
                    console.log(muiscSetRoot.name,":", value)
                    muisInfoList.model.write(muiscSetRoot.index,value)
                }
            }
        }
    }

}
