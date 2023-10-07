package simaqian

import (
	"github.com/goexl/simaqian/internal/core"
)

var _ = ParseLevel

// ParseLevel 解析日志级别
func ParseLevel(level string) core.Level {
	return core.ParseLevel(level)
}
