package test_test

import (
	"testing"
)

func TestExists(t *testing.T) {
	ok, err := engine.Sess().SF("select * from user where 1 = 1").Exists()
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("错误")
	}

	ok, err = engine.Sess().Table("user").SF("1 = 1").Exists()
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("错误")
	}

	ok, err = db.Table("user").Where("1 = 1").Exist()
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("错误")
	}
}

func TestMap(t *testing.T) {
	v, err := engine.Sess().SF("select * from user where 1 = 1").Map()
	if err != nil {
		t.Fatal(err)
	}
	if len(v) != 10 {
		t.Fatal("错误")
	}

	v, err = engine.Sess().Table("user").SF("1 = 1").Map()
	if err != nil {
		t.Fatal(err)
	}
	if len(v) != 10 {
		t.Fatal("错误")
	}

}
