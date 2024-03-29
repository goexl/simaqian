package simaqian_test

import (
	"encoding/json"
	"testing"

	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

type jsonField struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestZap(test *testing.T) {
	smq := simaqian.Default()
	smq.Info("测试字符串", field.New[string]("username", "storezhang"))
	smq.Info("测试字符串数组", field.New[[]string]("args", []string{"--clean", "test", "deploy"}))
	smq.Info("测试任意类型", field.New("args", []string{"--clean", "test", "deploy"}))
	smq.Info("测试JSON", field.New("json", jsonField{
		Username: "storezhang",
		Password: "password",
	}))
	smq.Info("测试JSON字符串", field.New("json.raw", json.RawMessage(`{"username":"storezhang  ","password":"password"}`)))
	_ = smq.Sync()
}
