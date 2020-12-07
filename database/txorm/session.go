package txorm

import (
	"context"
	"fmt"
	"xorm.io/xorm"
)

type Session struct {
	Id string
	// 事务相关
	context        context.Context
	contextCancel  context.CancelFunc
	beginTranslate bool
	// xorm session
	Sess *xorm.Session
	_db  *xorm.Engine

	mustCommit bool
	sql        string
	sqlwith    string
	query      map[string]interface{}
	args       []interface{}
	autoClose  bool
	tmps       map[string]interface{}
	withs      []string
}

func (this *Session) Session() *xorm.Session {
	return this.Sess
}

func (this *Session) Context() context.Context {
	if this.context == nil {
		this.context = context.Background()
	}
	ctx, cancel := context.WithCancel(this.context)
	this.context = ctx
	this.contextCancel = cancel

	return context.WithValue(this.context, CONTEXT_SESSION, this)
}

func (this *Session) Table(table interface{}) ISession {
	this.Sess.Table(table)
	return this
}

func (this *Session) With(name string) ISession {
	this.withs = append(this.withs, fmt.Sprintf("{{ %s .ctx }}", name))
	return this
}

func (this *Session) Begin() error {
	if this.beginTranslate {
		return nil
	}
	if err := this.Sess.Begin(); err != nil {
		return err
	}
	this.beginTranslate = true
	return nil
}

func (this *Session) Rollback() error {
	if !this.beginTranslate {
		return nil
	}
	if err := this.Sess.Rollback(); err != nil {
		return err
	}
	this.beginTranslate = false
	return this.Close()
}

func (this *Session) Commit() error {
	if !this.beginTranslate {
		return nil
	}

	closeFn := func() error {
		if err := this.Sess.Commit(); err != nil {
			return err
		}
		this.beginTranslate = false
		return this.Close()
	}

	if this.context != nil {
		if this.mustCommit {
			return closeFn()
		} else {
			go func() {
				select {
				case <-this.context.Done():
					_ = closeFn()
				}

			}()
		}

	} else {
		return closeFn()
	}
	return nil
}

func (this *Session) SetMustCommit(flag bool) ISession{
	this.mustCommit = flag
	return this
}

func (this *Session) Close() error {
	if this.autoClose {
		err := this.Sess.Close()
		if err != nil {
			return err
		}
	}
	this.beginTranslate = false
	err := this.ContextClose()
	if err != nil {
		return err
	}
	this.contextCancel = nil
	this.context = nil
	return nil
}

func (this *Session) ContextClose() error {
	if this.context != nil && this.contextCancel != nil {
		this.contextCancel()
	}
	return nil
}
