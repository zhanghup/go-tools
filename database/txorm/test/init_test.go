package test_test

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog/tmp"
	"time"
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

	//Bool2    *bool      `json:"bool2"`
	Byte     byte       `json:"byte"`
	Byte2    *byte      `json:"byte2"`
	Bytes    []byte     `json:"bytes"`
	Float32  float32    `json:"float_32"`
	Flat322  *float32   `json:"flat_322"`
	Float64  float64    `json:"float_64"`
	Float642 *float64   `json:"float_642"`
	Int64    int64      `json:"int_64"`
	Int642   *int64     `json:"int_642"`
	Int      int        `json:"int"`
	Int2     *int       `json:"int_2"`
	Time     *time.Time `json:"time"`
	Time2    time.Time  `json:"time2"`
}

func init() {
	e, err := txorm.NewXorm(txorm.Config{
		//Uri: "root:Zhang3611.@tcp(192.168.31.150:23306)/test2?charset=utf8",
		Uri:    "root:123@tcp(127.0.0.1)/test2?charset=utf8",
		Driver: "mysql",
		//Uri:    "./data.db",
		//Driver: "sqlite3",
		Debug: true,
	})
	if err != nil {
		tmp.Error(err.Error())
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

	e.DropTables(User{})
	err = engine.Sync(User{})
	if err != nil {
		tmp.Error(err.Error())
		return
	}

	for i := 0; i < 10; i++ {
		u := User{
			Id:    tools.IntToStr(i),
			Name:  tools.IntToStr(i),
			Age:   i,
			Kind:  &i,
			Time2: time.Now(),
		}
		t := time.Now()
		if i%2 == 0 {
			u.Time = &t
		}
		if i%3 == 0 {
			u.Int64 = time.Now().Unix()
			//u.Bool = true
			vv := make([]byte, 10, 16)
			vv[0] = 49
			vv[1] = 54
			vv[2] = 52
			vv[3] = 53
			vv[4] = 56
			vv[5] = 48
			vv[6] = 50
			vv[7] = 51
			vv[8] = 48
			vv[9] = 48
			u.Bytes = vv
		}
		err := engine.Sess().Insert(u)
		if err != nil {
			panic(err)
		}
	}
}
