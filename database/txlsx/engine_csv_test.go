package extraction_test

import (
	"fmt"
	extraction "github.com/zhanghup/go-tools/database/txlsx"
	"os"
	"testing"
)

func TestCsv(t *testing.T) {
	f, err := os.Open("./engine_csv.csv")
	if err != nil {
		t.Error(err)
		return
	}

	csvEngine := extraction.NewEngineCsv(nil)
	res, err := csvEngine.OpenIO(f)
	if err != nil {
		t.Error(err)
		return
	}

	for _, sh := range res {
		fmt.Println(sh.Column())
		fmt.Println(sh.ColumnCells())
		for _, row := range sh.DataMap() {
			fmt.Println(row)
		}
	}
}
