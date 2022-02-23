package txorm

import (
	"context"
	"github.com/zhanghup/go-tools"
	"xorm.io/xorm"
)

type ISession interface {
	Id() string
	SetId(id string)
	IsNew() bool
	Ctx() context.Context

	Table(bean interface{}) ISession
	Find(bean interface{}) error
	Get(bean interface{}) (bool, error)

	Insert(bean ...interface{}) error
	Update(bean interface{}, condiBean ...interface{}) error
	Delete(bean ...interface{}) error
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
	id      string
	context context.Context
	isNew   bool

	// xorm session
	sess    *xorm.Session
	_engine *Engine
	_db     *xorm.Engine

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

func (this *Session) IsNew() bool {
	return this.isNew
}
func (this *Session) Ctx() context.Context {
	if this.context == nil {
		this.context = context.Background()
	}
	return context.WithValue(this.context, CONTEXT_SESSION, this)
}

func (this *Session) AutoClose(fn func() error) error {
	err := fn()
	if err != nil {
		return err
	}
	this.tableName = ""

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
	if this.tableName != "" {
		return this
	}

	switch bean.(type) {
	case string:
		this.tableName = bean.(string)
	case *string:
		this.tableName = *(bean.(*string))
	default:
		tab := tools.RftTypeInfo(bean)
		this.tableName = this._db.GetTableMapper().Obj2Table(tab.Name)
	}

	return this
}
