package txorm

import (
	"context"
)

const CONTEXT_SESSION = "context-session"

func (this *Engine) NewSession(ctx ...context.Context) *Session {

	newSession := &Session{
		Sess:           this.DB.NewSession(),
		tmps:           this.tmps,
		autoClose:      false,
		beginTranslate: false,
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

func (this *Engine) TS(fn func(sess *Session) error) error {
	return this.NewSession().TS(fn)
}

// Engine直接调用，自动结束session
func (this *Engine) SF(sql string, querys ...map[string]interface{}) *Session {
	sess := this.NewSession()
	sess.autoClose = true
	return sess.SF(sql, querys...)
}

func (this *Engine) With(name string) *Session {
	return this.NewSession().With(name)
}
