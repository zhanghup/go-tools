package buntdb

import (
	"github.com/tidwall/buntdb"
	"time"
)

type ISession interface {
	IQuery

	Set(string, string, ...time.Duration) (string, bool, error)
	Delete(string) (string, error)
	DeleteAll() error
}

type Session struct {
	Query
	db        *buntdb.DB
	tx        *buntdb.Tx
	writeAble bool
}

func NewSessionTx(db *buntdb.DB) (ISession, error) {
	tx, err := db.Begin(false)
	if err != nil {
		return nil, err
	}
	return &Session{Query: Query{tx: tx}, db: db, tx: tx}, nil
}
