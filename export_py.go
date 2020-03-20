package tools

import "github.com/zhanghup/go-tools/pinyin"

type mypy struct{}

var Py = mypy{}

func (mypy) Py(str string) string {
	return pinyin.Py(str)
}
func (mypy) PY(str string) string {
	return pinyin.PY(str)
}
func (mypy) Pinyin(str string) string {
	return pinyin.Pinyin(str)
}
func (mypy) PINYIN(str string) string {
	return pinyin.Pinyin(str)
}
