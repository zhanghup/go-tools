package txorm

import (
	"context"
	"xorm.io/xorm"
)

type ISession interface {
	Id() string
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
	SF(sql string, querys ...map[string]interface{}) ISession
	/*
		sql = "select * from user where a = ? and b = ?"
		querys = []interface{}{"a","b"}
	 */
	SF2(sql string, querys ...interface{}) ISession
	Page(index, size int, count bool, bean interface{}) (int, error)
	Page2(index, size *int, count *bool, bean interface{}) (int, error)
	TS(fn func(sess ISession) error, commit ...bool) error
	Exec() error
}
