package sys

import (
	"github.com/shirou/gopsutil/cpu"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type CPUInfos struct {
	Cores     int `json:"cores"`      // CPU数量
	CoreNum   int `json:"core_num"`   // 总核心数
	ThreadNum int `json:"thread_num"` // 总线程数

	CoreInfos []CPUCore `json:"core_infos"`
}

type CPUInfo struct {
	Cores     int `json:"cores"`      // CPU数量
	CoreNum   int `json:"core_num"`   // 总核心数
	ThreadNum int `json:"thread_num"` // 总线程数

	CPUCore
}

type CPUCore struct {
	Index int    `json:"index"`
	Model string `json:"model"` // 型号
}

// CPU 比较耗性能
func CPU() (CPUInfo, error) {

	cs, err := CPUs()
	if err != nil {
		return CPUInfo{}, err
	}
	if len(cs.CoreInfos) == 0 {
		return CPUInfo{}, err
	}

	return CPUInfo{
		Cores:     1,
		CoreNum:   cs.CoreNum,
		ThreadNum: cs.ThreadNum,
		CPUCore:   cs.CoreInfos[0],
	}, nil

}

// CPUs 比较耗性能
func CPUs() (CPUInfos, error) {
	info := CPUInfos{}

	// 核心数
	{
		o, err := cpu.Counts(false)
		if err != nil {
			return info, err
		}
		info.CoreNum = o
	}

	// 线程数
	{
		o, err := cpu.Counts(true)
		if err != nil {
			return info, err
		}
		info.ThreadNum = o
	}

	infos, err := cpu.Info()
	if err != nil {
		return info, err
	}
	info.Cores = len(infos)

	for i, o := range infos {
		info.CoreInfos = append(info.CoreInfos, CPUCore{
			Index: i,
			Model: o.ModelName,
		})
	}

	return info, nil
}
