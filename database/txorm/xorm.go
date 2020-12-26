package txorm

import (
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

// 单例
var newengine *Engine

func NewEngine(db *xorm.Engine, flag ...bool) *Engine {
	if newengine != nil && (len(flag) == 0 || !flag[0]) {
		return newengine
	}
	newengine =  &Engine{DB: db, tmps: map[string]interface{}{}}
	return newengine
}
