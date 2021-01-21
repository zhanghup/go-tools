package texcel

import "github.com/zhanghup/go-tools"

type TCell string

func (this TCell) String() string {
	return string(this)
}

func (this TCell) Int() int {
	return tools.StrToInt(string(this))
}
func (this TCell) Int8() int8 {
	return tools.StrToInt8(string(this))
}
func (this TCell) Int16() int16 {
	return tools.StrToInt16(string(this))
}
func (this TCell) Int32() int32 {
	return tools.StrToInt32(string(this))
}
func (this TCell) Int64() int64 {
	return tools.StrToInt64(string(this))
}
func (this TCell) Float32() float32 {
	return tools.StrToFloat32(string(this))
}
func (this TCell) Float64() float64 {
	return tools.StrToFloat64(string(this))
}
