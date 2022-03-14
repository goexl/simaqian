package simaqian

var (
	_ = Types
	_ = Zap
	_ = Logrus
	_ = System

	_ option = (*optionType)(nil)
)

type optionType struct {
	logType logType
}

// Types 配置日志类型
func Types(_type string) *optionType {
	return &optionType{
		logType: *ParseType(_type),
	}
}

// Zap 使用Zap日志
func Zap() *optionType {
	return &optionType{
		logType: TypeZap,
	}
}

// Logrus 使用Logrus日志
func Logrus() *optionType {
	return &optionType{
		logType: TypeLogrus,
	}
}

// System 使用System日志
func System() *optionType {
	return &optionType{
		logType: TypeSystem,
	}
}

func (t *optionType) apply(options *options) {
	options.logType = t.logType
}
