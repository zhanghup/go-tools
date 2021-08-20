package loader

import (
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"sync"
	"time"
	"xorm.io/xorm"
)

func NewLoader(db *xorm.Engine) ILoader {
	return &Loader{
		db:        db,
		dbs:       txorm.NewEngine(db),
		objectMap: tools.CacheCreate(true),
	}
}

type ILoader interface {
	// LoadObject loader 方法的id，必须保证每个独立的使用方法使用不同的id
	LoadObject(id string, fetch ObjectFetch) IObject
	// LoadXorm 通过数据库查询data_loader数据,bean 查询的结果对象,sqlstr 查询语句,param 查询参数
	LoadXorm(bean interface{}, sqlstr string, fetch LoadXormFetch, param ...interface{}) IObject
	// LoadXormSess 通过Xorm Session的方式读取数据
	LoadXormSess(sess txorm.ISession, bean interface{}, sqlstr string, fetch LoadXormFetch, param ...interface{}) IObject
}

type Loader struct {
	db  *xorm.Engine
	dbs txorm.IEngine

	objectSync sync.Mutex
	objectMap  tools.ICache
}

func (this *Loader) LoadObject(id string, fetch ObjectFetch) IObject {
	this.objectSync.Lock()
	defer this.objectSync.Unlock()

	obj := this.objectMap.Get(id)
	if obj != nil {
		return obj.(IObject)
	}

	oo := NewObjectLoader(fetch)
	this.objectMap.Set(id, oo, time.Now().Unix()+86400)

	return oo
}
