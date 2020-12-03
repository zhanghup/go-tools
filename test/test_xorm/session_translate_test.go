package test_xorm

import (
	"context"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"testing"
	"time"
)

func TestSession_TS(t *testing.T) {
	ctx := context.Background()
	ctx, fn := context.WithCancel(ctx)

	sess := NewEngine().NewSession(ctx)
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

	users := make([]struct{Id string `xorm:"id"`},0)
	err = sess.SF("select * from user").Find(&users)
	if err != nil{
		panic(err)
	}

	sess = NewEngine().NewSession(ctx)
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