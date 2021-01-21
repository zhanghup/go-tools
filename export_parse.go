package tools

import (
	"strconv"
)

func StrToInt(s string) int {
	return int(StrToInt64(s))
}

func StrToInt8(s string) int8 {
	return int8(StrToInt64(s))
}

func StrToInt16(s string) int16 {
	return int16(StrToInt64(s))
}

func StrToInt32(s string) int32 {
	return int32(StrToInt64(s))
}

func StrToInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func StrToFloat32(s string) float32 {
	return float32(StrToFloat64(s))
}

func StrToFloat64(s string) float64 {
	n, _ := strconv.ParseFloat(s, 64)
	return n
}

func Float32ToStr(f float32, n ...int) string {
	return Float64ToStr(float64(f), n...)
}

func Float64ToStr(f float64, n ...int) string {
	prec := -1
	if len(n) > 0 {
		prec = n[0]
	}
	return strconv.FormatFloat(f, 'f', prec, 64)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int32ToStr(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int16ToStr(i int16) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int8ToStr(i int8) string {
	return strconv.FormatInt(int64(i), 10)
}

func IntToStr(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
