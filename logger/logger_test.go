package logger_test

import (
	"testing"

	"github.com/Git-So/white-noise/logger"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLog(t *testing.T) {
	Convey("日志功能测试", t, func() {
		logger.Debug("test Debug")
		logger.Info("test info")
		logger.Warn("test Warn")
		logger.Error("test Error")
	})
}
