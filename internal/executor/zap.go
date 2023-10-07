package executor

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/simaqian/internal/core"
	"github.com/goexl/simaqian/internal/param"
	"go.uber.org/zap"
)

var _ core.Executor = (*Zap)(nil)

type Zap struct {
	zap *zap.Logger
}

func NewZap(param *param.Zap) (logger *Zap, err error) {
	logger = new(Zap)
	config := zap.NewProductionConfig()
	outputsSize := len(param.Outputs)
	if 0 != outputsSize {
		config.OutputPaths = make([]string, 0, outputsSize)
		for _, output := range param.Outputs {
			config.OutputPaths = append(config.OutputPaths, string(output))
		}
	}

	errorsSize := len(param.Errors)
	if 0 != errorsSize {
		config.ErrorOutputPaths = make([]string, 0, errorsSize)
		for _, _error := range param.Errors {
			config.ErrorOutputPaths = append(config.OutputPaths, string(_error))
		}
	}
	logger.zap, err = config.Build()

	return
}

func (z *Zap) Debug(msg string, fields ...gox.Field[any]) {
	z.zap.Debug(msg, z.parse(fields...)...)
}

func (z *Zap) Info(msg string, fields ...gox.Field[any]) {
	z.zap.Info(msg, z.parse(fields...)...)
}

func (z *Zap) Warn(msg string, fields ...gox.Field[any]) {
	z.zap.Warn(msg, z.parse(fields...)...)
}

func (z *Zap) Error(msg string, fields ...gox.Field[any]) {
	z.zap.Error(msg, z.parse(fields...)...)
}

func (z *Zap) Panic(msg string, fields ...gox.Field[any]) {
	z.zap.Panic(msg, z.parse(fields...)...)
}

func (z *Zap) Fatal(msg string, fields ...gox.Field[any]) {
	z.zap.Fatal(msg, z.parse(fields...)...)
}

func (z *Zap) Sync() error {
	return z.zap.Sync()
}

func (z *Zap) parse(fields ...gox.Field[any]) (parsed []zap.Field) {
	parsed = make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		if "" == f.Key() || nil == f.Value() {
			continue
		}

		switch value := f.Value().(type) {
		case bool:
			parsed = append(parsed, zap.Bool(f.Key(), value))
		case *bool:
			parsed = append(parsed, zap.Boolp(f.Key(), value))
		case []bool:
			parsed = append(parsed, zap.Bools(f.Key(), value))
		case *[]bool:
			parsed = append(parsed, zap.Bools(f.Key(), *value))
		case int8:
			parsed = append(parsed, zap.Int8(f.Key(), value))
		case *int8:
			parsed = append(parsed, zap.Int8p(f.Key(), value))
		case int:
			parsed = append(parsed, zap.Int(f.Key(), value))
		case *int:
			parsed = append(parsed, zap.Intp(f.Key(), value))
		case []int:
			parsed = append(parsed, zap.Ints(f.Key(), value))
		case *[]int:
			parsed = append(parsed, zap.Ints(f.Key(), *value))
		case uint:
			parsed = append(parsed, zap.Uint(f.Key(), value))
		case *uint:
			parsed = append(parsed, zap.Uintp(f.Key(), value))
		case []uint:
			parsed = append(parsed, zap.Uints(f.Key(), value))
		case *[]uint:
			parsed = append(parsed, zap.Uints(f.Key(), *value))
		case int64:
			parsed = append(parsed, zap.Int64(f.Key(), value))
		case *int64:
			parsed = append(parsed, zap.Int64p(f.Key(), value))
		case []int64:
			parsed = append(parsed, zap.Int64s(f.Key(), value))
		case *[]int64:
			parsed = append(parsed, zap.Int64s(f.Key(), *value))
		case float32:
			parsed = append(parsed, zap.Float32(f.Key(), value))
		case *float32:
			parsed = append(parsed, zap.Float32p(f.Key(), value))
		case float64:
			parsed = append(parsed, zap.Float64(f.Key(), value))
		case *float64:
			parsed = append(parsed, zap.Float64p(f.Key(), value))
		case []float64:
			parsed = append(parsed, zap.Float64s(f.Key(), value))
		case *[]float64:
			parsed = append(parsed, zap.Float64s(f.Key(), *value))
		case *string:
			parsed = append(parsed, zap.Stringp(f.Key(), value))
		case []string:
			parsed = append(parsed, zap.Strings(f.Key(), value))
		case *[]string:
			parsed = append(parsed, zap.Strings(f.Key(), *value))
		case time.Time:
			parsed = append(parsed, zap.Time(f.Key(), value))
		case *time.Time:
			parsed = append(parsed, zap.Timep(f.Key(), value))
		case []time.Time:
			parsed = append(parsed, zap.Times(f.Key(), value))
		case time.Duration:
			parsed = append(parsed, zap.Duration(f.Key(), value))
		case *time.Duration:
			parsed = append(parsed, zap.Durationp(f.Key(), value))
		case []time.Duration:
			parsed = append(parsed, zap.Durations(f.Key(), value))
		case json.Marshaler, []json.Marshaler:
			// 一定要放在 fmt.Stringer 前面，保证优先使用 json 作为序列化器
			parsed = append(parsed, zap.Reflect(f.Key(), f.Value()))
		case fmt.Stringer:
			parsed = append(parsed, zap.Stringer(f.Key(), value))
		case []fmt.Stringer:
			parsed = append(parsed, zap.Stringers(f.Key(), value))
		case error:
			parsed = append(parsed, zap.Error(value))
		default:
			parsed = append(parsed, zap.Any(f.Key(), f.Value()))
		}
	}

	return
}
