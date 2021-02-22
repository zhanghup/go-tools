package tools

import (
	"bytes"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func UTF8ToISO8859_1(data []byte) ([]byte, error) {
	enc := charmap.ISO8859_1.NewEncoder()
	out, err := enc.Bytes(data)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func ISO8859_1ToUTF8(data []byte) ([]byte, error) {
	dec := charmap.ISO8859_1.NewDecoder()
	out, err := dec.Bytes(data)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func UTF8ToWindows1251(data []byte) ([]byte, error) {
	enc := charmap.Windows1251.NewEncoder()
	out, err := enc.Bytes(data)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func Windows1251ToUTF8(data []byte) ([]byte, error) {
	dec := charmap.Windows1251.NewDecoder()
	out, err := dec.Bytes(data)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func UTF8ToGBK(data []byte) ([]byte, error) {
	out, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(data), simplifiedchinese.GBK.NewEncoder()))
	return out, err
}
func GBKToUTF8(data []byte) ([]byte, error) {
	out, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(data), simplifiedchinese.GBK.NewDecoder()))
	return out, err
}
