package simaqian

var _ = Must

// New 创建新的日志
func New() *builder {
	return newBuilder()
}

// Must 必须返回日志
func Must() Logger {
	return Default()
}

// Default 默认日志
func Default() (logger Logger) {
	if _logger, err := New().Zap().Build(); nil != err {
		panic(err)
	} else {
		logger = _logger
	}

	return
}
