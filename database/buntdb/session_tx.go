package buntdb

import (
	"github.com/tidwall/buntdb"
	"github.com/zhanghup/go-tools"
	"time"
)

func (this *Session) TxWriteable() (err error) {

	if this.writeAble {
		return
	} else {
		err = this.tx.Rollback()
		if err != nil {
			return err
		}
	}
	this.writeAble = true
	this.tx, err = this.db.Begin(true)
	this.Query.tx = this.tx
	if err != nil {
		return
	}
	return nil
}

func (this *Session) Set(key, value string, ttl ...time.Duration) (previousValue string, replaced bool, err error) {
	var opt *buntdb.SetOptions
	if len(ttl) > 0 {
		opt = &buntdb.SetOptions{Expires: true, TTL: ttl[0]}
	}
	err = this.TxWriteable()
	if err != nil {
		return
	}
	return this.tx.Set(key, value, opt)
}

func (this *Session) SetJson(key string, value any, ttl ...time.Duration) (previousValue string, replaced bool, err error) {
	v := tools.JSONString(value)
	return this.Set(key, v, ttl...)
}

func (this *Session) DeleteAll() (err error) {
	err = this.TxWriteable()
	if err != nil {
		return
	}
	return this.tx.DeleteAll()
}

func (this *Session) Delete(key string) (val string, err error) {
	err = this.TxWriteable()
	if err != nil {
		return
	}
	return this.tx.Delete(key)
}

func (this *Session) Commit() (err error) {
	if this.writeAble {
		return this.tx.Commit()
	}
	return nil
}

func (this *Session) Rollback() (err error) {
	if this.writeAble {
		return this.tx.Rollback()
	}
	return nil
}
