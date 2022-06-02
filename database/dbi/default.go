package dbi

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/influx"
)

var defaultEngine influx.IEngine

func InitDefault(cfg ...[]byte) influx.IEngine {
	opt := struct {
		Influxdb influx.Option `json:"influxdb"`
	}{}

	for _, s := range cfg {
		err := tools.ConfOfByte(s, &opt)
		if err != nil {
			panic(err)
		}
	}
	defaultEngine = influx.NewEngine(opt.Influxdb)
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

//func Query() QueryString {
//	if defaultEngine == nil {
//		panic("【influxdb】 - defaultEngine 未初始化[3]")
//	}
//	return defaultEngine.Query()
//}
//
//func Len() int {
//	if defaultEngine == nil {
//		panic("【influxdb】 - defaultEngine 未初始化[3]")
//	}
//	return defaultEngine.(*influx.Engine).data.Len()
//}
