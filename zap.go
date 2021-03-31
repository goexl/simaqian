package glog

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
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
	for _, f := range fields {
		switch f.Value().(type) {
		case *field.Int8Field:
			zapFields = append(zapFields, zap.Int8(f.Key(), f.Value().(int8)))
		case *field.IntField:
			zapFields = append(zapFields, zap.Int(f.Key(), f.Value().(int)))
		case *field.UintField:
			zapFields = append(zapFields, zap.Uint(f.Key(), f.Value().(uint)))
		case *field.Int64Field:
			zapFields = append(zapFields, zap.Int64(f.Key(), f.Value().(int64)))
		case *field.Float32Field:
			zapFields = append(zapFields, zap.Float32(f.Key(), f.Value().(float32)))
		case *field.Float64Field:
			zapFields = append(zapFields, zap.Float64(f.Key(), f.Value().(float64)))
		case *field.StringField:
			zapFields = append(zapFields, zap.String(f.Key(), f.Value().(string)))
		case *field.ErrorField:
			zapFields = append(zapFields, zap.Error(f.Value().(error)))
		default:
			zapFields = append(zapFields, zap.Any(f.Key(), f.Value()))
		}
	}

	return
}
