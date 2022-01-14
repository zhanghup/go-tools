package buntdb

import (
	"context"
	"github.com/tidwall/buntdb"
)

const CONTEXT_SESSION = "context-buntdb-session"

type ISession interface {
	//Begin(writable ...bool) error
	//Rollback() error
	//Commit() error
	//
	//Set(key, val string, TTL ...time.Duration) (previousValue string, replaced bool, err error)
	//Get(key string) (val string, err error)
	//
	//SetJson(key string, val interface{}, TTL ...time.Duration) (replaced bool, err error)
	//GetJson(key string, result interface{}) (err error)
}

type Session struct {
	ctx  context.Context
	sess *buntdb.Tx
	db   *buntdb.DB

	autoCommit bool
}

// NewSession 若没有传入ctx就会自动提交session
func NewSession(db *buntdb.DB, ctxs ...context.Context) ISession {
	var ctx context.Context
	var s *Session
	if len(ctxs) > 0 {
		cc := ctxs[0].Value(CONTEXT_SESSION)
		if cc != nil {
			cc2, ok := cc.(*Session)
			if ok {
				return cc2
			}
		}
	}

	ctx = context.Background()
	autoCommit := true
	if len(ctxs) > 0 {
		ctx = ctxs[0]
		autoCommit = false
	}

	s = &Session{
		db:         db,
		autoCommit: autoCommit,
	}

	ctx = context.WithValue(ctx, CONTEXT_SESSION, s)
	s.ctx = ctx

	return s
}

//func (this *Session) Begin(writable ...bool) error {
//	update := true
//	if len(writable) > 0 && !writable[0] {
//		update = false
//	}
//
//	tx, err := this.db.Begin(update)
//	if err != nil {
//		return err
//	}
//	this.sess = tx
//	return nil
//}
//
//func (this *Session) Rollback() error {
//	err := this.sess.Rollback()
//	this.sess = nil
//	return err
//}
//
//func (this *Session) Commit() error {
//	err := this.sess.Commit()
//	this.sess = nil
//	return err
//}
//
//func (this *Session) newSess(writable bool) error {
//	if this.sess == nil {
//		return this.Begin(writable)
//	}
//	return nil
//}
//
//func (this *Session) Set(key, val string, TTL ...time.Duration) (previousValue string, replaced bool, err error) {
//	err = this.newSess(true)
//	if err != nil {
//		return
//	}
//
//	var opt *buntdb.SetOptions
//	if len(TTL) > 0 {
//		opt = &buntdb.SetOptions{Expires: true, TTL: TTL[0]}
//	}
//	previousValue, replaced, err = this.sess.Set(key, val, opt)
//	if err != nil {
//		return
//	}
//	if this.autoCommit {
//		err = this.Commit()
//	}
//	return
//}
//
//func (this *Session) Get(key string) (val string, err error) {
//	err = this.newSess(false)
//	if err != nil {
//		return
//	}
//	return this.sess.Get(key)
//}
//
//func (this *Session) SetJson(key string, val interface{}, TTL ...time.Duration) (replaced bool, err error) {
//	newVal := tools.JSONString(val)
//	_, replaced, err = this.Set(key, newVal, TTL...)
//	return
//}
//
//func (this *Session) GetJson(key string, result interface{}) (err error) {
//	val, err := this.Get(key)
//	if err != nil {
//		return
//	}
//	return json.Unmarshal([]byte(val), result)
//}
