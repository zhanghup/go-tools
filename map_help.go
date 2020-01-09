package tools

import (
	"encoding/json"
	"reflect"
)

type mapinf struct {
}

var omapinf = mapinf{}

func (this mapinf) MapOfInterface(o reflect.Value) map[string]interface{} {
	result := map[string]interface{}{}
	switch o.Kind() {
	case reflect.Ptr:
		return this.MapOfInterface(o.Elem())
	case reflect.Map:
		for _, k := range o.MapKeys() {
			result[k.String()] = o.MapIndex(k).Interface()
		}
		return result
	case reflect.Struct:
		bt, err := json.Marshal(o.Interface())
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bt, &result)
		if err != nil {
			panic(err)
		}
	}
	return result
}

func MapOfInterface(o interface{}) map[string]interface{} {
	return omapinf.MapOfInterface(reflect.ValueOf(o))
}
