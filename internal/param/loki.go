package param

import (
	"github.com/goexl/http"
	"github.com/goexl/simaqian/internal/config"
)

type Loki struct {
	Url      string
	Labels   map[string]string
	Batch    *config.Batch
	Username string
	Password string
	Http     *http.Client
}

func NewLoki() *Loki {
	return &Loki{
		Batch: config.NewBatch(),
		Http:  http.New().Build(),
	}
}
