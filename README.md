# 白噪音

移植小米白噪音到 PC

功能全部正常，理论支持全平台

## 一、编译

    make init
    make run # 仅编译
    make run # 编译并运行

> 确保安装 QT5.13.0 以及 golang 最新版及相关依赖
> docker 编译待写 🤪

## 二、协议

config/audio
config/icons
config/scene
qml/drawable
qml/icons
qml/scene

以上目录提取或修改自『小米白噪音』APP ，本人无版权
其他 LGPL

## 三、注

由于资源文件提取或修改自『小米白噪音』APP，故如需使用请自行编译
