package wxmp_test

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {

	//res,err := engine.Pay(&wxmp.PayOption{
	//	OutTradeNo:  "test_7FF1D9D9DD8CD383",
	//	NotifyUrl:   "http://zander123.cn/pay/test",
	//	Description: "测试订单",
	//	TotalPrice:  1,
	//	Openid:      "o_v1Z4y3jKyeeIG3B5_W52TRzLBQ",
	//})
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(tools.JSONString(res,true))
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

func Test2(t *testing.T) {
	str := `BvYaveAi0/g/rZtxFG3ETRcEv6uJbyINwJc9+J9j/xFYR24w1y6ZgTXhjtxTHIjtpVegmf37gbUztPy1SNPjWTD0ymHkuO9s566sGkiDp5fDA31Wb0WEnDb4NezmVfnXP5gdtEI8fpEkpCaw+OpiTjr74hZuCuvDGhf2ibi5ox95nyk7U4lLjXWOdF43gwQJmILywpA4DgNlAMeGsaPVzikZEuV36e0mI7Ji8H9TT7Y+txyEz+B1jWxZQ7f8IjUpCLLFmjJRgb+QCwd5K54KPK3nLJN7hXq4vIBBy/bYPWB3zagLgQCYjxXbGjIkuULmCJ0cDVFgSQLzP4w970ZT7WxLLtlAyHprq07XXKNb3YdlnW8RnY4F+uWb5mp0OSCEwYBn6x0Sv1rcmgw2G+maH3d1QIw3AxCxipnLi0bv1GyJqPMFJcP/Q0G3ypMe4CLnhgr76IhrsyJI2fXZuYM/hdjvBdeZMnelUbb3L7boDs8wulmjVn47i35wuQusDmmeRULeCP3tLqIqkcY0FPbJUwbqdsbxII8SQV7dnU/+5o1nWm+9OjzkdQpy1IFSix6GHQZVWJVY3Y9SIUMoT6dmqRISabmul7vz0vI+zNAmXWxwVRafLNiGrK10sJlRpqOgr/WNvCr+gvZMXglgCWBLHYdE4moso+vrGvL8cMVLrnvT5E8Or8rKioqj6H+JZXM9g7padB7puICuycHJn0domciwer8baVs1r/cMSYzAlnWNjWb3mhgH31mjM37W+BV3HKz011AmrHnA4MDLq8ZtRo5RbvGR51SQS68qJKuVzasuEM9FXVhPJagFQONP9UP7XxHE0nui3GOCZ0YwYFysbFTlBzqh9VX2gbxDpttFhV0pxjI8+zMSKxtYGK08BmnCQxM8AASYw6saIdTwqdFbFi6IdbXdHjI04e8vuqCmFqLNrniUSKQX9DlH/e5J1t6Gz9oVOudOSyn4y6e9OfKc29cmtHcNVWhDPTtJSiQaRfM9+vorSauafbTrMaILRv41yQ/0epHGSCamFjo44lCc+QhLrHZ6n0iBE/Ofiqjp8o4EdhUL/8aTUzRRCa5kWwU5GQtEypuLVK1ZatLcApBW3vAfLyiM1WQELjK39tDeICU6oyTxzRrb0L32X5ouikSS0x2Z2RdrijHJ7f1ESocTFPKQpN3DjpYNKW6MuDhtMDkX1ONggdfMu0sdWDUjtsyzcpMgb2LqALuwq2iG4EOVO3LLR9zF/lADJNRE+8AohjJkutKZXi0XI0YAGCPwo1Ab04G7S1fHihnyx60jqx1cmaVppQ8cZoVd9O6mIXjdY+Cd8JBZMSeuiR7v/SVgO5d7TuArA/CL8YBw0xz1X6sjv0veJ+aaoAfhU8g+JZnpk9SjoDieolKhoqfUqOzlzM4UkEXhmjAQ2CBSXgHCyLFKEstgSaBGAgcNNHh5ywxu72VF3xCYpDifHceHOIyFAgNEqKVGVrob0MHROv2khLdvgJcuKh9o/ZhJJgVUVnFoEj0WZybWh4gfBnNcLiNDrS41AMdOSIM4Q+WgiDgeqodBWveVQDvap3mrkFTIODRh18JFOjOhCFYaNbCzFS3kVGdLURkncmF4At9lyD8BpSYKiqecBz1CEjoYzij9ZjjmPJKVa82qPTGDQvmgYiHE/sizoAV1iKNEspd42AcnVPrunI367JVIthLCpedbMbFQFxz3v4wouLPTCVqzlQBr1Qo9kijOGMDELIknMedllLvPcc3JRoSVNuXJa871c84UE7U93yfiWzV5rzJ8+S5nKFACIzMhHQSa7c1qlzzN7Ubh8VzJqYT2OdhRbrqWgQhG3cba1XYTD1WNfaAJ1/D5AiX+sm5IE9tk2S5U4ndbri6KyreJqCGugg==`
	no := "b023e957f62e"
	Associated_data:="Associated_data"

	decodeBytes, _ := base64.StdEncoding.DecodeString(str)


	c, err := aes.NewCipher([]byte(""))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	nonceSize := gcm.NonceSize()
	if len(decodeBytes) < nonceSize {
		fmt.Println("密文证书长度不够")
		return
	}
	if Associated_data != "" {
		plaintext, _ := gcm.Open(nil, []byte(no), decodeBytes, []byte(Associated_data))
		fmt.Println(plaintext)
	} else {
		plaintext, _ := gcm.Open(nil, []byte(no), decodeBytes, nil)
		fmt.Println(plaintext)
	}
}
