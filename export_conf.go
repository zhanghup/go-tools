package tools

/*
	配置文件快速读取帮助方法
	依赖rice.go
*/

import (
	"errors"
	"gopkg.in/yaml.v2"
)

func ConfOfByte(dataByte []byte, data any) error {
	exception := func(s string, err error) error {
		return errors.New(StrFmt(`config.yml - %s - err: %s`, s, err.Error()))
	}
	if err := yaml.Unmarshal(dataByte, data); err != nil {
		return exception("yaml 格式化失败", err)
	}
	return nil
}
