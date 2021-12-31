package tgql

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"reflect"
	"regexp"
	"xorm.io/xorm"
)

type LoadXormFetch func(tempData interface{}) map[string]interface{}

func (this *Loader) SetDB(db *xorm.Engine) ILoader {
	this.db = db
	this.dbs = txorm.NewEngine(db)
	this.objectMap = tools.CacheCreate(true)
	return this
}

// LoadXorm 为了方便数据唯一，sqlstr可以给一个前缀, 例如 prefix_xxx select * from user => select * from user
func (this *Loader) LoadXorm(bean interface{}, sqlstr string, fetch LoadXormFetch, param ...interface{}) IObject {
	sess := this.dbs.Session()
	sess.SetId("None")
	return this.LoadXormSess(sess, bean, sqlstr, fetch, param...)
}

// LoadXormSess 为了方便数据唯一，sqlstr可以给一个前缀, 例如 prefix_xxx select * from user => select * from user
func (this *Loader) LoadXormSess(sess txorm.ISession, bean interface{}, sqlstr string, fetch LoadXormFetch, param ...interface{}) IObject {
	info := tools.RftTypeInfo(bean)

	key := fmt.Sprintf("sessId: %s, sql: %s, param: %s, bean.pkg: %s,bean.name: %s", sess.Id(), sqlstr, tools.JSONString(param), info.PkgPath, info.FullName)
	if info.Name == "" {
		key += ",bean.json: " + tools.JSONString(reflect.New(info.Type).Interface())
	}
	key = tools.MD5([]byte(key))
	re := regexp.MustCompile(`^prefix_\S+\s+`)

	return this.LoadObject(key, func(keys []string) (map[string]interface{}, error) {

		sqlstr = re.ReplaceAllString(sqlstr, "")

		s := sess.SF(sqlstr, append(param, map[string]interface{}{"keys": keys})...)

		switch bean.(type) {

		case []map[string]interface{}:
			maps, err := s.Map()
			if err != nil {
				return nil, err
			}
			return fetch(maps), nil

		default:
			data := reflect.New(reflect.TypeOf(bean))
			err := s.Find(data.Interface())
			if err != nil {
				return nil, err
			}
			return fetch(data.Elem().Interface()), nil
		}

	})
}
