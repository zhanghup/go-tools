package influx

import (
	"context"
	_ "embed"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
	"sync"
	"time"
)

//go:embed config-default.yml
var initConfigByte []byte

type Option struct {
	Bucket string `json:"bucket" yaml:"bucket"`
	Org    string `json:"org" yaml:"org"`
	Token  string `json:"token" yaml:"token"`
	Url    string `json:"url" yaml:"url"`
}

type Engine struct {
	opt Option

	client influxdb2.Client
	write  api.WriteAPIBlocking
	query  api.QueryAPI

	once sync.Once
	data tools.IQueue[*write.Point]
}

type IEngine interface {
	Client() influxdb2.Client

	Write(point ...*write.Point)
	WriteWithContext(ctx context.Context, point ...*write.Point) error

	Query(bucket string) QueryString
}

func InitEngine(cfg ...[]byte) IEngine {
	opt := Option{}

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
	return NewEngine(opt)
}

func NewEngine(opt Option) IEngine {
	client := influxdb2.NewClient(opt.Url, opt.Token)

	e := &Engine{
		opt:    opt,
		client: client,
		write:  client.WriteAPIBlocking(opt.Org, opt.Bucket),
		query:  client.QueryAPI(opt.Org),
		data:   tools.NewQueue[*write.Point](),
	}
	return e
}

func (this *Engine) Client() influxdb2.Client {
	return this.client
}

func (this *Engine) Write(point ...*write.Point) {
	this.data.Push(point...)

	this.once.Do(func() {
		go func() {
			for {
				tools.Run(func() {
					for {
						data := this.data.Pop(200)
						err := this.WriteWithContext(context.Background(), data...)
						if err != nil {
							tog.Infof("[influxdb] - 写入数据(%d)[错误1] - %s", len(data), err.Error())
							this.data.LPush(data...)
							time.Sleep(time.Second * 1)
						}

						if time.Now().Second() == 0 {
							tog.Infof("[influxdb] - 剩余数据长度(%d)", this.data.Len())
						}
					}
				}, func(res any) {
					tog.Infof("[influxdb] - 写入数据(%d)[错误3] - %v", this.data.Len(), res)
					time.Sleep(time.Second * 5)
				})
			}
		}()

	})
}
func (this *Engine) WriteWithContext(ctx context.Context, point ...*write.Point) error {
	return this.write.WritePoint(ctx, point...)
}
