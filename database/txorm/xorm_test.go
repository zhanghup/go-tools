package txorm_test

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"regexp"
	"testing"
)

var engine txorm.IEngine

func TestTemplate(t *testing.T) {
	err := engine.SF(`select * from corp where  corp = {{ ctx_corp }}`).With("users").Exec()
	if err != nil {
		t.Error(err)
	}
}

func TestQuery(t *testing.T) {
	m, err := engine.SF(`select * from corp `).Map()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tools.JSONString(m, true))
}

func TestStrTemp(t *testing.T) {
	str := ` 
		fjdsakjfksadjf
		{{ctx_1  }}
		{{ ctx_2  }} {{ cvx_4  }} {{ ctx_5  }}
		{{ ctx_3 "dajskd" }}
	`
	fmt.Println(str)

	r := regexp.MustCompile(`{{\s*ctx_\S+\s*}}`)
	r.FindAllString(str, -1)
	fmt.Println(r.FindAllString(str, -1))
}

func init() {
	e, err := txorm.NewXorm(txorm.Config{
		Uri:    "root:123@tcp(127.0.0.1)/nt?charset=utf8",
		Driver: "mysql",
		Debug:  true,
	})
	if err != nil {
		panic(err)
	}
	engine = txorm.NewEngine(e)

	engine.TemplateFunc("corp", func(n string) string {
		fmt.Println(n)
		return fmt.Sprintf("%s.corp = '{{ .corp }}'", n)
	})

	engine.TemplateFuncCtx("corp", func(ctx context.Context) string {
		return "'12jfkdasljfkla'"
	})

	engine.TemplateFuncWith("users", func(ctx context.Context) string {
		return "select 1"
	})
}
