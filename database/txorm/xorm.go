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
	tmps       map[string]any
	tmpWiths   map[string]any
	tmpCtxs    map[string]any
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
	TemplateFunc(name string, f any)                                   // template func
	TemplateFuncKeys() []string

	New(ctx ...context.Context) ISession
	Sess(ctx ...context.Context) ISession
	TS(ctx context.Context, fn func(ctx context.Context, sess ISession) error) error

	//Tables() []Table
	//Table(name string) Table
	//TableColumnExist(table, column string) bool
	DropTables(beans ...any) error
	Sync(beans ...any) error
}

// 单例
var newengine *Engine

func NewEngine(db *xorm.Engine, flag ...bool) IEngine {
	if newengine != nil && (len(flag) == 0 || !flag[0]) {
		return newengine
	}
	if len(flag) > 0 && flag[0] {
		return &Engine{
			DB:       db,
			tmps:     map[string]any{},
			tmpCtxs:  map[string]any{},
			tmpWiths: map[string]any{},
		}
	}

	newengine = &Engine{
		DB:       db,
		tmps:     map[string]any{},
		tmpCtxs:  map[string]any{},
		tmpWiths: map[string]any{},
	}
	return newengine
}

const CONTEXT_SESSION = "context-xorm-session"

func (this *Engine) New(ctx ...context.Context) ISession {
	return this.session(true, true, ctx...)
}
func (this *Engine) Sess(ctx ...context.Context) ISession {
	return this.session(true, false, ctx...)
}

func (this *Engine) Sync(beans ...any) error {
	return this.DB.Sync2(beans...)
}

/*
	session
	@autoClose: 是否自动关闭session
	@autoNew: 是否直接创建一个新的session，若false则在context中寻找一个旧的session，没有找到再创建一个新的
*/
func (this *Engine) session(autoClose, new bool, ctx ...context.Context) *Session {

	if !new && len(ctx) > 0 && ctx[0] != nil {
		v := ctx[0].Value(CONTEXT_SESSION)
		if v != nil {
			sessOld, ok := v.(*Session)
			if ok && !sessOld.sess.IsClosed() {
				sessOld.isNew = false
				return sessOld
			}
		}
	}

	newSession := &Session{
		id:        tools.UUID(),
		_engine:   this,
		_db:       this.DB,
		sess:      this.DB.NewSession(),
		tmps:      this.tmps,
		tmpCtxs:   this.tmpCtxs,
		tmpWiths:  this.tmpWiths,
		autoClose: autoClose,
		isNew:     true,
	}
	c := context.Background()
	if len(ctx) > 0 && ctx[0] != nil {
		c = context.WithValue(ctx[0], CONTEXT_SESSION, newSession)
	} else {
		c = context.WithValue(c, CONTEXT_SESSION, newSession)
	}
	newSession.context = c
	return newSession
}
