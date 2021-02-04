package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/zhanghup/go-tools/twindows"
)

func main() {
	f, err := os.Create("./m.text")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for {
		time.Sleep(time.Second)
		fmt.Print("hello world!!!\n")
		f.Write([]byte("hello world!!!\n"))
	}
}
