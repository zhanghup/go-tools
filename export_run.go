package tools

import (
	"context"
	"time"
)

func Run(fn func(), callback ...func(res interface{})) {
	defer func() {
		if r := recover(); r != nil {
			if len(callback) > 0 {
				callback[0](r)
			}
		}
	}()
	fn()
}

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
