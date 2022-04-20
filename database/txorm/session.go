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

	Table(bean any) ISession
	Find(bean any) error
	Get(bean any) (bool, error)

	Insert(bean ...any) error
	Update(bean any, condiBean ...any) error
	Delete(bean ...any) error
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
	SF(sql string, querys ...any) ISession
	Order(order ...string) ISession

	Page(index, size int, count bool, bean any) (int, error)
	Page2(index, size *int, count *bool, bean any) (int, error)
	Count() (int64, error)
	Int() (int, error)
	Int64() (int64, error)
	Float64() (float64, error)
	String() (string, error)
	Strings() ([]string, error)
	Exists() (bool, error)

	Map() ([]map[string]any, error)
	MapString() (v []map[string]string, err error)
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
	query     map[string]any
	args      []any

	autoClose bool

	tmps     map[string]any
	tmpWiths map[string]any
	tmpCtxs  map[string]any

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

func (this *Session) Table(bean any) ISession {
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
