package simaqian

import (
	"github.com/goexl/gox"
	"github.com/sirupsen/logrus"
)

var _ executor = (*_logrus)(nil)

type _logrus struct {
	logger *logrus.Logger
}

func newLogrus(_ *options) (logger *_logrus, err error) {
	logger = &_logrus{
		logger: logrus.StandardLogger(),
	}

	return
}

func (l *_logrus) debug(msg string, fields ...gox.Field[any]) {
	l.logger.WithFields(l.parse(fields...)).Debug(msg)
}

func (l *_logrus) info(msg string, fields ...gox.Field[any]) {
	l.logger.WithFields(l.parse(fields...)).Info(msg)
}

func (l *_logrus) warn(msg string, fields ...gox.Field[any]) {
	l.logger.WithFields(l.parse(fields...)).Warn(msg)
}

func (l *_logrus) error(msg string, fields ...gox.Field[any]) {
	l.logger.WithFields(l.parse(fields...)).Error(msg)
}

func (l *_logrus) panic(msg string, fields ...gox.Field[any]) {
	l.logger.WithFields(l.parse(fields...)).Panic(msg)
}

func (l *_logrus) fatal(msg string, fields ...gox.Field[any]) {
	l.logger.WithFields(l.parse(fields...)).Fatal(msg)
}

func (l *_logrus) parse(fields ...gox.Field[any]) (logrusFields logrus.Fields) {
	logrusFields = make(logrus.Fields, len(fields))
	for _, field := range fields {
		logrusFields[field.Key()] = field.Value()
	}

	return
}
