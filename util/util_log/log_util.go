package util_log

import (
	"fmt"
	"time"
)

// LogLevel 定义日志等级
type LogLevel string

const (
	LevelDebug LogLevel = "DEBUG"
	LevelInfo  LogLevel = "INFO"
	LevelWarn  LogLevel = "WARN"
	LevelError LogLevel = "ERROR"
)

// Log 日志函数，输出时间、等级和业务内容
func Log(level LogLevel, msgFormat string, a ...any) {
	msgContent := fmt.Sprintf(msgFormat, a...)
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", now, level, msgContent)
}
