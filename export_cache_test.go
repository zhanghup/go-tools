package tools_test

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkCacheSet(b *testing.B) {
	cache := tools.NewCache[int]()

	for n := 0; n < b.N; n++ {
		cache.Set(tools.IntToStr(n), n)
	}
}

func BenchmarkCacheGet(b *testing.B) {
	cache := tools.NewCache[int]()

	for n := 0; n < b.N; n++ {
		cache.Get(tools.IntToStr(n))
	}
}

func TestCacheBatch(t *testing.T) {
	cache := tools.NewCache[int]()

	go tools.Run(func() {
		for {
			v := rand.Int()
			cache.Set(tools.IntToStr(v%100000), v)
			if v%3 == 0 {
				cache.Set(tools.IntToStr(v%100000), v)
			} else if v%3 == 1 {
				cache.Get(tools.IntToStr(v % 100000))
			} else {
				cache.Delete(tools.IntToStr(v % 100000))
			}
		}
	})

	time.Sleep(time.Second * 5)

	fmt.Println(cache.Len())
}
