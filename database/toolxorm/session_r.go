package toolxorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"reflect"
	"regexp"
	"strings"
)

func (this *Session) SF(sql string, querys ...map[string]interface{}) *Session {
	query := map[string]interface{}{}
	if len(querys) > 0 && querys[0] != nil {
		query = querys[0]
	}
	this.query = query
	this.sql = tools.Str.Tmp(sql, query).String()
	this.sf_args()
	return this
}

func (this *Session) sf_args() *Session {
	r := regexp.MustCompile(`:\S+`)
	ss := r.FindAllString(this.sql, -1)
	for _, s := range ss {
		key := s[1:]
		value := this.query[key]
		this.sf_args_item(s, reflect.ValueOf(value))
	}
	return this
}

func (this *Session) sf_args_item(key string, value reflect.Value) *Session {
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

func (this *Session) Find(bean interface{}) error {
	err := this.Sess.SQL(this.sql, this.args...).Find(bean)
	if this.autoClose {
		this.Sess.Close()
	}
	return err
}
func (this *Session) Exec() error {
	sqls := []interface{}{this.sql}

	_, err := this.Sess.Exec(append(sqls, this.args...)...)
	if this.autoClose {
		this.Sess.Close()
	}
	return err
}
