package tools

import (
	"fmt"
	"testing"
)

func TestFilepath(t *testing.T) {
	path := Filepath("51e761c0-d4ff-478d-923a-14fb5b2bd0af", "202101", "application/json", "jpg")
	fmt.Println(path, len([]byte(path)))
}
