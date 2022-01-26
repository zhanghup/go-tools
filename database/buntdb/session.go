package buntdb

import (
	"github.com/tidwall/buntdb"
	"time"
)

type ISession interface {
	IQuery

	Set(string, string, ...time.Duration) (string, bool, error)
	SetJson(key string, value interface{}, ttl ...time.Duration) (previousValue string, replaced bool, err error)

	Delete(string) (string, error)
	DeleteAll() error
}

type Session struct {
	*Query
	db        *buntdb.DB
	tx        *buntdb.Tx
	writeAble bool
}

func NewSession(db *buntdb.DB) (ISession, error) {
	tx, err := db.Begin(false)
	if err != nil {
		return nil, err
	}
	return &Session{Query: &Query{tx: tx}, db: db, tx: tx}, nil
}
