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

    // 主页
    Component {
        id: homeView

        SHome {
            stackObj: stack
            aboutID: aboutView
            customizeID: customizeView
        }
    }

    // 自定义配置
    Component {
        id: customizeView

        SCustomize {
            stackObj: stack
            aboutID: aboutView
        }
    }

    // 关于
    Component {
        id: aboutView

        SAbout {
            stackObj: stack
        }
    }
}
