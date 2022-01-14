package badger

import "context"

func (this *Engine) Ts(fn func(sess ISession) error) error {
	s := NewSession(this.db).Begin()
	err := fn(s)
	if err != nil {
		s.Rollback()
		return nil
	}
	return s.Commit()
}

func (this *Engine) Sess(ctx ...context.Context) ISession {
	return NewSession(this.db, ctx...)
}
