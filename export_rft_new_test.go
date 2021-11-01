package tools_test

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"reflect"
	"testing"
	"xorm.io/xorm"
)

func TestRftInterfaceInfo(t *testing.T) {
	info := map[string]interface{}{
		"name": xorm.Engine{},
		"kind": "1",
		"sss":  nil,
		"haha": struct {
			Name string
			A    string
			b    string
		}{Name: "aaaa"},
	}

	tools.RftInterfaceInfo(info, func(field string, value interface{}, tag reflect.StructTag) bool {
		fmt.Println(field, value)
		return true
	})

}
