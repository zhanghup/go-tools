package tools

/*
	快速操作指针
*/

import (
	"errors"
	"reflect"
)

func PtrCheck(i interface{}) error {
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		return nil
	}
	return errors.New("数据类型异常，必须为指针类型")
}

func PtrOfUUID() *string {
	return PtrOfString(UUID())
}
func PtrOfString(i string) *string {
	return &i
}
func PtrOfInt(i int) *int {
	return &i
}
func PtrOfInt8(i int8) *int8 {
	return &i
}
func PtrOfInt16(i int16) *int16 {
	return &i
}
func PtrOfInt32(i int32) *int32 {
	return &i
}
func PtrOfInt64(i int64) *int64 {
	return &i
}
func PtrOfFloat32(i float32) *float32 {
	return &i
}
func PtrOfFloat64(i float64) *float64 {
	return &i
}
func PtrOfInterface(i interface{}) *interface{} {
	return &i
}

func PtrOfBool(i bool) *bool {
	return &i
}
