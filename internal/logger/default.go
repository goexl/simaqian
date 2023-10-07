package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/param"
)

var _ core.Logger = (*Default)(nil)

type Default struct {
	config   *param.Core
	executor core.Executor
}

func NewDefault(core *param.Core, executor core.Executor) *Default {
	return &Default{
		config:   core,
		executor: executor,
	}
}

func (d *Default) Level() core.Level {
	return d.config.Level
}

func (d *Default) Enable(lvl core.Level) {
	d.config.Level = lvl
}

func (d *Default) Enabled(lvl core.Level) bool {
	return d.config.Level.Rank() >= lvl.Rank()
}

func (d *Default) Debug(msg string, fields ...gox.Field[any]) {
	if d.config.Level.Rank() <= core.LevelDebug.Rank() {
		d.executor.Debug(msg, fields...)
	}
}

func (d *Default) Info(msg string, fields ...gox.Field[any]) {
	if d.config.Level.Rank() <= core.LevelInfo.Rank() {
		d.executor.Info(msg, fields...)
	}
}

func (d *Default) Warn(msg string, fields ...gox.Field[any]) {
	if d.config.Level.Rank() <= core.LevelWarn.Rank() {
		d.executor.Warn(msg, fields...)
	}
}

func (d *Default) Error(msg string, fields ...gox.Field[any]) {
	if d.config.Level.Rank() <= core.LevelError.Rank() {
		d.executor.Error(msg, fields...)
	}
}

func (d *Default) Panic(msg string, fields ...gox.Field[any]) {
	if d.config.Level.Rank() > core.LevelPanic.Rank() {
		return
	}

	d.addStacks(&fields)
	d.executor.Panic(msg, fields...)
}

func (d *Default) Fatal(msg string, fields ...gox.Field[any]) {
	if d.config.Level.Rank() > core.LevelFatal.Rank() {
		return
	}

	d.addStacks(&fields)
	d.executor.Fatal(msg, fields...)
}

func (d *Default) Sync() error {
	return d.executor.Sync()
}

func (d *Default) addStacks(fields *[]gox.Field[any]) {
	if nil == d.config.Stacktrace {
		return
	}

	callers := make([]uintptr, d.config.Stacktrace.Stack+1)
	count := runtime.Callers(d.config.Stacktrace.Skip+1, callers)
	frames := runtime.CallersFrames(callers[:count])

	stacks := make([]string, 0)
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
