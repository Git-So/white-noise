package logger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

/**
 * 日志直接输出
 */

var (
	// 日志前缀
	logPrefix = "info"

	// 记录属性
	logFlag = log.LstdFlags

	// 深度
	callerSkip = 2

	logObj *log.Logger
)

func init() {
	f := openLogFile(LevelInfo)

	logObj = log.New(f, logPrefix, logFlag)
}

// setLevel
func setLevelConf(level int) {
	// 输出前缀
	if _, file, line, ok := runtime.Caller(callerSkip); ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", LevelPrefix[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", LevelPrefix[level])
	}
	logObj.SetPrefix(logPrefix)

	// 输出文件
	logObj.SetOutput(openLogFile(level))
}
