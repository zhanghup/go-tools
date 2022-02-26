package tools

import (
	"fmt"
	"testing"
)

func TestStringMap_Get(t *testing.T) {
	s := StringMap("name:1,age:2")
	fmt.Println(s.Get("name"))

	s = s.Set("a", "2")
	fmt.Println(s)
}
