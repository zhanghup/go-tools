package config

import (
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-tools/str"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func exception(s string, err error) {
	panic(str.S(`config.yml - %s - err: %s`, s, err.Error()))
}

type cfg struct {
	file  *rice.File
	datas []byte
}

var Config *cfg

func Init(box *rice.Box) *cfg {
	if Config != nil {
		return Config
	}
	Config = &cfg{}
	f, err := box.Open("config.yml")
	if err != nil {
		exception("配置文件文件打开失败", err)
	}
	Config.file = f
	Config.datas, err = ioutil.ReadAll(f)
	if err != nil {
		exception("配置文件文件读取失败", err)
	}
	return Config
}

func (this *cfg) Unmarshal(data interface{}) {
	if err := yaml.Unmarshal(this.datas, data); err != nil {
		exception("yaml 格式化失败", err)
	}
}
