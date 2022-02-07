package txorm

import (
	"context"
	"github.com/zhanghup/go-tools"
)

func (this *Engine) TS(ctx context.Context, fn func(ctx context.Context, sess ISession) error) error {

	var sess *Session
	v := ctx.Value(CONTEXT_SESSION)
	if v != nil {
		sessOld, ok := v.(*Session)
		if ok {
			if !sessOld.sess.IsClosed() {
				sess = sessOld
			}
		}
	}

	commit := false
	if sess == nil {
		commit = true

		sess = &Session{
			id:        tools.UUID(),
			_engine:   this,
			_db:       this.DB,
			sess:      this.DB.NewSession(),
			tmps:      this.tmps,
			tmpCtxs:   this.tmpCtxs,
			tmpWiths:  this.tmpWiths,
			autoClose: false,
		}

		ctx = context.WithValue(ctx, CONTEXT_SESSION, sess)
		sess.context = ctx
		sess.sess.Begin()
	}

	err := fn(ctx, sess)
	if err != nil {
		_ = sess.sess.Rollback()
		return err
	}

	if commit {
		_ = sess.sess.Commit()
		return sess.Close()
	}

	return nil

}
