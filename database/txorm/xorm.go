package txorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type Config struct {
	Driver string `yaml:"driver"`
	Uri    string `yaml:"uri"`
	Debug  bool   `yaml:"debug"`
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

var newengine *Engine

func NewEngine(db *xorm.Engine, flag ...bool) *Engine {
	if newengine != nil && (len(flag) == 0 || !flag[0]) {
		return newengine
	}
	return &Engine{DB: db, tmps: map[string]interface{}{}}
}

func (this *Engine) TemplateFuncAdd(name string, f interface{}) {
	this.tmpsync.Lock()
	this.tmps[name] = f
	this.tmpsync.Unlock()
}

func (this *Engine) TemplateFuncKeys() []string {
	this.tmpsync.RLock()
	keys := make([]string, len(this.tmps))
	for k := range this.tmps {
		keys = append(keys, k)
	}
	this.tmpsync.RUnlock()
	return keys
}
