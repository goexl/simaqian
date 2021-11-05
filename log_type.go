package simaqian

const (
	// TypeZap Uber Zap日志
	TypeZap logType = "zap"
	// TypeLogrus Logrus日志
	TypeLogrus logType = "logrus"
	// TypeZero ZeroLog日志
	TypeZero logType = "zero"
	// TypeSystem 系统日志
	TypeSystem logType = "system"
)

type logType string

// ParseType 解析类型
func ParseType(_type string) (lt *logType) {
	lt = new(logType)
	switch logType(_type) {
	case TypeLogrus:
		*lt = TypeLogrus
	case TypeZap:
		*lt = TypeZap
	case TypeSystem:
		*lt = TypeSystem
	case TypeZero:
		*lt = TypeZero
	}

	return
}
