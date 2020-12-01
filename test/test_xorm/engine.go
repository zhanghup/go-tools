package test_xorm

import "github.com/zhanghup/go-tools/database/txorm"

func NewEngine() *txorm.Engine {
	e, err := txorm.NewXorm(txorm.Config{
		Driver: "mysql",
		Uri:    "root:123@/test?charset=utf8",
	})
	if err != nil {
		panic(err)
	}
	e.ShowSQL(true)
	return txorm.NewEngine(e)
}
