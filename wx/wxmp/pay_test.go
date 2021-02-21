package wxmp_test

import (
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
