package builder

import (
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/executor"
	"github.com/goexl/simaqian/internal/logger"
	"github.com/goexl/simaqian/internal/param"
)

type Builtin struct {
	config *param.Core
}

func NewBuiltin(config *param.Core) *Builtin {
	return &Builtin{
		config: config,
	}
}

func (b *Builtin) Build() core.Logger {
	return logger.NewDefault(b.config, executor.NewBuiltin())
}
