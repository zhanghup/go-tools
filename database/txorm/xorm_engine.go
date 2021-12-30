package txorm

import (
	"context"
	"github.com/zhanghup/go-tools"
)

const CONTEXT_SESSION = "context-session"

func (this *Engine) SessionAuto(ctx ...context.Context) ISession {
	return this._session(true, ctx...)
}

func (this *Engine) Session(ctx ...context.Context) ISession {
	return this._session(false, ctx...)
}

func (this *Engine) Sync(beans ...interface{}) error {
	return this.DB.Sync2(beans...)
}

func (this *Engine) _session(autoClose bool, ctx ...context.Context) *Session {

	if len(ctx) > 0 && ctx[0] != nil {
		c := ctx[0]
		v := c.Value(CONTEXT_SESSION)
		if v != nil {
			oldSession, ok := v.(*Session)
			if ok {
				if !oldSession.sess.IsClosed() {
					return oldSession
				}
			}
		}
	}

	newSession := &Session{
		id:             tools.UUID(),
		_engine:        this,
		_db:            this.DB,
		sess:           this.DB.NewSession(),
		tmps:           this.tmps,
		tmpCtxs:        this.tmpCtxs,
		tmpWiths:       this.tmpWiths,
		autoClose:      autoClose,
		beginTranslate: false,
	}
	if len(ctx) > 0 {
		newSession.context = ctx[0]
	} else {
		c := context.Background()
		c = context.WithValue(c, CONTEXT_SESSION, newSession)
		newSession.context = c
	}
	return newSession
}
