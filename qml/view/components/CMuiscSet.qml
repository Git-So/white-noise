import QtQuick 2.7
import QtQuick.Controls 1.4

Rectangle {
    color: "transparent"
    // color: "red"

    Column {
        width: parent.width - 50
        height: parent.height

        anchors.horizontalCenter: parent.horizontalCenter

        Rectangle {
            width: parent.width
            height: parent.height / 2
            color: "transparent"

            Text {
                text: modelData.name
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
                value: modelData.size
                maximumValue: 5
                minimumValue: 0.05
                stepSize: 0.01

                anchors.horizontalCenter: parent.horizontalCenter

                onValueChanged: {

                    console.log(modelData.name,":",value)

                }
            }
        }
    }

}
