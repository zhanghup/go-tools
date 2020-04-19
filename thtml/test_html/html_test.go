package test_html

import (
	"fmt"
	"github.com/zhanghup/go-tools/thtml"
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
	html, err := thtml.New(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(html.Charset())
	err = html.DecodeHtml()
	if err != nil{
		panic(err)
	}
	fmt.Println(html.Title())
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
	html, err := thtml.New(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(html.Charset())
	err = html.DecodeHtml()
	if err != nil{
		panic(err)
	}
	fmt.Println(html.Title())
}
