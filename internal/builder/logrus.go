package builder

import (
	"github.com/goexl/log"
	"github.com/goexl/simaqian/internal/executor"
	"github.com/goexl/simaqian/internal/logger"
	"github.com/goexl/simaqian/internal/param"
)

type Logrus struct {
	params *param.Logrus
}

func NewLogrus(params *param.Logrus) *Logrus {
	return &Logrus{
		params: params,
	}
}

func (l *Logrus) Build() log.Logger {
	return logger.NewDefault(l.params, executor.NewLogrus())
}
