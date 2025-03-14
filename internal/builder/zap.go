package builder

import (
	"github.com/goexl/log"
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/executor"
	"github.com/goexl/simaqian/internal/logger"
	"github.com/goexl/simaqian/internal/param"
)

type Zap struct {
	params *param.Zap
}

func NewZap() *Zap {
	return &Zap{
		params: param.NewZap(),
	}
}

func (z *Zap) Output(writer core.Writer) *Zap {
	z.params.Outputs = append(z.params.Outputs, writer)

	return z
}

func (z *Zap) Error(writer core.Writer) *Zap {
	z.params.Errors = append(z.params.Errors, writer)

	return z
}

func (z *Zap) Build() (_logger log.Logger, err error) {
	if zap, ne := executor.NewZap(z.params); nil != ne {
		err = ne
	} else {
		_logger = logger.NewDefault(z.core, zap)
	}

	return
}
