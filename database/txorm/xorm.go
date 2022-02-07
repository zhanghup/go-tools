package txorm

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhanghup/go-tools"
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
	DB         *xorm.Engine
	tmps       map[string]interface{}
	tmpWiths   map[string]interface{}
	tmpCtxs    map[string]interface{}
	tmpsync    sync.RWMutex
	sqliteSync sync.Mutex

	// 当前数据库所有表结构
	tables Tables
}

func NewXorm(cfg Config) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(cfg.Driver, cfg.Uri)
	if err != nil {
		return nil, err
	}
	if cfg.Debug {
		engine.Logger().SetLevel(log.LOG_INFO)
		//engine.SetLogger(log.NewSimpleLogger(tog.Toginfo))
		engine.ShowSQL(true)
	}

	return engine, err
}

type IEngine interface {
	TemplateFuncWith(name string, fn func(ctx context.Context) string) // sql_with_{{name}}
	TemplateFuncCtx(name string, fn func(ctx context.Context) string)  // ctx_{{name}}
	TemplateFunc(name string, f interface{})                           // template func
	TemplateFuncKeys() []string

	Sync(beans ...interface{}) error

	Sess(ctx ...context.Context) ISession
	TS(ctx context.Context, fn func(ctx context.Context, sess ISession) error) error

	Tables() []Table
	Table(name string) Table
	TableColumnExist(table, column string) bool
	DropTables(beans ...interface{}) error
}

// 单例
var newengine *Engine

func NewEngine(db *xorm.Engine, flag ...bool) IEngine {
	db.DBMetas()

	if newengine != nil && (len(flag) == 0 || !flag[0]) {
		return newengine
	}
	if len(flag) > 0 && flag[0] {
		return &Engine{
			DB:       db,
			tmps:     map[string]interface{}{},
			tmpCtxs:  map[string]interface{}{},
			tmpWiths: map[string]interface{}{},
		}
	}

	newengine = &Engine{
		DB:       db,
		tmps:     map[string]interface{}{},
		tmpCtxs:  map[string]interface{}{},
		tmpWiths: map[string]interface{}{},
	}
	return newengine
}

const CONTEXT_SESSION = "context-xorm-session"

func (this *Engine) Sess(ctx ...context.Context) ISession {
	sess := this.DB.Where("")
	newSession := &Session{
		id:        tools.UUID(),
		_engine:   this,
		_db:       this.DB,
		sess:      sess,
		tmps:      this.tmps,
		tmpCtxs:   this.tmpCtxs,
		tmpWiths:  this.tmpWiths,
		autoClose: true,
	}
	c := context.Background()
	if len(ctx) > 0 {
		c = ctx[0]
	}
	newSession.context = c
	return newSession
}

func (this *Engine) Sync(beans ...interface{}) error {
	return this.DB.Sync2(beans...)
}
