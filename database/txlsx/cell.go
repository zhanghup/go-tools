package extraction

import (
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
	"strconv"
	"strings"
	"time"
)

type Cell struct {
	Config    *Config  // Excel 数据配置
	Type      CellType // 数据类型
	Value     string   // Excel中的展示值
	Formatter string   // 数据格式化方式，一般用于时间格式化
}

type CellType string

const (
	CellTypeString CellType = "String"
	CellTypeFloat  CellType = "Float64"
	CellTypeInt    CellType = "Int64"
	CellTypeDate   CellType = "Date"

	TimeFormat = "2006-01-02 15:04:05"
)

func (this Cell) Interface() any {
	switch this.Type {
	case CellTypeString:
		return this.PtrString()
	case CellTypeDate:
		v, err := time.ParseInLocation(TimeFormat, this.TimeString(), time.Local)
		if err != nil {
			return nil
		}
		return tools.Ptr(v.Unix())
	case CellTypeInt:
		return this.PtrInt64()
	case CellTypeFloat:
		return this.PtrFloat64()
	default:
		return this.Value
	}
}

func (this Cell) TimeString() string {
	return time.Unix(this.Time(), 0).Format(TimeFormat)
}

func (this Cell) String() string {
	switch this.Type {
	case CellTypeDate:
		return this.TimeString()
	case CellTypeString:
		fallthrough
	default:
		return this.Value
	}
}

func (this Cell) PtrString() *string {
	if len(this.Value) == 0 {
		return nil
	}
	return tools.Ptr(this.String())
}

func (this Cell) PtrInt() *int {
	i, err := strconv.Atoi(this.Value)
	if err != nil {
		tog.Errorf("Excel数据转换异常,Error: %s", err.Error())
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
	i, err := strconv.ParseInt(this.Value, 10, 64)
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
	v, err := strconv.ParseFloat(this.Value, 64)
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
	if len(fmt) == 0 {
		f1 := this.Formatter
		{
			f1 = strings.ReplaceAll(f1, "ss", "05")     // 秒
			f1 = strings.ReplaceAll(f1, "mm", "04")     // 分钟
			f1 = strings.ReplaceAll(f1, "h", "15")      // 分钟
			f1 = strings.ReplaceAll(f1, "d", "02")      // 日
			f1 = strings.ReplaceAll(f1, "m", "01")      // 月
			f1 = strings.ReplaceAll(f1, "yyyy", "2006") // 年
			f1 = strings.ReplaceAll(f1, "yy", "06")     // 年
		}

		f2 := this.Formatter
		{
			f2 = strings.ReplaceAll(f2, "ss", "5")      // 秒
			f2 = strings.ReplaceAll(f2, "mm", "4")      // 分钟
			f2 = strings.ReplaceAll(f2, "h", "3")       // 分钟
			f2 = strings.ReplaceAll(f2, "d", "2")       // 日
			f2 = strings.ReplaceAll(f2, "m", "1")       // 月
			f2 = strings.ReplaceAll(f2, "yyyy", "2006") // 年
			f2 = strings.ReplaceAll(f2, "yy", "06")     // 年
		}

		v, err := time.ParseInLocation(f1, this.Value, time.Local)
		if err == nil {
			return tools.Ptr(v.Unix())
		}

		v, err = time.ParseInLocation(f2, this.Value, time.Local)
		if err == nil {
			return tools.Ptr(v.Unix())
		}
	}

	fs := fmt

	if len(fmt) == 0 {

		fs = []string{
			"2006-01-02 15:04:05",
			"2006-1-2 15:04:05",
			"2006-01-02",
			"2006-1-2",
			"2006/1/2",
			"2006/01/02 15:04:05",
			"2006/1/2 15:04:05",
			"2006年01月02日",
			"06年01月02日",
			"1/2/06 15:04",
		}
	}

	for _, f := range fs {

		if len(fmt) > 0 {
			f = fmt[0]
		}

		t, err := time.ParseInLocation(f, this.Value, time.Local)

		if err != nil {
			continue
		}

		return tools.Ptr(t.Unix())
	}

	return nil
}

func (this Cell) Time(fmt ...string) int64 {
	v := this.PtrTime(fmt...)
	if v == nil {
		return 0
	}
	return *v
}

func (this Cell) DictNameToValue(fmt string) string {
	return this.Config.GetDictValue(fmt, this.Value)
}
func (this Cell) DictPtrNameToValue(fmt string) *string {
	return this.Config.GetDictPtrValue(fmt, this.Value)
}
func (this Cell) DictValueToName(fmt string) string {
	return this.Config.GetDictName(fmt, this.Value)
}
func (this Cell) DictPtrValueToName(fmt string) *string {
	return this.Config.GetDictPtrName(fmt, this.Value)
}

func (this Cell) CustomValue(fmt string) string {
	return this.Config.GetCustomValue(fmt, this.Value)
}
func (this Cell) CustomPtrValue(fmt string) *string {
	return this.Config.GetCustomPtrValue(fmt, this.Value)
}
func (this Cell) CustomName(fmt string) string {
	return this.Config.GetCustomName(fmt, this.Value)
}
func (this Cell) CustomPtrName(fmt string) *string {
	return this.Config.GetCustomPtrName(fmt, this.Value)
}
