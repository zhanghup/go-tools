package txorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"reflect"
	"regexp"
	"strings"
)

func (this *Session) SF(sql string, querys ...map[string]interface{}) ISession {
	query := map[string]interface{}{}
	if len(querys) > 0 && querys[0] != nil {
		query = querys[0]
	}
	this.query = query
	this.sql = tools.Str.Tmp(sql, query).FuncMap(this.tmps).String()

	if len(this.withs) > 0 {
		// 去重
		withs := []string{"\n with recursive _ as (select 1)"}
		wmap := map[string]bool{}
		for _, w := range this.withs {
			wmap[w] = true
		}
		for k := range wmap {
			withs = append(withs, k)
		}

		this.sqlwith = strings.Join(withs, ",")
		this.sqlwith = tools.Str.Tmp(this.sqlwith, query).FuncMap(this.tmps).String()
	}

	this.args = make([]interface{}, 0)
	this.sf_args()
	return this
}

func (this *Session) sf_args() ISession {
	r := regexp.MustCompile(`:[0-1a-zA-Z]+`)
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

func (this *Session) Find(bean interface{}) error {
	if this.autoClose {
		defer this.sess.Close()
	}
	err := this.sess.SQL(this.sqlwith+" "+this.sql, this.args...).Find(bean)
	return err
}

func (this *Session) Page(index, size int, count bool, bean interface{}) (int, error) {
	if this.autoClose {
		defer this.sess.Close()
	}
	if size < 0 {
		err := this.sess.SQL(this.sql, this.args...).Find(bean)
		return 0, err
	} else if size == 0 {
		total := 0
		_, err := this.sess.SQL(fmt.Sprintf("%s select count(1) from (%s) _", this.sqlwith, this.sql), this.args...).Get(&total)
		return total, err
	} else {
		err := this.sess.SQL(fmt.Sprintf("%s limit ?,?", this.sql), append(this.args, (index-1)*size, size)...).Find(bean)
		if err != nil {
			return 0, err
		}
		if count {
			total := 0
			_, err := this.sess.SQL(fmt.Sprintf("%s select count(1) from (%s) _", this.sqlwith, this.sql), this.args...).Get(&total)
			return total, err
		}
	}

	return 0, nil
}

func (this *Session) Page2(index, size *int, count *bool, bean interface{}) (int, error) {
	if index == nil {
		index = tools.Ptr.Int(1)
	}
	if size == nil {
		size = tools.Ptr.Int(1)
	}
	if count == nil {
		count = tools.Ptr.Bool(false)
	}
	return this.Page(*index, *size, *count, bean)
}

func (this *Session) Exec() error {
	if this.autoClose {
		defer this.sess.Close()
	}

	sqls := []interface{}{this.sql}
	_, err := this.sess.Exec(append(sqls, this.args...)...)
	return err
}
