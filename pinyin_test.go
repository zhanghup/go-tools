package tools

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(Pinyin("你好"))
	fmt.Println(PINYIN("你好"))
	fmt.Println(Py("你好"))
	fmt.Println(PY("你好"))
}

// 性能测试
func BenchmarkPinyin1(b *testing.B) {
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		Pinyin("你好")
		PINYIN("你好")
		Py("你好")
		PY("你好")
	}
}

// 并发性能测试
func BenchmarkPinyin2(b *testing.B) {
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Pinyin("你好")
			PINYIN("你好")
			Py("你好")
			PY("你好")
		}
	})
}
