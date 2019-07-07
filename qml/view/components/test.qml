import QtQuick 2.7
import QtGraphicalEffects 1.13

Rectangle {
    width: 500
    height: 500
    color: "transparent"

    Rectangle {
        id: test1
        width: 300
        height: 300

        radius: 25
        clip:true
         smooth: true

        anchors.centerIn: parent

    }

    RadialGradient {
        id: musicInfo
        width: 300
        height: 300
        visible: false
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

    OpacityMask {
        anchors.fill: test1
        source: musicInfo
        maskSource: test1
    }

}
