package simaqian

type logrusBuilder struct {
	config *config
}

func newLogrusBuilder(config *config) *logrusBuilder {
	return &logrusBuilder{
		config: config,
	}
}

func (lb *logrusBuilder) Build() Logger {
	return newDefaultLogger(lb.config, newLogrus())
}
