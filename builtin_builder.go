package simaqian

type builtinBuilder struct {
	config *config
}

func newBuiltinBuilder(config *config) *builtinBuilder {
	return &builtinBuilder{
		config: config,
	}
}

func (bb *builtinBuilder) Build() Logger {
	return newDefaultLogger(bb.config, newBuiltin())
}
