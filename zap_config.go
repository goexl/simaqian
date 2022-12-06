package simaqian

type zapConfig struct {
	outputs []writer
	errors  []writer
}

func newZapConfig() *zapConfig {
	return &zapConfig{
		outputs: make([]writer, 0),
		errors:  make([]writer, 0),
	}
}
