package sys

import (
	"github.com/shirou/gopsutil/disk"
)

type DiskStat struct {
	Path        string  `json:"path"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

type DiskInfo struct {
	Device     string `json:"device"`
	MountPoint string `json:"mount_point"`
	Fstype     string `json:"fstype"`
	Opts       string `json:"opts"`

	Stat *DiskStat `json:"stat"`
}

// DiskInfos 比较耗性能
func DiskInfos() ([]DiskInfo, error) {
	ds, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	infos := make([]DiskInfo, 0)
	for _, o := range ds {
		info := DiskInfo{
			Device:     o.Device,
			MountPoint: o.Mountpoint,
			Fstype:     o.Fstype,
			Opts:       o.Opts,
		}

		d, err := DiskStats(o.Mountpoint)
		if err != nil {
			return nil, err
		}
		info.Stat = d

		infos = append(infos, info)
	}
	return infos, nil
}

// DiskStats 比较耗性能
func DiskStats(path ...string) (*DiskStat, error) {
	root := "/"
	if len(path) > 0 {
		root = path[0]
	}

	dk, err := disk.Usage(root)
	if err != nil {
		return nil, nil
	}

	info := &DiskStat{
		Path:        dk.Path,
		Total:       dk.Total,
		Free:        dk.Free,
		Used:        dk.Used,
		UsedPercent: dk.UsedPercent,
	}

	return info, nil
}
