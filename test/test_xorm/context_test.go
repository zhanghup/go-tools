package test_xorm

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()
	go task(ctx)
	time.Sleep(time.Second * 10)
}

func TestContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx,"a","123")
	ctx2 := context.WithValue(ctx,"a","12333")
	fmt.Println(ctx2.Value("a"))
	fmt.Println(ctx.Value("a"))

}

func task(ctx context.Context) {
	ch := make(chan struct{}, 0)
	go func() {
		// 模拟4秒耗时任务
		time.Sleep(time.Second * 4)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}
