package builder

import (
	"time"

	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/executor"
	"github.com/goexl/simaqian/internal/logger"
	"github.com/goexl/simaqian/internal/param"
)

type Loki struct {
	core   *param.Core
	params *param.Loki
}

func NewLoki(core *param.Core) *Loki {
	return &Loki{
		core:   core,
		params: param.NewLoki(),
	}
}

func (l *Loki) Url(url string) (loki *Loki) {
	l.params.Url = url
	loki = l

	return
}

func (l *Loki) Username(username string) (loki *Loki) {
	l.params.Username = username
	loki = l

	return
}

func (l *Loki) Password(password string) (loki *Loki) {
	l.params.Password = password
	loki = l

	return
}

func (l *Loki) Batch(size int, wait time.Duration) (loki *Loki) {
	l.params.Batch.Size = size
	l.params.Batch.Wait = wait
	loki = l

	return
}

func (l *Loki) Labels(labels map[string]string) (loki *Loki) {
	l.params.Labels = labels
	loki = l

	return
}

func (l *Loki) Label(key string, value string) (loki *Loki) {
	l.params.Labels[key] = value
	loki = l

	return
}

func (l *Loki) Build() (_logger core.Logger, err error) {
	if loki, ne := executor.NewLoki(l.params); nil != ne {
		err = ne
	} else {
		_logger = logger.NewDefault(l.core, loki)
	}

	return
}
