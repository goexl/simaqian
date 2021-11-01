package simaqian

const (
	// TypeZap Uber Zap日志
	TypeZap Type = "zap"
	// TypeLogrus Logrus日志
	TypeLogrus Type = "logrus"
	// TypeZeroLog ZeroLog日志
	TypeZeroLog Type = "zero"
	// TypeSystem 系统日志
	TypeSystem Type = "system"
)

// Type 日志类型
type Type string
