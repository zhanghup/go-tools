package texcel

import "github.com/zhanghup/go-tools"

type TCell string

func (this TCell) String() string {
	return string(this)
}

func (this TCell) Int() int {
	return tools.Parse.MustStrToInt(string(this))
}
func (this TCell) Int8() int8 {
	return tools.Parse.MustStrToInt8(string(this))
}
func (this TCell) Int16() int16 {
	return tools.Parse.MustStrToInt16(string(this))
}
func (this TCell) Int32() int32 {
	return tools.Parse.MustStrToInt32(string(this))
}
func (this TCell) Int64() int64 {
	return tools.Parse.MustStrToInt64(string(this))
}
func (this TCell) Float32() float32 {
	return tools.Parse.MustStrToFloat32(string(this))
}
func (this TCell) Float64() float64 {
	return tools.Parse.MustStrToFloat64(string(this))
}
