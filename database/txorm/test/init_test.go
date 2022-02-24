package test_test

import (
	"context"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

var engine txorm.IEngine
var db *xorm.Engine

type User struct {
	Id   string  `json:"id" xorm:"pk"`
	Name string  `json:"name" xorm:"index"`
	Age  int     `json:"age"`
	Kind *int    `json:"kind"`
	Send *int    `json:"send"`
	Kd   *string `json:"kd"`
}

func init() {
	e, err := txorm.NewXorm(txorm.Config{
		Uri:    "root:Zhang3611.@tcp(192.168.31.150:23306)/test2?charset=utf8",
		Driver: "mysql",
		Debug:  true,
	})
	if err != nil {
		tog.Error(err.Error())
		return
	}
	db = e
	e.SetMaxIdleConns(100)
	e.SetMaxOpenConns(100)
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
	err = engine.Sess().SF("delete from user").Exec()
	if err != nil {
		tog.Error(err.Error())
		return
	}

	for i := 0; i < 10; i++ {
		err := engine.Sess().Insert(User{
			Id:   tools.IntToStr(i),
			Name: tools.IntToStr(i),
			Age:  i,
			Kind: &i,
		})
		if err != nil {
			panic(err)
		}
	}
}
