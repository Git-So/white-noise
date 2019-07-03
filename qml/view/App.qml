import QtQuick 2.7
import QtQuick.Controls 2.5
import "./subview"

ApplicationWindow {
    title: qsTr("白噪音")

    property int block: 45

    width: 9 * block
    height: 16 * block

    visible: true

    StackView {
        id: stack
        initialItem: homeView

        anchors.fill: parent
    }

    Component {
        id: homeView

        SHome {}
    }
}
