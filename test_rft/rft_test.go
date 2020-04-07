package test_rft

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"reflect"
	"testing"
)

type Bean struct {
	Id      *string `json:"id" xorm:"Varchar(32) pk"`
	Created *int64  `json:"created" xorm:"created Int(14)"`
	Updated *int64  `json:"updated" xorm:"updated  Int(14)"`
	Weight  *int    `json:"weight" xorm:"weight  Int(9)"`
	Status  *int    `json:"status" xorm:"status  Int(1)"`
}

// 授权
type UserToken struct {
	Bean   `xorm:"extends"`
	Uid    *string `json:"uid"`
	Ops    *int64  `json:"ops"`    // 接口调用次数
	Type   *string `json:"type"`   // 授权类型 [pc]
	Expire *int64  `json:"expire"` // 到期时间
	Agent  *string `json:"agent"`  // User-Agent
}


func TestDeepSet(t *testing.T) {
	tok := UserToken{}
	tools.Rft.DeepSet(&tok, func(t reflect.Type, v reflect.Value, tf reflect.StructField) bool {
		fmt.Println(t.String(),v.Interface(),tf.Type)
		return true
	})
}
