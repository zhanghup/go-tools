package rft

import "reflect"

// 去除指针，获取真实的类型
func RealValue(o interface{}) interface{} {
	return realValue(reflect.ValueOf(o))
}
func realValue(o reflect.Value) reflect.Value {
	if o.Kind() == reflect.Ptr {
		o = o.Elem()
		return realValue(o)
	}
	return o
}
