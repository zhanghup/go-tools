package tools

import (
	"github.com/zhanghup/go-tools/pinyin"
	"strings"
)

func Pinyin(str string) string {
	args := pinyin.NewArgs()
	args.Separator = ""
	s := pinyin.Pinyin(str, args)
	result := ""
	for _, o := range s {
		for _, oo := range o {
			result += oo
		}
	}
	return result
}

func PINYIN(str string) string {
	return strings.ToUpper(Pinyin(str))
}

func Py(str string) string {
	args := pinyin.NewArgs()
	args.Separator = ""
	s := pinyin.Pinyin(str, args)
	result := ""
	for _, o := range s {
		for _, oo := range o {
			if len(oo) > 0 {
				result += oo[:1]
			}
		}
	}
	return result
}
func PY(str string) string {
	return strings.ToUpper(Py(str))
}
