package wxmp_test

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/wx/wxmp"
	"testing"
)

var engine wxmp.IEngine

func TestCode2Session(t *testing.T) {
	res, err := engine.Code2Session("081BCv1w3UvXPV2alX3w3xbsnq3BCv1D")
	if err != nil {
		panic(err)
	}
	fmt.Println(tools.JSONString(res))
}

func TestUserInfo(t *testing.T) { // DecryptUserInfo
	res,err := engine.UserInfoDecrypt(
		"LuBqfEjsP3BL/T+G17Zs5Q==",
		"{\"nickName\":\"zander\",\"gender\":1,\"language\":\"zh_CN\",\"city\":\"Jiaxing\",\"province\":\"Zhejiang\",\"country\":\"China\",\"avatarUrl\":\"https://thirdwx.qlogo.cn/mmopen/vi_32/DYAIOgq83eoqV4knJupHGCkJY6MUEia4A6Ye3WEUMUGyrfXunbwib0T5ZgNXR770aUvvvWHDDK6IpRiaUFl8DxYaQ/132\"}",
		"vT9EmN8mkHcsWN7XSuTOZb0tpn1I8NdNEqKxwMjNbvLUc2CmC19iVH+Vx8aSVMqLfk70n+mXBzzhphx4jrWmPtu37RyViwEoL8gwX2i4DVm5BIecFecHm1hfSWVQ57YPeRqIwkMZFeLwl2csQv5XdLgsIlZHGX9E1NYzoNUu1d/J+cE2xYLaNsnA9QFwQBA6oaK4xYg3ymP+ulpVfAxTXIVOzWCcM2knTw7RNcdJEs7w+BFILh29vofh1VXoxe7MddczYKznfcHjOJTSqatROVCunE38qLVPnzRw51Iolzp4Vpnn4jYYPt9VqNHyNU4/0HdrUm4sIeHmZgnH83Bo07c40Sz/ucj5VR4HOEZKdF8jYd1baw3m3AzmlZxY6Jg5pg3jZdXRJpcTJXG9b8bHXa6ecOWPjHLRX68CJi6GgQ8lbkvDgZFK9BBNXRhGl1PXjcYak3ONHIlvzD0ALfzytiL3fHjYY9MWePkIfxfKnm0=",
		"208a703640cd25aadb6fea59849465ddd83fd212",
		"rJiDUhsBsefUc/BkbIOC5g==",
	)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(tools.JSONString(res,true))
}

func TestUserMobile(t *testing.T) {
	res,err := engine.UserMobileDecrypt("sVCIMciiCTYqT+GG8iXTdw==","keuXD8sK1gjO6oc6xBDpGTyxf8PKAnHo58cLXuHRBQG4KWy5oSsbEXG4l6jtiZkt+fM1dxAI4trwOv2wQbEA/WduHkFMLJ+mz1BJCr4Pf0mOnRgJo28zRCR8S35pJE533ZPCLaXvqsFT6IfdSDsUg8LJoADgKTn+FBxxyK3V4NnBo6YP7QYZcBPd22T84StC84hoh6nCJesW50ey541bjQ==","cUY46qzDanC5D4z5gbSO4A==")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(tools.JSONString(res,true))
}

func init() {
	engine = wxmp.NewEngine(&wxmp.Option{
		Appsecret: "a0e2253fc4b5fd649a3b6a91ec6ee14b",
		Appid:     "wx1eb1fc2b333e11d0",
		Mchid:     "1606344047",
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
		MchSeriesNo: "7C74BEB297FF1D9D9DD8CD380CBD43DE79300C1B",
	})
}
