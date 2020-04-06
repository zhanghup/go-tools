package test_crypto

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestCBC(t *testing.T) {
	s := EncryptDES_CBC("hello world", "12345678")
	fmt.Println(s)
	fmt.Println(DecryptDES_CBC(s, "12345678"))

	s = tools.Crypto.DES("hello world", "12345678").CBCEncrypt()
	fmt.Println(s)
	fmt.Println(tools.Crypto.DES(s, "12345678").CBCDecrypt())
}

func TestECB(t *testing.T) {
	s := EncryptDES_ECB("hello world", "12345678")
	fmt.Println(s)
	fmt.Println(DecryptDES_ECB(s, "12345678"))

	s = tools.Crypto.DES("hello world", "12345678").ECBEncrypt()
	fmt.Println(s)
	fmt.Println(tools.Crypto.DES(s, "12345678").ECBDecrypt())
}
