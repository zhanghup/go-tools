package txorm

import (
	"context"
	"github.com/zhanghup/go-tools"
	"xorm.io/xorm"
)

const CONTEXT_SESSION = "context-session"

func (this *Engine) newSession(autoClose bool, ctx ...context.Context) ISession {
	return newSeesion(this.DB, autoClose, this.tmps, ctx...)
}

func (this *Engine) NewSession(ctx ...context.Context) ISession {
	return newSeesion(this.DB, false, this.tmps, ctx...)
}

func (this *Engine) TS(fn func(sess ISession) error) error {
	return this.newSession(true).TS(fn)
}

// Engine直接调用，自动结束session
func (this *Engine) SF(sql string, querys ...map[string]interface{}) ISession {
	sess := this.newSession(true)
	return sess.SF(sql, querys...)
}

func (this *Engine) With(name string) ISession {
	return this.newSession(true).With(name)
}

func (this *Engine) SessionContextClose(ctx context.Context) error {
	return this.newSession(true, ctx).ContextClose()
}

func newSeesion(db *xorm.Engine, autoClose bool, tmps map[string]interface{}, ctx ...context.Context) ISession {
	newSession := &Session{
		id:             tools.Str.Uid(),
		_db:            db,
		Sess:           db.NewSession(),
		tmps:           tmps,
		autoClose:      autoClose,
		beginTranslate: false,
		mustCommit:     false,
	}

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
				return oldSession
			}
		}
	} else {
		return newSession
	}
}
