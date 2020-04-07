package tools

/*
	快速操作指针
*/

import (
	"errors"
	"reflect"
)

type myptr struct{}

var Ptr = myptr{}

func (myptr) Check(i interface{}) error {
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		return nil
	}
	return errors.New("数据类型异常，必须为指针类型")
}

func (this myptr) Uid() *string {
	return this.String(Str.Uid())
}
func (myptr) String(i string) *string {
	return &i
}
func (myptr) Int(i int) *int {
	return &i
}
func (myptr) Int8(i int8) *int8 {
	return &i
}
func (myptr) Int16(i int16) *int16 {
	return &i
}
func (myptr) Int32(i int32) *int32 {
	return &i
}
func (myptr) Int64(i int64) *int64 {
	return &i
}
func (myptr) Float32(i float32) *float32 {
	return &i
}
func (myptr) Float64(i float64) *float64 {
	return &i
}
func (myptr) Interface(i interface{}) *interface{} {
	return &i
}
func (myptr) Bool(i bool) *bool {
	return &i
}

