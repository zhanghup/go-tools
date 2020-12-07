package txorm

import (
	"context"
	"xorm.io/xorm"
)

type ISession interface {
	Context() context.Context
	Session() *xorm.Session
	Table(table interface{}) ISession
	With(name string) ISession
	Begin() error
	Rollback() error
	Commit() error
	Close() error
	ContextClose() error
	SetMustCommit(flag bool) ISession
	Id() string

	Find(bean interface{}) error
	Insert(bean ...interface{}) error
	Update(bean interface{}, condiBean ...interface{}) error
	Delete(bean interface{}) error
	SF(sql string, querys ...map[string]interface{}) ISession
	Page(index, size int, count bool, bean interface{}) (int, error)
	Page2(index, size *int, count *bool, bean interface{}) (int, error)
	TS(fn func(sess ISession) error) error
	Exec() error
}
