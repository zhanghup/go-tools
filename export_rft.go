package tools

import "reflect"

type myrft struct{}

var Rft = myrft{}

// 去除指针，获取真实的类型
func (this myrft) RealValue(o interface{}) interface{} {
	v := this.realValue(reflect.ValueOf(o))
	return v.Interface()
}
func (this myrft) realValue(o reflect.Value) reflect.Value {
	if o.Kind() == reflect.Ptr {
		o = o.Elem()
		return this.realValue(o)
	}
	return o
}
