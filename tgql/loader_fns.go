package tgql

import (
	"github.com/zhanghup/go-tools/database/txorm"
	"xorm.io/xorm"
)

var loader ILoader = NewLoader(nil)

func LoadObject(id string, fetch ObjectFetch) IObject {
	return loader.LoadObject(id, fetch)
}

func LoadXorm(bean interface{}, sqlstr string, fetch LoadXormFetch, param ...interface{}) IObject {
	return loader.LoadXorm(bean, sqlstr, fetch, param...)
}

func LoadXormObject(sqlstr string, field string, param ...interface{}) IObject {
	return loader.LoadXormObject(sqlstr, field, param...)
}

func LoadXormSlice(sqlstr string, field string, param ...interface{}) IObject {
	return loader.LoadXormSlice(sqlstr, field, param...)
}

func LoadXormSess(sess txorm.ISession, bean interface{}, sqlstr string, fetch LoadXormFetch, param ...interface{}) IObject {
	return loader.LoadXormSess(sess, bean, sqlstr, fetch, param...)
}

func LoadXormSessObject(sess txorm.ISession, sqlstr string, field string, param ...interface{}) IObject {
	return loader.LoadXormSessObject(sess, sqlstr, field, param...)
}

func LoadXormSessSlice(sess txorm.ISession, sqlstr string, field string, param ...interface{}) IObject {
	return loader.LoadXormSessSlice(sess, sqlstr, field, param...)
}

func LoaderInitDB(db *xorm.Engine) {
	loader.SetDB(db)
}
