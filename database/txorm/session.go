package txorm

import (
	"context"
	"fmt"
	"xorm.io/xorm"
)

type Session struct {
	context   context.Context
	Sess      *xorm.Session
	sql       string
	sqlwith   string
	query     map[string]interface{}
	args      []interface{}
	autoClose bool
	tmps      map[string]interface{}
	withs     []string
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

func (this *Session) With(name string) *Session {
	this.withs = append(this.withs, fmt.Sprintf("{{ %s .ctx }}", name))
	return this
}
