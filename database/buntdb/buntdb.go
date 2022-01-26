package buntdb

import (
	"github.com/tidwall/buntdb"
)

type IEngine interface {
	IQuery

	Close() error

	Indexes() ([]string, error)
	IndexCreate(name, pattern string, indexes ...string) error
	IndexRectCreate(name, pattern string) error
	IndexJsonCreate(name, pattern string, indexes ...string) error
	IndexDrop(name string) error

	Ts(fn func(sess ISession) error) error
}

type Engine struct {
	*Query

	opt Option
	db  *buntdb.DB
}

func (this *Engine) Close() error {
	return this.db.Close()
}

type Option struct {
	Path string `json:"path" yaml:"path"` // :memory: 表示只使用内存暂存数据
}

func NewEngine(opt Option) (IEngine, error) {
	db, err := buntdb.Open(opt.Path)
	if err != nil {
		return nil, err
	}

	e := &Engine{
		Query: NewQuery(db),
		db:    db,
		opt:   opt,
	}

	return e, nil
}
