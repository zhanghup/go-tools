package tools

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type mycrypto struct{}

var Crypto = mycrypto{}

func (mycrypto) Password(password, salt string) string {
	sh := sha256.New()
	sh.Write([]byte(password + "  xxxx  "))
	bts := sh.Sum([]byte(salt))
	return fmt.Sprintf("%x", bts)
}

func (mycrypto) MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
