package executor

import (
	"github.com/goexl/gox"
	"github.com/goexl/simaqian/internal/core"
	"github.com/sirupsen/logrus"
)

var _ core.Executor = (*Logrus)(nil)

type Logrus struct {
	logger *logrus.Logger
}

func NewLogrus() *Logrus {
	return &Logrus{
		logger: logrus.StandardLogger(),
	}
}

func (l *Logrus) Debug(msg string, fields ...gox.Field[any]) {
	l.logger.WithFields(l.parse(fields...)).Debug(msg)
}

func (l *Logrus) Info(msg string, fields ...gox.Field[any]) {
	l.logger.WithFields(l.parse(fields...)).Info(msg)
}

func (l *Logrus) Warn(msg string, fields ...gox.Field[any]) {
	l.logger.WithFields(l.parse(fields...)).Warn(msg)
}

func (l *Logrus) Error(msg string, fields ...gox.Field[any]) {
	l.logger.WithFields(l.parse(fields...)).Error(msg)
}

func (l *Logrus) Panic(msg string, fields ...gox.Field[any]) {
	l.logger.WithFields(l.parse(fields...)).Panic(msg)
}

func (l *Logrus) Fatal(msg string, fields ...gox.Field[any]) {
	l.logger.WithFields(l.parse(fields...)).Fatal(msg)
}

func (l *Logrus) Sync() (err error) {
	return
}

func (l *Logrus) parse(fields ...gox.Field[any]) (parsed logrus.Fields) {
	parsed = make(logrus.Fields, len(fields))
	for _, field := range fields {
		parsed[field.Key()] = field.Value()
	}

	return
}
