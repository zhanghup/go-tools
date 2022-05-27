package tools

import (
	"bytes"
	"encoding/binary"
)

/*
	DataToBytes 将数据转换为[]byte
	支持的类型：
		bool/*bool/[]bool
		int8/*int8/[]int8
		uint8/*uint8/[]uint8
		int16/*int16/[]int16
		uint16/*uint16/[]uint16
		int32/*int32/[]int32
		uint32/*uint32/[]uint32
		int64/*int64/[]int64
		uint64/*uint64/[]uint64
		float32/*float32/[]float32
		float64/*float64/[]float64
*/
func DataToBytes[T any](n T) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

/*
	BytesToData 将数据转换为自定义类型
	支持的类型：
		bool/*bool/[]bool
		int8/*int8/[]int8
		uint8/*uint8/[]uint8
		int16/*int16/[]int16
		uint16/*uint16/[]uint16
		int32/*int32/[]int32
		uint32/*uint32/[]uint32
		int64/*int64/[]int64
		uint64/*uint64/[]uint64
		float32/*float32/[]float32
		float64/*float64/[]float64
*/
func BytesToData[T any](b []byte) T {
	bytesBuffer := bytes.NewBuffer(b)
	var x T
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}
