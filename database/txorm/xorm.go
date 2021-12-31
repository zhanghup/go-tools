package txorm

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhanghup/go-tools"
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
		engine.SetLogger(log.NewSimpleLogger(tog.Toginfo))
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

	SessionAuto(ctx ...context.Context) ISession
	Session(ctx ...context.Context) ISession

	Tables() ([]Table, error)
	Table(name string) (Table, error)
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

const CONTEXT_SESSION = "context-session"

func (this *Engine) SessionAuto(ctx ...context.Context) ISession {
	return this._session(true, ctx...)
}

func (this *Engine) Session(ctx ...context.Context) ISession {
	return this._session(false, ctx...)
}

func (this *Engine) Sync(beans ...interface{}) error {
	return this.DB.Sync2(beans...)
}

func (this *Engine) _session(autoClose bool, ctx ...context.Context) *Session {

	if len(ctx) > 0 && ctx[0] != nil {
		c := ctx[0]
		v := c.Value(CONTEXT_SESSION)
		if v != nil {
			oldSession, ok := v.(*Session)
			if ok {
				if !oldSession.sess.IsClosed() {
					return oldSession
				}
			}
		}
	}

	newSession := &Session{
		id:             tools.UUID(),
		_engine:        this,
		_db:            this.DB,
		sess:           this.DB.NewSession(),
		tmps:           this.tmps,
		tmpCtxs:        this.tmpCtxs,
		tmpWiths:       this.tmpWiths,
		autoClose:      autoClose,
		beginTranslate: false,
	}
	if len(ctx) > 0 {
		newSession.context = ctx[0]
	} else {
		c := context.Background()
		c = context.WithValue(c, CONTEXT_SESSION, newSession)
		newSession.context = c
	}
	return newSession
}
