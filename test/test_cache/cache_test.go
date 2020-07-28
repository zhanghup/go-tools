package test_cache

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func randomString() string {
	return fmt.Sprintf("%v", rand.Intn(1000))
}

func TestSetGet(t *testing.T) {
	mp := tools.CacheCreate()
	mp.Set("123", "123", time.Now().Unix()+2)
	data := mp.Get("123")
	fmt.Println(data)
	time.Sleep(time.Second * 3)
	data2 := mp.Get("123")
	fmt.Println(data2)
}

// 并发性能测试
func BenchmarkMapBatch(b *testing.B) {
	// 测试一个对象或者函数在多线程的场景下面是否安全
	mp := tools.CacheCreate()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if rand.Int()%2 == 0 {
				mp.Set(fmt.Sprintf("%v", randomString()), rand.Int(), time.Now().Unix()+1)
			} else {
				mp.Get(fmt.Sprintf("%v", randomString()))
			}
		}
	})
}

func TestPointer(t *testing.T) {
	catch := tools.CacheCreate()
	aa := new(int)
	fmt.Println(reflect.ValueOf(aa).Pointer())
	catch.Set("a", aa)
	bb := catch.Get("a")
	fmt.Println(reflect.ValueOf(bb).Pointer())
}
