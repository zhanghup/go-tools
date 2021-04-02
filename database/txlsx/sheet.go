package txlsx

import (
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
	"strconv"
	"time"
)

type Sheet struct {
	Header []string
	Rows   []Row
	ext    *ExcelExt
}

type Row map[string]Cell

type Cell struct {
	value string
	ext   *ExcelExt
}

func (this Row) Cell(name string) Cell {
	c, ok := this[name]
	if !ok {
		return Cell{value: "", ext: &ExcelExt{dictmap: map[string][]ExcelDictItem{}, custom: map[string]ExcelCustomList{}}}
	}
	return c
}

func (this Cell) String() string {
	return this.value
}
func (this Cell) PtrString() *string {
	if len(this.value) == 0 {
		return nil
	}
	return &this.value
}

func (this Cell) PtrInt() *int {
	i, err := strconv.Atoi(this.value)
	if err != nil {
		tog.Error("Excel数据转换异常,Error: %s", err.Error())
		return nil
	}
	return &i
}

func (this Cell) Int() int {
	i := this.PtrInt()
	if i == nil {
		return 0
	}
	return *i
}

func (this Cell) PtrInt64() *int64 {
	i, err := strconv.ParseInt(this.value, 10, 64)
	if err != nil {
		return nil
	}
	return &i
}

func (this Cell) Int64() int64 {
	i := this.PtrInt64()
	if i == nil {
		return 0
	}
	return *i
}

func (this Cell) PtrFloat64() *float64 {
	v, err := strconv.ParseFloat(this.value, 64)
	if err != nil {
		return nil
	}
	return &v
}

func (this Cell) Float64() float64 {
	i := this.PtrFloat64()
	if i == nil {
		return 0
	}
	return *i
}

func (this Cell) PtrTime(fmt ...string) *int64 {
	f := "2006-01-02 15:04:05"
	if len(fmt) > 0 {
		f = fmt[0]
	}
	t, err := time.ParseInLocation(f, this.value, time.Local)
	if err != nil {
		return nil
	}
	return tools.PtrOfInt64(t.Unix())

}

func (this Cell) Time(fmt ...string) int64 {
	v := this.PtrTime(fmt...)
	if v == nil {
		return 0
	}
	return *v
}

func (this Cell) DictValue(fmt string) string {
	return this.ext.GetDictValue(fmt, this.value)
}
func (this Cell) DictPtrValue(fmt string) *string {
	return this.ext.GetDictPtrValue(fmt, this.value)
}
func (this Cell) DictName(fmt string) string {
	return this.ext.GetDictName(fmt, this.value)
}
func (this Cell) DictPtrName(fmt string) *string {
	return this.ext.GetDictPtrName(fmt, this.value)
}

func (this Cell) CustomId(fmt string) string {
	return this.ext.GetCustomId(fmt, this.value)
}
func (this Cell) CustomPtrId(fmt string) *string {
	return this.ext.GetCustomPtrId(fmt, this.value)
}
func (this Cell) CustomName(fmt string) string {
	return this.ext.GetCustomName(fmt, this.value)
}
func (this Cell) CustomPtrName(fmt string) *string {
	return this.ext.GetCustomPtrName(fmt, this.value)
}
