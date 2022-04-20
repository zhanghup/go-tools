package tools

import (
	"fmt"
	"reflect"
)

/*
	RftTypeInfoView，解析对象的结果，

	FullName 表示实例的全名，例如: make([]string,0) => []string
	Name 表示实例名称，例如: make([]string,0) => string
	PkgPath 表示最终引用对象的包路径
	Type 表示最终引用对象的类型，可以反射出一个新的实例

	示例：RftTypeInfo(make([]*RftTypeInfoView,0))
	结果： {"name":"RftTypeInfoView","full_name":"[]*RftTypeInfoView","pkg_path":"github.com/zhanghup/go-tools"}

*/
type RftTypeInfoView struct {
	Name     string       `json:"name"`      // 对象名，例如： make([]string,0) => []string
	FullName string       `json:"full_name"` // 对象名，例如： make([]string,0) => []string
	PkgPath  string       `json:"pkg_path"`  // 对象的包路径，例如： make([]string,0) => string
	Type     reflect.Type `json:"type"`      // 对象类型
}

// RftTypeInfo 反射出对象的属性
func RftTypeInfo(obj any) RftTypeInfoView {
	if obj == nil {
		return RftTypeInfoView{}
	}

	ty := reflect.TypeOf(obj)
	name := ""

	for {
		switch ty.Kind() {
		case reflect.Ptr:
			name += "*"
			ty = ty.Elem()
		case reflect.Slice:
			name += "[]"
			ty = ty.Elem()
		default:
			return RftTypeInfoView{
				FullName: name + ty.Name(),
				Name:     ty.Name(),
				PkgPath:  ty.PkgPath(),
				Type:     ty,
			}
		}
	}
}

// RftInterfaceInfo 反射出一个对象的所有属性和值
func RftInterfaceInfo(obj any, fn func(field string, value any, tag reflect.StructTag) bool) {
	ty := reflect.TypeOf(obj)
	vl := reflect.ValueOf(obj)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
		vl = vl.Elem()
	}

	rftInterfaceInfo(ty, vl, "", fn)
}

// rftInterfaceInfo 反射具体属性并且回调
func rftInterfaceInfo(ty reflect.Type, vl reflect.Value, tag reflect.StructTag, fn func(field string, value any, tag reflect.StructTag) bool) {

	switch ty.Kind() {
	case reflect.Ptr:
		if vl.IsNil() {
			return
		}
		ty = ty.Elem()
		vl = vl.Elem()
		rftInterfaceInfo(ty, vl, tag, fn)
		return

	case reflect.Struct:
		for i := 0; i < vl.NumField(); i++ {
			tf := ty.Field(i)
			v := vl.Field(i)
			t := tf.Type

			if RftIsNil(v) {
				if !fn(tf.Name, nil, tf.Tag) {
					return
				}
			} else {
				if !v.IsZero() {
					if !fn(tf.Name, v.Interface(), tf.Tag) {
						return
					}
					rftInterfaceInfo(t, v, tf.Tag, fn)
				} else {
					if !fn(tf.Name, reflect.New(t).Elem().Interface(), tf.Tag) {
						return
					}
				}
			}
		}
	case reflect.Map:
		for _, o := range vl.MapKeys() {
			v := vl.MapIndex(o)
			field := InterfaceToString(o.Interface())

			if RftIsNil(v) {
				fn(field, nil, "")
			} else {
				targetValue := v.Interface()
				if !fn(field, targetValue, "") {
					return
				}
				rftInterfaceInfo(reflect.TypeOf(targetValue), reflect.ValueOf(targetValue), "", fn)
			}

		}
	}
}

func RftIsNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	}
	return false
}

// InterfaceToString 将基础类型数据转换为string
func InterfaceToString(o any) string {
	switch o.(type) {
	case string:
		return o.(string)
	case *string:
		return *(o.(*string))
	case bool:
		if o.(bool) {
			return "true"
		} else {
			return "false"
		}
	case *bool:
		if *(o.(*bool)) {
			return "true"
		} else {
			return "false"
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, complex64, complex128:
		return fmt.Sprintf("%v", o)
	case *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *complex64, *complex128:
		return fmt.Sprintf("%v", reflect.ValueOf(o).Elem().Interface())
	case float32, float64:
		return fmt.Sprintf("%f", o)
	case *float32, *float64:
		return fmt.Sprintf("%f", reflect.ValueOf(o).Elem().Interface())
	}
	return ""
}
