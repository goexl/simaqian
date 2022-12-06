package simaqian

type zapBuilder struct {
	config *config
	self   *zapConfig
}

func newZapBuilder(config *config) *zapBuilder {
	return &zapBuilder{
		config: config,
		self:   newZapConfig(),
	}
}

func (zb *zapBuilder) Output(writer writer) *zapBuilder {
	zb.self.outputs = append(zb.self.outputs, writer)

	return zb
}

func (zb *zapBuilder) Error(writer writer) *zapBuilder {
	zb.self.errors = append(zb.self.errors, writer)

	return zb
}

func (zb *zapBuilder) Build() (logger Logger, err error) {
	if executor, ze := newZap(zb.config, zb.self); nil != ze {
		err = ze
	} else {
		logger = newDefaultLogger(zb.config, executor)
	}

	return
}
