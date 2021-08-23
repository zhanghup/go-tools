package loader

import (
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"sync"
	"time"
	"xorm.io/xorm"
)

// NewLoader db可以为nil，只是不能使用LoadXorm与LoadXormSess
func NewLoader(db *xorm.Engine) ILoader {
	return &Loader{
		db:        db,
		dbs:       txorm.NewEngine(db),
		objectMap: tools.CacheCreate(true),
	}
}

/*
	ILoaderXorm
	关于 SQL的prefix，我们可以在sqlstr参数的前面加一个前缀, 例如 prefix_xxx select * from user => select * from user
	方法：LoadXorm，LoadXormSess 为通用数据查询方法，通过数据库查询data_loader数据,bean 查询的结果对象,sqlstr 查询语句,param 查询参数

 */
type ILoaderXorm interface {
	LoadXorm(bean interface{}, sqlstr string, fetch LoadXormFetch, param ...interface{}) IObject

	LoadXormObject(sqlstr string, field string, param ...interface{}) IObject

	// LoadXormSlice 通过数据库查询data_loader数据,bean 查询的结果对象,sqlstr 查询语句,param 查询参数
	// 为了方便数据唯一，sqlstr可以给一个前缀, 例如 prefix_xxx select * from user => select * from user
	LoadXormSlice(sqlstr string, field string, param ...interface{}) IObject

	// LoadXormSess 通用方法，通过Xorm Session的方式读取数据
	// 为了方便数据唯一，sqlstr可以给一个前缀, 例如 prefix_xxx select * from user => select * from user
	LoadXormSess(sess txorm.ISession, bean interface{}, sqlstr string, fetch LoadXormFetch, param ...interface{}) IObject

	// LoadXormSessObject 通过Xorm Session的方式读取数据
	// 为了方便数据唯一，sqlstr可以给一个前缀, 例如 prefix_xxx select * from user => select * from user
	LoadXormSessObject(sess txorm.ISession, sqlstr string, field string, param ...interface{}) IObject

	// LoadXormSessSlice 通过Xorm Session的方式读取数据
	// 为了方便数据唯一，sqlstr可以给一个前缀, 例如 prefix_xxx select * from user => select * from user
	LoadXormSessSlice(sess txorm.ISession, sqlstr string, field string, param ...interface{}) IObject
}

type ILoader interface {

	// LoadObject loader 方法的id，必须保证每个独立的使用方法使用不同的id
	LoadObject(id string, fetch ObjectFetch) IObject


	ILoaderXorm
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
