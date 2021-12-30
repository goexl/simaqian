package simaqian

type optionStacktrace struct {
	stacktrace *stacktrace
}

// Stacktrace 配置打印调用堆栈
func Stacktrace(skip int, stack int) *optionStacktrace {
	return &optionStacktrace{
		stacktrace: &stacktrace{
			skip:  skip,
			stack: stack,
		},
	}
}

func (s *optionStacktrace) apply(options *options) {
	options.stacktrace = s.stacktrace
}
