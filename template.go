package simaqian

import (
	`github.com/storezhang/gox`
)

type template struct {
	zap    executor
	system executor
	zero   executor
	logrus executor

	options *options
}

func newTemplate(opts ...option) (_template *template, err error) {
	_template = new(template)
	_template.options = defaultOptions()
	for _, opt := range opts {
		opt.apply(_template.options)
	}

	switch _template.options.logType {
	case TypeLogrus:
		_template.logrus, err = newLogrus(_template.options)
	case TypeZap:
		_template.zap, err = newZap(_template.options)
	case TypeSystem:
		_template.system, err = newSystem(_template.options)
	}

	return
}

func (t *template) Sets(opts ...option) {
	for _, opt := range opts {
		opt.apply(t.options)
	}
}

func (t *template) Level() level {
	return t.options.level
}

func (t *template) Debug(msg string, fields ...gox.Field) {
	if t.options.level.rank() > LevelDebug.rank() {
		return
	}

	switch t.options.logType {
	case TypeZap:
		t.zap.debug(msg, fields...)
	case TypeSystem:
		t.system.debug(msg, fields...)
	case TypeLogrus:
		t.logrus.debug(msg, fields...)
	case TypeZero:
		t.zero.debug(msg, fields...)
	}
}

func (t *template) Info(msg string, fields ...gox.Field) {
	if t.options.level.rank() > LevelInfo.rank() {
		return
	}

	switch t.options.logType {
	case TypeZap:
		t.zap.info(msg, fields...)
	case TypeSystem:
		t.system.info(msg, fields...)
	case TypeLogrus:
		t.logrus.info(msg, fields...)
	case TypeZero:
		t.zero.info(msg, fields...)
	}
}

func (t *template) Warn(msg string, fields ...gox.Field) {
	if t.options.level.rank() > LevelWarn.rank() {
		return
	}

	switch t.options.logType {
	case TypeZap:
		t.zap.warn(msg, fields...)
	case TypeSystem:
		t.system.warn(msg, fields...)
	case TypeLogrus:
		t.logrus.warn(msg, fields...)
	case TypeZero:
		t.zero.warn(msg, fields...)
	}
}

func (t *template) Error(msg string, fields ...gox.Field) {
	if t.options.level.rank() > LevelError.rank() {
		return
	}

	switch t.options.logType {
	case TypeZap:
		t.zap.error(msg, fields...)
	case TypeSystem:
		t.system.error(msg, fields...)
	case TypeLogrus:
		t.logrus.error(msg, fields...)
	case TypeZero:
		t.zero.error(msg, fields...)
	}
}

func (t *template) Panic(msg string, fields ...gox.Field) {
	if t.options.level.rank() > LevelPanic.rank() {
		return
	}

	switch t.options.logType {
	case TypeZap:
		t.zap.panic(msg, fields...)
	case TypeSystem:
		t.system.panic(msg, fields...)
	case TypeLogrus:
		t.logrus.panic(msg, fields...)
	case TypeZero:
		t.zero.panic(msg, fields...)
	}
}

func (t *template) Fatal(msg string, fields ...gox.Field) {
	if t.options.level.rank() > LevelFatal.rank() {
		return
	}

	switch t.options.logType {
	case TypeZap:
		t.zap.fatal(msg, fields...)
	case TypeSystem:
		t.system.fatal(msg, fields...)
	case TypeLogrus:
		t.logrus.fatal(msg, fields...)
	case TypeZero:
		t.zero.fatal(msg, fields...)
	}
}
