package load

import (
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"xorm.io/xorm"
)

var _db *xorm.Engine
var _dbs txorm.IEngine
var _cache = tools.NewCache[any](true)

func SetDB(db *xorm.Engine) {
	_db = db
	_dbs = txorm.NewEngine(db)
}
