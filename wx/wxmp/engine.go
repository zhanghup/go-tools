package wxmp

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Option struct {
	Appid             string          `yaml:"appid"`
	Appsecret         string          `yaml:"appsecret"`
	Mchid             string          `yaml:"mchid"`
	MchPrivateText    string          `yaml:"mch_private_text"`
	MchSeriesNo       string          `yaml:"mch_series_no"`
	MchApiKey         string          `yaml:"mch_api_key"`
	MchPublicKey      *rsa.PublicKey  `yaml:"-"`
	MchPrivateKey     *rsa.PrivateKey `yaml:"-"`
	MchPublicCertSn   string          `yaml:"-"`
	MchPublicCertTime int64           `yaml:"-"`
}

const HOST = "https://api.weixin.qq.com"

type IEngine interface {
	Code2Session(code string) (*ViewCode2Session, error)
	UserInfoDecrypt(ssk, rawData, encryptedData, signature, iv string) (*ViewUserInfo, error)
	UserMobileDecrypt(ssk, encryptedData, iv string) (*ViewUserMobile, error)

	Pay(charge *PayOption) (*PayRes, error)
	PayCancel(out_trade_no string) error
	PayDecrypt(data []byte) (*PayCallbackOption, error)
}

type Engine struct {
	opt *Option
}

func NewEngine(opt *Option) IEngine {
	e := Engine{opt: opt}
	block, _ := pem.Decode([]byte(e.opt.MchPrivateText))
	if block == nil {
		panic("私钥解码失败")
	}
	priKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	e.opt.MchPrivateKey = priKey.(*rsa.PrivateKey)
	return &e
}

func (this *Engine) get(url string) ([]byte, error) {
	res, err := resty.New().R().Get(HOST + url)
	if err != nil {
		return nil, this.error(err)
	}

	stru := struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}{}
	err = json.Unmarshal(res.Body(), &stru)
	if err != nil {
		return nil, this.error(err)
	}
	if stru.Errcode != 0 {
		return nil, this.error(fmt.Errorf("Errcode: %d, Errmsg: %s", stru.Errcode, stru.Errmsg))
	}
	return res.Body(), nil
}

func (this *Engine) error(err error) error {
	return fmt.Errorf("【微信小程序】 %s", err.Error())
}

func (this *Engine) errorStr(err string) error {
	return this.error(errors.New(err))
}
