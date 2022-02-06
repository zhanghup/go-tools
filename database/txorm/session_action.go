package txorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"reflect"
	"regexp"
	"strings"
)

func (this *Session) Order(order ...string) ISession {
	this.orderby = order
	return this
}

func (this *Session) SF(sql string, querys ...interface{}) ISession {
	sql = strings.TrimSpace(sql)

	// 重置排序功能
	this.orderby = []string{}

	// sql模板参数格式化
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

	// sql模板格式化
	this.withs = make([]string, 0)
	m1 := map[string]interface{}{
		"tmp": func(name string) string {
			this.withs = append(this.withs, name)
			return fmt.Sprintf("__sql_with_%s", name)
		},
		"ctx": func(name string) string {
			return fmt.Sprintf("{{ ctx_%s .ctx }}", name)
		},
	}
	// tmp模板
	sql = tools.StrTmp(sql, query).FuncMap(tools.MapMerge(m1, this.tmps)).String()
	// context 模板
	this.sql = tools.StrTmp(sql, map[string]interface{}{
		"ctx": this.context,
	}).FuncMap(this.tmpCtxs).String()

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
		if strings.HasPrefix(key, ":between_") {
			if len(args) == 2 {
				this.sql = strings.Replace(this.sql, key, "? and ?", 1)
				this.args = append(this.args, args...)
			}
		} else {
			this.sql = strings.Replace(this.sql, key, fmt.Sprintf("(%s)", strings.Join(ps, ",")), 1)
			this.args = append(this.args, args...)
		}

	default:
		this.sql = strings.Replace(this.sql, key, "?", 1)
		this.args = append(this.args, value.Interface())
	}
	return this
}

func (this *Session) _sql(orderFlag bool) string {
	if !orderFlag || len(this.orderby) == 0 {
		return this.sql
	}
	res := regexp.MustCompile(`\(.*\)`).ReplaceAllString(this.sql, "")
	match := regexp.MustCompile(`order\s+by\s+`).MatchString(res)

	orderBy := make([]string, 0)
	for _, s := range this.orderby {
		if regexp.MustCompile(`^-[a-zA-Z0-9_]+`).MatchString(s) {
			ss := strings.Replace(s, "-", "", 1)
			orderBy = append(orderBy, ss+" desc")
		} else if regexp.MustCompile(`[a-zA-Z0-9_]+`).MatchString(s) {
			orderBy = append(orderBy, s+" asc")
		} else {
			orderBy = append(orderBy, s+" ")
		}
	}
	if match {
		return this.sql + "," + strings.Join(orderBy, ",")
	} else {
		return this.sql + " order by " + strings.Join(orderBy, ",")
	}
}

func (this *Session) _sql_with() string {
	sqlwith := ""
	if len(this.withs) > 0 {
		// 去重
		with_header := "\n with recursive "
		withs := []string{}
		wmap := map[string]bool{}
		for _, w := range this.withs {
			wmap[w] = true
		}
		for k := range wmap {
			kk := tools.StrTmp(fmt.Sprintf("{{ tmp_%s .ctx }}", k), map[string]interface{}{"ctx": this.Ctx()}).FuncMap(this.tmpWiths).String()
			withs = append(withs, fmt.Sprintf("__sql_with_%s as (%s)", k, kk))
		}

		sqlwith = with_header + strings.Join(withs, ",")
		sqlwith = tools.StrTmp(sqlwith, map[string]interface{}{"ctx": this.Ctx()}).FuncMap(this.tmpWiths).String()
	}
	return sqlwith
}
