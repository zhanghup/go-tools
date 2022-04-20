package tgql

import (
	"context"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"sync"
	"time"
	"xorm.io/xorm"
)

// NewLoader db可以为nil，只是不能使用LoadXorm与LoadXormSess
func NewLoader(db *xorm.Engine) ILoader {
	if db == nil {
		return &Loader{objectMap: tools.NewCache[IObject](true)}
	}
	return &Loader{
		db:        db,
		dbs:       txorm.NewEngine(db),
		objectMap: tools.NewCache[IObject](true),
	}
}

/*
	ILoaderXorm
	关于 SQL的prefix，我们可以在sqlstr参数的前面加一个前缀, 例如 prefix_xxx select * from user => select * from user

	为通用数据查询方法，通过数据库查询data_loader数据,bean 查询的结果对象,sqlstr 查询语句,param 查询参数,方法：
		LoadXorm
		LoadXormSess

	通过数据库查询data_loader数据,bean 查询的结果对象,sqlstr 查询语句,param 查询参数，可以使用快速查询，例如：
		user.id
		=>
		prefix_id select user.* from user where user.id in :keys
		=>
		select user.* from user where user.id in :keys

		LoadXormSlice
		LoadXormSess
		LoadXormSessObject
		LoadXormSlice
*/
type ILoaderXorm interface {
	SetDB(db *xorm.Engine) ILoader

	LoadXormCtx(ctx context.Context, bean any, sqlstr string, fetch LoadXormFetch, param ...any) IObject

	LoadXormCtxObject(ctx context.Context, sqlstr string, field string, param ...any) IObject
	LoadXormCtxSlice(ctx context.Context, sqlstr string, field string, param ...any) IObject
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
	objectMap  tools.ICache[IObject]
}

func (this *Loader) LoadObject(id string, fetch ObjectFetch) IObject {
	this.objectSync.Lock()
	defer this.objectSync.Unlock()

	obj, ok := this.objectMap.Get(id)
	if ok {
		return obj.(IObject)
	}

	oo := NewObjectLoader(fetch)
	this.objectMap.Set(id, oo, time.Now().Unix()+86400)

	return oo
}
