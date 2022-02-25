package tools

import (
	"fmt"
	"testing"
)

func TestInt64ToBytes(t *testing.T) {
	v := Int64ToBytes(1645798603)
	fmt.Println(v)
	fmt.Println(BytesToInt64(Int64ToBytes(1645798603)))
}

func TestFloat32(t *testing.T) {
	v := Float32ToBytes(1645798603.123)
	fmt.Println(v)
	fmt.Println(BytesToFloat32(Float32ToBytes(1645798603.123)))
}

func TestFloat64(t *testing.T) {
	v := Float64ToBytes(1645798603.123)
	fmt.Println(v)
	fmt.Println(BytesToFloat64(Float64ToBytes(1645798603.123)))
}
