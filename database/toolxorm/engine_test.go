package toolxorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestSF(t *testing.T) {
	e, err := NewXorm(Config{
		Driver: "mysql",
		Uri:    "root:123@/test?charset=utf8",
	})
	if err != nil {
		panic(err)
	}
	e.ShowSQL(true)
	db := NewEngine(e)
	datas := make([]struct {
		Id      string `xorm:"id" json:"id"`
		Account string `json:"account"`
	}, 0)

	err = db.SF(`
		select * from user where account = :account
	`, map[string]interface{}{
		"account": "root",
		"id":      []string{"root", "ss"},
	}).Find(&datas)
	if err != nil {
		panic(err)
	}
	fmt.Println(tools.Str.JSONString(datas, true))
}
