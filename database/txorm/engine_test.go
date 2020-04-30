package txorm_test

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
	"testing"
	"time"
)

var db *txorm.Engine

type Bean struct {
	Id      *string `json:"id" xorm:"Varchar(32) pk"`
	Created *int64  `json:"created" xorm:"created Int(14)"`
	Updated *int64  `json:"updated" xorm:"updated  Int(14)"`
	Weight  *int    `json:"weight" xorm:"weight  Int(9)"`
	Status  *int    `json:"status" xorm:"status  Int(1)"`
}

// 数据字典
type Dict struct {
	Bean `xorm:"extends"`

	Code   *string `json:"code" xorm:"unique"`
	Name   *string `json:"name"`
	Remark *string `json:"remark"`
}

func TestSF(t *testing.T) {
	datas := make([]struct {
		Id      string `xorm:"id" json:"id"`
		Account string `json:"account"`
	}, 0)

	err := db.SF(`
		select * from user where account = :account
	`, map[string]interface{}{
		"account": "root",
		"id":      []string{"root", "ss"},
	}).Find(&datas)
	if err != nil {
		panic(err)
	}
	tog.Info(tools.Str.JSONString(datas))
}

func TestSession_Exec(t *testing.T) {
	err := db.SF("update user_token set status = 0 ").Exec()
	if err != nil {
		panic(err)
	}
}

func TestSession_TS(t *testing.T) {
	ctx := context.Background()
	ctx, fn := context.WithCancel(ctx)

	sess := db.NewSession(ctx)
	ctx = sess.Context()
	err := sess.TS(func(sess *txorm.Session) error {
		_, err := sess.Sess.Table("user").Insert(map[string]interface{}{
			"id":     tools.Str.Uid(),
			"status": 1,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	sess = db.NewSession(ctx)
	err = sess.TS(func(sess *txorm.Session) error {
		_, err := sess.Sess.Table("user").Insert(map[string]interface{}{
			"id":     tools.Str.Uid(),
			"status": 1,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	fn()

	time.Sleep(time.Second)
}

func TestPage(t *testing.T) {
	dict := make([]Dict, 0)
	n, err := db.SF("select * from dict").Page(2, 2, true, &dict)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(tools.Str.JSONString(dict, true))
}

func init() {
	e, err := txorm.NewXorm(txorm.Config{
		Driver: "mysql",
		Uri:    "root:123@/test?charset=utf8",
	})
	if err != nil {
		panic(err)
	}
	e.ShowSQL(true)
	db = txorm.NewEngine(e)
}
