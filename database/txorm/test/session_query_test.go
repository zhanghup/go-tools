package test_test

import (
	"testing"
)

func TestExists(t *testing.T) {
	ok, err := engine.Session().SF("select * from user where 1 = 1").Exists()
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("错误")
	}

	ok, err = engine.Session().Table("user").SF("where 1 = 1").Exists()
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
	v, err := engine.Session().SF("select * from user where 1 = 1").Map()
	if err != nil {
		t.Fatal(err)
	}
	if len(v) != 10 {
		t.Fatal("错误")
	}

	v, err = engine.Session().Table("user").SF("where 1 = 1").Map()
	if err != nil {
		t.Fatal(err)
	}
	if len(v) != 10 {
		t.Fatal("错误")
	}

}
