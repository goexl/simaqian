package simaqian

// New 创建新的日志
func New(opts ...option) (logger Logger, err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	switch _options.logType {
	case TypeLogrus:
		logger, err = newLogrusLogger(_options)
	case TypeZap:
		logger, err = newZapLogger(_options)
	case TypeSystem:
		logger, err = newSystemLogger(_options)
	}

	return
}

// Must 必须返回日志
func Must(opts ...option) (logger Logger) {
	var err error
	if logger, err = New(opts...); nil != err {
		panic(err)
	}

	return
}
