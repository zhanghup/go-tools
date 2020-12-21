package texcel

import (
	"fmt"
	"os"
	"testing"
)

func TestOpenExcel(t *testing.T) {
	f, err := os.Open("./TestOrder.xlsx")
	if err != nil {
		panic(err)
	}
	e, err := ExcelIO(f)
	if err != nil {
		panic(err)
	}
	err = e.ReadRow(func(row int, cell []TCell) {
		fmt.Println(cell)
	})
	if err != nil {
		panic(err)
	}
}

func TestExcelData(t *testing.T) {
	f, err := os.Open("./TestOrder.xlsx")
	if err != nil {
		panic(err)
	}
	e, err := ExcelIO(f)
	if err != nil {
		panic(err)
	}

	res,err := e.ReadMapBySheetName(0,4,5)

	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
