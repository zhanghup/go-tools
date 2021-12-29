package sys

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestMem(t *testing.T) {
	mem, err := Mem()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tools.JSONString(mem))
	fmt.Println(mem.Total.String())
	fmt.Println(mem.Used.String())
	fmt.Println(mem.Free.String())
	fmt.Println(1 << 10)
}
