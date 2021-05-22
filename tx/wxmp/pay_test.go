package wxmp_test

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/wx/wxmp"
	"testing"
)

func TestOrder(t *testing.T) {
	res, err := engine.Pay(&wxmp.PayOption{
		OutTradeNo:  "test_7FF1D9D9DD8CD383",
		NotifyUrl:   "http://zander123.cn/pay/test",
		Description: "测试订单",
		TotalPrice:  1,
		Openid:      "o_v1Z4y3jKyeeIG3B5_W52TRzLBQ",
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(tools.JSONString(res, true))
}

func TestSign(t *testing.T) {
	err := engine.PayCancel("test_7FF1D9D9DD8CD383")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}


func TestDecrypt(t *testing.T) {
	str := `
		{
"id":"936cd613-269a-58e5-9620-4b1c3013c1bb",
"create_time":"2021-02-21T21:44:04+08:00",
"resource_type":"encrypt-resource",
"event_type":"TRANSACTION.SUCCESS",
"summary":"支付成功",
"resource":{
	"original_type":"transaction",
	"algorithm":"AEAD_AES_256_GCM",
	"ciphertext":"KQy9qxIBP/OrMUftFQB5ppNeb3ZPS+TBXZto+1x4Qk2mDvZK/xLTdubo9qJZi3C8DH8jEys9GZ14U/vOOSE3/PiCMpgn0fuQXZNTSXAr/6ywkpi7Iosn1B3lt3+u7yqbUluWZ/7kNNioK/0MFt3R0sP30drSLMi7RvuATtdWZ8Mx1KOS6RaFP+H7Rms9GWOehG4Jm27TipTQsAReaM8Bth+Fd5BMC7X45FDiMU4FlkX0x5mi1geDRVYy3dsYwvLKfg3HQKD/xKp/gYTgs4QB3FgoKn27358mFzvntlIqiPrO1M43FUcxz7blIIfmCSd3VBDrHHhVddRS4iiHdWWl4royE4/9wqqOlTWYLnjcQloVbJXBIVKrfoPbz6I6uqY/hP+UMZAS64o+T0kiV+CocL7ewolk+gc+M4+PfNNoeN6MY/Bcb93czB8PJL9nDPg3fz7VBmJsZCLmkEM8I+0iSGjKX+Ud6hwV8SwPrxX3TBDquW5MIalLIEctYdlP6EdZHtn08jA78yvxaVGblp4Jh8UUF+6XVQNTZ8VmK6weD7DxdnF3yO5QHSIq0AowIev2ESSqyYJSlvBq962EPbZO",
	"associated_data":"transaction",
	"nonce":"tZRj9BWUAfIq"
}}	
	`

	fmt.Println(engine.PayDecrypt([]byte(str)))
}

func CertificateDecryption(apiv3key string) (interface{}, error) {
	ciphertext := `tMq/ExWoIzLQMSIlKyzPCDVJULm4RInSR2SPxeXIG/kGHPnI8GgRrRObnA6kh4WzLFmAPDWALaAs5sYLDUGuZEgdJfZQ1aLR2S7TwjZ40xFUZFQH7fOn0vlHh5u3l7KxElrkHoINP5INphD8d/tSeSb8aRgkz0H3bT54eTE2ZUE+BLbO7H9TAvACW8MJW8yd9GsmRSZUkeJMIG1kAmbBPW75DdbBoBk+fJ94Al97Zn6zvri6Zc9mndkuAw/JIXn17Zo0dJNNV+J/bBtSr/9LFJ1dDCJpIcNmKQh5J1qkyKrAsOIA1s6ZmNbSXF9AsPBTfJ9C9EWUgiMkcV6CBoVLrUd2zUSqqH9Hp2kBhQ0gH4h0L3iAwUU4VX+4LO1WaQ3KxWtoxHs4huw6KDNTdL9PBGNVSmv2gVQEqkAx97tDIuNTaJPLVYNpm+eDrCs0CJqsWyhBsEJhNHFXXRNi/0szkmsf0b+NQKcG3VKzwpVZ7gdVgAPDtnvNamwtKe3RNH9shxB5wjTARY+JeFjPsqBMGBcHoYHRI4YVrbv0fH8nxm4fFRjw8yS1BuZ1fqlHV8QNts8++jRdoVeaOOSZTE4LuBDxIda1YQmsI6uiGM95IpAGcqDIvdcKG7udd5uimR7+W1DHyHHwwDFOYcabnRfhwlO3s6dgv4VBV73SJnUtMrAfGk8yUAQ/gHl9Bb94I1VicUkby4xK4ZBis8CIO7KYzPCl50Js222XFV8Lc6s+wVMnC2goINsv+rmvNs0ho8Egbn4ruV3Qpc3TcoHMUwa9a0dc+UxAQoyJ9ym3yIojByRdOEdnxxRKFp/axv54aelgkKqEwlHGT+KU+Huu21Id//Pmodg9DD080j9Xz8V3pHRz8Md5i3u0c3BzK2n4gVFX7omtHdM+hy54GfOusdr63xXO5TSq53Imh7zA8kse8QXX3l5NE8AuyS4Cq0iTrRUNhVD9B4ibON8Q7M7lUusGGTuDJzKb4T8NAwYTYkym7GY/mYRnacTMVjrDLgL0EXoc1K1KusbP4T+UIfS7jwVu3I/28kqWsLvTNnsjyyHHQgJFMFBtJLdl89VSQrnCOLYg4ZKqdBl4sp6v9FxfnHS1OTGljiuFRJRIns2Khji5SIIHeSaxILDykPsR/lD2rt5fWHGZ41DAGcZ3pqHGSGpnxTOr1vtx2mVNqd4NvQKyeDFVGgzcNklwm5RqjFR3agtIDdO3neCARQXgn7ftWRKw1BSfQLPFfc4Br7MMpAhSQf2zri5LVPQcyhZwHd1P76r5kPQZEh6s0N6d2vufg8mG34eCXUb/9HCt7uh/n3JAb1asvc0krhmalC1L7bDKgUwx3Qulj4rXnWHERGzsd/QlSxJaVmk9h8ZVuwMEBi8LFl+TJZc5oFmt0pleNkLPi8O1SnDVzflQOwKIlLZ+Zk8Mp22tG/tNO37CSRJzaT6CiI98cYo9ZIZr0vW4qOirE4qmi8kLrFjN/oDEEDBZ0H9pYW4OjXk5QgFufyiIUGlQ/nDrIuVxpSTaVJnFFDyroXtyjWwKK0vvlsPzNxkNnHjupUUQmrjfg+pKFlTHQsEXGPNssmvtqk+oElpP+3MVAViR7DG/6DHm2ByVwDD7HsESRQhtz2J43hXITUih/q3MCDnRx32wu/ChTGTjIZx6ZIskAkmS5fkILtkC7xIe/e70ohTQPFPkA0qCCg37kkxrZr2jg58hW5UxICdUSS3GcUvmy6I2ErxUAvTnVtmLhiA1CZpAHKJM0n3+7yk16MY9L36eILF8+3RVwtAmYFaIUN0EDWfZ7x6IK6+JBDpE4ojPLVzn4iaSo5Du+0N7PvIeC7QbyeHi+cw5m1KLD3Qzk97qib5uOX4Uwli4bhFbStCLHxXAwPHizg==`
	nonce := "a3e65f1b5ca1"
	associated_data := "certificate"

	// 对编码密文进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	c, err := aes.NewCipher([]byte(apiv3key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(decodeBytes) < nonceSize {
		return nil, errors.New("密文证书长度不够")
	}
	plaintext, err := gcm.Open(nil, []byte(nonce), decodeBytes, []byte(associated_data))
	if err != nil {
		return nil, err
	}
	fmt.Println(string(plaintext))
	return nil,nil
}
