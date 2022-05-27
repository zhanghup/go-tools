package db

import (
	"context"
	_ "embed"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

var defaultEngine txorm.IEngine
var defaultDB *xorm.Engine

func Init(ymlData ...[]byte) {
	e, err := txorm.InitXorm(ymlData...)
	if err != nil {
		panic(err)
	}
	defaultDB = e
	defaultEngine = txorm.NewEngine(e, true)
}

func DB(ctx ...context.Context) txorm.ISession {
	if defaultEngine == nil {
		panic("")
	}
	return defaultEngine.New(ctx...)
}

func Sess(ctx ...context.Context) txorm.ISession {
	return defaultEngine.Sess(ctx...)
}

func Ts(ctx context.Context, fn func(ctx context.Context, sess txorm.ISession) error) error {
	return defaultEngine.TS(ctx, fn)
}

func Initialized(bean any) bool {
	t, err := defaultDB.TableInfo(bean)
	if err != nil {
		tog.Error(err.Error())
		return false
	}
	return t != nil
}

func Sync(o ...any) error {
	return defaultEngine.Sync(o...)
}

func DropTables(o ...any) error {
	return defaultEngine.DropTables(o...)
}
