package executor

import (
	"context"

	"github.com/goexl/simaqian/internal/executor/internal/config"
	"github.com/goexl/simaqian/internal/internal/loki"
	"github.com/goexl/simaqian/internal/param"
	"go.uber.org/zap"
)

func NewLoki(params *param.Loki) (logger *Zap, err error) {
	logger = new(Zap)
	lokiConfig := new(loki.Config)
	lokiConfig.Url = params.Url
	lokiConfig.Batch = params.Batch
	lokiConfig.Http = params.Http

	if 0 != len(params.Labels) {
		lokiConfig.Labels = params.Labels
	}
	if "" != params.Username {
		lokiConfig.Username = params.Username
	}
	if "" != params.Password {
		lokiConfig.Password = params.Password
	}
	pusher := loki.New(context.Background(), lokiConfig)
	logger.zap, err = pusher.Build(config.DefaultZap(), zap.WithCaller(false))

	return
}
