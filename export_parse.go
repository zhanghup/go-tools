package tools

import (
	"strconv"
)

func StrToInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](s string) T {
	n, _ := strconv.ParseInt(s, 10, 64)
	return T(n)
}

func StrToFloat[T ~float32 | ~float64](s string) T {
	n, _ := strconv.ParseFloat(s, 64)
	return T(n)
}

func FloatToStr[T ~float32 | ~float64](f T, n ...int) string {
	prec := -1
	if len(n) > 0 {
		prec = n[0]
	}
	return strconv.FormatFloat(float64(f), 'f', prec, 64)
}

func IntToStr[T ~int | ~int8 | ~int16 | ~int32 | ~int64](i T) string {
	return strconv.FormatInt(int64(i), 10)
}
