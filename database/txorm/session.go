package txorm

import (
	"context"
	"sync"
	"xorm.io/xorm"
)

type ISession interface {
	Id() string
	SetId(id string)

	Ctx() context.Context
	Begin()
	Rollback() error
	Commit() error
	AutoClose() error

	Find(bean interface{}) error
	Get(bean interface{}) (bool, error)

	Insert(bean ...interface{}) error
	Update(bean interface{}, condiBean ...interface{}) error
	Delete(bean interface{}) error
	TS(fn func(sess ISession) error, commit ...bool) error
	Exec() error
	/*
		示例1：
			sql = "select * from user where a = ? and b = ?"
			querys = []interface{}{"a","b"}
		示例2：
			sql = "select * from user where a = :a and b = ?"
			querys = []interface{}{"b",map[string]interface{}{"a":"a"}}
	*/
	SF(sql string, querys ...interface{}) ISession
	Order(order ...string) ISession

	Page(index, size int, count bool, bean interface{}) (int, error)
	Page2(index, size *int, count *bool, bean interface{}) (int, error)
	Count() (int64, error)
	Int() (int, error)
	Int64() (int64, error)
	Float64() (float64, error)
	String() (string, error)
	Strings() ([]string, error)
	Exists() (bool, error)
	Map() ([]map[string]interface{}, error)
}

type Session struct {
	id             string
	context        context.Context
	beginTranslate bool
	// xorm session
	sess  *xorm.Session
	_db   *xorm.Engine
	_sync sync.Mutex

	sql       string
	query     map[string]interface{}
	args      []interface{}
	autoClose bool

	tmps     map[string]interface{}
	tmpWiths map[string]interface{}
	tmpCtxs  map[string]interface{}

	withs   []string
	orderby []string
}

func (this *Session) Ctx() context.Context {
	if this.context == nil {
		this.context = context.Background()
	}
	return context.WithValue(this.context, CONTEXT_SESSION, this)
}


func (this *Session) begin() error {
	if err := this.sess.Begin(); err != nil {
		return err
	}
	return nil
}

func (this *Session) Begin() {
	this._sync.Lock()
	defer this._sync.Unlock()
	if this.beginTranslate {
		return
	}

	this.beginTranslate = true
	return
}

func (this *Session) Rollback() error {
	if !this.beginTranslate {
		return nil
	}
	if err := this.sess.Rollback(); err != nil {
		return err
	}
	this.beginTranslate = false
	return this.AutoClose()
}

func (this *Session) Commit() error {
	if !this.beginTranslate {
		return nil
	}

	if err := this.sess.Commit(); err != nil {
		return err
	}
	this.beginTranslate = false
	return nil

}

// 由engine直接进入的方法，需要自动关闭session
func (this *Session) AutoClose() error {
	err := this.Commit()
	if err != nil {
		return err
	}
	if this.sess.IsClosed() {
		return nil
	}
	return this.sess.Close()
}

func (this *Session) Id() string {
	return this.id
}

func (this *Session) SetId(id string) {
	this.id = id
}
