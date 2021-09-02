package extraction_test

import (
	"fmt"
	extraction "github.com/zhanghup/go-tools/database/txlsx"
	"os"
	"testing"
)

func TestXls(t *testing.T) {
	f, err := os.Open("./engine_xls.xls")
	if err != nil {
		t.Error(err)
		return
	}

	csvEngine := extraction.NewEngineXls(nil)
	res, err := csvEngine.OpenIO(f)
	if err != nil {
		t.Error(err)
		return
	}

	for _, sh := range res {
		for _, row := range sh.DataMap() {
			fmt.Println(row)
		}
	}
}
