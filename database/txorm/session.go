package txorm

import (
	"context"
	"sync"
	"xorm.io/xorm"
)

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

func (this *Session) S() *xorm.Session {
	return this.sess
}

func (this *Session) E() *xorm.Engine {
	return this._db
}

func (this *Session) Ctx() context.Context {
	if this.context == nil {
		this.context = context.Background()
	}
	return context.WithValue(this.context, CONTEXT_SESSION, this)
}

func (this *Session) Table(table interface{}) ISession {
	this.sess.Table(table)
	return this
}

func (this *Session) With(name string) ISession {
	this.withs = append(this.withs, name)
	return this
}

func (this *Session) Begin() error {
	this._sync.Lock()
	defer this._sync.Unlock()
	if this.beginTranslate {
		return nil
	}
	if err := this.sess.Begin(); err != nil {
		return err
	}
	this.beginTranslate = true
	return nil
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
