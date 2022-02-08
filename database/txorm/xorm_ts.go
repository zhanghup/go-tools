package txorm

import (
	"context"
)

func (this *Engine) TS(ctx context.Context, fn func(ctx context.Context, sess ISession) error) error {

	sess := this.session(false, false, ctx)
	commit := sess.isNew
	if commit {
		_ = sess.sess.Begin()
	}
	err := fn(sess.context, sess)
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
