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
	})
}
