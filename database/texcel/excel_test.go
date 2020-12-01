package texcel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"testing"
)

func TestOpenExcel(t *testing.T) {
	f, err := excelize.OpenFile("D:\\workspace-go\\work\\jiahua\\boc-gateway-doc\\中国银行\\银企直连\\嘉化cbs收款项目20191202.xlsx")
	if err != nil {
		panic(err)
	}
	rows, err := f.Rows("目标任务及流程")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			panic(err)
		}
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
