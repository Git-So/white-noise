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

	// LevelFileName 日志目录名
	LevelFileName = map[int]string{
		LevelDebug: "debug",
		LevelInfo:  "info",
		LevelWarn:  "warn",
		LevelError: "error",
	}

	// LevelFile 日志文件实例
	LevelFile = make(map[int]*os.File, 4)
)

// GetLogFilePath 获取某等级日志文件路径
func GetLogFilePath(level int) (path string) {
	return fmt.Sprintf("%s%s.%s", LogSavePath, LevelFileName[level], LogFileExt)
}

// initLogSavePath
func openLogFile(level int) *os.File {
	// 已打开文件
	if f, ok := LevelFile[level]; ok {
		return f
	}

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

	// 保存文件实例
	LevelFile[level] = f

	return f
}
