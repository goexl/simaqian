package core

var (
	_ = Stdout
	_ = Stderr
)

type Writer string

// Stdout 标准输出流
func Stdout() Writer {
	return "stdout"
}

// Stderr 标准错误流
func Stderr() Writer {
	return "stderr"
}
