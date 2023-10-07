package simaqian

import (
	"github.com/goexl/simaqian/internal/builder"
	"github.com/goexl/simaqian/internal/core"
)

var _ = Must

// New 创建新的日志
func New() *builder.Core {
	return builder.NewBuilder()
}

// Must 必须返回日志
func Must() core.Logger {
	return Default()
}

// Default 默认日志
func Default() (logger core.Logger) {
	if zap, err := New().Zap().Build(); nil != err {
		panic(err)
	} else {
		logger = zap
	}

	return
}
