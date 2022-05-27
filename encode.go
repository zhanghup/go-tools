package tools

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
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
	"io"
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

func Base64Enc(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
func Base64Encode(data string) string {
	return Base64Enc([]byte(data))
}
func Base64Dec(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
func Base64Decode(data string) (string, error) {
	b, err := Base64Dec(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func SHA256WithRSA(signContent string, privateKey string) []byte {
	return RsaSign(signContent, privateKey, crypto.SHA256)
}

func SHA256(data []byte) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, bytes.NewBuffer(data)); err != nil {
		return "", err
	}

	checksum := fmt.Sprintf("%X", h.Sum(nil))
	return checksum, nil
}

func RsaSign(signContent string, privateKey string, hash crypto.Hash) []byte {
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
		priKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return priKey.(*rsa.PrivateKey), nil
	}

	priKey, err := ParsePrivateKey(privateKey)
	if err != nil {
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, hash, hashed)
	if err != nil {
		panic(err)
	}
	return signature
}

func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)

	return base64.StdEncoding.EncodeToString(cryted)

}

func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
