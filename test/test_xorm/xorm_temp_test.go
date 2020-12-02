package test_xorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestTemplate(t *testing.T) {
	db := NewEngine()
	s := make([]struct {
		Id string `xorm:"id" json:"id"`
	}, 0)
	err := db.With("users").SF(`select u.* from user u join user uu on u.id = uu.id`).Find(&s)
	if err != nil {
		panic(err)
	}
	fmt.Print(tools.Str.JSONString(s))
}
