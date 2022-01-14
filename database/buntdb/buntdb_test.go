package buntdb

import (
	"fmt"
	"testing"
)

func TestBuntdb(t *testing.T) {
	e, err := NewEngine(Option{
		Path: "./data.db",
	})
	if err != nil {
		panic(err)
	}

	sess := e.Sess()
	if err != nil {
		panic(err)
	}
	pri, replace, err := sess.Set("a", "a222")
	fmt.Println("pri:", pri, "replace:", replace)

	v, err := sess.Get("a")
	if err != nil {
		panic(err)
	}
	fmt.Println("value:", v)
}
