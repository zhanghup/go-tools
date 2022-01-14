package buntdb

import (
	"context"
)

func (this *Engine) Ts(fn func(sess ISession) error) error {
	//return this.db.Update(func(tx *buntdb.Tx) error {
	//
	//})
	//
	//
	//
	//s := NewSession(this.db)
	//err := s.Begin()
	//if err != nil {
	//	return err
	//}
	//err = fn(s)
	//if err != nil {
	//	return s.Rollback()
	//}
	//return s.Commit()
	return nil
}

func (this *Engine) Sess(ctx ...context.Context) ISession {
	return NewSession(this.db, ctx...)
}
