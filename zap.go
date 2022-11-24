package simaqian

import (
	"fmt"
	"time"

	"github.com/goexl/gox"
	"go.uber.org/zap"
)

var _ executor = (*_zap)(nil)

type _zap struct {
	logger *zap.Logger
}

func newZap(options *options) (logger *_zap, err error) {
	logger = new(_zap)

	config := zap.NewDevelopmentConfig()
	outputsSize := len(options.outputs)
	if 0 != outputsSize {
		config.OutputPaths = make([]string, 0, outputsSize)
		for _, output := range options.outputs {
			config.OutputPaths = append(config.OutputPaths, string(output))
		}
	}

	errorsSize := len(options.errors)
	if 0 != errorsSize {
		config.ErrorOutputPaths = make([]string, 0, errorsSize)
		for _, _error := range options.errors {
			config.ErrorOutputPaths = append(config.OutputPaths, string(_error))
		}
	}

	zapOptions := []zap.Option{
		zap.AddCallerSkip(options.skip),
	}
	if logger.logger, err = config.Build(zapOptions...); nil != err {
		return
	}
	defer func() {
		_ = logger.logger.Sync()
	}()

	return
}

func (z *_zap) debug(msg string, fields ...gox.Field[any]) {
	z.logger.Debug(msg, z.parse(fields...)...)
}

func (z *_zap) info(msg string, fields ...gox.Field[any]) {
	z.logger.Info(msg, z.parse(fields...)...)
}

func (z *_zap) warn(msg string, fields ...gox.Field[any]) {
	z.logger.Warn(msg, z.parse(fields...)...)
}

func (z *_zap) error(msg string, fields ...gox.Field[any]) {
	z.logger.Error(msg, z.parse(fields...)...)
}

func (z *_zap) panic(msg string, fields ...gox.Field[any]) {
	z.logger.Panic(msg, z.parse(fields...)...)
}

func (z *_zap) fatal(msg string, fields ...gox.Field[any]) {
	z.logger.Fatal(msg, z.parse(fields...)...)
}

func (z *_zap) parse(fields ...gox.Field[any]) (zapFields []zap.Field) {
	zapFields = make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		if `` == f.Key() || nil == f.Value() {
			continue
		}

		switch value := f.Value().(type) {
		case bool:
			zapFields = append(zapFields, zap.Bool(f.Key(), value))
		case *bool:
			zapFields = append(zapFields, zap.Boolp(f.Key(), value))
		case []bool:
			zapFields = append(zapFields, zap.Bools(f.Key(), value))
		case *[]bool:
			zapFields = append(zapFields, zap.Bools(f.Key(), *value))
		case int8:
			zapFields = append(zapFields, zap.Int8(f.Key(), value))
		case *int8:
			zapFields = append(zapFields, zap.Int8p(f.Key(), value))
		case int:
			zapFields = append(zapFields, zap.Int(f.Key(), value))
		case *int:
			zapFields = append(zapFields, zap.Intp(f.Key(), value))
		case []int:
			zapFields = append(zapFields, zap.Ints(f.Key(), value))
		case *[]int:
			zapFields = append(zapFields, zap.Ints(f.Key(), *value))
		case uint:
			zapFields = append(zapFields, zap.Uint(f.Key(), value))
		case *uint:
			zapFields = append(zapFields, zap.Uintp(f.Key(), value))
		case []uint:
			zapFields = append(zapFields, zap.Uints(f.Key(), value))
		case *[]uint:
			zapFields = append(zapFields, zap.Uints(f.Key(), *value))
		case int64:
			zapFields = append(zapFields, zap.Int64(f.Key(), value))
		case *int64:
			zapFields = append(zapFields, zap.Int64p(f.Key(), value))
		case []int64:
			zapFields = append(zapFields, zap.Int64s(f.Key(), value))
		case *[]int64:
			zapFields = append(zapFields, zap.Int64s(f.Key(), *value))
		case float32:
			zapFields = append(zapFields, zap.Float32(f.Key(), value))
		case *float32:
			zapFields = append(zapFields, zap.Float32p(f.Key(), value))
		case float64:
			zapFields = append(zapFields, zap.Float64(f.Key(), value))
		case *float64:
			zapFields = append(zapFields, zap.Float64p(f.Key(), value))
		case []float64:
			zapFields = append(zapFields, zap.Float64s(f.Key(), value))
		case *[]float64:
			zapFields = append(zapFields, zap.Float64s(f.Key(), *value))
		case *string:
			zapFields = append(zapFields, zap.Stringp(f.Key(), value))
		case []string:
			zapFields = append(zapFields, zap.Strings(f.Key(), value))
		case *[]string:
			zapFields = append(zapFields, zap.Strings(f.Key(), *value))
		case fmt.Stringer:
			zapFields = append(zapFields, zap.Stringer(f.Key(), value))
		case []fmt.Stringer:
			zapFields = append(zapFields, zap.Stringers(f.Key(), value))
		case time.Time:
			zapFields = append(zapFields, zap.Time(f.Key(), value))
		case *time.Time:
			zapFields = append(zapFields, zap.Timep(f.Key(), value))
		case []time.Time:
			zapFields = append(zapFields, zap.Times(f.Key(), value))
		case time.Duration:
			zapFields = append(zapFields, zap.Duration(f.Key(), value))
		case *time.Duration:
			zapFields = append(zapFields, zap.Durationp(f.Key(), value))
		case []time.Duration:
			zapFields = append(zapFields, zap.Durations(f.Key(), value))
		case error:
			zapFields = append(zapFields, zap.Error(value))
		default:
			zapFields = append(zapFields, zap.Any(f.Key(), f.Value()))
		}
	}

	return
}
