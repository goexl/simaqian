package glog

import (
	`github.com/sirupsen/logrus`
	`github.com/storezhang/gox`
)

// LogrusLogger Logrus日志
type LogrusLogger struct {
	logger *logrus.Logger
}

// NewLogrusLogger 创建Logrus日志记录器
func NewLogrusLogger(logger *logrus.Logger) *LogrusLogger {
	return &LogrusLogger{
		logger: logger,
	}
}

func (ll *LogrusLogger) Debug(msg string, fields ...gox.Field) {
	ll.logger.WithFields(ll.parse(fields...)).Debug(msg)
}

func (ll *LogrusLogger) Info(msg string, fields ...gox.Field) {
	ll.logger.WithFields(ll.parse(fields...)).Info(msg)
}

func (ll *LogrusLogger) Warn(msg string, fields ...gox.Field) {
	ll.logger.WithFields(ll.parse(fields...)).Warn(msg)
}

func (ll *LogrusLogger) Error(msg string, fields ...gox.Field) {
	ll.logger.WithFields(ll.parse(fields...)).Error(msg)
}

func (ll *LogrusLogger) Panic(msg string, fields ...gox.Field) {
	ll.logger.WithFields(ll.parse(fields...)).Panic(msg)
}

func (ll *LogrusLogger) Fatal(msg string, fields ...gox.Field) {
	ll.logger.WithFields(ll.parse(fields...)).Fatal(msg)
}

func (ll *LogrusLogger) parse(fields ...gox.Field) (logrusFields logrus.Fields) {
	logrusFields = make(logrus.Fields, len(fields))
	for _, field := range fields {
		logrusFields[field.Key()] = field.Value()
	}

	return
}
