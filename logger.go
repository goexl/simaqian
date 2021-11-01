package simaqian

import (
	`github.com/storezhang/gox`
)

// Logger 日志接口
type Logger interface {
	// Debug 记录调试日志
	Debug(msg string, fields gox.Fields)

	// Info 记录普通信息日志
	Info(msg string, fields gox.Fields)

	// Warn 记录警告日志
	Warn(msg string, fields gox.Fields)

	// Error 记录错误日志
	Error(msg string, fields gox.Fields)

	// Panic 记录异常日志，程序会退出，可以使用recover机制来阻止程序退出
	Panic(msg string, fields gox.Fields)

	// Fatal 记录致命错误日志，程序会退出
	Fatal(msg string, fields gox.Fields)
}
