package executor

import (
	"context"
	"time"

	"github.com/goexl/simaqian/internal/param"
	"github.com/paul-milne/zap-loki"
	"go.uber.org/zap"
)

func NewLoki(params *param.Loki) (logger *Zap, err error) {
	logger = new(Zap)
	config := zap.NewProductionConfig()
	loki := zaploki.New(context.Background(), zaploki.Config{
		Url:          params.Url,
		BatchMaxSize: 1000,
		BatchMaxWait: 10 * time.Second,
		Labels:       params.Labels,
	})
	logger.zap, err = loki.WithCreateLogger(config)

	return
}
