package simaqian

var (
	_ = Stdout
	_ = Stderr
)

type writer string

// Stdout 标准输出流
func Stdout() writer {
	return `stdout`
}

// Stderr 标准错误流
func Stderr() writer {
	return `stderr`
}
