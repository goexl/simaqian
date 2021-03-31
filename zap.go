package glog

import (
	`github.com/storezhang/gox`
	`go.uber.org/zap`
)

// ZapLogger Zap日志
type ZapLogger struct {
	logger *zap.Logger
}

// NewZapLogger 创建Zap日志记录器
func NewZapLogger(logger *zap.Logger) *ZapLogger {
	return &ZapLogger{
		logger: logger,
	}
}

func (zl *ZapLogger) Debug(msg string, fields ...gox.Field) {
	zl.logger.Debug(msg, zl.parse(fields...)...)
}

func (zl *ZapLogger) Info(msg string, fields ...gox.Field) {
	zl.logger.Info(msg, zl.parse(fields...)...)
}

func (zl *ZapLogger) Warn(msg string, fields ...gox.Field) {
	zl.logger.Warn(msg, zl.parse(fields...)...)
}

func (zl *ZapLogger) Error(msg string, fields ...gox.Field) {
	zl.logger.Error(msg, zl.parse(fields...)...)
}

func (zl *ZapLogger) Panic(msg string, fields ...gox.Field) {
	zl.logger.Panic(msg, zl.parse(fields...)...)
}

func (zl *ZapLogger) Fatal(msg string, fields ...gox.Field) {
	zl.logger.Fatal(msg, zl.parse(fields...)...)
}

func (zl *ZapLogger) parse(fields ...gox.Field) (zapFields []zap.Field) {
	zapFields = make([]zap.Field, 0, len(fields))
	for _, field := range fields {
		switch field.Value().(type) {
		case gox.IntField:
			zapFields = append(zapFields, zap.Int(field.Key(), field.Value().(int)))
		case gox.StringField:
			zapFields = append(zapFields, zap.String(field.Key(), field.Value().(string)))
		default:
			zapFields = append(zapFields, zap.Any(field.Key(), field.Value()))
		}
	}

	return
}
