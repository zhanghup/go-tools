package txorm_test

import (
	"context"
	"github.com/zhanghup/go-tools/database/txorm"
	"testing"
	//_ "github.com/mattn/go-sqlite3"
)

var engine txorm.IEngine

func TestWithTemplate(t *testing.T) {
	err := engine.Session().SF(`select * from {{ tmp "users" }} as u where u.id = '44bbb8ef-c72f-4f66-a294-d651be5948f4'
	`).Exec()
	if err != nil {
		t.Error(err)
	}
}

func TestSessionContextTemplate(t *testing.T) {
	err := engine.Session().SF(`select * from {{ tmp "users" }} as u where u.id = '44bbb8ef-c72f-4f66-a294-d651be5948f4' 
	and u.corp = {{ ctx "corp" }}
	{{ if .ty }} and u.corp = {{ ctx "corp" }} {{ end }}
	{{ if .t }} and u.corp = ? {{ end }}
	`, map[string]interface{}{
		"ty": true,
		"t":  true,
	}, "ceaaeb6d-9f47-4ecb-ab4b-3247091229b7").Exec()
	if err != nil {
		t.Error(err)
	}
}

func init() {
	e, err := txorm.NewXorm(txorm.Config{
		Uri:    "root:Zhang3611.@tcp(192.168.31.150:23306)/test2?charset=utf8",
		Driver: "mysql",
		Debug:  true,
	})
	if err != nil {
		panic(err)
	}
	engine = txorm.NewEngine(e)

	engine.TemplateFuncWith("users", func(ctx context.Context) string {
		return "select * from user"
	})

	engine.TemplateFuncCtx("corp", func(ctx context.Context) string {
		return "'ceaaeb6d-9f47-4ecb-ab4b-3247091229b7'"
	})
}
