package builder

import (
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/internal"
	"github.com/goexl/simaqian/internal/internal/builder"
	"github.com/goexl/simaqian/internal/param"
)

type Core struct {
	config *param.Core
}

func NewBuilder() *Core {
	return &Core{
		config: param.NewCore(),
	}
}

func (c *Core) Debug() *Core {
	return c.Level(core.LevelDebug)
}

func (c *Core) Info() *Core {
	return c.Level(core.LevelInfo)
}

func (c *Core) Warn() *Core {
	return c.Level(core.LevelWarn)
}

func (c *Core) Error() *Core {
	return c.Level(core.LevelError)
}

func (c *Core) Fatal() *Core {
	return c.Level(core.LevelFatal)
}

func (c *Core) Level(lvl core.Level) *Core {
	c.config.Level = lvl

	return c
}

func (c *Core) Skip(skip int) *Core {
	c.config.Skip = skip

	return c
}

func (c *Core) Stacktrace(skip int, stack int) *Core {
	c.config.Stacktrace = &internal.Stacktrace{
		Skip:  skip,
		Stack: stack,
	}

	return c
}

func (c *Core) Zap() *builder.Zap {
	return builder.NewZap(c.config)
}

func (c *Core) Logrus() *builder.Logrus {
	return builder.NewLogrus(c.config)
}

func (c *Core) Builtin() *builder.Builtin {
	return builder.NewBuiltin(c.config)
}

func (c *Core) Loki() *builder.Loki {
	return builder.NewLoki(c.config)
}
