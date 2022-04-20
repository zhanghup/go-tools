package wxmp

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/zhanghup/go-tools"
	"time"
)

// Pay 支付
type PayOption struct {
	OutTradeNo  string
	NotifyUrl   string
	Openid      string
	TotalPrice  int // 支付金额
	Description string

	TimeExpire *int64
	Attach     *string
	Currency   *string // 支付货币
	GoodsTag   *string
}
type PayRes struct {
	Appid     string `json:"appid"`
	Timestamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

func (this *Engine) PayPublicKeyRefresh() error {
	if this.opt.MchPublicKey != nil && time.Now().Unix() < this.opt.MchPublicCertTime {
		return nil
	}

	res, err := resty.New().R().SetHeaders(this.PayHeader("GET", "/v3/certificates", nil)).
		Get("https://api.mch.weixin.qq.com/v3/certificates")
	if err != nil {
		return err
	}

	stru := struct {
		Data []struct {
			SerialNo           string `json:"serial_no"`
			EncryptCertificate struct {
				AssociatedData string `json:"associated_data"`
				Ciphertext     string `json:"ciphertext"`
				Nonce          string `json:"nonce"`
			} `json:"encrypt_certificate"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(res.Body(), &stru)
	if err != nil {
		return err
	}

	if len(stru.Data) == 0 {
		return this.errorStr("当前商户未配置证书")
	}

	item := stru.Data[0]

	// 对编码密文进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(item.EncryptCertificate.Ciphertext)
	if err != nil {
		return this.error(err)
	}

	c, err := aes.NewCipher([]byte(this.opt.MchApiKey))
	if err != nil {
		return this.error(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return this.error(err)
	}

	nonceSize := gcm.NonceSize()
	if len(decodeBytes) < nonceSize {
		return this.errorStr("密文证书长度不够")
	}
	plaintext, err := gcm.Open(nil, []byte(item.EncryptCertificate.Nonce), decodeBytes, []byte(item.EncryptCertificate.AssociatedData))
	if err != nil {
		return this.error(err)
	}

	this.opt.MchPublicCertSn = item.SerialNo
	block_pub, _ := pem.Decode(plaintext)

	if block_pub == nil || block_pub.Type != "CERTIFICATE" {
		return this.errorStr("解码包含平台公钥的PEM块失败！")
	}
	pu, err := x509.ParseCertificate(block_pub.Bytes)
	if err != nil {
		return this.error(err)
	}

	this.opt.MchPublicKey = pu.PublicKey.(*rsa.PublicKey)
	this.opt.MchPublicCertTime = time.Now().Unix() + 3600
	return nil
}

func (this *Engine) PayHeader(method, path string, param map[string]any) map[string]string {
	p := ""
	if param != nil {
		p = tools.JSONString(param)
	}
	nonce_str := tools.StrOfRand(32)
	timestamp := time.Now().Unix()

	header := map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Authorization": fmt.Sprintf(`WECHATPAY2-SHA256-RSA2048 mchid="%s",serial_no="%s",nonce_str="%s",timestamp="%d",signature="%s"`,
			this.opt.Mchid, this.opt.MchSeriesNo, nonce_str, timestamp,
			this.PaySign(method, path, tools.Int64ToStr(timestamp), nonce_str, p)),
	}
	return header
}

func (this *Engine) Pay(charge *PayOption) (*PayRes, error) {
	var m = map[string]any{
		"appid":        this.opt.Appid,
		"mchid":        this.opt.Mchid,
		"out_trade_no": charge.OutTradeNo,
		"description":  charge.Description,
		"notify_url":   charge.NotifyUrl,
		"payer": map[string]any{
			"openid": charge.Openid,
		},
	}
	if charge.TimeExpire != nil {
		m["time_expire"] = time.Unix(*charge.TimeExpire, 0).Format("2006-01-02T15:04:05+") + "08:00"
	}
	if charge.Attach != nil {
		m["attach"] = *charge.Attach
	}
	if charge.GoodsTag != nil {
		m["goods_tag"] = *charge.GoodsTag
	}
	amount := map[string]any{
		"total": charge.TotalPrice,
	}
	if charge.Currency != nil {
		amount["currency"] = *charge.Currency
	}
	m["amount"] = amount

	res, err := resty.New().R().SetBody(m).SetHeaders(this.PayHeader("POST", "/v3/pay/transactions/jsapi", m)).
		Post("https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi")

	if err != nil {
		return nil, err
	}
	rest := struct {
		PrepayId string `json:"prepay_id"`
		Message  string `json:"message"`
	}{}
	err = json.Unmarshal(res.Body(), &rest)
	if err != nil {
		return nil, err
	}
	if rest.PrepayId == "" {
		return nil, errors.New(rest.Message)
	}
	timestamp := time.Now().Unix()
	nonce_str := tools.StrOfRand(32)
	pk := fmt.Sprintf("prepay_id=" + rest.PrepayId)
	return &PayRes{
		Appid:     this.opt.Appid,
		Timestamp: tools.Int64ToStr(timestamp),
		NonceStr:  nonce_str,
		Package:   pk,
		SignType:  "RSA",
		PaySign:   tools.Base64Enc(tools.SHA256WithRSA(fmt.Sprintf("%s\n%d\n%s\n%s\n", this.opt.Appid, timestamp, nonce_str, pk), this.opt.MchPrivateText)),
	}, nil
}

func (this *Engine) PayCancel(out_trade_no string) error {
	m := map[string]any{"mchid": this.opt.Mchid}
	path := "/v3/pay/transactions/out-trade-no/" + out_trade_no + "/close"
	_, err := resty.New().
		R().
		SetBody(m).
		SetHeaders(this.PayHeader("POST", path, m)).
		Post("https://api.mch.weixin.qq.com" + path)
	return err
}

type PayCallbackOption struct {
	Mchid          string `json:"mchid"`
	Appid          string `json:"appid"`
	OutTradeNo     string `json:"out_trade_no"`
	TransactionId  string `json:"transaction_id"`
	TradeType      string `json:"trade_type"`
	TradeState     string `json:"trade_state"`
	TradeStateDesc string `json:"trade_state_desc"`
	BankType       string `json:"bank_type"`
	Attach         string `json:"attach"`
	SuccessTime    string `json:"success_time"`
	Amount         struct {
		Total         int    `json:"total"`
		PayerTotal    int    `json:"payer_total"`
		Currency      string `json:"currency"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
	Payer struct {
		Openid string `json:"openid"`
	} `json:"payer"`
}

func (this *Engine) PayDecrypt(data []byte) (*PayCallbackOption, error) {
	err := this.PayPublicKeyRefresh()
	if err != nil {
		return nil, err
	}
	dataStu := struct {
		Resource struct {
			Ciphertext     string `json:"ciphertext"`
			AssociatedData string `json:"associated_data"`
			Nonce          string `json:"nonce"`
		} `json:"resource"`
	}{}
	err = json.Unmarshal(data, &dataStu)
	if err != nil {
		return nil, err
	}

	// 对编码密文进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(dataStu.Resource.Ciphertext)
	if err != nil {
		return nil, this.error(err)
	}

	c, err := aes.NewCipher([]byte(this.opt.MchApiKey))
	if err != nil {
		return nil, this.error(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, this.error(err)
	}

	nonceSize := gcm.NonceSize()
	if len(decodeBytes) < nonceSize {
		return nil, this.errorStr("密文证书长度不够")
	}
	plaintext, err := gcm.Open(nil, []byte(dataStu.Resource.Nonce), decodeBytes, []byte(dataStu.Resource.AssociatedData))
	if err != nil {
		return nil, this.error(err)
	}

	result := PayCallbackOption{}
	err = json.Unmarshal(plaintext, &result)
	return &result, err
}

func (this *Engine) PaySign(method, url, t, nonce_str, body string) string {
	content := tools.StrFmt("{{.method}}\n{{.url}}\n{{.t}}\n{{.rand}}\n{{.body}}\n", map[string]any{
		"method": method,
		"url":    url,
		"t":      t,
		"rand":   nonce_str,
		"body":   body,
	})
	s := tools.SHA256WithRSA(content, this.opt.MchPrivateText)
	ss := tools.Base64Enc(s)
	return ss
}
