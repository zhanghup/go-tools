package sys

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestMem(t *testing.T) {

	m, err := mem.VirtualMemory()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tools.JSONString(m, true))

	mm, err := mem.SwapMemory()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tools.JSONString(mm, true))

	mmm,err := mem.SwapDevices()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tools.JSONString(mmm, true))
}
