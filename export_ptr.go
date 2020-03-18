package tools

/*
	快速操作指针
 */

import (
	"errors"
	"reflect"
)

func PtrCheck(i interface{}) error {
	if reflect.TypeOf(i).Kind() == reflect.Ptr{
		return nil
	}
	return errors.New("数据类型异常，必须为指针类型")
}
func PtrString(i string) *string {
	return &i
}
func PtrInt(i int) *int {
	return &i
}
func PtrInt8(i int8) *int8 {
	return &i
}
func PtrInt16(i int16) *int16 {
	return &i
}
func PtrInt32(i int32) *int32 {
	return &i
}
func PtrInt64(i int64) *int64 {
	return &i
}
func PtrFloat32(i float32) *float32 {
	return &i
}
func PtrFloat64(i float64) *float64 {
	return &i
}
func PtrInterface(i interface{}) *interface{} {
	return &i
}