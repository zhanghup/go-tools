package tools

/*
	配置文件快速读取帮助方法
	依赖rice.go
*/

import (
	"errors"
	rice "github.com/giter/go.rice"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Conf(box *rice.Box, data interface{}) error {
	err := Ptr.Check(data)
	if err != nil {
		return err
	}

	exception := func(s string, err error) error {
		return errors.New(S.Str(`config.yml - %s - err: %s`, s, err.Error()))
	}

	f, err := box.Open("config.yml")
	if err != nil {
		return exception("配置文件文件打开失败", err)
	}
	datas, err := ioutil.ReadAll(f)
	if err != nil {
		return exception("配置文件文件读取失败", err)
	}
	if err := yaml.Unmarshal(datas, data); err != nil {
		return exception("yaml 格式化失败", err)
	}
	return nil
}
