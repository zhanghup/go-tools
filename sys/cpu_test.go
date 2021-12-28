package sys

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestCpus(t *testing.T) {
	o, err := CPUs()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tools.JSONString(o))
}

func TestCpu(t *testing.T) {
	o, err := CPU()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tools.JSONString(o))
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CPUs()
	}
}
