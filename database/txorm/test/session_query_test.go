package test_test

import (
	"encoding/json"
	"fmt"
	"github.com/zhanghup/go-tools"
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

func TestMap2(t *testing.T) {
	v, err := engine.Sess().SF("select * from user where 1 = 1 limit 1").Map()
	if err != nil {
		t.Fatal(err)
	}
	if len(v) != 1 {
		t.Fatal("错误")
	}

	fmt.Println(tools.JSONString(v[0]))

	bs, err := json.Marshal(v[0])
	if err != nil {
		t.Fatal(err)
	}

	info := User{}
	err = json.Unmarshal(bs, &info)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tools.JSONString(info))

}

func TestMap(t *testing.T) {
	v, err := engine.Sess().SF("select * from user where 1 = 1").Map()
	if err != nil {
		t.Fatal(err)
	}
	if len(v) != 10 {
		t.Fatal("错误")
	}
	fmt.Println(tools.JSONString(v))

	v, err = engine.Sess().Table("user").SF("1 = 1").Map()
	if err != nil {
		t.Fatal(err)
	}
	if len(v) != 10 {
		t.Fatal("错误")
	}

}
