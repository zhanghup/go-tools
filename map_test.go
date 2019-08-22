package tools

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func randomString() string {
	return fmt.Sprintf("%v", rand.Intn(100000000))
}

func TestMap(t *testing.T) {
	mp := NewMap()
	for i := 0; i < 10; i++ {
		go func() {
			for {
				if rand.Int()%2 == 0 {
					mp.Set2(fmt.Sprintf("%v", randomString()), rand.Int(), time.Now().Unix()+1)
				} else {
					mp.Get(fmt.Sprintf("%v", randomString()))
				}
			}
		}()
	}

	time.Sleep(time.Minute * 2)
}
