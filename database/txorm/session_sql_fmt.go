package txorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"reflect"
	"regexp"
	"strings"
)

func (this *Session) SF(sql string, querys ...map[string]interface{}) ISession {
	// 重置排序功能
	this.orderby = []string{}

	query := map[string]interface{}{}
	if len(querys) > 0 && querys[0] != nil {
		query = querys[0]
	}
	this.query = query
	this.sql = tools.StrTmp(sql, query).FuncMap(this.tmps).String()

	this.args = make([]interface{}, 0)
	this.sf_args()
	return this
}

func (this *Session) SF2(sql string, querys ...interface{}) ISession {
	// 重置排序功能
	this.orderby = []string{}

	query := map[string]interface{}{}

	for i := range querys {
		switch querys[i].(type) {
		case map[string]interface{}:
			for k, v := range querys[i].(map[string]interface{}) {
				query[k] = v
			}
		default:
			uid := strings.ReplaceAll(tools.UUID(), "-", "")
			sql = strings.Replace(sql, "?", ":"+uid, 1)
			query[uid] = querys[i]
		}
	}

	this.query = query
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

func (this *Session) sf_args_item(key string, value reflect.Value) ISession {
	ty := value.Type()
	switch ty.Kind() {
	case reflect.Ptr:
		return this.sf_args_item(key, value.Elem())
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
