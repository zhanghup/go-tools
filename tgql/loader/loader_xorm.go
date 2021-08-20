package loader

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"reflect"
)

type LoadXormFetch func(tempData interface{}) map[string]interface{}

func (this *Loader) LoadXorm(bean interface{}, sqlstr string, fetch LoadXormFetch, param ...interface{}) IObject {
	sess := this.dbs.NewSession(true)
	sess.SetId("None")
	return this.LoadXormSess(sess, bean, sqlstr, fetch, param...)
}

func (this *Loader) LoadXormSess(sess txorm.ISession, bean interface{}, sqlstr string, fetch LoadXormFetch, param ...interface{}) IObject {
	info := tools.RftTypeInfo(bean)

	key := fmt.Sprintf("sessId: %s, sql: %s, param: %s, bean.pkg: %s,bean.name: %s", sess.Id(), sqlstr, tools.JSONString(param), info.PkgPath, info.FullName)
	if info.Name == "" {
		key += ",bean.json: " + tools.JSONString(reflect.New(info.Type).Interface())
	}
	key = tools.MD5([]byte(key))

	return this.LoadObject(key, func(keys []string) (map[string]interface{}, error) {

		data := reflect.New(reflect.TypeOf(bean))
		err := sess.SF(sqlstr, append(param, map[string]interface{}{"keys": keys})...).Find(data.Interface())
		if err != nil {
			return nil, err
		}
		return fetch(data.Elem().Interface()), nil
	})
}
