package extraction

import "hub.ffcode.net/ffcode/framework/utils"

type Sheet struct {
	Name     string
	Config   *Config
	DataList [][]Cell

	headerIdx int
	columnIdx int
	dataIdx   *int
}

func NewSheet(cfg *Config) *Sheet {
	s := &Sheet{
		Config: cfg,
	}

	return s
}

type Row struct {
	data map[string]Cell
}

func (this Row) Cell(name string) Cell {
	return this.data[name]
}

func (this Row) Data() map[string]string {
	data := map[string]string{}

	for k, v := range this.data {
		data[k] = v.String()
	}

	return data
}

func (this *Sheet) SetHeader(i int) *Sheet {
	this.headerIdx = i
	return this
}

func (this *Sheet) SetColumn(i int) *Sheet {
	this.columnIdx = i
	return this
}

func (this *Sheet) SetDataIdx(i int) *Sheet {
	this.dataIdx = &i
	return this
}

func (this *Sheet) Header() []string {
	if this.headerIdx >= len(this.DataList) {
		return nil
	}

	result := make([]string, 0)
	for _, c := range this.DataList[this.headerIdx] {
		result = append(result, c.Value)
	}

	return result
}

func (this *Sheet) Column() []string {
	cells := this.ColumnCells()
	ls := make([]string, 0)

	for _, o := range cells {
		ls = append(ls, o.Value)
	}

	return ls
}

func (this *Sheet) ColumnCells() []Cell {
	if this.columnIdx >= len(this.DataList) {
		return nil
	}
	typeMap := map[string]CellType{}
	if this.Config != nil {
		typeMap = this.Config.typeMap
	}

	result := make([]Cell, 0)
	for _, c := range this.DataList[this.columnIdx] {
		ctype, ok := typeMap[c.Value]
		if ok {
			c.Type = ctype
		}
		result = append(result, c)
	}

	return result
}

func (this *Sheet) DataMap() []Row {
	// 属性映射
	typeMap := map[string]CellType{}
	if this.Config != nil {
		typeMap = this.Config.typeMap
	}

	// 数据所在行设定
	idx := this.dataIdx
	if idx == nil {
		idx = utils.PtrOfInt(1)
	}

	// 列定义
	headers := this.Header()

	result := make([]Row, 0)
	for i := *idx; i < len(this.DataList); i++ {
		item := map[string]Cell{}

		for j := 0; j < len(this.DataList[i]); j++ {
			if j < len(headers) {
				key := headers[j]
				cell := this.DataList[i][j]
				ctype, ok := typeMap[key]
				if ok {
					cell.Type = ctype
					item[key] = cell
				} else {
					item[key] = cell
				}
			}
		}

		result = append(result, Row{
			data: item,
		})
	}

	return result
}
