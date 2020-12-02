package texcel

import (
	"os"
	"testing"
)

func TestOpenExcel(t *testing.T) {
	f, err := os.Open("C:\\Users\\Administrator\\Downloads\\表具档案 (1).xlsx")
	if err != nil {
		panic(err)
	}
	e, err := ExcelIO(f)
	if err != nil {
		panic(err)
	}
	err = e.ReadMapBySheetName(0, 0, 2)
	if err != nil {
		panic(err)
	}
}
