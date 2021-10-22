package tools_test

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
	"time"
)

func TestSyncMutex(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func(n int) {
			tools.Mutex("1").Lock()
			fmt.Println(n, "111111111")
			tools.Mutex("1").Unlock()
		}(i)
	}

	for i := 0; i < 100; i++ {
		go func(n int) {
			tools.Mutex("2").Lock()
			fmt.Println(n, "222222222")
			tools.Mutex("2").Unlock()
		}(i)
	}

	time.Sleep(time.Second)
}
