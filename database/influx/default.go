package influx

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/zhanghup/go-tools"
)

var defaultEngine IEngine

func InitDefault(cfg ...[]byte) IEngine {
	opt := struct {
		Influxdb Option `json:"influxdb"`
	}{}

	err := tools.ConfOfByte(initConfigByte, &opt)
	if err != nil {
		panic(err)
	}

	for _, s := range cfg {
		err = tools.ConfOfByte(s, &opt)
		if err != nil {
			panic(err)
		}
	}
	defaultEngine = NewEngine(opt.Influxdb)
	return defaultEngine
}

func Write(point ...*write.Point) {
	if defaultEngine == nil {
		panic("【influxdb】 - defaultEngine 未初始化[1]")
	}
	defaultEngine.Write(point...)
}

func WriteWithContext(ctx context.Context, point ...*write.Point) error {
	if defaultEngine == nil {
		panic("【influxdb】 - defaultEngine 未初始化[2]")
	}
	return defaultEngine.WriteWithContext(ctx, point...)
}

func Query(bucket string) QueryString {
	if defaultEngine == nil {
		panic("【influxdb】 - defaultEngine 未初始化[3]")
	}
	return defaultEngine.Query(bucket)
}

func Len() int {
	if defaultEngine == nil {
		panic("【influxdb】 - defaultEngine 未初始化[3]")
	}
	return defaultEngine.(*Engine).data.Len()
}
