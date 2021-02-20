package wxmp_test

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/wx/wxmp"
	"testing"
)

func TestSign(t *testing.T) {
	engine = wxmp.NewEngine(&wxmp.Option{
		Appsecret:     "a0e2253fc4b5fd649a3b6a91ec6ee14b",
		Appid:         "wx1eb1fc2b333e11d0",
		Mchid:         "1606344047",
		MchPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDp7m5IYn4Ps7RA
6w0q/2GGoXTocCfez3WYXC6N4Sz4uKbTiaCT5vaFaN1gBouMF+Y+KymjsbbCPyKb
zPOXZchO7PS/x8MqPVNoUhsmQ0Za64eE4rhyeTXk8DC+mUGIq/dr+1IDCxk3Jcwy
FFPvWsxMeIyHUknTfHDX7OiZIwua0j3yGtpXJSl8b3+rLqyl5/mvHODo2m7zuLGw
HetwsoczfG+S3Tst2ObpTwwmnsHqNki1eYdNRvMzzathoTGekX4CbAbaA/3Uyiif
ikPsqzUlF5TNhmdx2xeBKS3NnzCiH7QJ9uhziw/14PV+4QPW9QaLSWyzo5zEZmnH
keI6qvN5AgMBAAECggEAOVb4MUXzIGYsCL2UvLVccmOhBFh5bWPsryvpPV20aELa
oX9anOZABvNtT9xK2EwDY7mwuy8UXQicsxPJoZKRRsdONrQAYpoll6yIexiCZkbV
MP1huK19SGCXkiB+r8F3JEC2GYje5WWeZ6gT6teWvCfQSlshPXWDM2oemWq9rXTj
Oy+tCuiwZFSTa44Jtn+sxhUs9ypkdZgiZvSmhjAr5YzQR773+HLju4BmwYf23Dio
bHd2DYHRHjS9RpTmFQ6KrrSaTXNSmvRalY/FzCEauEeXckx67IsOPGCAhES2kHcx
1sRsL59S2Tsn1g4xhhcV3zoZgw+KI6TfxdQr3utPAQKBgQD6AorptQjuiPc09Kzm
swZf8dM7ORyG3rUeBBvGgX5oFX7pIMhWvmVN3ETP6BBYci++pYiusot1EEMtz2ix
xIiP+sCn+ThRv7wPSeDmMYWxscHTuJQg00Qn0CWiwvS9exKtymgTHv9R8Fa8m38y
Rn0ek9OFgHSFsrq+Q1v4L0GkhwKBgQDviUTYPc/PpuWOBowkuqLSw2BKt4dpGir4
qJTMSJhLpcGDDugFNacqKQWcFgjsdj+UsJSAz+XXYamrisqn+tqHD9cVaT9uqNUm
b4EnOwlbRORK3lFtfNZBKqjR2aBVPMKIk1i7ALqehdOjtJ/puBuq7QG/d+Jiv1Pn
At9y62yn/wKBgAWuIxvWOiK5R+yTFo6TSLTLWMJCtOw3iSPqcfsbnBSfUGfZj9Ow
tbqEI6gZnK11wrHxLt7RPavmN7CFwtovHe8vgksOtYHd+lbaldqFC4WTBVVbHzpz
slu5NfGxvj/D2RPLwnuUu7ZP4Jiea9Bnm5YjQ64H0h3rhqSmASPtZu9nAoGBAO0V
boChXWhoBlk2fct0tufo3QvW7z3F2rZXFT/EsohdPVVckaVmX1hJVfYRkS+KMpAW
3kVIgHNXhLn3G3J7xYNc2EOm8lOy45WxU6HiuvYND/BSb0HxB5dkg8eAoUL8aocH
YBFnPU7dooYrpwOLaEcbYlmCbR3TxVWm8EcsYVU9AoGAa1S53FQ4NrE/dnzJlPjJ
FN3FLVzF7l6PyZElRpwfKJYqGXcI9X5rcRnmp4azMS4EtFRLMwt2eusEdOijO1Op
cjZZI/bYxZ98pkBk2seRf/dJjsQ2D0bjoknS322Hwge6UFH/e1l6SOAAzhhT+CQt
KpbNjzdogCMvbtQ0cq75+kU=
-----END PRIVATE KEY-----
`,
		MchSeriesNo:   "7C74BEB297FF1D9D9DD8CD380CBD43DE79300C1B",
	})
	res,err := engine.Pay(&wxmp.PayOption{
		OutTradeNo:  "test_7FF1D9D9DD8CD38",
		NotifyUrl:   "http://zander123.cn/pay/test",
		Description: "测试订单",
		TotalPrice:  1,
		Openid:      "o_v1Z4y3jKyeeIG3B5_W52TRzLBQ",
	})
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(tools.JSONString(res,true))
}

//func TestEngy(t *testing.T) {
//	secretMessage := []byte("send reinforcements, we're going to advance")
//	rng := rand.Reader
//
//	cipherdata, err := EncryptOAEP(sha1.New(), rng, rsaPublicKey, secretMessage, nil)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
//		return
//	}
//
//	ciphertext := base64.StdEncoding.EncodeToString(cipherdata)
//	fmt.Printf("Ciphertext: %s\n", ciphertext)
//}