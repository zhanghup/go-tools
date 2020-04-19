package tools

import "golang.org/x/text/encoding/charmap"

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
func (this myCharset) ISO8859_1decode(data []byte) ([]byte, error) {
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
func (this myCharset) Windows1251decode(data []byte) ([]byte, error) {
	dec := charmap.Windows1251.NewDecoder()
	out, err := dec.Bytes(data)
	if err != nil {
		return nil, err
	}
	return out, nil
}
