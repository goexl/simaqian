package glog

import (
	`github.com/sirupsen/logrus`
	`github.com/storezhang/gox`
)

// Logrus日志
type logrusLogger struct {
	logger *logrus.Logger
}

// 创建Logrus日志记录器
func newLogrusLogger(_ *options) (logger *logrusLogger, err error) {
	logger = &logrusLogger{
		logger: logrus.StandardLogger(),
	}

	return
}

func (ll *logrusLogger) Debug(msg string, fields ...gox.Field) {
	ll.logger.WithFields(ll.parse(fields...)).Debug(msg)
}

func (ll *logrusLogger) Info(msg string, fields ...gox.Field) {
	ll.logger.WithFields(ll.parse(fields...)).Info(msg)
}

func (ll *logrusLogger) Warn(msg string, fields ...gox.Field) {
	ll.logger.WithFields(ll.parse(fields...)).Warn(msg)
}

func (ll *logrusLogger) Error(msg string, fields ...gox.Field) {
	ll.logger.WithFields(ll.parse(fields...)).Error(msg)
}

func (ll *logrusLogger) Panic(msg string, fields ...gox.Field) {
	ll.logger.WithFields(ll.parse(fields...)).Panic(msg)
}

func (ll *logrusLogger) Fatal(msg string, fields ...gox.Field) {
	ll.logger.WithFields(ll.parse(fields...)).Fatal(msg)
}

func (ll *logrusLogger) parse(fields ...gox.Field) (logrusFields logrus.Fields) {
	logrusFields = make(logrus.Fields, len(fields))
	for _, field := range fields {
		logrusFields[field.Key()] = field.Value()
	}

	return
}
