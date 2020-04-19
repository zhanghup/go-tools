package test_html

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"io/ioutil"
	"os"
	"testing"
)

func TestWindow1521(t *testing.T) {
	f, err := os.Open("./windows1521.html")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	tools.Html(data).DecodeHtml()
}

func TestGbk(t *testing.T) {
	f, err := os.Open("./gbk.html")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	h := tools.Html(data)
	h.DecodeHtml()
	fmt.Println(string(h.Title()))
}
