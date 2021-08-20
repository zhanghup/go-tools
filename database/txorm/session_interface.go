package txorm

import (
	"context"
	"xorm.io/xorm"
)

type ISession interface {
	Id() string
	SetId(id string)

	E() *xorm.Engine
	S() *xorm.Session
	Ctx() context.Context
	With(name string) ISession
	Begin() error
	Rollback() error
	Commit() error
	AutoClose() error

	Table(table interface{}) ISession
	Order(order ...string) ISession
	Find(bean interface{}) error
	Get(bean interface{}) (bool, error)
	Insert(bean ...interface{}) error
	Update(bean interface{}, condiBean ...interface{}) error
	Delete(bean interface{}) error
	/*
		示例1：
			sql = "select * from user where a = ? and b = ?"
			querys = []interface{}{"a","b"}
		示例2：
			sql = "select * from user where a = :a and b = ?"
			querys = []interface{}{"b",map[string]interface{}{"a":"a"}}
	*/
	SF(sql string, querys ...interface{}) ISession
	Page(index, size int, count bool, bean interface{}) (int, error)
	Page2(index, size *int, count *bool, bean interface{}) (int, error)
	TS(fn func(sess ISession) error, commit ...bool) error
	Exec() error
	Count() (int64, error)
	Int() (int, error)
	Int64() (int64, error)
	Float64() (float64, error)
	String() (string, error)
	Strings() ([]string, error)
	Exists() (bool, error)
	Map() ([]map[string][]byte, error)
}
