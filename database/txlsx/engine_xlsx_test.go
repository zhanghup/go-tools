package extraction_test

import (
	"fmt"
	extraction "github.com/zhanghup/go-tools/database/txlsx"
	"os"
	"testing"
)

func TestXlsx(t *testing.T) {
	f, err := os.Open("./engine_xlsx.xlsx")
	if err != nil {
		t.Error(err)
		return
	}

	csvEngine := extraction.NewEngineXlsx(extraction.NewConfig().CellTypeMap(map[string]extraction.CellType{
		"上期行至": extraction.CellTypeDate,
	}))

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
