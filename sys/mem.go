package sys

import (
	"github.com/shirou/gopsutil/mem"
)

type MemInfo struct {
	Total       SizeNumber `json:"total"`
	Used        SizeNumber `json:"used"`
	Free        SizeNumber `json:"free"`
	UsedPercent float64    `json:"used_percent"`
}

func Mem() (*MemInfo, error) {

	mm, err := mem.SwapMemory()
	if err != nil {
		return nil, err
	}

	return &MemInfo{
		Total:       SizeNumber(mm.Total),
		Used:        SizeNumber(mm.Used),
		Free:        SizeNumber(mm.Free),
		UsedPercent: mm.UsedPercent,
	}, nil
}
