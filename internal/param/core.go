package param

import (
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/internal"
)

type Core struct {
	Level      core.Level
	Skip       int
	Stacktrace *internal.Stacktrace
}

func NewCore() *Core {
	return &Core{
		Level:      core.LevelInfo,
		Stacktrace: internal.NewStacktrace(),
	}
}
