package tog_test

import (
	"github.com/zhanghup/go-tools/tog"
	"testing"
)

func TestMyLogger(t *testing.T) {
	tog.Info("ddddddddddddddddd dsjkdj")
	tog.Error("ddddddddddddddddd dsjkdj")
	tog.Warn("ddddddddddddddddd dsjkdj")

	tog.WidthField(map[string]interface{}{"a": 1, "b": 2})
	tog.Warn("ddddddddddddddddd dsjkdj")

	tog.WidthField(map[string]interface{}{"c": 1, "d": 2})
	tog.Error("ddddddddddddddddd dsjkdj")
}
