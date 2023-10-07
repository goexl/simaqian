package builder

import (
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/executor"
	"github.com/goexl/simaqian/internal/logger"
	"github.com/goexl/simaqian/internal/param"
)

type Zap struct {
	core   *param.Core
	config *param.Zap
}

func NewZap(core *param.Core) *Zap {
	return &Zap{
		core:   core,
		config: param.NewZap(),
	}
}

func (z *Zap) Output(writer core.Writer) *Zap {
	z.config.Outputs = append(z.config.Outputs, writer)

	return z
}

func (z *Zap) Error(writer core.Writer) *Zap {
	z.config.Errors = append(z.config.Errors, writer)

	return z
}

func (z *Zap) Build() (_logger core.Logger, err error) {
	if zap, ne := executor.NewZap(z.config); nil != ne {
		err = ne
	} else {
		_logger = logger.NewDefault(z.core, zap)
	}

	return
}
