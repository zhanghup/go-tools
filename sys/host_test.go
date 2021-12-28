package sys

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestHost(t *testing.T) {
	info, err := host.Info()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tools.JSONString(info, true))
}
