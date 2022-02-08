package tgql_test

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgql"
	"testing"
	"time"
	"xorm.io/xorm"
)

func TestLoader(t *testing.T) {
	lod := tgql.NewLoader(nil)

	for i := 0; i < 10; i++ {
		go func(o int) {
			var res []string
			lod.LoadObject("2", func(keys []string) (map[string]interface{}, error) {
				return map[string]interface{}{
					"1": []string{"11111111"},
					"2": []string{"222222222222"},
					"3": []string{"333333333"},
					"4": []string{"4444444444"},
					"5": []string{"5555555"},
				}, nil
			}).Load(tools.IntToStr(o), &res)
			fmt.Println(res)
		}(i)
	}

	time.Sleep(time.Second)
}

func TestObject(t *testing.T) {
	loder := tgql.NewObjectLoader(func(keys []string) (map[string]interface{}, error) {
		fmt.Println(keys)
		return nil, nil
	})
	for i := 0; i < 600; i++ {
		go func(o int) {
			_, err := loder.Load(tools.IntToStr(o), nil)
			if err != nil {
				panic(err)
			}
		}(i)
	}

	time.Sleep(time.Second * 3)
}

func TestLoaderXorm(t *testing.T) {
	lod := tgql.NewLoader(enginedb)

	type Account struct {
		Id   string `json:"id"`
		Corp string `json:"corp"`
	}

	ids := []string{"44bbb8ef-c72f-4f66-a294-d651be5948f4", "44bbb8ef-c72f-4f66-a294-d651be5948f4", ""}

	for i := 0; i < 3; i++ {
		go func(ii int) {
			fmt.Println(ii)
			oo := Account{}
			ok, err := lod.LoadXormCtx(nil, make([]Account, 0), `
				select * from account where id in :keys and corp =  {{ ctx "corp"}}
				`, func(tempData interface{}) map[string]interface{} {
				data := tempData.([]Account)
				result := map[string]interface{}{}
				for i, o := range data {
					result[o.Id] = data[i]
				}
				return result
			}).Load(ids[ii], &oo)
			if err != nil || !ok {
				fmt.Printf("没有获取到数据 %s \n", tools.IntToStr(ii))
			} else {
				fmt.Println(tools.JSONString(oo))
			}

		}(i)
	}
	time.Sleep(time.Second)
}

var enginedb *xorm.Engine
var engine txorm.IEngine

func init() {
	e, err := txorm.NewXorm(txorm.Config{
		Uri:    "root:Zhang3611.@tcp(192.168.31.150:23306)/test2?charset=utf8",
		Driver: "mysql",
		Debug:  true,
	})
	if err != nil {
		panic(err)
	}
	enginedb = e
	engine = txorm.NewEngine(e)

	engine.TemplateFunc("corp", func(n string) string {
		if n == "" {
			return fmt.Sprintf("corp = {{ ctx_corp }}")
		}
		return fmt.Sprintf("%s.corp = {{ ctx_corp }}", n)
	})

	engine.TemplateFuncCtx("corp", func(ctx context.Context) string {
		return "'0000'"
	})

	engine.TemplateFuncWith("users", func(ctx context.Context) string {
		return "select 1"
	})
}
