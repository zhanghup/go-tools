package tools

import (
	"reflect"
)

func RftStructDeep(obj interface{}, fn func(t1 reflect.Type, v1 reflect.Value, tg reflect.StructTag, fieldName string) bool, uncheck ...bool) {
	if fn == nil || obj == nil {
		panic("参数不能为空")
	}
	ty := reflect.TypeOf(obj)
	vl := reflect.ValueOf(obj)

	uck := false
	if len(uncheck) > 0 && uncheck[0] {
		uck = true
	}
	rftSelfDeep(ty, vl, "", "", fn, uck)
}

func rftSelfDeep(ty reflect.Type, vl reflect.Value, tg reflect.StructTag, fieldName string, fn func(rty reflect.Type, rvl reflect.Value, rtg reflect.StructTag, fieldName string) bool, uncheck bool) {
	switch ty.Kind() {
	case reflect.Ptr:
		if vl.Pointer() == 0 && vl.CanSet() {
			if !fn(ty, vl, tg, "") {
				return
			}
		}
		ty = ty.Elem()
		vl = vl.Elem()
		rftSelfDeep(ty, vl, tg, fieldName, fn, uncheck)

	case reflect.Struct:
		if !vl.CanSet() && !uncheck {
			return
		}

		if !fn(ty, vl, tg, "") {
			return
		}
		for i := 0; i < ty.NumField(); i++ {
			t := ty.Field(i).Type

			tag := ty.Field(i).Tag
			v := vl.Field(i)
			rftSelfDeep(t, v, tag, ty.Field(i).Name, fn, uncheck)
		}
	case reflect.Map:
	case reflect.Slice:
	case reflect.Func:
	case reflect.Chan:
	case reflect.Interface:
		ty = reflect.TypeOf(vl.Interface())
		vl = reflect.ValueOf(vl.Interface())
		tg = ""
		if ty.Kind() == reflect.Struct {
			rftSelfDeep(ty, vl, tg, fieldName, fn, uncheck)
		}
	default:
		if vl.CanSet() {
			fn(ty, vl, tg, fieldName)
		}
	}

}
