package glog

import (
	`go.uber.org/zap`
)

// NewLogger 创建新的日志
func NewLogger(config Config) (logger Logger, err error) {
	switch config.Type {
	case TypeLogrus:
		logger = NewLogrusLogger(config.logrusLogger())
	case TypeZap:
		var zapLogger *zap.Logger

		if zapLogger, err = config.zapLogger(); nil != err {
			return
		}

		logger = NewZapLogger(zapLogger)
	}

	return
}
