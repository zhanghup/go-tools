package badger

import (
	"context"
	"github.com/dgraph-io/badger/v3"
)

type IEngine interface {
	Sess(ctx ...context.Context) ISession
	Ts(fn func(sess ISession) error) error
	Close() error
}

type Option struct {
	Path string `json:"path" yaml:"path"` // :memory: 表示只使用内存暂存数据
}

type Engine struct {
	opt Option
	db  *badger.DB
}

func NewEngine(opt Option) (e IEngine, err error) {
	var badgerOption badger.Options
	if opt.Path == ":memory:" {
		badgerOption = badger.DefaultOptions("").WithInMemory(true)
	} else {
		badgerOption = badger.DefaultOptions(opt.Path)
	}
	db, err := badger.Open(badgerOption)
	if err != nil {
		return
	}

	e = &Engine{
		opt: opt,
		db:  db,
	}

	return
}

func (this *Engine) Close() error {
	return this.db.Close()
}
