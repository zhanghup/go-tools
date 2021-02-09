package tools

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"reflect"
	"regexp"
	"time"
)

/*
	快速操作字符串
*/

var strfmtregex = regexp.MustCompile("{{.*?}}")

func StrFmt(format string, args ...interface{}) string {
	if strfmtregex.MatchString(format) && len(args) > 0 && reflect.TypeOf(args[0]).Kind() == reflect.Map {
		return StrTmp(format, args...).String()
	}

	params := make([]interface{}, 0)
	for _, p := range args {
		params = append(params, Rft.RealValue(p))
	}
	return fmt.Sprintf(format, params...)
}

// bson 的 UUID
func UUID() string {
	id := uuid.NewV4()
	return id.String()
}

// 以json格式输出struct对象
func JSONString(obj interface{}, format ...bool) string {
	var datas []byte
	if len(format) > 0 && format[0] {

		r, err := json.MarshalIndent(obj, "", "\t")
		if err != nil {
			datas = []byte("数据格式化异常")
		} else {
			datas = r
		}
	} else {
		r, err := json.Marshal(obj)
		if err != nil {
			datas = []byte("数据格式化异常")
		} else {
			datas = r
		}
	}
	return string(datas)
}

// 判断字符串是否包含在数组中
func StrContains(src []string, tag string) bool {
	for _, s := range src {
		if s == tag {
			return true
		}
	}
	return false
}

// 取固定长度的随机字符串
// flag 自否可包含特殊字符
func StrOfRand(l int, flag ...bool) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	res := make([]byte, l)
	for i := 0; i < l; i++ {
		b := 0
		if len(flag) == 0 || !flag[0] {
			switch r.Int() % 3 {
			case 0:
				b = r.Intn(10) + 48
			case 1:
				b = r.Intn(26) + 65
			case 2:
				b = r.Intn(26) + 97
			}
		} else {
			b = r.Intn(90) + 33
		}
		res[i] = byte(b)
	}
	return string(res)
}
