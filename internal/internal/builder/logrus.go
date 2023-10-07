package builder

import (
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/executor"
	"github.com/goexl/simaqian/internal/logger"
	"github.com/goexl/simaqian/internal/param"
)

type Logrus struct {
	config *param.Core
}

func NewLogrus(config *param.Core) *Logrus {
	return &Logrus{
		config: config,
	}
}

func (l *Logrus) Build() core.Logger {
	return logger.NewDefault(l.config, executor.NewLogrus())
}
