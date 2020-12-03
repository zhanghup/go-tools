package test_xorm

type Bean struct {
	Id      *string `json:"id" xorm:"Varchar(32) pk"`
	Created *int64  `json:"created" xorm:"created Int(14)"`
	Updated *int64  `json:"updated" xorm:"updated  Int(14)"`
	Weight  *int    `json:"weight" xorm:"weight  Int(9)"`
	Status  *int    `json:"status" xorm:"status  Int(1)"`
}

// 数据字典
type Dict struct {
	Bean `xorm:"extends"`

	Code   *string `json:"code" xorm:"unique"`
	Name   *string `json:"name"`
	Remark *string `json:"remark"`
}
