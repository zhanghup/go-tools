package test_str_test

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestStr(t *testing.T) {
	fmt.Println(tools.Str.Fmt("dasfj %s", "666"))
}

func TestUid(t *testing.T) {
	fmt.Println(tools.Str.Uid())
}

func TestJSONString(t *testing.T) {
	fmt.Println(tools.Str.JSONString(map[string]interface{}{"a": 1, "b": 2}))
	fmt.Println(tools.Str.JSONString(map[string]interface{}{"a": 1, "b": 2, "c": map[string]interface{}{"d": 4}}, true))
}

func TestContains(t *testing.T) {
	fmt.Println(tools.Str.Contains([]string{"a", "b", "c"}, "c"))
	fmt.Println(tools.Str.Contains([]string{"a", "b", "c"}, "d"))
}

func TestRandom(t *testing.T) {
	fmt.Println(tools.Str.RandString(100))
	fmt.Println(tools.Str.RandString(100, true))
}
