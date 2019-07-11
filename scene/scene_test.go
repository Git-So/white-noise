package scene_test

import (
	"testing"

	"github.com/Git-So/white-noise/scene"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoad(t *testing.T) {
	Convey("测试加载场景数据:", t, func() {
		list, err := scene.Load("../config/scene/prebuilt_scene_config.json")

		Convey("无错执行", func() {
			So(err, ShouldBeNil)
		})

		Convey("标题数据效验", func() {
			So(list[0].Title, ShouldEqual, "夏雨")
			So(list[1].Title, ShouldEqual, "森林")
			So(list[2].Title, ShouldEqual, "炉火")
		})

		Convey("其他数据效验", func() {
			So(list[1].Desp, ShouldEqual, "把你的秘密藏进森林")
			So(list[2].ImagePath, ShouldEqual, "炉火.jpg")
			So(list[2].CropPicSmall.CenterLine, ShouldEqual, 0.5405)
			So(list[3].CropPicLarge.CenterLine, ShouldEqual, 0.2345)
		})
	})
}

func TestLoadAudio(t *testing.T) {
	Convey("测试加载场景音频数据:", t, func() {
		list := scene.LoadAudio("夏雨")

		Convey("标题数据效验", func() {
			So(list[0].Name, ShouldEqual, "河流")
			So(list[1].Name, ShouldEqual, "夏雨")
			So(list[2].Name, ShouldEqual, "雨打树叶")
		})

		Convey("其他数据效验", func() {
			So(list[1].IsLine, ShouldEqual, true)
			So(list[2].IsPoint, ShouldEqual, false)
			So(list[2].TotalEndTime, ShouldEqual, 0)
			So(list[3].Volume, ShouldEqual, 1.0)
		})
	})
}
