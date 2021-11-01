package simaqian

type options struct {
	// 记录调用方法时增加方法过滤
	skip int
	// 日志类型
	logType Type
}

func defaultOptions() *options {
	return &options{
		skip:    1,
		logType: TypeZap,
	}
}
