package param

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian/internal/config"
)

type Loki struct {
	Url      string
	Labels   map[string]string
	Batch    *config.Batch
	Username string
	Password string
	Http     *resty.Client
}

func NewLoki() *Loki {
	return &Loki{
		Batch: config.NewBatch(),
		Http:  resty.New(),
	}
}
