package txorm

import (
	"context"
	"sync"
	"xorm.io/xorm"
)

type Engine struct {
	DB      *xorm.Engine
	tmps    map[string]interface{}
	tmpsync sync.RWMutex
}

const CONTEXT_SESSION = "context-session"

func (this *Engine) NewSession(ctx ...context.Context) *Session {
	var sess *Session
	if len(ctx) > 0 && ctx[0] != nil {
		c := ctx[0]
		v := c.Value(CONTEXT_SESSION)
		if v == nil {
			sess = this.session()
			sess.context = c
		} else {
			s, ok := v.(*Session)
			if !ok {
				sess = this.session()
			}
			sess = s
		}
	} else {
		sess = this.session()
	}
	return sess
}

func (this *Engine) session() *Session {
	return &Session{Sess: this.DB.NewSession(), tmps: this.tmps, autoClose: false}
}

func (this *Engine) TS(fn func(sess *Session) error) error {
	return this.NewSession().TS(fn)
}

func (this *Engine) SF(sql string, querys ...map[string]interface{}) *Session {
	sess := this.NewSession()
	sess.autoClose = true
	return sess.SF(sql, querys...)
}

func (this *Engine) With(name string) *Session {
	return this.NewSession().With(name)
}
