package load

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"reflect"
	"regexp"
	"strings"
	"time"
	"xorm.io/xorm"
)

func Load[T any](id string, fetch func(keys []string) (map[string]T, error)) IObject[T] {
	sncKey := "51e761c0-d4ff-478d-923a-14fb5b2bd0af,f3fe7357-2908-4758-8652-1778bb764b27"
	snc := tools.Mutex(sncKey)
	snc.Lock()
	defer snc.Unlock()

	ty := reflect.TypeOf(new(T))
	key := fmt.Sprintf("%s,%s,%s,%s,%s,%s", sncKey, ty.PkgPath(), ty.Name(), ty.String(), ty.Kind().String(), id)

	obj, ok := _cache.Get(key)
	if ok {
		return obj.(IObject[T])
	}

	oo := NewObjectLoader[T](fetch)
	_cache.Set(id, oo, time.Now().Unix()+86400)
	return oo
}

var sqlFormatRegexp = regexp.MustCompile(`^\w+$`)

func SliceDB[Result any](db *xorm.Engine, ctx context.Context, beanNameOrSql string, field string, param ...any) IObject[[]Result] {
	info := tools.RftTypeInfo(make([]Result, 0))
	sess := txorm.NewEngine(db, true).Sess(ctx)
	if sess.IsNew() {
		sess.SetId("None")
	}

	key := fmt.Sprintf("sessId: %s, sql: %s, param: %s, bean.pkg: %s,bean.name: %s", sess.Id(), beanNameOrSql, tools.JSONString(param), info.PkgPath, info.FullName)
	if info.Name == "" {
		key += ",bean.json: " + tools.JSONString(reflect.New(info.Type).Interface())
	}
	key = tools.MD5([]byte(key))

	return Load[[]Result](key, func(keys []string) (map[string][]Result, error) {
		res := make([]struct {
			Info Result `xorm:"extends"`
			Nid  string `xorm:"_B51e761c0"`
		}, 0)
		err := sess.SF(sqlFormat(beanNameOrSql, field), append(param, map[string]any{"keys": keys})...).Find(&res)

		result := map[string][]Result{}

		if err != nil {
			return result, nil
		}

		for _, o := range res {
			result[o.Nid] = append(result[o.Nid], o.Info)
		}

		return result, err
	})
}

func Slice[Result any](ctx context.Context, beanNameOrSql string, field string, param ...any) IObject[[]Result] {
	return SliceDB[Result](_db, ctx, beanNameOrSql, field, param...)
}

func InfoDB[Result any](db *xorm.Engine, ctx context.Context, beanNameOrSql string, field string, param ...any) IObject[Result] {
	info := tools.RftTypeInfo(make([]Result, 0))
	sess := txorm.NewEngine(db, true).Sess(ctx)
	if sess.IsNew() {
		sess.SetId("None")
	}

	key := fmt.Sprintf("sessId: %s, sql: %s, param: %s, bean.pkg: %s,bean.name: %s", sess.Id(), beanNameOrSql, tools.JSONString(param), info.PkgPath, info.FullName)
	if info.Name == "" {
		key += ",bean.json: " + tools.JSONString(reflect.New(info.Type).Interface())
	}
	key = tools.MD5([]byte(key))

	return Load[Result](key, func(keys []string) (map[string]Result, error) {
		res := make([]struct {
			Info Result `xorm:"extends"`
			Nid  string `xorm:"_B51e761c0"`
		}, 0)
		err := sess.SF(sqlFormat(beanNameOrSql, field), append(param, map[string]any{"keys": keys})...).Find(&res)

		result := map[string]Result{}

		if err != nil {
			return result, nil
		}

		for _, o := range res {
			result[o.Nid] = o.Info
		}

		return result, err
	})
}

func Info[Result any](ctx context.Context, beanNameOrSql string, field string, param ...any) IObject[Result] {
	return InfoDB[Result](_db, ctx, beanNameOrSql, field, param...)
}

func sqlFormat(sqlstr, field string) string {
	sqlstr = regexp.MustCompile(`^prefix_\S+\s+`).ReplaceAllString(sqlstr, "")

	if strings.Index(sqlstr, "select") == -1 && sqlFormatRegexp.MatchString(sqlstr) {
		sqlstr = tools.StrTmp(`
			select {{ .table }}.*,{{ .table }}.{{ .field }} _B51e761c0 from {{ .table }} where {{ .table }}.{{ .field }} in :keys
		`, map[string]any{
			"table": sqlstr,
			"field": field,
		}).String()
	} else {
		sqlstr = fmt.Sprintf(`select s.*,s.%s _B51e761c0 from (%s) s`, field, sqlstr)
	}
	return sqlstr
}
