package simaqian

type builder struct {
	config *config
}

func newBuilder() *builder {
	return &builder{
		config: newConfig(),
	}
}

func (b *builder) Debug() *builder {
	return b.Level(LevelDebug)
}

func (b *builder) Info() *builder {
	return b.Level(LevelInfo)
}

func (b *builder) Warn() *builder {
	return b.Level(LevelWarn)
}

func (b *builder) Error() *builder {
	return b.Level(LevelError)
}

func (b *builder) Fatal() *builder {
	return b.Level(LevelFatal)
}

func (b *builder) Level(lvl level) *builder {
	b.config.level = lvl

	return b
}

func (b *builder) Skip(skip int) *builder {
	b.config.skip = skip

	return b
}

func (b *builder) Stacktrace(skip int, stack int) *builder {
	b.config.stacktrace = &stacktrace{
		skip:  skip,
		stack: stack,
	}

	return b
}

func (b *builder) Zap() *zapBuilder {
	return newZapBuilder(b.config)
}

func (b *builder) Logrus() *logrusBuilder {
	return newLogrusBuilder(b.config)
}

func (b *builder) Builtin() *builtinBuilder {
	return newBuiltinBuilder(b.config)
}
