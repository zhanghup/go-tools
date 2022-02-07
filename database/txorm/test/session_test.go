package test_test

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"testing"
	"time"
)

func TestSessionNew(t *testing.T) {
	ctx := context.Background()
	sess := engine.Session(true)
	fmt.Println(sess.Id())
	ctx = context.WithValue(ctx, txorm.CONTEXT_SESSION, sess)
	ss := engine.Session(true, ctx)
	fmt.Println(ss.Id())
	ss3 := engine.Session(false, ctx)
	fmt.Println(ss3.Id())
	ss2 := engine.Session(true)
	fmt.Println(ss2.Id())
}

func TestSessionInsert(t *testing.T) {
	err := engine.Session(true).Insert(User{
		Id:   tools.UUID(),
		Name: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
}

// TestSessionTx Session中若只是查询，则不开启事务逻辑，若包含操作逻辑则开启
func TestSessionTx(t *testing.T) {
	sess := engine.Session(true)

	t.Run("只查询", func(t *testing.T) {
		users := make([]User, 0)
		err := sess.TS(func(ctx context.Context, sess txorm.ISession) error {
			err := sess.SF(`limit 1`).Find(&users)
			if err != nil {
				return err
			}
			return nil

		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(tools.JSONString(users))
	})

	t.Run("包含操作", func(t *testing.T) {
		users := make([]User, 0)
		err := sess.TS(func(ctx context.Context, sess txorm.ISession) error {
			err := sess.Find(&users)
			if err != nil {
				return err
			}

			err = sess.Insert(User{Id: tools.UUID(), Name: "test"})
			return err
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(tools.JSONString(users))
	})
}

func Benchmark并发插入(b *testing.B) {
	for i := 0; i < 100; i++ {
		go func() {
			sess := engine.Session(false)
			err := sess.TS(func(ctx context.Context, sess txorm.ISession) error {
				err := sess.Insert(User{Id: tools.UUID(), Name: "test"})
				return err
			})
			if err != nil {
				b.Fatal(err)
			}
			sess.Close()
		}()
	}

	time.Sleep(time.Second * 10)
}

func Benchmark并发查询(b *testing.B) {
	for i := 0; i < 100; i++ {
		sess := engine.Session(false)
		go func(n int) {
			nn := tools.IntToStr(n)
			user := make([]User, 0)
			err := sess.SF("select * from user v" + nn).Order("v" + nn + ".name").Find(&user)
			if err != nil {
				panic(err)
			}
		}(i)
	}

	time.Sleep(time.Second * 10)
}
