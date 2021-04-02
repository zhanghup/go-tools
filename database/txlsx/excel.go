package txlsx

import "github.com/tealeg/xlsx"

func (this Engine) Excel(data []byte, columnIdx, dataIdx int) (Excel, error) {
	xlFile, err := xlsx.OpenBinary(data)
	if err != nil {
		return Excel{}, err
	}

	columnIdx -= 1
	dataIdx -= 1

	result := Excel{origin: xlFile, ext: this.ext, Data: map[string]Sheet{}}

	for _, sheet := range xlFile.Sheets {
		result.Data[sheet.Name] = Sheet{
			Header: make([]string, 0),
			Rows:   make([]Row, 0),
			ext: this.ext,
		}

		if len(sheet.Rows) > columnIdx {
			row := sheet.Rows[columnIdx]
			sh := result.Data[sheet.Name]
			for _, cell := range row.Cells {
				sh.Header = append(sh.Header, cell.String())
			}
			result.Data[sheet.Name] = sh
		}
	}

	for _, sheet := range xlFile.Sheets {
		sh := result.Data[sheet.Name]
		for i, row := range sheet.Rows {
			if i < dataIdx || row == nil {
				continue
			}
			rowmap := Row{}
			for j, h := range result.Data[sheet.Name].Header {
				if j >= len(row.Cells) || len(h) == 0 {
					continue
				}
				rowmap[h] = Cell{row.Cells[j].String(), this.ext}
			}
			sh.Rows = append(sh.Rows, rowmap)
		}
		result.Data[sheet.Name] = sh
	}

	return result, nil
}

