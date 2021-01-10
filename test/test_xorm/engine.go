package test_xorm

import (
	"context"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
)

func NewEngine() txorm.IEngine {
	e, err := txorm.NewXorm(txorm.Config{
		Driver: "mysql",
		Uri:    "root:123@/test?charset=utf8",
	})
	if err != nil {
		panic(err)
	}
	e.ShowSQL(true)
	dbs := txorm.NewEngine(e)
	dbs.TemplateFuncAdd("users", func(ctx context.Context) string {
		return tools.Str.Tmp(`
			select id from user
		`, map[string]interface{}{}).String()
	})

	return dbs
}
