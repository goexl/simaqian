package glog

import (
	`fmt`
	`log`

	`github.com/storezhang/gox`
)

// SystemLogger System日志
type SystemLogger struct {
	logger *log.Logger
}

// NewSystemLogger 创建内置系统日志记录器
func NewSystemLogger(logger *log.Logger) *SystemLogger {
	return &SystemLogger{
		logger: logger,
	}
}

func (sl *SystemLogger) Debug(msg string, fields ...gox.Field) {
	sl.logger.Println(sl.parse(LogTypeDebug, msg, fields...))
}

func (sl *SystemLogger) Info(msg string, fields ...gox.Field) {
	sl.logger.Println(sl.parse(LogTypeInfo, msg, fields...))
}

func (sl *SystemLogger) Warn(msg string, fields ...gox.Field) {
	sl.logger.Println(sl.parse(LogTypeWarn, msg, fields...))
}

func (sl *SystemLogger) Error(msg string, fields ...gox.Field) {
	sl.logger.Println(sl.parse(LogTypeError, msg, fields...))
}

func (sl *SystemLogger) Panic(msg string, fields ...gox.Field) {
	sl.logger.Println(sl.parse(LogTypePanic, msg, fields...))
}

func (sl *SystemLogger) Fatal(msg string, fields ...gox.Field) {
	sl.logger.Println(sl.parse(LogTypeFatal, msg, fields...))
}

func (sl *SystemLogger) parse(logType LogType, msg string, fields ...gox.Field) (args []interface{}) {
	args = make([]interface{}, 0, len(fields)+1)
	args = append(args, logType)
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
