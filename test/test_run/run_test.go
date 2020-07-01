package test_run

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
	"time"
)

func TestRun(t *testing.T) {

}

func TestRunWithContext(t *testing.T) {
	ctx, c := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 10)
		c()
	}()
	tools.RunWithContext(time.Second, ctx, func() {
		fmt.Println("............")
	})
}
