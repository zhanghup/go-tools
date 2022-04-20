package tools

import (
	"context"
	"time"
)

func RunWithError(fn func() error) error {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	return fn()
}

func Run(fn func(), callback ...func(res any)) {
	defer func() {
		if r := recover(); r != nil {
			if len(callback) > 0 {
				callback[0](r)
			}
		}
	}()
	fn()
}

// RunWithContext 循环执行任务，直到context关闭
func RunWithContext(interval time.Duration, ctx context.Context, fn func()) {
	flag := false
	go func() {
		<-ctx.Done()
		flag = true
	}()
	task := time.NewTicker(interval)
	for {
		select {
		case <-task.C:
			go Run(fn)
		}
		if flag {
			break
		}
	}
	task.Stop()
}
