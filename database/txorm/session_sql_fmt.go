package txorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"reflect"
	"regexp"
	"strings"
)

func (this *Session) SF(sql string, querys ...interface{}) ISession {
	// 重置排序功能
	this.orderby = []string{}

	query := map[string]interface{}{}

	for i := range querys {
		ty := reflect.TypeOf(querys[i])
		if ty.Kind() == reflect.Map {
			vl := reflect.ValueOf(querys[i])
			for _, key := range vl.MapKeys() {
				v := vl.MapIndex(key)
				query[key.String()] = v.Interface()
			}
		} else {
			uid := strings.ReplaceAll(tools.UUID(), "-", "_")
			sql = strings.Replace(sql, "?", ":"+uid, 1)
			query[uid] = querys[i]
		}
	}

	this.query = query

	sql = this.sql_ctx(sql)
	this.sql = tools.StrTmp(sql, query).FuncMap(this.tmps).String()

	this.args = make([]interface{}, 0)
	this.sf_args()
	return this
}

func (this *Session) sf_args() ISession {
	r := regexp.MustCompile(`:[0-9a-zA-Z_]+`)
	ss := r.FindAllString(this.sql, -1)
	for _, s := range ss {
		key := s[1:]
		value := this.query[key]
		if value == nil {
			continue
		}
		this.sf_args_item(s, reflect.ValueOf(value))
	}
	return this
}

func (this *Session) sql_ctx(sql string) string {
	r := regexp.MustCompile(`{{\s*ctx_\S+\s*}}`)
	ss := r.FindAllString(sql, -1)

	for _, s := range ss {
		rdata := tools.StrTmp(strings.Replace(s, "}}", " .ctx }}", 1), map[string]interface{}{"ctx": this.Ctx()}).FuncMap(this.tmpCtxs).String()
		sql = strings.Replace(sql, s, rdata, -1)
	}
	return sql
}

func (this *Session) sf_args_item(key string, value reflect.Value) ISession {
	ty := value.Type()
	switch ty.Kind() {
	case reflect.Ptr:
		if value.Pointer() == 0 {
			this.sql = strings.Replace(this.sql, key, "?", 1)
			this.args = append(this.args, nil)
		} else {
			return this.sf_args_item(key, value.Elem())
		}
	case reflect.Array, reflect.Slice:
		ps := []string{}
		args := []interface{}{}
		for i := 0; i < value.Len(); i++ {
			v := value.Index(i)
			ps = append(ps, "?")
			args = append(args, v.Interface())
		}
		this.sql = strings.Replace(this.sql, key, fmt.Sprintf("(%s)", strings.Join(ps, ",")), 1)
		this.args = append(this.args, args...)
	default:
		this.sql = strings.Replace(this.sql, key, "?", 1)
		this.args = append(this.args, value.Interface())
	}
	return this
}
