package txorm

import (
	"context"
	"xorm.io/xorm"
)

type Session struct {
	context   context.Context
	Sess      *xorm.Session
	sql       string
	query     map[string]interface{}
	args      []interface{}
	autoClose bool
}

func (this *Session) Context() context.Context {
	if this.context == nil {
		return context.Background()
	}

	return context.WithValue(this.context, CONTEXT_SESSION, this)
}

func (this *Session) Table(table interface{}) *Session {
	this.Sess.Table(table)
	return this
}
