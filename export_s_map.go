package tools

import (
	"fmt"
	"strings"
)

/*
	StringMap
	将一个map对象格式化成一个string，例如：
	map[string]string{"name":"test","age": "123"} => name:test,age:123
*/
type StringMap string

func (this StringMap) toMap() map[string]string {
	ss := strings.Split(string(this), ",")
	res := map[string]string{}

	for _, sss := range ss {
		s := strings.Split(sss, ":")
		if len(s) == 2 {
			res[s[0]] = s[1]
		}
	}
	return res
}

func (this StringMap) Get(key string) string {
	return this.toMap()[key]
}

func (this StringMap) Set(key, val string) StringMap {
	m := this.toMap()
	m[key] = val

	strs := make([]string, 0)
	for k, v := range m {
		strs = append(strs, fmt.Sprintf("%s:%s", k, v))
	}
	return StringMap(strings.Join(strs, ","))
}
