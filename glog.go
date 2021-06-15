package glog

// New 创建新的日志
func New(opts ...option) (logger Logger, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	switch options.logType {
	case TypeLogrus:
		logger, err = newLogrusLogger(options)
	case TypeZap:
		logger, err = newZapLogger(options)
	case TypeSystem:
		logger, err = newSystemLogger(options)
	}

	return
}
