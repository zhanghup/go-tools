package tgql_test

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

type DictItem struct {
	Id      string  `json:"id" xorm:"Varchar(128) pk"`
	Corp    string  `json:"corp" xorm:"notnull default('')"`
	Status  string  `json:"status" xorm:"status notnull Varchar(2) default('1')"`
	Weight  float64 `json:"weight" xorm:"weight default(0)"`
	Created int64   `json:"created" xorm:"created notnull Int(14)"`
	Updated int64   `json:"updated" xorm:"updated notnull Int(14)"`
	Code    string  `json:"code"`
	Name    string  `json:"name"`
	Value   string  `json:"value"`
}
