package tgql_test

import (
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgql"
	"testing"
	"time"
)

type Dict struct {
	Id       string  `json:"id" xorm:"Varchar(128) pk"`
	Corp     string  `json:"corp" xorm:"notnull default('')"`
	Status   string  `json:"status" xorm:"status notnull Varchar(2) default('1')"`
	Weight   float64 `json:"weight" xorm:"weight default(0)"`
	Created  int64   `json:"created" xorm:"created notnull Int(14)"`
	Updated  int64   `json:"updated" xorm:"updated notnull Int(14)"`
	Code     string  `json:"code" xorm:"'code' notnull"`         // 字典编码
	Name     string  `json:"name" xorm:"'name' notnull"`         // 字典名称
	Type     string  `json:"type" xorm:"'type' notnull"`         // 字典类型
	Disabled string  `json:"disabled" xorm:"'disabled' notnull"` // 是否禁止修改
	Remark   *string `json:"remark" xorm:"'remark'"`             // 备注
}

func TestObj(t *testing.T) {
	e, err := txorm.NewXorm(txorm.Config{
		Driver: "mysql",
		Uri:    "root:123@/test",
		Debug:  true,
	})
	if err != nil {
		panic(err)
	}
	if err := e.Ping(); err != nil {
		panic(err)
	}
	load := tgql.NewDataLoaden(e)

	for j := 0; j < 50; j++ {
		keys := []string{"SYS001", "SYS002", "SYS003", "SYS004", "SYS005", "SYS006"}

		for _, k := range keys {
			go func(i string) {
				dict := new(Dict)
				_, err := load.Object(dict, "select * from dict where code in :keys", nil, "Code", "").Load(i, dict)
				//fmt.Println( i)
				if err != nil {
					panic(err)
				}

				dict2 := new(Dict)
				_, err = load.Object(dict, "select * from dict where 1 = 1 and code in :keys", nil, "Code", "").Load(i, dict2)
				//fmt.Println( i)
				if err != nil {
					panic(err)
				}
			}(k)
		}
		time.Sleep(time.Microsecond)
	}

	time.Sleep(time.Second * 3)

}
