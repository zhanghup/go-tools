package wxmp_test

import (
	"github.com/zhanghup/go-tools/wx/wxmp"
	"testing"
)

func TestSign(t *testing.T) {
	engine = wxmp.NewEngine(&wxmp.Option{
		Appsecret: "a0e2253fc4b5fd649a3b6a91ec6ee14b",
		Appid:     "wx1eb1fc2b333e11d0",
	})
	engine.Pay(nil)
}
