package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	// LogSavePath 文件保存目录
	LogSavePath = "config/log/"

	// LogFileExt log 后缀
	LogFileExt = "log"

	// LevelFile 错误目录名
	LevelFile = map[int]string{
		LevelDebug: "debug",
		LevelInfo:  "info",
		LevelWarn:  "warn",
		LevelError: "error",
	}
)

// GetLogFilePath 获取某等级日志文件路径
func GetLogFilePath(level int) (path string) {
	return fmt.Sprintf("%s%s.%s", LogSavePath, LevelFile[level], LogFileExt)
}

// initLogSavePath
func openLogFile(level int) *os.File {
	// 文件状态
	file := GetLogFilePath(level)
	_, err := os.Stat(file)
	switch {
	case os.IsNotExist(err):
		// 新建文件
		if err = os.MkdirAll(LogSavePath, os.ModePerm); err != nil {
			log.Panicf("创建文件出错: %v", err)
		}
	case os.IsPermission(err):
		log.Panicf("权限出错: %v", err)
	}

	// 打开日志文件
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicf("权限出错: %v", err)
	}

	return f
}
