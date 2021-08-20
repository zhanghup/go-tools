package tools

import (
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
func RftTypeInfo(obj interface{}) RftTypeInfoView {
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
