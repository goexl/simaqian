package simaqian

type options struct {
	level   level
	skip    int
	logType logType
	stack   int
}

func defaultOptions() *options {
	return &options{
		level:   LevelInfo,
		skip:    1,
		logType: TypeZap,
		stack:   0,
	}
}
