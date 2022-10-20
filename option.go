package simaqian

var _ = NewOptions

type (
	option interface {
		apply(options *options)
	}

	options struct {
		level      level
		skip       int
		logType    logType
		stacktrace *stacktrace

		outputs []writer
		errors  []writer
	}
)

// NewOptions 创建选项，因为option接口不对外暴露，如果用户想在外面创建option并赋值将无法完成，特意提供创建option的快捷方式
func NewOptions(opts ...option) []option {
	return opts
}

func defaultOptions() *options {
	return &options{
		level:   LevelInfo,
		skip:    2,
		logType: TypeZap,
	}
}
