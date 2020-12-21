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

func (this *TExcel) sheetByName(name ...string) (*excelize.Rows, error) {
	sn := ""
	if len(name) > 0 {
		sn = name[0]
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

	return this.excel.Rows(sn)
}

func (this *TExcel) ReadRow(fn func(row int, cols []TCell), sheetName ...string) error {
	rows, err := this.sheetByName(sheetName...)
	if err != nil {
		return err
	}
	for i := 0; rows.Next(); i++ {
		cols, err := rows.Columns()
		if err != nil {
			return err
		}
		tcells := make([]TCell, 0)
		for j := 0; j < len(cols); j++ {
			tcells = append(tcells, TCell(cols[j]))
		}
		fn(i, tcells)
	}
	return nil
}

/*
	规则表格解析
	@sheetName: 工作表名称 - 默认读取第一张工作表
	@startCol: 读取第几列以后的数据,从0开始数
	@rowHeader: 作为map对象的key是在工作表的第几行,行数从0开始数
	@rowData: 从工作表的第几行开始读数据,行数从0开始数。 注意：@rowData的值一定比@rowHeader的值小
*/
func (this *TExcel) ReadMapBySheetName(startCol, rowHeader, rowData int, sheetNames ...string) (TSheet, error) {

	keys := make([]TCell, 0)
	resultmap := make(TSheet, 0)

	err := this.ReadRow(func(i int, cols []TCell) {
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
					result[keys[j].String()] = cols[j]
				}
			}
			resultmap = append(resultmap, result)
		}
	}, sheetNames...)

	return resultmap, err

}

/*
	规则表格解析
	@sheetIndex: 工作表序号（第几张工作表） - 默认读取第一张工作表
	@startCol: 读取第几列以后的数据,从0开始数
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
