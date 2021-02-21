package wxmp

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
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

func (this *Engine) PayHeader(path string, param map[string]interface{}) map[string]string {
	nonce_str := tools.StrOfRand(32)
	timestamp := time.Now().Unix()
	header := map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Authorization": fmt.Sprintf(`WECHATPAY2-SHA256-RSA2048 mchid="%s",serial_no="%s",nonce_str="%s",timestamp="%d",signature="%s"`,
			this.opt.Mchid, this.opt.MchSeriesNo, nonce_str, timestamp,
			this.PaySign("POST", path, tools.Int64ToStr(timestamp), nonce_str, tools.JSONString(param))),
	}
	return header
}

func (this *Engine) Pay(charge *PayOption) (*PayRes, error) {
	var m = map[string]interface{}{
		"appid":        this.opt.Appid,
		"mchid":        this.opt.Mchid,
		"out_trade_no": charge.OutTradeNo,
		"description":  charge.Description,
		"notify_url":   charge.NotifyUrl,
		"payer": map[string]interface{}{
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
	amount := map[string]interface{}{
		"total": charge.TotalPrice,
	}
	if charge.Currency != nil {
		amount["currency"] = *charge.Currency
	}
	m["amount"] = amount

	res, err := resty.New().R().SetBody(m).SetHeaders(this.PayHeader("/v3/pay/transactions/jsapi", m)).
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
		PaySign:   tools.Base64Enc(tools.SHA256WithRSA(fmt.Sprintf("%s\n%d\n%s\n%s\n", this.opt.Appid, timestamp, nonce_str, pk), this.opt.MchPrivateKey)),
	}, nil
}

func (this *Engine) PayCancel(out_trade_no string) error {
	m := map[string]interface{}{"mchid": this.opt.Mchid}
	path := "/v3/pay/transactions/out-trade-no/" + out_trade_no + "/close"
	_, err := resty.New().
		R().
		SetBody(m).
		SetHeaders(this.PayHeader(path, m)).
		Post("https://api.mch.weixin.qq.com" + path)
	return err
}

type PayCallbackOption struct {
	TransactionId string `json:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no"`
	Amount        struct {
		PayerTotal    int    `json:"payer_total"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
}

func (this *Engine) PayDecrypt(data []byte) (*PayCallbackOption, error) {
	dataStu := struct {
		Resource struct {
			Ciphertext     string `json:"ciphertext"`
			AssociatedData string `json:"associated_data"`
			Nonce          string `json:"nonce"`
		} `json:"resource"`
	}{}
	err := json.Unmarshal(data, &dataStu)
	if err != nil {
		return nil, err
	}

	ct, _ := tools.Base64Dec(dataStu.Resource.Ciphertext)
	nc := []byte(dataStu.Resource.Nonce)
	block, err := aes.NewCipher([]byte(this.opt.MchPrivateKey))
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nc, ct, []byte(dataStu.Resource.AssociatedData))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(plaintext))
	return nil, nil
}

func (this *Engine) PaySign(method, url, t, nonce_str, body string) string {
	content := tools.StrFmt("{{.method}}\n{{.url}}\n{{.t}}\n{{.rand}}\n{{.body}}\n", map[string]interface{}{
		"method": method,
		"url":    url,
		"t":      t,
		"rand":   nonce_str,
		"body":   body,
	})
	s := tools.SHA256WithRSA(content, this.opt.MchPrivateKey)
	return tools.Base64Enc(s)
}
