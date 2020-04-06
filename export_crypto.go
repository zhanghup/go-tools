package tools

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
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

type mydes struct {
	src     []byte
	key     []byte
	block   cipher.Block
	padding []byte
}

func (mycrypto) DES(src, key string) *mydes {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	dd := &mydes{
		block: block,
		src:   data,
		key:   keyByte,
	}
	return dd.pKCS5Padding()
}

func (this *mydes) CBCEncrypt() string {
	//获取CBC加密模式
	iv := this.key //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCEncrypter(this.block, iv)
	out := make([]byte, len(this.padding))
	mode.CryptBlocks(out, this.padding)
	return fmt.Sprintf("%X", out)
}

func (this *mydes) CBCDecrypt() string {
	dist, err := hex.DecodeString(string(this.src))
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(this.block, this.key)
	plaintext := make([]byte, len(dist))
	mode.CryptBlocks(plaintext, dist)
	plaintext = this.pKCS5UnPadding(plaintext)
	return string(plaintext)
}

func (this *mydes) ECBEncrypt() string {
	out := make([]byte, len(this.padding))
	dst := out
	bs := this.block.BlockSize()
	for len(this.padding) > 0 {
		//对明文按照blocksize进行分块加密
		//必要时可以使用go关键字进行并行加密
		this.block.Encrypt(dst, this.padding[:bs])
		this.padding = this.padding[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%X", out)
}

func (this *mydes) ECBDecrypt() string {
	data, err := hex.DecodeString(string(this.src))
	if err != nil {
		panic(err)
	}

	bs := this.block.BlockSize()
	if len(data)%bs != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		this.block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = this.pKCS5UnPadding(out)
	return string(out)
}

//明文补码算法
func (this *mydes) pKCS5Padding() *mydes {
	padding := this.block.BlockSize() - len(this.src)%this.block.BlockSize()
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	this.padding = append(this.src, padtext...)
	return this
}

//明文减码算法
func (this *mydes) pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
