package test_xorm

import (
	"testing"
)

type BizWorkflow struct {
	Id      *string `json:"id" xorm:"Varchar(128) pk"`
	Created *int64  `json:"created" xorm:"created Int(14)"`
	Updated *int64  `json:"updated" xorm:"updated  Int(14)"`
	Weight  *int    `json:"weight" xorm:"weight  Int(9)"`
	Status  *string `json:"status" xorm:"status  Int(1)"`
	Name    *string `json:"name"`                 // 任务名称
	Remark  *string `json:"remark" xorm:"remark"` // 备注
	Uid     *string `json:"uid"`                  // 负责人
	State   *string `json:"state"`                // 任务状态 dict:SUB003
	Cstate  *string `json:"cstate"`               // 当前状态 dict:SUB004

}

func TestOrderBy(t *testing.T) {
	db := NewEngine()
	datas := make([]BizWorkflow, 0)
	err := db.SF(`
		select * from (select  * from (select * from biz_workflow order by state) s order by state desc) s  order by s.state
	`).Order("id", "-name").Find(&datas)
	if err != nil {
		panic(err)
	}
	//fmt.Println(tools.Str.JSONString(datas,true))
}
