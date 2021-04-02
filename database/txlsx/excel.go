package txlsx

import "github.com/tealeg/xlsx/v3"

func (this Engine) Excel(data []byte, columnIdx, dataIdx int) (Excel, error) {
	xlFile, err := xlsx.OpenBinary(data)
	if err != nil {
		return Excel{}, err
	}

	result := Excel{origin: xlFile, ext: this.ext, Data: map[string]Sheet{}}

	for _, sheet := range xlFile.Sheets {
		result.Data[sheet.Name] = Sheet{
			Header: make([]string, 0),
			Rows:   make([]Row, 0),
			ext:    this.ext,
		}

		if sheet.MaxRow < columnIdx {
			return result, nil
		}

		err := sheet.ForEachRow(func(row *xlsx.Row) error {
			if row.GetCoordinate() != columnIdx {
				return nil
			}

			sh := result.Data[sheet.Name]
			err := row.ForEachCell(func(cell *xlsx.Cell) error {
				sh.Header = append(sh.Header, cell.String())
				return nil
			})
			if err != nil {
				return err
			}
			result.Data[sheet.Name] = sh
			return nil
		})
		if err != nil {
			return result, err
		}
	}

	for _, sheet := range xlFile.Sheets {
		sh := result.Data[sheet.Name]

		err := sheet.ForEachRow(func(row *xlsx.Row) error {
			if row.GetCoordinate() < dataIdx {
				return nil
			}
			rowmap := Row{data: map[string]Cell{}}
			err := row.ForEachCell(func(cell *xlsx.Cell) error {
				n, _ := cell.GetCoordinates()
				if n < len(sh.Header) {
					rowmap.data[sh.Header[n]] = Cell{value: cell.String(), ext: this.ext, info: cell}
				}

				return nil
			})
			if err != nil {
				return err
			}
			sh.Rows = append(sh.Rows, rowmap)
			return nil
		})
		if err != nil {
			return result, err
		}

		result.Data[sheet.Name] = sh
	}

	return result, nil
}
