package txlsx

import "github.com/tealeg/xlsx"

type Engine struct {
	excel *xlsx.File
	ext   *ExcelExt
}

func NewEngine() *Engine {
	result := Engine{
		ext: &ExcelExt{
			custom:  map[string]ExcelCustomList{},
			dictmap: map[string][]ExcelDictItem{},
		},
	}
	return &result
}

func (this *Engine) SetDicts(dicts []ExcelDictItem) {
	this.ext.dictmap = map[string][]ExcelDictItem{}
	for _, o := range dicts {
		if _, ok := this.ext.dictmap[o.Code]; ok {
			this.ext.dictmap[o.Code] = append(this.ext.dictmap[o.Code], o)
		} else {
			this.ext.dictmap[o.Code] = []ExcelDictItem{o}
		}
	}
}

func (this *Engine) SetCustom(name string, item []ExcelCustomItem) {
	this.ext.custom[name] = item
}

type Excel struct {
	origin *xlsx.File
	Data   map[string]Sheet
	ext    *ExcelExt
}
