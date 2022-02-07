package tgql

import (
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"regexp"
	"strings"
)

var sqlFormatRegexp = regexp.MustCompile(`^[a-zA-Z_]+$`)

func (this *Loader) SqlFormat(sqlstr, field string) string {
	if strings.Index(sqlstr, "select") == -1 && sqlFormatRegexp.MatchString(sqlstr) {

		sqlstr = tools.StrTmp(`
			select {{ .table }}.* from {{ .table }} where {{ .table }}.{{ .field }} in :keys
		`, map[string]interface{}{
			"table": sqlstr,
			"field": field,
		}).String()
	}

	return sqlstr
}

func (this *Loader) LoadXormSessObject(sess txorm.ISession, sqlstr string, field string, param ...interface{}) IObject {

	return this.LoadXormSess(sess, []map[string]interface{}{}, this.SqlFormat(sqlstr, field), func(tempData interface{}) map[string]interface{} {
		data := tempData.([]map[string]interface{})
		result := map[string]interface{}{}
		for i, o := range data {
			key, ok := o[field].(string)
			if !ok {
				continue
			}
			result[key] = data[i]
		}
		return result
	}, param...)
}

func (this *Loader) LoadXormSessSlice(sess txorm.ISession, sqlstr string, field string, param ...interface{}) IObject {

	return this.LoadXormSess(sess, []map[string]interface{}{}, this.SqlFormat(sqlstr, field), func(tempData interface{}) map[string]interface{} {
		data := tempData.([]map[string]interface{})
		tmp := map[string][]map[string]interface{}{}
		for i, o := range data {
			key, ok := o[field].(string)
			if !ok {
				continue
			}
			tmp[key] = append(tmp[key], data[i])
		}

		result := map[string]interface{}{}
		for k, v := range tmp {
			result[k] = v
		}
		return result
	}, param...)
}

func (this *Loader) LoadXormObject(sqlstr string, field string, param ...interface{}) IObject {
	sess := this.dbs.Session(true)
	sess.SetId("None")
	return this.LoadXormSessObject(sess, sqlstr, field, param...)
}

func (this *Loader) LoadXormSlice(sqlstr string, field string, param ...interface{}) IObject {
	sess := this.dbs.Session(true)
	sess.SetId("None")
	return this.LoadXormSessSlice(sess, sqlstr, field, param...)
}
