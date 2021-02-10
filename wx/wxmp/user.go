package wxmp

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/zhanghup/go-tools"
)

type ViewCode2Session struct {
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	SessionKey string `json:"session_key"`
}

func (this *Engine) Code2Session(code string) (*ViewCode2Session, error) {
	query := map[string]interface{}{
		"appid":     this.opt.Appid,
		"appsecret": this.opt.Appsecret,
		"code":      code,
	}
	res, err := this.get(tools.StrFmt(`/sns/jscode2session?appid={{ .appid }}&secret={{ .appsecret }}&js_code={{ .code }}&grant_type=authorization_code`, query))
	if err != nil {
		return nil, this.error(err)
	}
	resu := ViewCode2Session{}
	err = json.Unmarshal(res, &resu)
	if err != nil {
		return nil, this.error(err)
	}
	return &resu, err
}

type ViewUserInfo struct {
	Openid    string `json:"openId"`
	Nickname  string `json:"nickName"`
	Gender    int `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	Unionid   string `json:"unionId"`
	Language string `json:"language"`
	Watermark struct {
		Appid     string `json:"appid"`
		Timestamp int64  `json:"timestamp"`
	} `json:"watermark"`
}

func (this *Engine) UserInfoDecrypt(ssk, rawData, encryptedData, signature, iv string) (*ViewUserInfo, error) {

	raw := sha1.Sum([]byte(rawData + ssk))
	if signature != hex.EncodeToString(raw[:]) {
		return nil, this.errorStr("解密用户信息失败")
	}
	resdata, err := this.DecryptUserData(ssk, encryptedData, iv)
	if err != nil {
		return nil, this.error(err)
	}

	result := ViewUserInfo{}
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, this.error(err)
	}
	return &result, nil
}

func (this *Engine) DecryptUserData(sskStr, ciphertextStr, ivStr string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(sskStr)
	if err != nil {
		return nil, err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextStr)
	if err != nil {
		return nil, err
	}

	iv, err := base64.StdEncoding.DecodeString(ivStr)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	size := aes.BlockSize
	iv = iv[:size]

	if len(ciphertext) < size {
		return nil, this.errorStr("ciphertext too short")
	}

	if len(ciphertext)%size != 0 {
		return nil, this.errorStr("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	plaintext := ciphertext[:]
	ln := len(plaintext)
	pad := int(plaintext[ln-1])
	if pad < 1 || pad > 32 {
		pad = 0
	}
	result := plaintext[:(ln - pad)]
	return result, nil

}
