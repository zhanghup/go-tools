package txorm

import (
	"context"
	"github.com/zhanghup/go-tools"
	"sync"
	"xorm.io/xorm"
)

type ISession interface {
	Id() string
	SetId(id string)

	Ctx() context.Context
	Begin()
	Rollback() error
	Commit() error
	Close() error

	Table(bean interface{}) ISession
	Find(bean interface{}) error
	Get(bean interface{}) (bool, error)

	Insert(bean ...interface{}) error
	Update(bean interface{}, condiBean ...interface{}) error
	Delete(bean ...interface{}) error
	TS(fn func(sess ISession) error) error
	Exec() error
	/*
		示例1：
			sql = "select * from user where a = ? and b = ?"
			querys = []interface{}{"a","b"}
		示例2：
			sql = "select * from user where a = :a and b = ?"
			querys = []interface{}{"b",map[string]interface{}{"a":"a"}}
		示例3：
			sql = "where a = ?"
			querys = []interface{}{"b"}
			bean = models.User
			>>> select user.* from user where a = ?

	*/
	SF(sql string, querys ...interface{}) ISession
	Order(order ...string) ISession

	Page(index, size int, count bool, bean interface{}) (int, error)
	Page2(index, size *int, count *bool, bean interface{}) (int, error)
	Count() (int64, error)
	Int() (int, error)
	Int64() (int64, error)
	Float64() (float64, error)
	String() (string, error)
	Strings() ([]string, error)
	Exists() (bool, error)
	Map() ([]map[string]interface{}, error)
}

type Session struct {
	id             string
	context        context.Context
	beginTranslate bool // 配置是否需要开启事务
	openTranslate  bool // 是否包含有事务类操作，然后主动开启了事务

	// xorm session
	sess    *xorm.Session
	_engine *Engine
	_db     *xorm.Engine
	_sync   sync.Mutex

	tableName string
	sql       string
	query     map[string]interface{}
	args      []interface{}

	autoClose bool

	tmps     map[string]interface{}
	tmpWiths map[string]interface{}
	tmpCtxs  map[string]interface{}

	withs   []string
	orderby []string
}

func (this *Session) Ctx() context.Context {
	if this.context == nil {
		this.context = context.Background()
	}
	return context.WithValue(this.context, CONTEXT_SESSION, this)
}

// begin 当事务总包含有操作雷逻辑的时候，自动开启事务（前提是需要开启）
func (this *Session) begin(fn func() error) error {
	this._sync.Lock()
	defer this._sync.Unlock()

	// 判断是否需要开启事务
	if this.beginTranslate && !this.openTranslate {
		this._engine.lock()
		if err := this.sess.Begin(); err != nil {
			this._engine.unlock()
			return err
		}
		this.openTranslate = true

	}

	// 执行逻辑
	err := fn()
	if err != nil {
		return err
	}

	// 判断是否需要关闭session
	if this.autoClose {
		return this.Close()
	}
	return nil
}

func (this *Session) Begin() {
	// 防止并发开启事务
	this._sync.Lock()
	defer this._sync.Unlock()

	// 若当前已经开启事务，则无需再次开启
	if this.beginTranslate {
		return
	}

	// 准备使用事务，但此时并未开启，需要等到真的有事务的时候再开启
	this.beginTranslate = true
	return
}

func (this *Session) Rollback() error {
	// 若事务并没有开启，跳出
	if !this.beginTranslate {
		return nil
	}
	// 若没有执行任何事务操作，代表本次事务中没有需要开启事务的操作，跳出
	if !this.openTranslate {
		return nil
	}

	// 还原操作
	if err := this.sess.Rollback(); err != nil {
		return err
	}

	// 关闭事务开启状态
	this._engine.unlock()
	this.beginTranslate = false
	this.openTranslate = false
	return nil
}

func (this *Session) Commit() error {
	// 若事务并没有开启，跳出
	if !this.beginTranslate {
		return nil
	}
	// 若没有执行任何事务操作，代表本次事务中没有需要开启事务的操作，跳出
	if !this.openTranslate {
		return nil
	}

	// 提交事务
	if err := this.sess.Commit(); err != nil {
		return err
	}

	// 关闭事务开启状态
	this._engine.unlock()
	this.beginTranslate = false
	this.openTranslate = false
	return nil

}

func (this *Session) AutoClose(fn func() error) error {
	err := fn()
	if err != nil {
		return err
	}

	if this.autoClose {
		return this.Close()
	}
	return nil
}

// Close 自动关闭session
func (this *Session) Close() error {
	if this.sess.IsClosed() {
		return nil
	}
	return this.sess.Close()
}

func (this *Session) Id() string {
	return this.id
}

func (this *Session) SetId(id string) {
	this.id = id
}

func (this *Session) Table(bean interface{}) ISession {
	switch bean.(type) {
	case string:
		this.tableName = bean.(string)
	case *string:
		this.tableName = *(bean.(*string))
	default:
		tab := tools.RftTypeInfo(bean)
		newTable := this._db.GetTableMapper().Obj2Table(tab.Name)
		if newTable != "" {
			this.tableName = this._db.GetTableMapper().Obj2Table(tab.Name)
		}
	}

	return this
}
