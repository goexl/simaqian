package glog

const (
	// LogTypeDebug 调试日志
	LogTypeDebug LogType = "debug"
	// LogTypeInfo 普通信息日志
	LogTypeInfo LogType = "debug"
	// LogTypeWarn 警告日志
	LogTypeWarn LogType = "debug"
	// LogTypeError 错误日志
	LogTypeError LogType = "debug"
	// LogTypePanic 异常日志，程序会退出，可以使用recover机制来阻止程序退出
	LogTypePanic LogType = "debug"
	// LogTypeFatal 致命错误日志，程序会退出
	LogTypeFatal LogType = "debug"
)

// LogType 日志类型
type LogType string
