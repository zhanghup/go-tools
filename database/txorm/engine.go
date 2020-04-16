package txorm

import (
	"context"
	"xorm.io/xorm"
)

type Engine struct {
	DB *xorm.Engine
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
	return &Session{Sess: this.DB.NewSession(), autoClose: false}
}

func (this *Engine) TS(fn func(sess *Session) error) error {
	sess := &Session{Sess: this.DB.NewSession(), autoClose: false}
	return sess.TS(fn)
}

func (this *Engine) SF(sql string, querys ...map[string]interface{}) *Session {
	sess := this.DB.NewSession()
	return (&Session{Sess: sess, autoClose: true}).SF(sql, querys...)
}