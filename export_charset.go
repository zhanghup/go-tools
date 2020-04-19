package tools

import (
	"bytes"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

type myCharset struct{}

var Charset = myCharset{}

func (this myCharset) ISO8859_1Encode(data []byte) ([]byte, error) {
	enc := charmap.ISO8859_1.NewEncoder()
	out, err := enc.Bytes(data)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (this myCharset) ISO8859_1Decode(data []byte) ([]byte, error) {
	dec := charmap.ISO8859_1.NewDecoder()
	out, err := dec.Bytes(data)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (this myCharset) Windows1251Encode(data []byte) ([]byte, error) {
	enc := charmap.Windows1251.NewEncoder()
	out, err := enc.Bytes(data)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (this myCharset) Windows1251Decode(data []byte) ([]byte, error) {
	dec := charmap.Windows1251.NewDecoder()
	out, err := dec.Bytes(data)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (this myCharset) GBKEncode(data []byte) ([]byte, error) {
	out, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(data)), simplifiedchinese.GBK.NewEncoder()))
	return out, err
}
func (this myCharset) GBKDecode(data []byte) ([]byte, error) {
	out, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(data)), simplifiedchinese.GBK.NewDecoder()))
	return out, err
}
