package tools

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
	"xorm.io/xorm"
)

var __fileEngine *xorm.Engine
var __fileRegex = regexp.MustCompile(`^\w+\d{6}-`)
var ErrFileNotExist = errors.New("[uploader] 读取文件不存在")

func FileInfo(id string) (io.Reader, error) {
	id = strings.Split(id, ".")[0]

	v := __fileRegex.FindAllString(id, 1)
	if len(v) == 0 {
		return nil, ErrFileNotExist
	}
	vv := v[0]
	ym := vv[len(vv)-7 : len(vv)-1]
	ftype := vv[:len(vv)-7]
	path := ""
	if ftype == "stream" {
		path = fmt.Sprintf("./upload/%s/%s/%s", ym[:4], ym[4:], ftype)
	} else {
		path = fmt.Sprintf("./upload/%s/%s/%s/%s.%s", ym[:4], ym[4:], ftype, id, ftype)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

/*
	FileUploadIO 文件上传
	reader: 文件流
	name: 文件名，例如1.jpg
*/
func FileUploadIO(reader io.Reader, filename string) (string, error) {
	now := time.Now()
	id := fmt.Sprintf("%04d%02d-%s", now.Year(), now.Month(), UUID())
	ftype := ""

	// 文件后缀解析
	if len(strings.Split(filename, ".")) > 1 {
		ftype = strings.Split(filename, ".")[1]
	}

	// 资源文件路径定义
	path := ""
	{
		dir := fmt.Sprintf("upload/%d/%02d/%s", now.Year(), now.Month(), ftype)
		if _, err := os.Stat(dir); err != nil {
			if os.IsNotExist(err) {
				err = os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					return "", err
				}
			} else {
				return "", err
			}
		}

		if ftype != "" {
			id = ftype + id
			path = fmt.Sprintf("%s/%s.%s", dir, id, ftype)
		} else {
			ftype = "stream"
			path = fmt.Sprintf("%s/%s", dir, id)
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return "", err
	}

	// 写入文件数据
	_, err = io.Copy(f, reader)
	if err != nil {
		return "", err
	}

	return id, nil
}
