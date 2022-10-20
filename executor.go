package simaqian

import (
	"github.com/goexl/gox"
)

type executor interface {
	debug(msg string, fields ...gox.Field)

	info(msg string, fields ...gox.Field)

	warn(msg string, fields ...gox.Field)

	error(msg string, fields ...gox.Field)

	panic(msg string, fields ...gox.Field)

	fatal(msg string, fields ...gox.Field)
}
