package logger

const (
	// LevelDebug debug
	LevelDebug = iota
	// LevelInfo info
	LevelInfo
	// LevelWarn warn
	LevelWarn
	// LevelError error
	LevelError
)

// LevelPrefix 日志等级前缀
var LevelPrefix = map[int]string{
	LevelDebug: "Debug",
	LevelInfo:  "Info",
	LevelWarn:  "Warn",
	LevelError: "Error",
}

// IsDebug 调试状态,通过传参确定
var IsDebug = false

// Debug func
func Debug(v ...interface{}) {
	setLevelConf(LevelDebug)
	logObj.Println(v...)
}

// Info func
func Info(v ...interface{}) {
	setLevelConf(LevelInfo)
	logObj.Println(v...)
}

// Warn func
func Warn(v ...interface{}) {
	setLevelConf(LevelWarn)
	logObj.Println(v...)
}

// Error func
func Error(v ...interface{}) {
	setLevelConf(LevelError)
	logObj.Println(v...)
}
