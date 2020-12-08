package test_xorm

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"testing"
	"time"
)

func TestSession_TS(t *testing.T) {
	ctx := context.Background()
	ctx,fn := context.WithTimeout(ctx,time.Second)
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

	sess.Close()

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

	sess.Commit()
	sess.Close()
	fmt.Println("11")
	fn()
	fmt.Println("22")
	//sess.ContextClose()
	time.Sleep(time.Second * 1)
}
