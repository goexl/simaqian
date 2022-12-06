package simaqian

type config struct {
	level      level
	skip       int
	stacktrace *stacktrace
}

func newConfig() *config {
	return &config{
		level: LevelInfo,
		skip:  2,
	}
}
