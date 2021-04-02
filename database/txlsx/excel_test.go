package txlsx_test

import (
	"fmt"
	"github.com/zhanghup/go-tools/database/txlsx"
	"io/ioutil"
	"os"
	"testing"
)

func TestExcel(t *testing.T) {
	path := "C:\\Users\\Administrator\\Downloads\\用户档案20210401141102.xlsx"
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	e := txlsx.NewEngine()

	e.SetDicts([]txlsx.ExcelDictItem{
		{Code: "S0003", Name: "流量计", Value: "Value"},
	})

	ex, err := e.Excel(data, 2, 3)
	if err != nil {
		panic(err)
	}
	for _, sheet := range ex.Data {
		for _, row := range sheet.Rows {
			//fmt.Println(row)
			fmt.Println(row.Cell("o_device.type").DictPtrValue("S0003"))
		}
	}

	//fmt.Println(ex)
}
