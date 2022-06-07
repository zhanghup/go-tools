package tools

import (
	"fmt"
	"os"
	"testing"
)

func TestFileUploadIO(t *testing.T) {

	f, err := os.Open("C:\\Users\\Administrator\\Desktop\\图片s\\img\\f.png")
	if err != nil {
		t.Fatal(err)
	}
	stat, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	id, err := FileUploadIO(f, stat.Name())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(id)
}

func TestFileGet(t *testing.T) {
	id := "png202205-9c2b52e8-02fe-40bd-beab-78568658535a"
	fmt.Println(FileInfo(id))
}

func TestName(t *testing.T) {
	v := fmt.Sprintf("%08d", 1231)
	fmt.Println(v, len([]byte(v)))
}
