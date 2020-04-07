package toolxorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"reflect"
	"regexp"
	"strings"
	"xorm.io/xorm"
)

type Engine struct {
	DB *xorm.Engine
}
type Session struct {
	DB    *xorm.Engine
	sql   string
	query map[string]interface{}
	args  []interface{}
}

func (this *Engine) SF(sql string, querys ...map[string]interface{}) *Session {
	return (&Session{DB: this.DB}).SF(sql, querys...)
}

func (this *Session) SF(sql string, querys ...map[string]interface{}) *Session {
	query := map[string]interface{}{}
	if len(querys) > 0 {
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
	return this.DB.SQL(this.sql, this.args...).Find(bean)
}
