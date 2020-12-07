package test_xorm

import (
	"context"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"testing"
	"time"
)

func TestSession_TS(t *testing.T) {
	ctx,cancel := context.WithCancel(context.Background())
	sess := NewEngine().NewSession(ctx)
	err := sess.TS(func(sess txorm.ISession) error {
		_, err := sess.Session().Table("user").Insert(map[string]interface{}{
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

	users := make([]struct {
		Id string `xorm:"id"`
	}, 0)
	err = sess.SF("select * from user").Find(&users)
	if err != nil {
		panic(err)
	}

	sess = NewEngine().NewSession(sess.Context())
	err = sess.TS(func(sess txorm.ISession) error {
		_, err := sess.Session().Table("user").Insert(map[string]interface{}{
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

	//sess.ContextClose()
	cancel()
	time.Sleep(time.Second * 1)
}
