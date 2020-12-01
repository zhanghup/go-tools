package test_xorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestEngineTemplate(t *testing.T) {
	str := tools.Str.Tmp(`
		{{ mwith 1}}
			asdjfklasjfd
		select * from user
	`).FuncMap(map[string]interface{}{
		"mwith": func(i int) string{
			return "aaa"
		},
	}).String()

	fmt.Println(str)
}
