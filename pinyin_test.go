package tools

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(Pinyin("你好"))
	fmt.Println(PINYIN("你好"))
	fmt.Println(Py("你好"))
	fmt.Println(PY("你好"))
}
