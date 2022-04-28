package tools

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()

	for i := 0; i < 10; i++ {
		go func() {
			for {
				q.Push(rand.Int())
			}
		}()
	}
	time.Sleep(time.Second * 1)
	fmt.Println(q.Len())
	fmt.Println(q.Pop(10))
}
