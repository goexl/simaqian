package executor

import (
	"context"

	"github.com/goexl/simaqian/internal/internal/loki"
	"github.com/goexl/simaqian/internal/param"
	"go.uber.org/zap"
)

func NewLoki(params *param.Loki) (logger *Zap, err error) {
	logger = new(Zap)
	config := loki.Config{
		Url:   params.Url,
		Batch: params.Batch,
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
	pusher := loki.New(context.Background(), config)
	logger.zap, err = pusher.Build(zap.NewProductionConfig(), zap.WithCaller(false))

	return
}
