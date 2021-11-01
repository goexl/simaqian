package simaqian

import (
	`fmt`
	`log`

	`github.com/storezhang/gox`
)

// System日志
type systemLogger struct {
	logger *log.Logger
}

// 创建内置系统日志记录器
func newSystemLogger(_ *options) (logger *systemLogger, err error) {
	logger = &systemLogger{
		logger: log.Default(),
	}

	return
}

func (sl *systemLogger) Debug(msg string, fields gox.Fields) {
	sl.logger.Println(sl.parse(LogTypeDebug, msg, fields))
}

func (sl *systemLogger) Info(msg string, fields gox.Fields) {
	sl.logger.Println(sl.parse(LogTypeInfo, msg, fields))
}

func (sl *systemLogger) Warn(msg string, fields gox.Fields) {
	sl.logger.Println(sl.parse(LogTypeWarn, msg, fields))
}

func (sl *systemLogger) Error(msg string, fields gox.Fields) {
	sl.logger.Println(sl.parse(LogTypeError, msg, fields))
}

func (sl *systemLogger) Panic(msg string, fields gox.Fields) {
	sl.logger.Println(sl.parse(LogTypePanic, msg, fields))
}

func (sl *systemLogger) Fatal(msg string, fields gox.Fields) {
	sl.logger.Println(sl.parse(LogTypeFatal, msg, fields))
}

func (sl *systemLogger) parse(logType LogType, msg string, fields gox.Fields) (args []interface{}) {
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
