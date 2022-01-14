package buntdb

import (
	"github.com/tidwall/buntdb"
	"time"
)

type ISessionTx interface {
	Set(key, value string, ttl ...time.Duration) (previousValue string, replaced bool, err error)
	DeleteAll() error
}

type SessionTx struct {
	sess *buntdb.Tx
}

func (this *SessionTx) Set(key, value string, ttl ...time.Duration) (previousValue string, replaced bool, err error) {
	var opt *buntdb.SetOptions
	if len(ttl) > 0 {
		opt = &buntdb.SetOptions{Expires: true, TTL: ttl[0]}
	}
	return this.sess.Set(key, value, opt)
}

func (this *SessionTx) DeleteAll() error {
	return this.sess.DeleteAll()
}
