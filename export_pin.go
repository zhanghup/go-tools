package tools

import "github.com/zhanghup/go-tools/pinyin"

/* 中文转拼音，例如：“你好” => “nh”  */
func Py(str string) string {
	return pinyin.Py(str)
}

/* 中文转拼音，例如：“你好” => “NH”  */
func PY(str string) string {
	return pinyin.PY(str)
}

/* 中文转拼音，例如：“你好” => “nihao”  */
func Pinyin(str string) string {
	return pinyin.Pinyin(str)
}

/* 中文转拼音，例如：“你好” => “NIHAO”  */
func PINYIN(str string) string {
	return pinyin.Pinyin(str)
}
