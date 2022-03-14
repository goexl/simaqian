package simaqian

var (
	_ = Level
	_ = Levels

	_ option = (*optionLevel)(nil)
)

type optionLevel struct {
	level level
}

// Level 配置日志级别
func Level(level level) *optionLevel {
	return &optionLevel{
		level: level,
	}
}

// Levels 配置日志级别
func Levels(level string) *optionLevel {
	return &optionLevel{
		level: *ParseLevel(level),
	}
}

func (l *optionLevel) apply(options *options) {
	options.level = l.level
}
