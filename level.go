package simaqian

const (
	LevelDebug level = "debug"
	LevelInfo  level = "info"
	LevelWarn  level = "warn"
	LevelError level = "error"
	LevelPanic level = "panic"
	LevelFatal level = "fatal"
)

var _ = ParseLevel

type level string

// ParseLevel 解析日志级别
func ParseLevel(_level string) (lvl level) {
	switch level(_level) {
	case LevelDebug:
		lvl = LevelDebug
	case LevelInfo:
		lvl = LevelInfo
	case LevelWarn:
		lvl = LevelWarn
	case LevelError:
		lvl = LevelError
	case LevelPanic:
		lvl = LevelPanic
	case LevelFatal:
		lvl = LevelFatal
	}

	return
}

func (l level) rank() (rank int) {
	switch l {
	case LevelDebug:
		rank = 10
	case LevelInfo:
		rank = 20
	case LevelWarn:
		rank = 30
	case LevelError:
		rank = 40
	case LevelPanic:
		rank = 50
	case LevelFatal:
		rank = 60
	}

	return
}
