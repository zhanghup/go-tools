package tools

/*
	http快速请求帮助方法
*/

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
)

type myhttp struct {
	header map[string]string
}

func H() myhttp {
	return myhttp{}
}
func (this myhttp) Header(header map[string]string) myhttp {
	this.header = header
	return this
}

// 支持 text/template 字符串格式化
func (this myhttp) GetF(url string, param map[string]interface{}) (*http.Response, error) {
	return this.Get(S.T(url, param).String())
}

func (this myhttp) GetI(url string, param map[string]interface{}, result interface{}) error {
	ty := reflect.TypeOf(result)
	if ty.Kind() != reflect.Ptr {
		return errors.New("传入的result参数必须未指针值")
	}

	res, err := this.GetF(url, param)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

func (this myhttp) Get(url string) (*http.Response, error) {
	return this.Request(http.MethodGet, url, nil)
}

func (this myhttp) PostI(url string, param interface{}, result interface{}) error {
	ty := reflect.TypeOf(result)
	if ty.Kind() != reflect.Ptr {
		return errors.New("传入的result参数必须未指针值")
	}

	res, err := this.Post(url, param)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

func (this myhttp) Post(url string, param interface{}) (*http.Response, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	return this.Request(http.MethodPost, url, bytes.NewBuffer(data))
}

func (this myhttp) Request(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil
	}

	if this.header != nil {
		for k, v := range this.header {
			req.Header.Add(k, v)
		}
	}

	cli := http.Client{}
	return cli.Do(req)
}
