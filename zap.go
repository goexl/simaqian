package simaqian

import (
	`time`

	`github.com/goexl/gox`
	`github.com/goexl/gox/field`
	`go.uber.org/zap`
)

var _ executor = (*_zap)(nil)

type _zap struct {
	logger *zap.Logger
}

func newZap(options *options) (logger *_zap, err error) {
	logger = new(_zap)

	zapOptions := []zap.Option{
		zap.AddCallerSkip(options.skip),
	}
	// 日志输出时，因为用glog封装了一层，需要在寻找调用链的时候跳过，不然会一直输出glog的调用点
	if logger.logger, err = zap.NewProduction(zapOptions...); nil != err {
		return
	}
	defer func() {
		_ = logger.logger.Sync()
	}()

	return
}

func (z *_zap) debug(msg string, fields ...gox.Field) {
	z.logger.Debug(msg, z.parse(fields...)...)
}

func (z *_zap) info(msg string, fields ...gox.Field) {
	z.logger.Info(msg, z.parse(fields...)...)
}

func (z *_zap) warn(msg string, fields ...gox.Field) {
	z.logger.Warn(msg, z.parse(fields...)...)
}

func (z *_zap) error(msg string, fields ...gox.Field) {
	z.logger.Error(msg, z.parse(fields...)...)
}

func (z *_zap) panic(msg string, fields ...gox.Field) {
	z.logger.Panic(msg, z.parse(fields...)...)
}

func (z *_zap) fatal(msg string, fields ...gox.Field) {
	z.logger.Fatal(msg, z.parse(fields...)...)
}

func (z *_zap) parse(fields ...gox.Field) (zapFields []zap.Field) {
	zapFields = make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		switch f.Value().(type) {
		case *field.BoolField:
			zapFields = append(zapFields, zap.Boolp(f.Key(), f.Value().(*bool)))
		case *field.BoolsField:
			zapFields = append(zapFields, zap.Bools(f.Key(), f.Value().([]bool)))
		case *field.Int8Field:
			zapFields = append(zapFields, zap.Int8p(f.Key(), f.Value().(*int8)))
		case *field.IntField:
			zapFields = append(zapFields, zap.Intp(f.Key(), f.Value().(*int)))
		case *field.IntsField:
			zapFields = append(zapFields, zap.Ints(f.Key(), f.Value().([]int)))
		case *field.UintField:
			zapFields = append(zapFields, zap.Uintp(f.Key(), f.Value().(*uint)))
		case *field.UintsField:
			zapFields = append(zapFields, zap.Uints(f.Key(), f.Value().([]uint)))
		case *field.Int64Field:
			zapFields = append(zapFields, zap.Int64p(f.Key(), f.Value().(*int64)))
		case *field.Int64sField:
			zapFields = append(zapFields, zap.Int64s(f.Key(), f.Value().([]int64)))
		case *field.Float32Field:
			zapFields = append(zapFields, zap.Float32p(f.Key(), f.Value().(*float32)))
		case *field.Float64Field:
			zapFields = append(zapFields, zap.Float64p(f.Key(), f.Value().(*float64)))
		case *field.Float64sField:
			zapFields = append(zapFields, zap.Float64s(f.Key(), f.Value().([]float64)))
		case *field.StringField:
			zapFields = append(zapFields, zap.Stringp(f.Key(), f.Value().(*string)))
		case *field.StringsField:
			zapFields = append(zapFields, zap.Strings(f.Key(), f.Value().([]string)))
		case *field.TimeField:
			zapFields = append(zapFields, zap.Timep(f.Key(), f.Value().(*time.Time)))
		case *field.TimesField:
			zapFields = append(zapFields, zap.Times(f.Key(), f.Value().([]time.Time)))
		case *field.DurationField:
			zapFields = append(zapFields, zap.Durationp(f.Key(), f.Value().(*time.Duration)))
		case *field.DurationsField:
			zapFields = append(zapFields, zap.Durations(f.Key(), f.Value().([]time.Duration)))
		case *field.ErrorField:
			zapFields = append(zapFields, zap.Error(f.Value().(error)))
		default:
			zapFields = append(zapFields, zap.Any(f.Key(), f.Value()))
		}
	}

	return
}
