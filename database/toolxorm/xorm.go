package toolxorm

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Config struct {
	Driver string `yaml:"driver"`
	Uri    string `yaml:"uri"`
}

func NewEngine(cfg Config) (*Engine, error) {
	engine, err := xorm.NewEngine(cfg.Driver, cfg.Uri)
	if err != nil {
		return nil, err
	}

	e := &Engine{DB: engine}
	return e, err
}
