package buntdb

import (
	"context"
	"github.com/tidwall/buntdb"
)

type IEngine interface {
	Close() error
	Ts(fn func(sess ISession) error) error
	Sess(ctx ...context.Context) ISession

	CreateStringIndex(name, pattern string, desc ...bool) error
	CreateBinaryIndex(name, pattern string, desc ...bool) error
	CreateIntIndex(name, pattern string, desc ...bool) error
	CreateUintIndex(name, pattern string, desc ...bool) error
	CreateFloatIndex(name, pattern string, desc ...bool) error
}

type Option struct {
	Path string `json:"path" yaml:"path"` // :memory: 表示只使用内存暂存数据
}

type Engine struct {
	opt Option
	db  *buntdb.DB
}

func NewEngine(opt Option) (IEngine, error) {
	db, err := buntdb.Open(opt.Path)
	if err != nil {
		return nil, err
	}

	e := &Engine{
		db:  db,
		opt: opt,
	}

	return e, nil
}

func (this *Engine) Close() error {
	return this.db.Close()
}
