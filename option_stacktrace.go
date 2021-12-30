package simaqian

type optionStacktrace struct {
	stack int
}

// Stacktrace 配置打印调用堆栈
func Stacktrace(stack int) *optionStacktrace {
	return &optionStacktrace{
		stack: stack,
	}
}

func (s *optionStacktrace) apply(options *options) {
	options.stack = s.stack
}
