package core

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
	LevelPanic Level = "panic"
	LevelFatal Level = "fatal"
)

type Level string

func ParseLevel(level string) (lvl Level) {
	switch Level(level) {
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

func (l Level) Rank() (rank int) {
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
