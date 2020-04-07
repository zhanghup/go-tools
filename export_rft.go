package tools

import (
	"reflect"
)

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

func (this myrft) DeepSet(o interface{}, fn func(t reflect.Type, v reflect.Value, tf reflect.StructField) bool) {
	ty := reflect.TypeOf(o)
	vl := reflect.ValueOf(o)
	if ty.Kind() != reflect.Ptr {
		panic("DeepSet 输入必须为指针")
	}
	ty = ty.Elem()
	vl = vl.Elem()
	this.deepSet(ty, vl, reflect.StructField{}, fn)
}

func (this myrft) deepSet(ty reflect.Type, vl reflect.Value, tf reflect.StructField, fn func(t reflect.Type, v reflect.Value, tf reflect.StructField) bool) {
	switch ty.Kind() {
	case reflect.Ptr:
		if !vl.CanSet() {
			return
		}
		if vl.Pointer() == 0 {
			if !fn(ty, vl, tf) {
				return
			}
		}
		ty = ty.Elem()
		vl = vl.Elem()
		this.deepSet(ty, vl, tf, fn)
	case reflect.Struct:
		for i := 0; i < vl.NumField(); i++ {
			tf := ty.Field(i)
			v := vl.Field(i)
			t := tf.Type
			this.deepSet(t, v, tf, fn)
		}
	default:
		if !vl.CanSet() {
			return
		}
		fn(ty, vl, tf)
	}

}
