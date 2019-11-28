package tools

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
)

type htp struct {
	header map[string]string
}

func Http() htp {
	return htp{}
}
func (this htp) Header(header map[string]string) htp {
	this.header = header
	return this
}

// 支持 text/template 字符串格式化
func (this htp) GetF(url string, param map[string]interface{}) (*http.Response, error) {
	str, err := Str().Template(url, param, nil)
	if err != nil {
		return nil, err
	}
	return this.Get(str)
}

func (this htp) GetI(url string, param map[string]interface{}, result interface{}) error {
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

func (this htp) Get(url string) (*http.Response, error) {
	return this.Request(http.MethodGet, url, nil)
}

func (this htp) PostI(url string,param map[string]interface{},result interface{}) error{
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

func (this htp) Post(url string, param map[string]interface{}) (*http.Response, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	return this.Request(http.MethodPost, url, bytes.NewBuffer(data))
}

func (this htp) Request(method, url string, body io.Reader) (*http.Response, error) {
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
