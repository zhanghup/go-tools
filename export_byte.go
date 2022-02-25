package tools

import (
	"bytes"
	"encoding/binary"
)

func Float32ToBytes(n float32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

func BytesToFloat32(b []byte) float32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x float32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func Float64ToBytes(n float64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

func BytesToFloat64(b []byte) float64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x float64
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func Int64ToBytes(n int) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

func BytesToInt64(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func IntToBytes(n int) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}
