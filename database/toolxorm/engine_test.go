package toolxorm_test

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/toolxorm"
	"testing"
	"time"
)

var db *toolxorm.Engine

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
	fmt.Println(tools.Str.JSONString(datas, true))
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
	err := sess.TS(func(sess *toolxorm.Session) error {
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
	err = sess.TS(func(sess *toolxorm.Session) error {
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

func init() {
	e, err := toolxorm.NewXorm(toolxorm.Config{
		Driver: "mysql",
		Uri:    "root:123@/test?charset=utf8",
	})
	if err != nil {
		panic(err)
	}
	e.ShowSQL(true)
	db = toolxorm.NewEngine(e)
}
