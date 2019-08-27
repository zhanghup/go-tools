package pinyin

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	hans := "中国人"
	a := NewArgs()
	fmt.Println(Pinyin(hans, a))
}
