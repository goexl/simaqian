package simaqian_test

import (
	"testing"

	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

func TestZap(test *testing.T) {
	smq := simaqian.Must()
	smq.Info(`测试字符串`, field.New[string](`username`, `storezhang`))
	smq.Info(`测试字符串数组`, field.New[[]string](`args`, []string{`--clean`, `test`, `deploy`}))
	smq.Info(`测试字任意类型`, field.New(`args`, []string{`--clean`, `test`, `deploy`}))
}
