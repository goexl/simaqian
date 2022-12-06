package simaqian

import (
	"fmt"
	"log"

	"github.com/goexl/gox"
)

var _ executor = (*builtin)(nil)

type builtin struct {
	logger *log.Logger
}

func newBuiltin() *builtin {
	return &builtin{
		logger: log.Default(),
	}
}

func (b *builtin) debug(msg string, fields ...gox.Field[any]) {
	b.logger.Println(b.parse(LevelDebug, msg, fields...))
}

func (b *builtin) info(msg string, fields ...gox.Field[any]) {
	b.logger.Println(b.parse(LevelInfo, msg, fields...))
}

func (b *builtin) warn(msg string, fields ...gox.Field[any]) {
	b.logger.Println(b.parse(LevelWarn, msg, fields...))
}

func (b *builtin) error(msg string, fields ...gox.Field[any]) {
	b.logger.Println(b.parse(LevelError, msg, fields...))
}

func (b *builtin) panic(msg string, fields ...gox.Field[any]) {
	b.logger.Println(b.parse(LevelPanic, msg, fields...))
}

func (b *builtin) fatal(msg string, fields ...gox.Field[any]) {
	b.logger.Println(b.parse(LevelFatal, msg, fields...))
}

func (b *builtin) sync() (err error) {
	return
}

func (b *builtin) parse(level level, msg string, fields ...gox.Field[any]) (args []interface{}) {
	args = make([]interface{}, 0, len(fields)+1)
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
