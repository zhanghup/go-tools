package toolxorm

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Config struct {
	Driver string `yaml:"driver"`
	Uri    string `yaml:"uri"`
}

func NewXorm(cfg Config) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(cfg.Driver, cfg.Uri)
	if err != nil {
		return nil, err
	}

	return engine, err
}

func NewEngine(db *xorm.Engine) *Engine {
	return &Engine{DB: db}
}
