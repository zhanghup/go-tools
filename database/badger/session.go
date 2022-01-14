package badger

import (
	"context"
	"github.com/dgraph-io/badger/v3"
)

const CONTEXT_SESSION = "context-badger-session"

type ISession interface {
	Begin(writable ...bool) ISession
	Rollback() ISession
	Commit() error

	Set(key, val []byte) error
	SetString(key, val string) error

	Get(key []byte, val func(v []byte) error) error
	GetString(key string, val func(v string) error) error
}

type Session struct {
	ctx  context.Context
	sess *badger.Txn
	db   *badger.DB

	autoCommit bool
}

func NewSession(db *badger.DB, ctxs ...context.Context) ISession {
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

func (this *Session) Begin(writable ...bool) ISession {
	update := true
	if len(writable) > 0 && !writable[0] {
		update = false
	}

	this.sess = this.db.NewTransaction(update)
	return this
}

func (this *Session) Rollback() ISession {
	this.sess.Discard()
	return this
}

func (this *Session) Commit() error {
	return this.sess.Commit()
}

func (this *Session) newSess(writable bool) {
	if this.sess == nil {
		this.Begin(writable)
	}
}

func (this *Session) Set(key, val []byte) (err error) {
	this.newSess(true)
	err = this.sess.Set(key, val)
	if err != nil {
		return err
	}
	if this.autoCommit {
		err = this.Commit()
		this.sess = nil
	}
	return
}

func (this *Session) SetString(key, val string) error {
	return this.Set([]byte(key), []byte(val))
}

func (this *Session) Get(key []byte, val func(v []byte) error) error {
	this.newSess(false)

	v, err := this.sess.Get(key)
	if err != nil {
		return err
	}
	return v.Value(val)
}

func (this *Session) GetString(key string, val func(v string) error) error {
	return this.Get([]byte(key), func(v []byte) error {
		return val(string(v))
	})
}
