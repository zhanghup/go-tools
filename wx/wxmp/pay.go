package wxmp

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/zhanghup/go-tools"
	"time"
)

// Pay 支付
type PayOption struct {
	OutTradeNo string
	NotifyUrl  string
	Openid     string
	TotalPrice int // 支付金额

	Description *string
	TimeExpire  *int64
	Attach      *string
	Currency    *string // 支付货币
	GoodsTag    *string
}
type PayRes struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	PrepayId  string `json:"prepay_id"`
}

func (this *Engine) Pay(charge *PayOption) (*PayRes, error) {
	var m = map[string]interface{}{
		"appid":        this.opt.Appid,
		"mchid":        this.opt.Mchid,
		"out_trade_no": charge.OutTradeNo,
		"notify_url":   charge.NotifyUrl,
		"paper": map[string]interface{}{
			"openid": charge.Openid,
		},
	}
	if charge.TimeExpire != nil {
		m["time_expire"] = time.Unix(*charge.TimeExpire, 0).Format("2006-01-02T15:04:05+") + "08:00"
	}
	if charge.Description != nil {
		m["description"] = *charge.Description
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
	m["payer"] = map[string]interface{}{
		"openid": charge.Openid,
	}

	timestamp := time.Now().Unix()
	nonce_str := tools.StrOfRand(32)
	res, err := resty.New().
		R().
		SetBody(m).
		SetHeaders(map[string]string{
			"Authorization": fmt.Sprintf(`WECHATPAY2-SHA256-RSA2048 mchid="%s",serial_no="%s",nonce_str="%s",timestamp="%d",signature="%s"`,
				this.opt.Mchid, this.opt.MchSeriesNo, nonce_str, timestamp,
				this.PaySign("POST", "/v3/pay/transactions/jsapi", tools.Int64ToStr(timestamp), nonce_str, tools.JSONString(m))),
		}).Post("https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi")

	if err != nil {
		return nil, err
	}
	fmt.Println(res.String())
	rest := struct {
		PrepayId string `json:"prepay_id"`
	}{}
	err = json.Unmarshal(res.Body(), &rest)
	if err != nil {
		return nil, err
	}
	return &PayRes{
		Appid:     this.opt.Appid,
		Timestamp: timestamp,
		NonceStr:  nonce_str,
		PrepayId:  rest.PrepayId,
	}, nil
}

func (this *Engine) PaySign(method, url, t, nonce_str, body string) string {
	content := tools.StrFmt(`{{.method}}\n{{.url}}\n{{.t}}\n{{.rand}}\n{{.body}}\n`, map[string]interface{}{
		"method": method,
		"url":    url,
		"t":      t,
		"rand":   nonce_str,
		"body":   body,
	})
	return content
}
