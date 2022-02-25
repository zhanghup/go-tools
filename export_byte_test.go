package tools

import (
	"fmt"
	"testing"
	"time"
)

func TestInt64ToBytes(t *testing.T) {
	v := Int64ToBytes(time.Now().Unix())
	fmt.Println(v)
	//vv := Int64ToBytes(1645798603)
	vv := make([]byte, 10, 16)
	vv[0] = 49
	vv[1] = 54
	vv[2] = 52
	vv[3] = 53
	vv[4] = 56
	vv[5] = 48
	vv[6] = 50
	vv[7] = 51
	vv[8] = 48
	vv[9] = 48

	fmt.Println(BytesToInt64(Int64ToBytes(time.Now().Unix())))
	fmt.Println(BytesToInt64(vv))
	fmt.Println(string(vv))
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
