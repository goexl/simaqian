package simaqian

type optionSkip struct {
	skip int
}

// Skip 配置方法调用过滤层级
func Skip(skip int) *optionSkip {
	return &optionSkip{
		skip: skip,
	}
}

func (t *optionSkip) apply(options *options) {
	options.skip = t.skip
}
