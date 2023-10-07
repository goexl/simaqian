package param

import (
	"github.com/goexl/simaqian/internal/core"
)

type Zap struct {
	Outputs []core.Writer
	Errors  []core.Writer
}

func NewZap() *Zap {
	return &Zap{
		Outputs: make([]core.Writer, 0),
		Errors:  make([]core.Writer, 0),
	}
}
