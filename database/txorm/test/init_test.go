package test_test

import (
	"context"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

var engine txorm.IEngine
var db *xorm.Engine

type User struct {
	Id      string  `json:"id" xorm:"pk"`
	Corp    string  `json:"corp" xorm:"notnull default('')"`
	Status  string  `json:"status" xorm:"status notnull Varchar(2) default('1')"`
	Weight  float64 `json:"weight" xorm:"weight default(0)"`
	Created int64   `json:"created" xorm:"created notnull Int"`
	Updated int64   `json:"updated" xorm:"updated notnull Int"`

	System string `json:"system" xorm:"'system' notnull"`           // 是否位系统账户
	Login  string `json:"login" xorm:"'login' notnull default('')"` // 允许登录的平台，等个的话，逗号隔开
	Type   string `json:"type" xorm:"'type' notnull"`               // 用户类型
	Dept   string `json:"dept" xorm:"'dept' notnull"`               // 所属部门
	Name   string `json:"name" xorm:"'name' notnull"`               // 用户名称
	Py     string `json:"py" xorm:"'py' notnull"`                   // 拼音
	Pinyin string `json:"pinyin" xorm:"'pinyin' notnull"`           // 拼音
	Mobile string `json:"mobile" xorm:"'mobile' notnull"`           // 联系电话

	Username string `json:"username" xorm:"'username' notnull"` // 用户名
	Password string `json:"password" xorm:"'password' notnull"` // 密码
	Salt     string `json:"-" xorm:"'salt' notnull"`            // 加盐
	Reset    string `json:"reset" xorm:"'reset' notnull"`       // 登录后是否需要修改密码

	Avatar      *string `json:"avatar" xorm:"'avatar'"`               // 头像
	IdCard      *string `json:"id_card" xorm:"'id_card'"`             // 身份证ID
	IdCardPhoto *string `json:"id_card_photo" xorm:"'id_card_photo'"` // 身份证ID
	Birth       *int64  `json:"birth" xorm:"'birth'"`                 // 出生日期
	Sex         *string `json:"sex" xorm:"'sex'"`                     // 人物性别
	Sn          *string `json:"sn" xorm:"'sn'"`                       // 工号
	Remark      *string `json:"remark" xorm:"'remark'"`               // 备注
}

func init() {
	e, err := txorm.NewXorm(txorm.Config{
		Uri:    "root:Zhang3611.@tcp(192.168.31.150:23306)/test2?charset=utf8",
		Driver: "mysql",
		Debug:  true,
	})
	if err != nil {
		tog.Error(err.Error())
		return
	}
	db = e
	engine = txorm.NewEngine(e)

	engine.TemplateFuncWith("users", func(ctx context.Context) string {
		return "select * from user"
	})

	engine.TemplateFuncCtx("corp", func(ctx context.Context) string {
		return "'ceaaeb6d-9f47-4ecb-ab4b-3247091229b7'"
	})

	err = engine.Sync(User{})
	if err != nil {
		tog.Error(err.Error())
		return
	}
	err = engine.Session(true).SF("delete from user").Exec()
	if err != nil {
		tog.Error(err.Error())
		return
	}

	for i := 0; i < 10; i++ {
		err := engine.Session(true).Insert(User{
			Id:   tools.IntToStr(i),
			Name: tools.IntToStr(i),
		})
		if err != nil {
			panic(err)
		}
	}
}
