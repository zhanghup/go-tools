package tog_test

import (
	"github.com/zhanghup/go-tools/tog"
	"testing"
)

func TestMyLogger(t *testing.T) {
	for i := 0; i < 100;i++{
		tog.Info("ddddddddddddddddd dsjkdj")
		tog.Error("ddddddddddddddddd dsjkdj")
		tog.Warn("ddddddddddddddddd dsjkdj")

		tog.Warn("ddddddddddddddddd dsjkdj")

		tog.Error("ddddddddddddddddd dsjkdj")

		tog.InfoAsJson(map[string]interface{}{"a":1,"b":2})
		tog.InfoAsJson(map[string]interface{}{"a":1,"b":2},true)
	}
}
