package txorm

import (
	"context"
	"fmt"
	"xorm.io/xorm"
)

type Session struct {
	context        context.Context
	Sess           *xorm.Session
	sql            string
	sqlwith        string
	query          map[string]interface{}
	args           []interface{}
	autoClose      bool
	tmps           map[string]interface{}
	withs          []string
	beginTranslate bool
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
	if this.autoClose {
		return this.Sess.Close()
	}
	return nil
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
		if this.autoClose {
			return this.Sess.Close()
		}
		return nil
	}

	if this.context != nil {
		go func() {
			<-this.context.Done()
			_ = closeFn()
		}()
	} else {
		return closeFn()
	}
	return nil
}
