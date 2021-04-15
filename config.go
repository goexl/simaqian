package glog

import (
	`github.com/sirupsen/logrus`
	`go.uber.org/zap`
)

// Config 配置
type Config struct {
	// Type 类型
	Type Type `json:"type" yaml:"type" validate:"required,oneof=zap zero system logrus"`
}

func (c *Config) zapLogger() (logger *zap.Logger, err error) {
	logger, err = zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	return
}

func (c *Config) logrusLogger() (logger *logrus.Logger) {
	logger = logrus.StandardLogger()

	return
}
