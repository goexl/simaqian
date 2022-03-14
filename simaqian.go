package simaqian

var _ = Must

// New 创建新的日志
func New(opts ...option) (logger Logger, err error) {
	return newTemplate(opts...)
}

// Must 必须返回日志
func Must(opts ...option) (logger Logger) {
	var err error
	if logger, err = New(opts...); nil != err {
		panic(err)
	}

	return
}
