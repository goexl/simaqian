package simaqian

var (
	_ = Error

	_ option = (*optionError)(nil)
)

type optionError struct {
	error writer
}

// Error 错误流
func Error(error writer) *optionError {
	return &optionError{
		error: error,
	}
}

func (e *optionError) apply(options *options) {
	options.errors = append(options.errors, e.error)
}
