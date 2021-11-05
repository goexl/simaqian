package simaqian

import (
	`fmt`
	`log`

	`github.com/storezhang/gox`
)

var _ executor = (*system)(nil)

type system struct {
	logger *log.Logger
}

func newSystem(_ *options) (logger *system, err error) {
	logger = &system{
		logger: log.Default(),
	}

	return
}

func (s *system) debug(msg string, fields ...gox.Field) {
	s.logger.Println(s.parse(LevelDebug, msg, fields...))
}

func (s *system) info(msg string, fields ...gox.Field) {
	s.logger.Println(s.parse(LevelInfo, msg, fields...))
}

func (s *system) warn(msg string, fields ...gox.Field) {
	s.logger.Println(s.parse(LevelWarn, msg, fields...))
}

func (s *system) error(msg string, fields ...gox.Field) {
	s.logger.Println(s.parse(LevelError, msg, fields...))
}

func (s *system) panic(msg string, fields ...gox.Field) {
	s.logger.Println(s.parse(LevelPanic, msg, fields...))
}

func (s *system) fatal(msg string, fields ...gox.Field) {
	s.logger.Println(s.parse(LevelFatal, msg, fields...))
}

func (s *system) parse(level level, msg string, fields ...gox.Field) (args []interface{}) {
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
