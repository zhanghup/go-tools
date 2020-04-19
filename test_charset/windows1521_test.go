package test_charset

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"io/ioutil"
	"os"
	"testing"
)

func TestWindows1521Encoding(t *testing.T) {
	f, err := os.Open("./index.html")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	res, err := tools.Charset.Windows1251decode(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
