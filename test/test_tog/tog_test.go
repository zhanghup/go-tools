package test_tog

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
	"net/http"
	"testing"
	"time"
)

func TestMyLogger(t *testing.T) {
	go func() {
		err := tgin.NewGin(tgin.Config{Port: "8899"}, func(g *gin.Engine) error {
			g.GET("/test", func(ctx *gin.Context) {
				ctx.String(200, "hello world")
			})
			return nil
		})
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		for {
			http.Get("http://127.0.0.1:8899/test")
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		e, err := txorm.NewXorm(txorm.Config{
			Driver: "mysql",
			Uri:    "root:123@/test?charset=utf8",
			Debug:  true,
		})
		if err != nil {
			panic(err)
		}
		db := txorm.NewEngine(e)
		for {
			user := make([]struct {
				Id      *string `json:"id" xorm:"Varchar(128) pk"`
				Created *int64  `json:"created" xorm:"created Int(14)"`
				Updated *int64  `json:"updated" xorm:"updated  Int(14)"`
				Weight  *int    `json:"weight" xorm:"weight  Int(9)"`
				Status  *string `json:"status" xorm:"status  Int(1)"`
			},0)
			err := db.SF(`select * from user`).Find(&user)
			if err != nil{
				panic(err)
			}
			time.Sleep(time.Millisecond)
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}
