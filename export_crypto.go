package tools

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

func Password(password, salt string) string {
	sh := sha256.New()
	sh.Write([]byte(password + "  xxxx  "))
	bts := sh.Sum([]byte(salt))
	return fmt.Sprintf("%x", bts)
}

func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
func Base64Decode(data string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func SHA256WithRSA(signContent string, privateKey string) string {
	return RsaSign(signContent, privateKey, crypto.SHA256)
}

func RsaSign(signContent string, privateKey string, hash crypto.Hash) string {
	shaNew := hash.New()
	shaNew.Write([]byte(signContent))
	hashed := shaNew.Sum(nil)

	ParsePrivateKey := func(privateKey string) (*rsa.PrivateKey, error) {
		PEM_BEGIN := "-----BEGIN RSA PRIVATE KEY-----\n"
		PEM_END := "\n-----END RSA PRIVATE KEY-----"
		if !strings.HasPrefix(privateKey, "-----") {
			privateKey = PEM_BEGIN + privateKey
		}
		if !strings.HasSuffix(privateKey, "-----") {
			privateKey = privateKey + PEM_END
		}

		// 2、解码私钥字节，生成加密对象
		block, _ := pem.Decode([]byte(privateKey))
		if block == nil {
			return nil, errors.New("私钥信息错误！")
		}
		// 3、解析DER编码的私钥，生成私钥对象
		priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return priKey, nil
	}

	priKey, err := ParsePrivateKey(privateKey)
	if err != nil {
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, hash, hashed)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(signature)
}
