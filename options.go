package simaqian

type options struct {
	level   level
	skip    int
	logType logType
}

func defaultOptions() *options {
	return &options{
		level:   LevelInfo,
		skip:    1,
		logType: TypeZap,
	}
}
