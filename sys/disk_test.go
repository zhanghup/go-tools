package sys

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestDiskStat(t *testing.T) {
	info, err := DiskStats()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tools.JSONString(info))
}

func TestDiskInfos(t *testing.T) {
	info, err := DiskInfos()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tools.JSONString(info, true))
}

func BenchmarkDiskStat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DiskStats()
	}
}

func BenchmarkDiskInfos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DiskInfos()
	}
}
