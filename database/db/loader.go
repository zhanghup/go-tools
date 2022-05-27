package db

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/loader"
	"reflect"
	"regexp"
	"strings"
	"xorm.io/xorm"
)

var sqlFormatRegexp = regexp.MustCompile(`^\w+$`)

// SliceDB 查找数据库对象,ctx可以为nil
func SliceDB[Result any](db *xorm.Engine, ctx context.Context, beanNameOrSql string, field string, param ...any) loader.IObject[[]Result] {
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

	return loader.Load[[]Result](key, func(keys []string) (map[string][]Result, error) {
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

// Slice 查找数据库对象,ctx可以为nil
func Slice[Result any](ctx context.Context, beanKey, beanNameOrSql string, field string, param ...any) ([]Result, error) {
	l := SliceDB[Result](defaultDB, ctx, beanNameOrSql, field, param...)
	res, ok, err := l.Load(beanKey)
	if err != nil || !ok {
		return nil, err
	}
	return res, nil
}

// InfoDB 查找数据库对象,ctx可以为nil
func InfoDB[Result any](db *xorm.Engine, ctx context.Context, beanNameOrSql string, field string, param ...any) loader.IObject[Result] {
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

	return loader.Load[Result](key, func(keys []string) (map[string]Result, error) {
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

// Info 查找数据库对象,ctx可以为nil
func Info[Result any](ctx context.Context, beanKey, beanNameOrSql string, field string, param ...any) (*Result, error) {
	l := InfoDB[Result](defaultDB, ctx, beanNameOrSql, field, param...)
	res, ok, err := l.Load(beanKey)
	if err != nil || !ok {
		return nil, err
	}
	return &res, nil
}

// InfoId 根据id查找数据库对象,ctx可以为nil
func InfoId[Result any](ctx context.Context, beanKey, beanNameOrSql string, param ...any) (*Result, error) {
	return Info[Result](ctx, beanKey, beanNameOrSql, "id", param...)
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
