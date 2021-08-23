package txorm

import (
	"context"
	"github.com/zhanghup/go-tools"
	"xorm.io/xorm"
)

const CONTEXT_SESSION = "context-session"

func (this *Engine) NewSession(autoClose bool, ctx ...context.Context) ISession {
	return this._newClearSession(autoClose, ctx...)
}

func (this *Engine) Session(ctx ...context.Context) ISession {
	return this._newSeesion(false, ctx...)
}

func (this *Engine) TS(fn func(sess ISession) error) error {
	return this.NewSession(true).TS(fn)
}

// Engine直接调用，自动结束session
func (this *Engine) SF(sql string, querys ...interface{}) ISession {
	sess := this.NewSession(true)
	return sess.SF(sql, querys...)
}

func (this *Engine) Engine() *xorm.Engine {
	return this.DB
}

func (this *Engine) _newSeesion(autoClose bool, ctx ...context.Context) ISession {
	newSession := this._newClearSession(autoClose, ctx...)

	if len(ctx) > 0 && ctx[0] != nil {
		c := ctx[0]
		v := c.Value(CONTEXT_SESSION)
		newSession.context = c
		if v == nil {
			return newSession
		} else {
			oldSession, ok := v.(*Session)
			if !ok {
				return newSession
			} else {
				oldSession.context = c
				return oldSession
			}
		}
	} else {
		return newSession
	}
}

func (this *Engine) _newClearSession(autoClose bool, ctx ...context.Context) *Session {
	s := &Session{
		id:             tools.UUID(),
		_db:            this.DB,
		sess:           this.DB.NewSession(),
		tmps:           this.tmps,
		tmpCtxs:        this.tmpCtxs,
		tmpWiths:       this.tmpWiths,
		autoClose:      autoClose,
		beginTranslate: false,
	}
	if len(ctx) > 0 {
		s.context = ctx[0]
	}
	return s
}
