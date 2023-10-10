package param

import (
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/internal/constant"
)

type Core struct {
	Level      core.Level
	Skip       int
	Stacktrace int
}

func NewCore() *Core {
	return &Core{
		Level:      core.LevelInfo,
		Stacktrace: constant.DefaultStacktrace,
	}
}
