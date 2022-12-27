package simaqian

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/goexl/gox"
	"go.uber.org/zap"
)

var _ executor = (*_zap)(nil)

type _zap struct {
	logger *zap.Logger
}

func newZap(config *config, self *zapConfig) (logger *_zap, err error) {
	logger = new(_zap)

	_config := zap.NewDevelopmentConfig()
	outputsSize := len(self.outputs)
	if 0 != outputsSize {
		_config.OutputPaths = make([]string, 0, outputsSize)
		for _, output := range self.outputs {
			_config.OutputPaths = append(_config.OutputPaths, string(output))
		}
	}

	errorsSize := len(self.errors)
	if 0 != errorsSize {
		_config.ErrorOutputPaths = make([]string, 0, errorsSize)
		for _, _error := range self.errors {
			_config.ErrorOutputPaths = append(_config.OutputPaths, string(_error))
		}
	}

	zapOptions := []zap.Option{
		zap.AddCallerSkip(config.skip),
	}
	if logger.logger, err = _config.Build(zapOptions...); nil != err {
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

func (z *_zap) sync() error {
	return z.logger.Sync()
}

func (z *_zap) parse(fields ...gox.Field[any]) (zfs []zap.Field) {
	zfs = make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		if "" == f.Key() || nil == f.Value() {
			continue
		}

		switch value := f.Value().(type) {
		case bool:
			zfs = append(zfs, zap.Bool(f.Key(), value))
		case *bool:
			zfs = append(zfs, zap.Boolp(f.Key(), value))
		case []bool:
			zfs = append(zfs, zap.Bools(f.Key(), value))
		case *[]bool:
			zfs = append(zfs, zap.Bools(f.Key(), *value))
		case int8:
			zfs = append(zfs, zap.Int8(f.Key(), value))
		case *int8:
			zfs = append(zfs, zap.Int8p(f.Key(), value))
		case int:
			zfs = append(zfs, zap.Int(f.Key(), value))
		case *int:
			zfs = append(zfs, zap.Intp(f.Key(), value))
		case []int:
			zfs = append(zfs, zap.Ints(f.Key(), value))
		case *[]int:
			zfs = append(zfs, zap.Ints(f.Key(), *value))
		case uint:
			zfs = append(zfs, zap.Uint(f.Key(), value))
		case *uint:
			zfs = append(zfs, zap.Uintp(f.Key(), value))
		case []uint:
			zfs = append(zfs, zap.Uints(f.Key(), value))
		case *[]uint:
			zfs = append(zfs, zap.Uints(f.Key(), *value))
		case int64:
			zfs = append(zfs, zap.Int64(f.Key(), value))
		case *int64:
			zfs = append(zfs, zap.Int64p(f.Key(), value))
		case []int64:
			zfs = append(zfs, zap.Int64s(f.Key(), value))
		case *[]int64:
			zfs = append(zfs, zap.Int64s(f.Key(), *value))
		case float32:
			zfs = append(zfs, zap.Float32(f.Key(), value))
		case *float32:
			zfs = append(zfs, zap.Float32p(f.Key(), value))
		case float64:
			zfs = append(zfs, zap.Float64(f.Key(), value))
		case *float64:
			zfs = append(zfs, zap.Float64p(f.Key(), value))
		case []float64:
			zfs = append(zfs, zap.Float64s(f.Key(), value))
		case *[]float64:
			zfs = append(zfs, zap.Float64s(f.Key(), *value))
		case *string:
			zfs = append(zfs, zap.Stringp(f.Key(), value))
		case []string:
			zfs = append(zfs, zap.Strings(f.Key(), value))
		case *[]string:
			zfs = append(zfs, zap.Strings(f.Key(), *value))
		case json.Marshaler, []json.Marshaler:
			// 一定要放在 fmt.Stringer 前面，保证优先使用 json 作为序列化器
			zfs = append(zfs, zap.Reflect(f.Key(), f.Value()))
		case fmt.Stringer:
			zfs = append(zfs, zap.Stringer(f.Key(), value))
		case []fmt.Stringer:
			zfs = append(zfs, zap.Stringers(f.Key(), value))
		case time.Time:
			zfs = append(zfs, zap.Time(f.Key(), value))
		case *time.Time:
			zfs = append(zfs, zap.Timep(f.Key(), value))
		case []time.Time:
			zfs = append(zfs, zap.Times(f.Key(), value))
		case time.Duration:
			zfs = append(zfs, zap.Duration(f.Key(), value))
		case *time.Duration:
			zfs = append(zfs, zap.Durationp(f.Key(), value))
		case []time.Duration:
			zfs = append(zfs, zap.Durations(f.Key(), value))
		case error:
			zfs = append(zfs, zap.Error(value))
		default:
			zfs = append(zfs, zap.Any(f.Key(), f.Value()))
		}
	}

	return
}
