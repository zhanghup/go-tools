package tools

import (
	"bytes"
	"strings"
	"text/template"
)

// string template 格式化
type myStringTemplate struct {
	tpl   *template.Template
	str   string
	fns   template.FuncMap
	param interface{}
}

func (this myString) T(str string, param ...interface{}) myStringTemplate {
	tt := template.New(this.Uid())
	fmap := template.FuncMap{
		"title": strings.Title,
	}

	result := myStringTemplate{
		tpl: tt,
		str: str,
		fns: fmap,
	}
	if param != nil && len(param) > 0 {
		result.param = param[0]
	}
	return result
}

func (this myStringTemplate) FuncMap(param map[string]interface{}) myStringTemplate {
	if param == nil {
		return this
	}
	for k, v := range param {
		this.fns[k] = v
	}
	return this
}

func (this myStringTemplate) String() string {
	data := bytes.NewBuffer(nil)
	err := this.tpl.Execute(data, this.param)
	if err != nil {
		return S.Str("模板格式化异常,error:%s", err.Error())
	}
	return data.String()

}
