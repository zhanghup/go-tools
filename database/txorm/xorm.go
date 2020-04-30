package txorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhanghup/go-tools/tog/logger"
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
		engine.ShowSQL(true)
	}

	logopt := logger.OptionStdout()
	logopt.ShowLine = false
	logopt.TimeKey = ""
	logopt.LevelKey = ""
	engine.SetLogger(log.NewSimpleLogger(logger.NewLogger(logopt)))

	return engine, err
}

func NewEngine(db *xorm.Engine) *Engine {
	return &Engine{DB: db}
}
