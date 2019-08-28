package tools

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func randomString() string {
	return fmt.Sprintf("%v", rand.Intn(1000))
}

// 并发性能测试
func BenchmarkMapBatch(b *testing.B) {
	// 测试一个对象或者函数在多线程的场景下面是否安全
	mp := NewCache()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if rand.Int()%2 == 0 {
				mp.Set2(fmt.Sprintf("%v", randomString()), rand.Int(), time.Now().Unix()+1)
			} else {
				mp.Get(fmt.Sprintf("%v", randomString()))
			}
		}
	})
}

