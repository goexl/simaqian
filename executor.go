package simaqian

import (
	"github.com/goexl/gox"
)

type executor interface {
	debug(msg string, fields ...gox.Field[any])

	info(msg string, fields ...gox.Field[any])

	warn(msg string, fields ...gox.Field[any])

	error(msg string, fields ...gox.Field[any])

	panic(msg string, fields ...gox.Field[any])

	fatal(msg string, fields ...gox.Field[any])
}
