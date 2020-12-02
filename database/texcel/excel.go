package texcel

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io"
)

type TExcel struct {
	excel *excelize.File
}
type TSheet []TSheetItem
type TSheetItem map[string]TCell

func ExcelIO(read io.Reader) (*TExcel, error) {
	f, err := excelize.OpenReader(read)
	if err != nil {
		return nil, err
	}
	obj := &TExcel{excel: f}
	return obj, nil
}

/*
	@sheetName: 工作表名称 - 默认读取第一张工作表
	@startCol: 读取第几列以后的数据
	@rowHeader: 作为map对象的key是在工作表的第几行,行数从0开始数
	@rowData: 从工作表的第几行开始读数据,行数从0开始数。 注意：@rowData的值一定比@rowHeader的值小
*/
func (this *TExcel) ReadMapBySheetName(startCol, rowHeader, rowData int, sheetNames ...string) (TSheet, error) {
	sn := ""
	if len(sheetNames) > 0 {
		sn = sheetNames[0]
		flag := false
		for _, s := range this.excel.GetSheetList() {
			if s == sn {
				flag = true
				break
			}
		}
		if !flag {
			return nil, errors.New("[2]工作表不存在")
		}
	} else {
		sn = this.excel.GetSheetList()[0]
	}

	rows, err := this.excel.Rows(sn)
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0)
	resultmap := make(TSheet, 0)

	for i := 0; rows.Next(); i++ {
		cols, err := rows.Columns()
		if err != nil {
			return nil, err
		}

		if i == rowHeader {
			for j := range cols {
				if j < startCol {
					continue
				}
				keys = append(keys, cols[j])
			}
		}

		if i >= rowData {
			result := TSheetItem{}
			for j := range cols {
				if j < startCol {
					continue
				}
				if j < len(keys) {
					result[keys[j]] = TCell(cols[j])
				}
			}
			resultmap = append(resultmap, result)
		}

	}
	return resultmap, nil

}

/*
	@sheetIndex: 工作表序号（第几张工作表） - 默认读取第一张工作表
	@startCol: 读取第几列以后的数据
	@rowHeader: 作为map对象的key是在工作表的第几行,行数从0开始数
	@rowData: 从工作表的第几行开始读数据,行数从0开始数。 注意：@rowData的值一定比@rowHeader的值小
*/
func (this *TExcel) ReadMapBySheetIndex(startCol, rowHeader, rowData int, sheetIndexs ...int) (TSheet, error) {
	sheetIndex := 0
	if len(sheetIndexs) > 0 {
		sheetIndex = sheetIndexs[0]
		if sheetIndex >= len(this.excel.GetSheetList()) {
			return nil, errors.New("[1] 工作表不存在")
		}
	}

	return this.ReadMapBySheetName(startCol, rowHeader, rowData, this.excel.GetSheetList()[sheetIndex])
}
