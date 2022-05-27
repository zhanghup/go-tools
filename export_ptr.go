package tools

/*
	快速操作指针
*/

import (
	"errors"
	"reflect"
)

func PtrCheck(i any) error {
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		return nil
	}
	return errors.New("数据类型异常，必须为指针类型")
}

func Ptr[T any](v T) *T {
	return &v
}

func PtrOfUUID() *string {
	return Ptr(UUID())
}
