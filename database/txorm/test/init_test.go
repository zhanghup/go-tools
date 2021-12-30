package test_test

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

var engine txorm.IEngine
var db *xorm.Engine

type User struct {
	Id   string `json:"id" xorm:"pk"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func init() {
	e, err := txorm.NewXorm(txorm.Config{
		Uri:    "./data.db",
		Driver: "sqlite3",
		Debug:  true,
	})
	if err != nil {
		tog.Error(err.Error())
		return
	}
	db = e
	engine = txorm.NewEngine(e)

	engine.TemplateFuncWith("users", func(ctx context.Context) string {
		return "select * from user"
	})

	engine.TemplateFuncCtx("corp", func(ctx context.Context) string {
		return "'ceaaeb6d-9f47-4ecb-ab4b-3247091229b7'"
	})

	err = engine.Sync(User{})
	if err != nil {
		tog.Error(err.Error())
		return
	}
	err = engine.SessionAuto().SF("delete from user").Exec()
	if err != nil {
		tog.Error(err.Error())
		return
	}

	for i := 0; i < 10; i++ {
		engine.SessionAuto().Insert(User{
			Id:   tools.IntToStr(i),
			Name: tools.IntToStr(i),
			Age:  i,
		})
	}
}
