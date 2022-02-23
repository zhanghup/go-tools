package test_test

import (
	"context"
	"github.com/zhanghup/go-tools/database/txorm"
	"testing"
)

func TestTs(t *testing.T) {
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

		return err
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewSession(t *testing.T) {
	err := engine.New().Table("user").SF("id = ?", "1").Update(map[string]interface{}{
		"id": "111",
	})
	if err != nil {
		t.Fatal(err)
	}

	err = engine.Sess().Table("user").SF("id = ?", "1").Update(map[string]interface{}{
		"id": "111",
	})
	if err != nil {
		t.Fatal(err)
	}

	err = engine.TS(context.Background(), func(ctx context.Context, sess txorm.ISession) error {
		err := sess.Table("user").SF("id = ?", "1").Update(map[string]interface{}{
			"id": "111",
		})
		if err != nil {
			return err
		}

		err = engine.Sess(ctx).Table("user").SF("id = ?", "1").Update(map[string]interface{}{
			"id": "111",
		})
		if err != nil {
			return err
		}
		return nil
	})
}
