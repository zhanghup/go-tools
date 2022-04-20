package buntdb

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"reflect"
	"testing"
)

func TestBuntdb(t *testing.T) {
	e, err := NewEngine(Option{
		Path: ":memory:",
	})
	if err != nil {
		panic(err)
	}

	err = e.IndexCreate("key", "key:*A", "string")
	if err != nil {
		panic(err)
	}
	err = e.IndexJsonCreate("json", "key:*A", "age")
	if err != nil {
		panic(err)
	}

	err = e.Ts(func(sess ISession) error {
		e.Get("123")

		for i := 0; i < 1000; i++ {
			_, _, _ = sess.Set(fmt.Sprintf("key:%dA", i+1000), fmt.Sprintf(`{"age":%d,"type":"%d"}`, i+1000, i%10))
		}

		for i := 0; i < 1000; i++ {
			_, _, _ = sess.Set(fmt.Sprintf("key:%dB", i+1000), fmt.Sprintf(`{"age":%d,"type":"%d"}`, i+1000, i%10))
		}

		err = sess.List("key", ListParam{}, func(key, value string) bool {
			return true
		})
		if err != nil {
			return err
		}

		result := make([]struct {
			Age  int    `json:"age"`
			Type string `json:"type"`
		}, 0)

		err = sess.ListJson("json", ListJsonParam{Index: 1, Size: 4}, &result)
		if err != nil {
			return err
		}

		fmt.Println(tools.JSONString(result))

		result = make([]struct {
			Age  int    `json:"age"`
			Type string `json:"type"`
		}, 0)

		err = sess.ListJson("", ListJsonParam{Index: 1, Size: 4, ListParam: ListParam{
			Query: `{"age":1002,"type":"2"}`,
		}}, &result)
		if err != nil {
			return err
		}

		fmt.Println(tools.JSONString(result))

		return nil
	})
	if err != nil {
		panic(err)
	}
}

func TestReflectSlice(t *testing.T) {
	var v *[]string
	r := reflect.TypeOf(v)
	fmt.Println(r.String())
	r = r.Elem()
	fmt.Println(r.String())
	r = r.Elem()
	fmt.Println(r.String())
}
