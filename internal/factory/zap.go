package factory

import (
	"github.com/goexl/simaqian/internal/executor"
	"github.com/goexl/simaqian/internal/param"
)

type Zap struct {
	params *param.Zap
}

func NewZap(params *param.Zap) *Zap {
	return &Zap{
		params: params,
	}
}

func (z *Zap) New() (*executor.Zap, error) {
	return executor.NewZap(z.params)
}
