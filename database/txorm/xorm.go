package txorm

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhanghup/go-tools/tog"
	"sync"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type Config struct {
	Driver string `yaml:"driver"`
	Uri    string `yaml:"uri"`
	Debug  bool   `yaml:"debug"`
}

type Engine struct {
	DB      *xorm.Engine
	tmps    map[string]interface{}
	tmpsync sync.RWMutex
}

func NewXorm(cfg Config) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(cfg.Driver, cfg.Uri)
	if err != nil {
		return nil, err
	}
	if cfg.Debug {
		engine.Logger().SetLevel(log.LOG_INFO)
		engine.SetLogger(log.NewSimpleLogger(tog.Toginfo))
		engine.ShowSQL(true)
	}

	return engine, err
}

type IEngine interface {
	TemplateFuncAdd(name string, f interface{})
	TemplateFuncKeys() []string
	NewSession(autoClose bool, ctx ...context.Context) ISession
	Session(ctx ...context.Context) ISession
	TS(fn func(sess ISession) error) error
	SF(sql string, querys ...map[string]interface{}) ISession
	With(name string) ISession
	Engine() xorm.EngineInterface
}

// 单例
var newengine *Engine

func NewEngine(db *xorm.Engine, flag ...bool) IEngine {
	if newengine != nil && (len(flag) == 0 || !flag[0]) {
		return newengine
	}
	if len(flag) > 0 && flag[0] {
		return &Engine{DB: db, tmps: map[string]interface{}{}}
	}

	newengine = &Engine{DB: db, tmps: map[string]interface{}{}}
	return newengine
}
