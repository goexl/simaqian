package simaqian

import (
	`testing`

	`github.com/storezhang/gox`
	`go.uber.org/zap`
)

func TestZapDebug(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	defer func() {
		_ = zapLogger.Sync()
	}()

	var logger Logger = newZapLogger(zapLogger)
	logger.Warn("测试", gox.NewStringField("username", "storezhang"), gox.NewInt8Field("age", 18))
}
