package tools

/*
	配置文件快速读取帮助方法
	依赖rice.go
*/

import (
	"embed"
	"errors"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func Conf(box *rice.Box, data interface{}) error {
	err := PtrCheck(data)
	if err != nil {
		return err
	}

	exception := func(s string, err error) error {
		return errors.New(StrFmt(`config.yml - %s - err: %s`, s, err.Error()))
	}

	f, err := os.Open(fmt.Sprintf("%s/%s", box.Name(), "config.yml"))
	datas := make([]byte, 0)
	if err != nil {
		f, err := box.Open("config.yml")
		if err != nil {
			return exception("[1] 配置文件文件打开失败", err)
		}
		datas, err = ioutil.ReadAll(f)
		if err != nil {
			return exception("[2] 配置文件文件读取失败", err)
		}
	} else {
		datas, err = ioutil.ReadAll(f)
		if err != nil {
			return exception("[3] 配置文件文件读取失败", err)
		}
	}

	return ConfOfByte(datas, data)
}

func ConfOfEmbed(localPath string, ebd embed.FS, data interface{}) error {
	err := PtrCheck(data)
	if err != nil {
		return err
	}

	exception := func(s string, err error) error {
		return errors.New(StrFmt(`config.yml - %s - err: %s`, s, err.Error()))
	}

	f, err := os.Open(localPath)
	datas := make([]byte, 0)
	if err != nil {
		datas, err = ebd.ReadFile("config.yml")
		if err != nil {
			return exception("[1] 配置文件文件读取失败", err)
		}
	} else {
		datas, err = ioutil.ReadAll(f)
		if err != nil {
			return exception("[2] 配置文件文件读取失败", err)
		}
	}

	return ConfOfByte(datas, data)
}

func ConfOfByte(dataByte []byte, data interface{}) error {
	exception := func(s string, err error) error {
		return errors.New(StrFmt(`config.yml - %s - err: %s`, s, err.Error()))
	}
	if err := yaml.Unmarshal(dataByte, data); err != nil {
		return exception("yaml 格式化失败", err)
	}
	return nil
}
