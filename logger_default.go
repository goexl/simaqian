package simaqian

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

var _ Logger = (*defaultLogger)(nil)

type defaultLogger struct {
	config   *config
	executor executor
}

func newDefaultLogger(config *config, executor executor) *defaultLogger {
	return &defaultLogger{
		config:   config,
		executor: executor,
	}
}

func (dl *defaultLogger) Level() level {
	return dl.config.level
}

func (dl *defaultLogger) Enabled(lvl level) bool {
	return dl.config.level.rank() >= lvl.rank()
}

func (dl *defaultLogger) Debug(msg string, fields ...gox.Field[any]) {
	if dl.config.level.rank() > LevelDebug.rank() {
		return
	}

	dl.addStacks(&fields)
	dl.executor.debug(msg, fields...)
}

func (dl *defaultLogger) Info(msg string, fields ...gox.Field[any]) {
	if dl.config.level.rank() > LevelInfo.rank() {
		return
	}

	dl.addStacks(&fields)
	dl.executor.info(msg, fields...)
}

func (dl *defaultLogger) Warn(msg string, fields ...gox.Field[any]) {
	if dl.config.level.rank() > LevelWarn.rank() {
		return
	}

	dl.addStacks(&fields)
	dl.executor.warn(msg, fields...)
}

func (dl *defaultLogger) Error(msg string, fields ...gox.Field[any]) {
	if dl.config.level.rank() > LevelError.rank() {
		return
	}

	dl.addStacks(&fields)
	dl.executor.error(msg, fields...)
}

func (dl *defaultLogger) Panic(msg string, fields ...gox.Field[any]) {
	if dl.config.level.rank() > LevelPanic.rank() {
		return
	}

	dl.addStacks(&fields)
	dl.executor.panic(msg, fields...)
}

func (dl *defaultLogger) Fatal(msg string, fields ...gox.Field[any]) {
	if dl.config.level.rank() > LevelFatal.rank() {
		return
	}

	dl.addStacks(&fields)
	dl.executor.fatal(msg, fields...)
}

func (dl *defaultLogger) Sync() error {
	return dl.executor.sync()
}

func (dl *defaultLogger) addStacks(fields *[]gox.Field[any]) {
	if nil == dl.config.stacktrace {
		return
	}

	pc := make([]uintptr, dl.config.stacktrace.stack+1)
	count := runtime.Callers(dl.config.skip+1+dl.config.stacktrace.skip, pc)
	frames := runtime.CallersFrames(pc[:count])

	stacks := make([]string, 0, 0)
	for {
		frame, more := frames.Next()
		stacks = append(stacks, fmt.Sprintf("%s[%s]:%d", filepath.Base(frame.File), frame.Function, frame.Line))
		if !more {
			break
		}
	}
	sort.SliceStable(stacks, func(i, j int) bool {
		return true
	})
	*fields = append(*fields, field.New[string]("stacks", strings.Join(stacks, " -> ")))
}
