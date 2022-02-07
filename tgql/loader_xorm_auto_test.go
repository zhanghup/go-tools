package tgql_test

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tgql"
	"testing"
	"time"
)

// TestXormSessionObject
// 测试查询单个map对象并且转换为struct对象
func TestXormSessionObject(t *testing.T) {

	lod := tgql.NewLoader(enginedb)
	sess := engine.Sess()

	ids := []string{"12d07c3d-8133-48ab-b353-124da316b0d7", "14e7d395-5a6a-45e1-bbad-ceb4e8ff64aa", "4803d264-2b5a-4f70-8c05-be5712a694f3"}
	for i := range []int{0, 1, 2} {
		go func(ii int) {
			info := Dict{}
			ok, err := lod.LoadXormSessObject(sess, `select * from dict where dict.id in :keys`, "id").Load(ids[ii], &info)
			if err != nil || !ok {
				t.Fatal(err)
			} else {
				fmt.Println(tools.JSONString(info, true))
			}
		}(i)
	}

	time.Sleep(time.Second)
}

// TestXormSessionSlice
// 测试查询多个map对象并且转换为[]struct数组
func TestXormSessionSlice(t *testing.T) {

	lod := tgql.NewLoader(enginedb)
	sess := engine.Sess()

	ids := []string{"12d07c3d-8133-48ab-b353-124da316b0d7", "14e7d395-5a6a-45e1-bbad-ceb4e8ff64aa", "4803d264-2b5a-4f70-8c05-be5712a694f3"}
	for i := range []int{0, 1, 2} {
		go func(ii int) {
			info := make([]DictItem, 0)
			ok, err := lod.LoadXormSessSlice(sess, `select * from dict_item where code in :keys`, "code").Load(ids[ii], &info)
			if err != nil || !ok {
				t.Fatal(err)
			} else {
				fmt.Println(tools.JSONString(info, true))
			}
		}(i)
	}

	time.Sleep(time.Second)
}

// TestXormSessionObjectAuto  智能SQL转换
// 测试查询单个map对象并且转换为struct对象
func TestXormSessionObjectAuto(t *testing.T) {

	lod := tgql.NewLoader(enginedb)
	sess := engine.Sess()

	ids := []string{"12d07c3d-8133-48ab-b353-124da316b0d7", "14e7d395-5a6a-45e1-bbad-ceb4e8ff64aa", "4803d264-2b5a-4f70-8c05-be5712a694f3"}
	for i := range []int{0, 1, 2} {
		go func(ii int) {
			info := Dict{}
			ok, err := lod.LoadXormSessObject(sess, `dict`, "id").Load(ids[ii], &info)
			if err != nil || !ok {
				t.Fatal(err)
			} else {
				fmt.Println(tools.JSONString(info, true))
			}
		}(i)
	}

	time.Sleep(time.Second)
}

// TestXormSessionSliceAuto 智能SQL转换
// 测试查询多个map对象并且转换为[]struct数组
func TestXormSessionSliceAuto(t *testing.T) {

	lod := tgql.NewLoader(enginedb)
	sess := engine.Sess()

	ids := []string{"12d07c3d-8133-48ab-b353-124da316b0d7", "14e7d395-5a6a-45e1-bbad-ceb4e8ff64aa", "4803d264-2b5a-4f70-8c05-be5712a694f3"}
	for i := range []int{0, 1, 2} {
		go func(ii int) {
			info := make([]DictItem, 0)
			ok, err := lod.LoadXormSessSlice(sess, `dict_item`, "code").Load(ids[ii], &info)
			if err != nil || !ok {
				t.Fatal(err)
			} else {
				fmt.Println(tools.JSONString(info, true))
			}
		}(i)
	}

	time.Sleep(time.Second)
}
