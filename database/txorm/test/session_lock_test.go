package test_test

import (
	"context"
	"github.com/zhanghup/go-tools/database/txorm"
	"testing"
	"time"
)

func TestSessionLock1(t *testing.T) {
	go func() {
		err := engine.TS(context.Background(), func(ctx context.Context, sess txorm.ISession) error {
			err := sess.Table("user").SF("id = ?", "1").Update(map[string]interface{}{
				"id": "1",
			})
			if err != nil {
				return err
			}
			time.Sleep(time.Second * 5)

			return nil
		})
		if err != nil {
			t.Fatal(err)
		}

	}()
	time.Sleep(time.Second)
	//err := engine.SessionAuto().Table("user").SF("id = ?", "1").Update(map[string]interface{}{
	//	"age": 111,
	//})
	//if err != nil {
	//	t.Fatal(err)
	//}

	{
		err := engine.TS(context.Background(), func(ctx context.Context, sess txorm.ISession) error {
			err := sess.Table("user").SF("id = ?", "1").Update(map[string]interface{}{
				"id": "1",
			})
			if err != nil {
				return err
			}
			time.Sleep(time.Second * 5)

			return nil
		})

		if err != nil {
			t.Fatal(err)
		}
	}

	time.Sleep(4 * time.Second)
}

func TestSessionLock2(t *testing.T) {

	{
		err := engine.TS(context.Background(), func(ctx context.Context, sess txorm.ISession) error {
			err := sess.Table("user").SF("id = ?", "1").Update(map[string]interface{}{
				"id": "111",
			})
			if err != nil {
				return err
			}

			err = engine.TS(ctx, func(ctx context.Context, sess txorm.ISession) error {
				err := sess.Table("user").SF("id = ?", "1").Update(map[string]interface{}{
					"id": "111",
				})
				if err != nil {
					return err
				}
				return nil
			})

			if err != nil {
				t.Fatal(err)
			}

			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
	}

}

func TestSessionLock3(t *testing.T) {

	{
		sess := db.NewSession()
		sess.Begin()
		_, err := sess.Table(User{}).Where("id = ?", "1").Update(map[string]interface{}{
			"id": 111,
		})
		if err != nil {
			panic(err)
		}
		_, err = db.Table(User{}).Where("id = ?", "1").Update(map[string]interface{}{
			"id": 111,
		})
		sess.Commit()

	}

	{
		err := engine.TS(context.Background(), func(ctx context.Context, sess txorm.ISession) error {
			err := sess.Table("user").SF("id = ?", "1").Update(map[string]interface{}{
				"id": 111,
			})
			if err != nil {
				return err
			}

			err = engine.Sess(ctx).Table("user").SF("id = ?", "1").Update(map[string]interface{}{
				"id": 111,
			})
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
	}

}

func TestSessionLock4(t *testing.T) {

	{
		err := engine.TS(context.Background(), func(ctx context.Context, sess txorm.ISession) error {
			err := sess.Table("user").SF("id = ?", "1").Update(map[string]interface{}{
				"id": 111,
			})
			if err != nil {
				return err
			}

			user := User{}
			_, err = engine.Sess().Table("user").SF("id = ?", "1").Get(&user)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
	}

}
