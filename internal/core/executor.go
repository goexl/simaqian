package core

import (
	"github.com/goexl/gox"
)

type Executor interface {
	Debug(msg string, fields ...gox.Field[any])

	Info(msg string, fields ...gox.Field[any])

	Warn(msg string, fields ...gox.Field[any])

	Error(msg string, fields ...gox.Field[any])

	Panic(msg string, fields ...gox.Field[any])

	Fatal(msg string, fields ...gox.Field[any])

	Sync() error
}
