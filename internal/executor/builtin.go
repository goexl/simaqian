package executor

import (
	"fmt"
	"log"

	"github.com/goexl/gox"
	"github.com/goexl/simaqian/internal/core"
)

var _ core.Executor = (*Builtin)(nil)

type Builtin struct {
	logger *log.Logger
}

func NewBuiltin() *Builtin {
	return &Builtin{
		logger: log.Default(),
	}
}

func (b *Builtin) Debug(msg string, fields ...gox.Field[any]) {
	b.logger.Println(b.parse(core.LevelDebug, msg, fields...))
}

func (b *Builtin) Info(msg string, fields ...gox.Field[any]) {
	b.logger.Println(b.parse(core.LevelInfo, msg, fields...))
}

func (b *Builtin) Warn(msg string, fields ...gox.Field[any]) {
	b.logger.Println(b.parse(core.LevelWarn, msg, fields...))
}

func (b *Builtin) Error(msg string, fields ...gox.Field[any]) {
	b.logger.Println(b.parse(core.LevelError, msg, fields...))
}

func (b *Builtin) Panic(msg string, fields ...gox.Field[any]) {
	b.logger.Println(b.parse(core.LevelPanic, msg, fields...))
}

func (b *Builtin) Fatal(msg string, fields ...gox.Field[any]) {
	b.logger.Println(b.parse(core.LevelFatal, msg, fields...))
}

func (b *Builtin) Sync() (err error) {
	return
}

func (b *Builtin) parse(level core.Level, msg string, fields ...gox.Field[any]) (args []any) {
	args = make([]any, 0, len(fields)+1)
	args = append(args, level)
	args = append(args, msg)
	if 0 != len(fields) {
		args = append(args, "[")
		for _, field := range fields {
			args = append(args, fmt.Sprintf("{%s = %v}", field.Key(), field.Value()))
		}
		args = append(args, "]")
	}

	return
}
