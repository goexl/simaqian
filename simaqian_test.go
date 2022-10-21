package simaqian_test

import (
	"testing"

	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

func TestZap(t *testing.T) {
	smq := simaqian.Must()
	smq.Info(`测试字符串`, field.String(`username`, `storezhang`))
	smq.Info(`测试字符串数组`, field.Strings(`args`, `--clean`, `test`, `deploy`))
	smq.Info(`测试字任意类型`, field.Any(`args`, []string{`--clean`, `test`, `deploy`}))
}
