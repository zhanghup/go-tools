package test_config

import (
	"fmt"
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-tools"
	"testing"
)

type Mysql struct {
	Host string `yaml:"host"`
}
type Redis struct {
	Database int `yaml:"database"`
}
type Mongo struct {
	Host string `yaml:"host"`
}
type conf struct {
	// 使用匿名字段， 这样 config 会拥有所有的 字段
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
	Mongo Mongo `yaml:"mongo"`
}

func TestConfig(t *testing.T) {
	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	cfg := conf{}
	err = tools.Conf(box, &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(tools.S.JSONString(cfg, true))
}
