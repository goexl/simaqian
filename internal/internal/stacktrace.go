package internal

type Stacktrace struct {
	Skip  int
	Stack int
}

func NewStacktrace() *Stacktrace {
	return &Stacktrace{
		Skip: 2,
	}
}
