package tools

import (
	"strconv"
)

type myParse struct{}

var Parse = myParse{}

func (this myParse) MustStrToInt(s string) int {
	return int(this.MustStrToInt64(s))
}

func (this myParse) MustStrToInt16(s string) int16 {
	return int16(this.MustStrToInt64(s))
}

func (this myParse) MustStrToInt32(s string) int32 {
	return int32(this.MustStrToInt64(s))
}

func (this myParse) MustStrToInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func (this myParse) MustStrToFloat32(s string) float32 {
	return float32(this.MustStrToFloat64(s))
}

func (this myParse) MustStrToFloat64(s string) float64 {
	n, _ := strconv.ParseFloat(s, 64)
	return n
}

func (this myParse) Float32ToStr(f float32, n ...int) string {
	return this.Float64ToStr(float64(f), n...)
}

func (this myParse) Float64ToStr(f float64, n ...int) string {
	prec := -1
	if len(n) > 0 {
		prec = n[0]
	}
	return strconv.FormatFloat(f, 'f', prec, 64)
}

func (this myParse) Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func (this myParse) Int32ToStr(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func (this myParse) Int16ToStr(i int16) string {
	return strconv.FormatInt(int64(i), 10)
}

func (this myParse) IntToStr(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
