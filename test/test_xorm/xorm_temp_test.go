package test_xorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
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



func TestSF(t *testing.T) {
	datas := make([]struct {
		Id      string `xorm:"id" json:"id"`
		Account string `json:"account"`
	}, 0)

	err := NewEngine().SF(`
		select * from user where account = :account
	`, map[string]interface{}{
		"account": "root",
		"id":      []string{"root", "ss"},
	}).Find(&datas)
	if err != nil {
		panic(err)
	}

	datas = make([]struct {
		Id      string `xorm:"id" json:"id"`
		Account string `json:"account"`
	}, 0)

	err = NewEngine().SF(`
		select * from user where account = :account
	`, map[string]interface{}{
		"account": "root",
		"id":      []string{"root", "ss"},
	}).Find(&datas)
	if err != nil {
		panic(err)
	}
	tog.Info(tools.Str.JSONString(datas))
}

func TestSession_Exec(t *testing.T) {
	err := NewEngine().SF("update user_token set status = 0 ").Exec()
	if err != nil {
		panic(err)
	}
}



func TestPage(t *testing.T) {
	dict := make([]Dict, 0)
	n, err := NewEngine().SF("select * from dict").Page(2, 2, true, &dict)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(tools.Str.JSONString(dict, true))
}

func TestSF2(t *testing.T) {
	dict := make([]Dict, 0)
	n, err := NewEngine().SF2("select * from dict where status = ?",1).Page(2, 2, true, &dict)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(tools.Str.JSONString(dict, true))
}