package param

import (
	"github.com/goexl/simaqian/internal/internal"
)

type Loki struct {
	Url      string
	Labels   map[string]string
	Batch    *internal.Batch
	Username string
	Password string
}

func NewLoki() *Loki {
	return &Loki{
		Batch: internal.NewBatch(),
	}
}
