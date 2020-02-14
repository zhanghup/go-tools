package secret

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Password(password, slat string) string {
	sh := sha256.New()
	sh.Write([]byte(password))
	bts := sh.Sum([]byte(slat))
	return fmt.Sprintf("%x", bts)
}
func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
