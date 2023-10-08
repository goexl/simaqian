package simaqian

import (
	"github.com/goexl/simaqian/internal/core"
)

const (
	LevelDebug = core.LevelDebug
	LevelInfo  = core.LevelInfo
	LevelWarn  = core.LevelWarn
	LevelError = core.LevelError
	LevelPanic = core.LevelPanic
	LevelFatal = core.LevelFatal
)

var (
	_ = ParseLevel
	_ = LevelDebug
	_ = LevelInfo
	_ = LevelWarn
	_ = LevelError
	_ = LevelPanic
	_ = LevelFatal
)

// ParseLevel 解析日志级别
func ParseLevel(level string) core.Level {
	return core.ParseLevel(level)
}
