package wxmp

import (
	"fmt"
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
}
type PayRes struct {
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
	amount := map[string]interface{}{
		"total": charge.TotalPrice,
	}
	if charge.Currency != nil {
		amount["currency"] = *charge.Currency
	}
	m["amount"] = amount

	fmt.Println(this.PaySign("POST", "/test", "{}"))

	return nil, nil
}

func (this *Engine) PaySign(method, url, body string) string {
	content := tools.StrFmt(`{{.method}}\n{{.url}}\n{{.t}}\n{{.rand}}\n{{.body}}\n`, map[string]interface{}{
		"method": method,
		"url":    url,
		"t":      time.Now().Unix(),
		"rand":   tools.StrOfRand(32),
		"body":   body,
	})
	return content
}
