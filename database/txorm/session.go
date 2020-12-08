package txorm

import (
	"context"
	"fmt"
	"xorm.io/xorm"
)

type Session struct {
	id             string
	context        context.Context
	beginTranslate bool
	// xorm session
	sess *xorm.Session
	_db  *xorm.Engine

	sql       string
	sqlwith   string
	query     map[string]interface{}
	args      []interface{}
	autoClose bool
	tmps      map[string]interface{}
	withs     []string
}

func (this *Session) Session() *xorm.Session {
	return this.sess
}

func (this *Session) Context() context.Context {
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
	this.withs = append(this.withs, fmt.Sprintf("{{ %s .ctx }}", name))
	return this
}

func (this *Session) Begin() error {
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
	return this.Close()
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

func (this *Session) Close() error {
	return this.sess.Close()
}

func (this *Session) Id() string {
	return this.id
}
