package executor

import (
	"context"

	"github.com/goexl/simaqian/internal/param"
	"github.com/paul-milne/zap-loki"
	"go.uber.org/zap"
)

func NewLoki(params *param.Loki) (logger *Zap, err error) {
	logger = new(Zap)
	config := zaploki.Config{
		Url:          params.Url,
		BatchMaxSize: params.Batch.Size,
		BatchMaxWait: params.Batch.Wait,
	}
	if 0 != len(params.Labels) {
		config.Labels = params.Labels
	}
	if "" != params.Username {
		config.Username = params.Username
	}
	if "" != params.Password {
		config.Password = params.Password
	}
	loki := zaploki.New(context.Background(), config)
	logger.zap, err = loki.WithCreateLogger(zap.NewProductionConfig())

	return
}
