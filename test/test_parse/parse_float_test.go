package test_parse

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestFloat(t *testing.T) {
	s := tools.Float64ToStr(1.111111111111111111111111111111111111111)
	fmt.Println(s)
}
